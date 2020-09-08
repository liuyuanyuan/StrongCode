# 分布式锁

[TOC]



## 1 幂等性处理

### 幂等性概念：

幂等一词源自数学概念，在程序中如果相同条件下多次请求对资源的影响表现一致则称请求为幂等请求，对应的接口为幂等接口。

### 幂等性处理:

##### 1 使用数据库主键、唯一索引，防止重复数据插入；

特点：实现简单，但通用性差，也无法解决并发请求对应用服务的消耗；

##### 2 使用分布式锁（可通过redis实现）

分布式锁一般适用于需要长时间处理的任务，在任务处理期间防止重复请求，如数据导出、复杂计算等，由于这些操作本身就要求串行处理，所以加锁对性能地影响有限（锁粒度为请求条件）

##### 3 缓存请求URI并设定超时时长

URI做为请求Token再加上过期时间，比如 `PUT /user/001` 幂等有效时间30秒，则在30秒内同一个URI请求都视为重复直接过滤，这种做法可简化请求方操作但仅限于REST请求且符合REST规范。

##### 4 将请求放到MQ然后再进行处理

主流的 MQ 实现在 `autocommit=true` 时天然实现了幂等；但考虑业务处理可能出错的情况我们一般会将 autocommit 设置成 false ，在业务处理成功后再提交，这时就需要使用上述幂等方案了：在接收到消息时写入请求Token以实现去重判断（Token可为Topic+Offset）提交后删除Token，整体上可以做到对业务透明。

### dos攻击预防：



## 2 分布式锁

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

#### 直接使用Redis的命令：

创建锁对象使用命令 SETNX，没有原子化的值比较命令，无法原子化确认占用锁的是否是当前实例的当前线程，导致比较难实现重入锁；

特点：保证了锁的原子性、超时限制；但是没法保证释放锁的一致性，不能对锁续期；

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

#### 使用Lua脚本实现Redis锁的原子操作：

获取锁时：使用Lua脚本将【 SETNX 、超时设置Expire、GET锁对象并一致性确认】合为一个原子操作；

释放锁时：使用Lua脚本将【GET锁对象并一致性确认、DEL删除】合为一个原子操作。

特点：获取锁和释放锁实现了原子性、超时限制和一致性，但是锁使用中超时没有续期处理。

#### 使用Redisson（推荐，具有最高的A）：

在通过Lua脚本实现获取锁、释放锁的原子操作的基础上，增加定时任务实现锁续期。

优点：在Redis单节点服务中，在实现了CP的同时，具有最高的A，单机Redis可以达到 10w QPS；

缺点：在Redis集群服务中，由于创建锁时，主备结点的数据复制是异步的，所以在负载均衡和主备间切换时，存在数据复制不及时的风险，可能造成边界锁对象不可用。

#### 使用Redlock（不推荐，A差）

在 Redisson 实现的基础上，对Redis分布式集群做了处理（创建锁对象时，要同步在所有结点创建完毕后才返回锁对象）；

特点：这样保证了Redis集群结点之间数据的强一致性，但是锁对象的创建成本增大，系统的可用性降低。



### 方案3：Zookeeper

可使用Zookeeper的持久节点（PERSISTENT）、临时节点（EPHEMERAL）、时序节点（SEQUENTIAL ）的特性组合及 watcher 接口实现，这一方案可保证最为严格的数据一致性、在性能及高可用也有着比较好的表现，推荐对一致性高要求极高、并发量大的场景使用。



## 编码实践

### Redisson + SpringBoot 实践

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

应用代码：

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

#### 

### Zookeeper 实践