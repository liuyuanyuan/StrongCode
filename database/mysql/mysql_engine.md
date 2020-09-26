## 架构

![参考马士兵的MySQL培训](/Users/liuyuanyuan/github/CodeBetter/database/mysql/img/mysql_structure.png)

MySQL Query Cache 是缓存我们所执行的SELECT语句以及该语句的结果集;

MySQL 在实现 Query Cache 的具体技术细节上类似典型的 KV 存储，就是将 SELECT 语句和该查询语句的结果集做了一个 HASH 映射并保存在一定的内存区域中。当客户端发起 SQL 查询时，Query Cache 的查找逻辑是，先对 SQL 进行相应的权限验证，接着就通过 Query Cache 来查找结果（注意必须是完全相同，即使多一个空格或者大小写不同都认为不同，即使完全相同的SQL，如果使用不同的字符集、不同的协议等也会被认为是不同的查询，而分别进行缓存）。它不需要经过Optimizer模块进行执行计划的分析优化，更不需要发生同任何存储引擎的交互，减少了大量的磁盘IO和CPU运 算，所以有时候效率非常高。



## InnoDB SQL执行的 BufferPool 缓存机制

![img](/Users/liuyuanyuan/github/CodeBetter/database/mysql/img/mysql_innodb_sql_bufferpool.png)

为什么MySQL不能直接更新磁盘上的数据而设置这么一套复杂的机制来执行SQL了? 因为来一个请求就直接对磁盘文件进行随机读写，然后更新磁盘文件里的数据性能可能相当差。

- 因为磁盘随机读写的性能是非常差的，所以直接更新磁盘文件是不能让数据库抗住很高并发的。

- MySQL这套机制看起来复杂，但它可以保证每个更新请求都是更新内存BufferPool，然后顺序写日志文件，同时还能保证各种异常情况下的数据一致性。 

  更新内存的性能是极高的，然后顺序写磁盘上的日志文件的性能也是非常高的，要远高于随机读写磁盘文件。 正是通过这套机制，才能让MySQL数据库在较高配置的机器上可以抗下每秒几千的读写请求。





## MySQL 存储引擎

MySQL 存储引擎是表级别的，每张表可以定义自己的存储引擎。

### InnoDB (MySQL 5.5之后，是默认存储引擎)

##### 表的存储

使用InnoDB 的表test，在data目录下有2个文件：

- test.**frm**  存储表的**结构**

- test.**ibd**   存储表的**数据和索引**

  表数据本身就按照 B+tree 来组织的，形成一个B+tree索引结构的文件。

##### 建议使用InnoDB的表必须建主键（整型、自增字段），原因是什么？

因为InnoDB是采用聚集索引的，整型、自增的key效率更高。

##### 操作

页（page）作为磁盘和内存之间交互的基本单位；

一个页内可能有一行或多行数据。

##### 应用场景

事务性、安全性操作较多的情况；



### MyISAM

表 test  (使用存储引擎MyISAM)

data目录下有3个相关文件：

- test.**frm**    存储表的**结构**

- test.**MYD**  存储表的**数据**
- test.**MYI**    存储表的**索引**

##### 特点

不提供对事务的支持，不支持行级锁，不支持外键；

##### 应用场景

执行大量SELECT查询的情况。



### Memory



## MySQL的日志系统

- MySQL -  binlog(包括查询和DML)：增量日志，超限新增，用于主从复制和灾难恢复；

- MySQL InnoDB - undo log/redo log（仅DML）：记录事务提交前和提交后的数据快照，作用于两阶段提交；

