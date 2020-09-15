# 分布式锁

[TOC]

## 单体架构下的锁

锁是用于保护资源避免并发竞争的工具，在JVM的发展过程中锁的演化更是历久弥新：当下，可以用公平锁实现顺序排队、用非公平锁实现优先级管理、用可重入锁减少死锁的几率、用读写锁提升读取性能，JVM更是实现了从偏向锁到轻量级锁再到重量级锁的逐渐膨胀优化，更多的情况下我们还可以基于 CAS 的原子操作包（java.util.concurrent.atomic）实现无锁（乐观锁）编程。

单机环境下，锁只作用于同一JVM；在分布式系统中锁要跨JVM作用于多个节点，实现的难度及成本都要远高于前者。

在谈分布式锁之前我们必须明，如非必要切勿用锁：

- 一方面锁会将并行逻辑转成串行严重影响性能，

- 另一方面还要考虑锁的容错，处理不当可能导致死锁。

如果可能笔者更推荐使用如下方案：

- Set化后的MQ替代分布式锁，比如上面的例子，我们可以按用户ID做Set（用户ID % Set数）进而分成多个组，为不同的组创建不同的MQ队列，这样一个用户同一时间只在一个队列中，一个队列的处理是串行化的，实现了锁的功能，同时又有多个Set来完成并行化，在性能上会好于分布式锁，并且代码上没有太多改动

- 使用乐观锁，再如上面的例子，为account创建一个更新版本字段（update_version）,每次更新时版本加1，更新的条件是版本号要等于传入版本号：

  ```java
    var (balance,currentVersion) = db.account.getBalanceAndVersion(id)
    if(balance < amount){
      return error("余额少于扣款金额")
    }
    // 此操作对应的SQL: UPDATE account SET balance = balance  - <amount> , update_verison = update_verison + 1 WHERE id = <id> AND update_version = <currentVersion>
    if(db.account.updateBalance(id,-amount, currentVersion) == 0){
    return error("扣款失败") // 或递归执行此代码进行重试
    }
  ```

但在一些情况下我还是会不得不用锁，比如如果要同时加锁诸如用户、订单、商品、SKU等多个对象时用锁反而可能是更好的选择（当然前提要反思这样的业务操作是否合理、设计架构是否有缺陷，对此本节不展开讨论），再如在并发量很高的情况下很可能用悲观锁会比乐观锁效率更高。如需要使用分布式锁，我们必要注意这些问题：



## 分布式锁

### 实现原理

- 锁获取需要：要线程阻塞、原子性、有超时限制；
- 锁在使用中：锁超时的续期处理；
- 锁释放需要：验证锁的一致性(同一线程获得和释放的是同一把锁)；
- 在锁对象存储的分布式集群中：集群结点中锁的可用性，也就是集群结点间数据的一致性问题，这个是可以平衡的因素；

### 方案比较

- Zookeeper方案：实现了CAP的平衡，在集群服务中能保证部分结点数据同步复制；服务的可用性稍有降低、但是一致性得到了提升。

- 缓存(redis)方案：实现了CAP的平衡，只是在Redis集群存在数据异步复制引发的锁对象不可用；服务具有最高的可用性，但是存在小概率的一致性风险；

- 关系型数据库方案：CP

从性能角度（从高到低）：缓存 > Zookeeper >= 数据库

从可靠性角度（从高到低）：Zookeeper > 缓存 > 数据库

从实现的复杂性角度（从低到高）：Zookeeper >= 缓存 > 数据库

从理解的难易程度角度（从低到高）：数据库 > 缓存 > Zookeeper

### 方案1：关系型数据库

由关系型数据库的某些特性来实现，比如使用主键唯一性约束及数据一致来确保同一时间只有一个请求能获得锁，这一方案实现简单，但对高并发场景或可重入时存在比较大的性能瓶颈。

### 方案2：Redis 缓存

可使用Redis单线程、原子化操作（setnx）来实现，这一方案也很简单；

实现方式：

#### 直接使用Redis命令：

使用命令 SETNX 创建锁对象，没有原子化的值比较命令，无法原子化确认占用锁的是否是当前实例的当前线程，导致比较难实现重入锁；

特点：保证了锁的原子性、超时限制；但是没法保证释放锁的一致性，不能对锁续期；

