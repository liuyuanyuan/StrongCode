# 关系型数据库事务原理

[TOC]

## 事务的4大特性

如果说数据库支持事务，就是说数据库具备事务的四大特征。数据库中事务具有四大特性（ACID）：

### 1 原子性 Atomicity

事务包含的所有操作，要么全部成功，要么全部失败会滚。

> mysql的原子性是基于undo/redo log来实现的。

### 2 一致性 Consistency

一致性是指事务必须使数据库从一个一致性状态变换到另一个一致性状态，也就是说一个事务执行之前和执行之后都必须处于一致性状态。（银行存钱、取钱）

### 3 隔离性 Isolation

隔离性是当多个用户并发访问数据库时，比如操作同一张表时，数据库为每一个用户开启的事务，不能被其他事务的操作所干扰，多个并发事务之间要相互隔离。

### 4 持久性 Durability

持久性是指已被提交的事务对数据库的修改应该永久保存在数据库中。

即一个事务一旦被提交了，那么对数据库中的数据的改变就是永久性的，即便是在数据库系统遇到故障的情况下也不会丢失提交事务的操作。

> 检查点checkpoint：



## 数据库的锁

### 1 共享锁/读锁 Shared Lock

### 2 排他锁/写锁 Exclusive Lock

读数据时加读锁，其他事务都可以读，但是不可以写；

写数据时加写锁，其他事务不能写，也不能读；

注意：数据库的增、删、改操作会加排它锁；而查询操作并不会加锁，只能通过在 select 语句后显式加 lock in share mode 来加共享锁 或者 for update 来加排它锁。

### 3 行锁 Record Lock

### 4 表锁 Table Lock

行锁和表锁，是从锁的粒度上进行划分的；

行锁只锁定当前数据行，锁的粒度小，加锁慢，发生锁冲突的概率小，并发度高；

表锁则锁的粒度大，加锁快，开销小，但是锁冲突的概率大，并发度低；

注意：update 有 where 条件时，若是有索引的条件，则直接锁定指定行；若是没有索引的条件下，就获取所有行，都加上行锁，然后Mysql会再次过滤符合条件的行并释放锁，只有符合条件的行才会继续持有锁；

> mysql中，InnoDB 支持行锁并且支持事务，MyISAM 不支持事务也不支持行锁。

### 5 间隙锁(范围锁) Gap Lock 

间隙锁则分为两种：Gap Locks 和 Next-Key Locks。

- Gap Locks会锁住两个索引之间的区间；

  比如 select * from User where id>3 and id<5 for update，就会在区间 (3，5) 之间加上Gap Locks。

-  Next-Key Locks 是 Gap Locks + Record Locks 形成闭区间锁(加区间锁并给区间内所有行加行锁)；

  比如 select * from User where id>=3 and id=<5 for update，就会在区间 [3,5] 之间加上Next-Key Locks。



## 事务的4种隔离级别

以下四种隔离级别的，隔离性由弱到强，并发处理性能由强到弱。

![在这里插入图片描述](img/pg_transaction_isolation.png)

### 1 读未提交 READ UNCOMMITTED

读未提交，即可以读取其他事务未提交的数据。

最低的隔离级别，事务之间没什么隔离性限制，但是并发性能最高。

> **脏读问题**：
>
> 事务A事先读取了数据，事务B紧接着更新数据并未提交，而事务A再次读取该数据时，数据已经发生了改变。这就造成了脏读；

### 2 读已提交 READ COMMITTED

读已提交，即只能读到其他事务已经提交数据(未提交的读取不到)。

解决了脏读的问题；

> **不可重复读（虚读）**：
>
> 事务A事先读取了数据，事务B紧接着更新数据并提交事务，而事务A再次读取该数据时，数据已经发生了改变(主要针对update操作)。这就造成了不可重复读（虚读）。
>
> 不可重复读即在一个事务内两次读到的数据是不一样的，即原始读取不可重复。

### 3 可重复读 REPEATABLE READ 

解决了不可重复读(虚读)；

> **幻读问题：**
>
> 事务A对某行数据进行了修改( 比如delete了x行)；跟接着事务B也对该行数据进行了修改(比如insert了该行)并提交；事务A再读取时，却发现数据跟预期的不同（删除的数据还存在，就像发生幻觉）。这就造成了幻读；幻读主要是由 insert/delete 引起的。

### 4 串行化 SERIALIZABLE

串行化，即所有的事务串行执行，一个事务完成(commit/rollback)后另一个事务才能执行操作；

最高的隔离级别，并行处理性能极差。



### 数据库的隔离级别设置

#### Oracle提供了2种隔离级别（从高到低）：

- **READ COMMITTED (默认级别)**
- SERIALIZABLE 

#### PostgreSQL提供了4种隔离级别：

- READ UNCOMMITTED
- **READ COMMITTED（默认级别）**
- REPEATABLE READ
- SERIALIZABLE

语法：

```sql
SET TRANSACTION transaction_mode [, ...]
SET TRANSACTION SNAPSHOT snapshot_id
SET SESSION CHARACTERISTICS AS TRANSACTION transaction_mode [, ...]

where transaction_mode is one of:
ISOLATION LEVEL { SERIALIZABLE | REPEATABLE READ | READ COMMITTED | READ UNCOMMITTED }
READ WRITE | READ ONLY
[ NOT ] DEFERRABLE
```
使用方法：

