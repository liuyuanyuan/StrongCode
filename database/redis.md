# Redis - 内存数据库/分布式缓存中间件

外部缓存的重要代表，就是Redis，Memcache这样的分布式缓存中间件。当然外部缓存，你要把文件系统等划分进来，也不是不行，只要可以满足对缓存的定义即可。

## Redis 概述

Redis （Remote Dictionary Server）

官网：redis.io

中文网：redis.cn

安装部署：

- Linux安装包从官网下载，Win安装包官网已经停更很久了只能从github下载。官方推荐在Linux下安装使用Redis。
- 下载、解压、双击启动即可。

使用方法：

- 安装图形化客户端

- 或者使用命令行工具：

  ```java
  #启动服务
  redis-server，
  
  #命令客户端
  redis-cli
  
  #切换数据库用
  select 0
  ```

  

特点：

- 默认端口号6379，默认有0-15共16个数据库；

- 支持 key-value 存储，同时还提供 list，set，zset，hash 等数据结构的存储。

- 支持事务：Redis 单条命令有原子性，但是，Redis事务不保证原子性，也没有事务隔离级别(发起execute命令的时候直接执行)；

  - 开启事务： multi
  - 命令入队
  - 执行事务：提交exec / 放弃discard

- 支持持久化（RDB，AOF）：断电数据持久化；

  - .rdb 文件

  - .aof 文件：

- 内存存储，读写速度极高，可用于高速缓存，单节点查询达到11w+QPS，写是81000次/s；

- 使用单线程：

  - 官方表示Redis是基于内存操作，CPU不是Redis性能瓶颈，Redis的性能瓶颈来自机器内存和网络带宽；另外，多线程是先通过cpu核数分配，不足则对cpu分片进行时间片轮转，多线程的cpu上下文切换是耗时操作。

  - 可以通过单线程实现就用了单线程；

- Lua 脚本：多个原子操作通过 Lua 脚本组合成一个原子操作；

- 用于实现分布式锁：一般用 Redisson 中间件，基于Redis 实现 Java 的分布式锁；

- 用于实现发布/订阅：

- LRU 驱动事件：

- 多种集群方案：

  **replication 主-从复制模式**（1主多从）：主机负责写，从机负责读；初次连接时做全量同步，后续操作做增量同步；

  - 实现读写分离：分担主库的读写压力
  - 方便做容灾恢复：主机故障切换到从机；
  - 

  **replication-sentinal 哨兵模式（适用于普通的读写分离+高可用+自动failover）：**

  当主结点中断服务后，可以将一个从结点升级为主服务器，以便继续提供服务，但这个过程需要人工手动来操作，就费事费力还会导致服务在一段时间内不可用。 

  为此，Redis 2.8 中提供了哨兵模式，来实现自动化系统监控和故障恢复功能。

  - Redis 提供了哨兵命令，哨兵是一个独立的进程，它会独立运行。哨兵通过发送命令等待 Redis 服务器响应，从而监控多个 Redis 实例的运行状况；

  - 当哨兵监测到 master 宕机，会自动将 slave 切换成 master ，然后通过发布/订阅模式，通知其他的slave 修改配置文件，让它们切换主机。

  - 然而，一个哨兵进程对Redis服务器进行监控，可能会出现误判问题；为此，可以使用多个哨兵进行监控；各个哨兵之间还会进行监控，这样就形成了多哨兵模式。哨兵1先检测到这个结果，系统并不会马上进行failover过程，仅仅是哨兵1主观的认为主服务器不可用，这个现象成为**主观下线**。

    当后面的哨兵也检测到主服务器不可用，并且数量达到一定值时，那么哨兵之间就会进行一次投票，投票的结果由一个哨兵发起，进行failover操作。切换成功后，就会通过发布订阅模式，让各个哨兵把自己监控的从服务器实现切换主机，这个过程称为**客观下线**。这样对于客户端而言，一切都是透明的。

    ```bash
    # 启动Redis服务器进程
    ./redis-server ../redis.conf
    # 启动哨兵进程
    ./redis-sentinel ../sentinel.conf
    ```

  **cluster 分片存储集群模式（主要针对海量数据+高并发+高可用的场景）：**

  redis cluster是Redis的分布式解决方案，在3.0版本推出后有效地解决了redis分布式存储需求，也就是说每台redis节点上存储不同的内容；

  - 自动将数据进行分片，每个master上放一部分数据

  - 提供内置的高可用支持，部分master不可用时，还是可以继续工作的

  - 支撑N个redis master node，每个master node都可以挂载多个slave node

  - 高可用，因为每个master都有salve节点，那么如果mater挂掉，redis cluster这套机制，就会自动将某个slave切换成master
  - 数据分布算法：哈希算法：计算 key的 hash 值，然后均匀的映射到到 N 个 redis上：hash(key)%N。

- 自带压力测试工具 redis-benchmark ，可用于性能压测；

- Jedis ，java连接redis的基础中间件；

  Jedis<Spring 封装为RedisTemplate<自行封装为RedisHelper，操作对象必须序列化，否则报错。

  参考：

  https://github.com/xetorthio/jedis
  https://redis.io/clients#java

  First download Jedis.jar and build to path
  Then code like:

  	public static void main(String[] args)
  	{
  		// 连接本地的 Redis 服务
  		Jedis jedis = new Jedis("192.168.102.52", 6379);
  		System.out.println("连接成功");
  		// 查看服务是否运行
  		System.out.println("ping: " + jedis.ping());
  		System.out.println("foo: " + jedis.get("foo"));
  		
  	}

  



## 安装使用

#### 1 Install Redis (Centos7)

``` 
# download
$ wget http://download.redis.io/releases/redis-4.0.11.tar.gz
$ tar xzf redis-4.0.11.tar.gz
$ cd redis-4.0.11
$ make

#test
$ cd src
$ make test

#start server and enter cli
$ src/redis-server [redis.conf]

#shutdown server
$ src/redis -cli shutdown

#Enter cli
$ src/redis-cli 
#Enter cli and avoid unidentifiable Chinese code 
$ src/redis -cli --raw 
```

#### 2 Config to allow Remote Connection

```
#modify redis.conf
vim redis.conf
   bind 0.0.0.0
   protected-mode no

#turn off os firewall
sudo service iptables stop
```


