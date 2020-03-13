# redis - 分布式(外部)缓存中间件

外部缓存的重要代表，就是Redis，Memcache这样的分布式缓存中间件。当然外部缓存，你要把文件系统等划分进来，也不是不行，只要可以满足对缓存的定义即可。



与传统数据库不同的是， redis 的数据是存在内存中的，所以存写速度非常快， 因此 redis 被广泛应用于缓存方向。另外，redis 也经常用来做分布式锁。redis 提供了多种数据类型来支持不同的业务场景。此外，redis 支持事务 、持久化、LUA脚本、LRU驱动事件、多种集群方案。

# Redis & Jedis

[Redis](https://redis.io) is an open source (BSD licensed), in-memory data structure store, used as a database, cache and message broker. It supports data structures such as strings, hashes, lists, sets, sorted sets with range queries, bitmaps, hyperloglogs, geospatial indexes with radius queries and streams. Redis has built-in replication, Lua scripting, LRU eviction, transactions and different levels of on-disk persistence, and provides high availability via Redis Sentinel and automatic partitioning with Redis Cluster.

## 1 Install Redis (Centos7)

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



## 2 Config to allow Remote Connection

```
#modify redis.conf
vim redis.conf
   bind 0.0.0.0
   protected-mode no

#turn off os firewall
sudo service iptables stop
```



## 3 Java access Redis by Jedis

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