```sql
BEGIN TRANSACTION ISOLATION LEVEL REPEATABLE READ;
SELECT pg_export_snapshot();
 pg_export_snapshot
---------------------
 00000003-0000001B-1
(1 row)

--或者：--
BEGIN TRANSACTION ISOLATION LEVEL REPEATABLE READ;
SET TRANSACTION SNAPSHOT '00000003-0000001B-1';
```



#### MySQL提供了4种隔离级别（从高到低）：

- READ UNCOMMITTED：最低级别，任何情况都无法保证。
- READ COMMITTED：可避免脏读的发生。
- **REPEATABLE READ（默认级别）**：可避免脏读、不可重复读的发生。
- SERIALIZABLE：可避免脏读、不可重复读、幻读的发生。

MySQL设置数据库隔离级别：

```
#查看隔离级别
SELECT @@tx_isolation;
#设置隔离级别
SET [GLOBAL | SESSION] TRANSACTION ISOLATION LEVEL 隔离级别名称；
SET tx_isolation='隔离级别名称';
```

**注意：**

- **设置数据库的隔离级别一定要是在开启事务之前，否则无效！**

- **隔离级别的设置只对当前连接有效。**

> - 对于使用MySQL命令窗口而言，一个窗口就相当于一个连接，当前窗口设置的隔离级别只对当前窗口中的事务有效；
>
> - 对于JDBC操作数据库来说，一个Connection对象相当于一个链接，而对于Connection对象设置的隔离级别只对该Connection对象有效，与其他链接Connection对象无关。



### JDBC的隔离级别设置：

```java
    /**
     * A constant indicating that transactions are not supported.
     */
    int TRANSACTION_NONE             = 0;

    /**
     * A constant indicating that
     * dirty reads, non-repeatable reads and phantom reads can occur.
     * This level allows a row changed by one transaction to be read
     * by another transaction before any changes in that row have been
     * committed (a "dirty read").  If any of the changes are rolled back,
     * the second transaction will have retrieved an invalid row.
     */
    int TRANSACTION_READ_UNCOMMITTED = 1;

    /**
     * A constant indicating that
     * dirty reads are prevented; non-repeatable reads and phantom
     * reads can occur.  This level only prohibits a transaction
     * from reading a row with uncommitted changes in it.
     */
    int TRANSACTION_READ_COMMITTED   = 2;

    /**
     * A constant indicating that
     * dirty reads and non-repeatable reads are prevented; phantom
     * reads can occur.  This level prohibits a transaction from
     * reading a row with uncommitted changes in it, and it also
     * prohibits the situation where one transaction reads a row,
     * a second transaction alters the row, and the first transaction
     * rereads the row, getting different values the second time
     * (a "non-repeatable read").
     */
    int TRANSACTION_REPEATABLE_READ  = 4;

    /**
     * A constant indicating that
     * dirty reads, non-repeatable reads and phantom reads are prevented.
     * This level includes the prohibitions in
     * <code>TRANSACTION_REPEATABLE_READ</code> and further prohibits the
     * situation where one transaction reads all rows that satisfy
     * a <code>WHERE</code> condition, a second transaction inserts a row that
     * satisfies that <code>WHERE</code> condition, and the first transaction
     * rereads for the same condition, retrieving the additional
     * "phantom" row in the second read.
     */
    int TRANSACTION_SERIALIZABLE     = 8;

```

JDBC设置事务隔离级别：

```java
Connection  conn = null;
try{
    conn = JdbcUtil.getConnection(jdbcInfo);
    //设置事务隔离级别
    conn.setTransactionIsolation(Connection.TRANSACTION_SERIALIZABLE);
    //开启事务（非自动提交事务）
    conn.setAutoCommit(false);
  
    //...
  
    //提交事务
    conn.commit(); 
}finally{
    if(conn != null){
       //还原事务为自动提交
       conn.setAutoCommit(true);
       //及时关闭连接
       conn.close();
    }
}
```



## 事务的多版本并发控制(MVCC)

在 MySQL 的 InnoDB 中，每个事务都有一个自己的事务id，并且是唯一的，递增的 。

对于MySQL中的每一个数据行，都有可能存在多个版本；在每次事务更新数据的时候，都会生成一个新的数据版本，并且把自己的数据 id 赋值给当前版本的 row 的trx_id。

InnoDB 最基本行记录(row)中包含一些额外的存储信息：DATA_TRX_ID，DATA_ROLL_PTR，DB_ROW_ID，DELETE BIT 。

| 列名          | 长度  | 备注                                                         |
| ------------- | ----- | ------------------------------------------------------------ |
| DATA_TRX_ID   | 6字节 | 标记了最新更新这条行记录的transaction id，每处理一个事务，值自动+1； |
| DATA_ROLL_PTR | 7字节 | 指向当前记录项的rollback segment的undo log记录，找之前版本的数据就是通过这个指针 |
| DB_ROW_ID     | 6字节 | InnoDB自动产生聚集索引时，聚集索引包括这一列，否则聚集索引中不包括这个值。 |
| DELETE BIT    |       | 位用于标识该记录是否被删除，这里的不是真正的删除数据，而是标志出来的删除。真正的删除是在commit的时候。 |