> 参考[Redis 官方文档](http://www.redis.cn/documentation.html)
>
> **SETNX命令（SET if Not eXists）**
> 语法：SETNX key value
> 功能：原子性操作，当且仅当 key 不存在，将 key 的值设为 value ，并返回1；若给定的 key 已经存在，则 SETNX 不做任何动作，并返回0。
>
> **Expire命令**
> 语法：expire(key, expireTime)
> 功能：key设置过期时间
>
> **DEL命令**
> 语法：DEL key [KEY …]
> 功能：删除给定的一个或多个 key ,不存在的 key 会被忽略。

#### 使用Lua脚本实现Redis原子操作：

获取锁时：使用Lua脚本将【 SETNX 、超时设置Expire、GET锁对象并一致性确认】合为一个原子操作；

释放锁时：使用Lua脚本将【GET锁对象并一致性确认、DEL删除】合为一个原子操作。

特点：获取锁和释放锁实现了原子性、超时限制和一致性，但是锁使用中超时没有续期处理。

#### 使用Redisson（推荐，具有最高性能）：

在通过Lua脚本实现获取锁、释放锁的原子操作的基础上，增加定时任务实现锁续期。

优点：在Redis单节点服务中，在实现了CP的同时，具有最高的性能和A，单机Redis可以达到 10w QPS；

缺点：在Redis集群服务中，由于创建锁时，主备结点的数据复制是异步的，所以在负载均衡和主备间切换时，存在数据复制不及时的风险，可能造成边界锁对象不可用。

#### 使用Redlock（不推荐，性能差）

在 Redisson 实现的基础上，对Redis分布式集群做了处理（创建锁对象时，要同步在所有结点创建完毕后才返回锁对象）；

特点：这样保证了Redis集群结点之间数据的强一致性，但是锁对象的创建成本增大，系统的性能和A降低。



### 方案3：分布式

#### Zookeeper

可使用Zookeeper的持久节点（PERSISTENT）、临时节点（EPHEMERAL）、时序节点（SEQUENTIAL ）的特性组合及 watcher 接口实现，这一方案可保证最为严格的数据一致性、在性能及高可用也有着比较好的表现，推荐对一致性高要求极高、并发量大的场景使用。



#### etcd

etcd 是一个分布式的、可靠的、 key-value 存储的分布式系统。它不仅仅用于存储，还提供配置共享及服务发现。etcd与zookeeper类似，算是后起之秀。





### 编码实践

#### Redisson + SpringBoot 实践

pom.xml 引入

```xml
       <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-data-redis</artifactId>
            <version>2.2.6.RELEASE</version>
        </dependency>
        <dependency>
            <groupId>org.redisson</groupId>
            <artifactId>redisson</artifactId>
            <version>3.11.4</version>
        </dependency>
```

application.properties 配置

```properties
## redisson (此处仅为单节点配置)##
# redis链接地址
spring.redisson.address=dev-redis.ttsingops.com
spring.redisson.port=6379
spring.redisson.password=DevNewl23Olio
```

配置文件：

```java
import org.redisson.config.Config;
import org.springframework.context.annotation.Bean;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;
import org.redisson.api.*;
import org.redisson.Redisson;

/**
 * Redisson 连接配置
 * @author liuyuanyuan
 * @version 1.0.0
 * @create 2020/9/8
 */
@ConfigurationProperties(prefix = "spring.redisson")
@Configuration
public class RedissonConfig {

    private String host;
    private String port;
    private String password;

    @Bean
    public RedissonClient getRedisson(){
        Config config = new Config();
        //单节点配置
        config.useSingleServer()
                .setAddress("redis://" + host + ":" + port).setPassword(password);
        //添加主从配置	    	
        //config.useMasterSlaveServers()
        //      .setMasterAddress("").setPassword("").addSlaveAddress(new String[]{"",""});
        return Redisson.create(config);
    }
}
```

**使用 RLock 实现分布式锁操作：**

```java
import org.redisson.api.RedissonClient;
import org.redisson.api.RLock;

@Component
public  class TestLock{
    @Autowired
    private RedissonClient redissonClient;
    public void sync(int i){
        String lockKey = "order" + i;
        RLock lock = redissonClient.getLock(lockKey);
        try{
            lock.lock();
            System.out.println(Thread.currentThread());
            // bussiness execution start
            Thread.sleep(10000);
            // bussiness execution end
        }catch (Exception ex){
            ex.printStackTrace();
        }finally{
            lock.unlock();
        }
    }
}
```

**使用 RAtomicLong 实现 Redis 原子操作：**

RAtomicLong 是 Java 中 AtomicLong 类的分布式“替代品”，用于在并发环境中保存长值。以下是 RAtomicLong 的用法：

```java
import org.redisson.Redisson;
import org.redisson.api.RAtomicLong;
import org.redisson.api.RedissonClient;

public class AtomicLongExamples {
    public static void main(String[] args) {
        // 默认连接上127.0.0.1:6379
        RedissonClient client = Redisson.create();
        
        RAtomicLong atomicLong = client.getAtomicLong("myLong");
        System.out.println("Init value: " + atomicLong.get());
        atomicLong.incrementAndGet();
        System.out.println("Current value: " + atomicLong.get());
        atomicLong.addAndGet(10L);
        System.out.println("Final value: " + atomicLong.get());

        client.shutdown();
    }
}
```




#### Zookeeper 实践