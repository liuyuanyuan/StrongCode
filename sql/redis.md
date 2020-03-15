# redis - 分布式(外部)缓存中间件

外部缓存的重要代表，就是Redis，Memcache这样的分布式缓存中间件。当然外部缓存，你要把文件系统等划分进来，也不是不行，只要可以满足对缓存的定义即可。

与传统数据库不同的是， redis 的数据是存在内存中的，所以存写速度非常快， 因此 redis 被广泛应用于缓存方向。另外，redis 也经常用来做分布式锁。redis 提供了多种数据类型来支持不同的业务场景。此外，redis 支持事务 、持久化、LUA脚本、LRU驱动事件、多种集群方案。



# Redis & Jedis

[Redis](https://redis.io)（REmote DIctionary Server）是 Salvatore Sanfilippo 使用ANSI C语言编写的、完全开源免费(遵守BSD协议)的、支持网络、可基于内存亦可持久化的日志型、Key-Value数据库，并提供多种语言的API。

它通常被称为数据结构服务器，因为值（value）支持的数据结构有：字符串(String)，哈希(Hash)， 列表(list)，集合(sets) 和 有序集合(sorted sets)等类型，支持的索引有：范围查询(range query)，位图(bitmap)，超级日志(hyperloglogs)，地理空间( geospatial)索引。Redis具有内置的复制、Lua脚本、LRU驱逐、事务和不同级别的磁盘持久性，并通过Redis Sentinel和Redis Cluster自动分区提供了高可用性。

#### Redis 特点：

- Redis 支持数据的持久化，可以将内存中的数据保存在磁盘中，重启的时候可以再次加载进行使用。
- Redis 不仅支持简单的key-value类型的数据，同时还提供list，set，zset，hash等数据结构的存储。
- Redis 支持数据的备份，即 master-slave模式的数据备份。

#### Redis 优势

- 性能极高 – Redis能读的速度是110000次/s,写的速度是81000次/s 。
- 丰富的数据类型 – Redis支持二进制案例的 Strings, Lists, Hashes, Sets 及 Ordered Sets 数据类型操作。
- 原子 – Redis的所有操作都是原子性的，意思就是要么成功执行要么失败完全不执行。单个操作是原子性的。多个操作也支持事务，即原子性，通过MULTI和EXEC指令包起来。
- 丰富的特性 – Redis还支持 publish/subscribe, 通知, key 过期等等特性。







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



#### 3 Java access Redis by Jedis

reference: 

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


