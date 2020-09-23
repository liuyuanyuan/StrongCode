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

### 1按性能分：乐观锁(用版本比对实现) & 悲观锁

- 乐观锁：不加锁，通过版本比对，不符合时执行失败，失败后可以重试；

- 悲观锁：

### 2 按操作分：读锁/共享锁Shared Lock & 写锁/排他锁 Exclusive Lock

- 读锁，对其他事务不阻塞读，但阻塞写；(读数据时加读锁)

- 写锁，对其他事务阻塞写，也阻塞读；(写数据时加写锁)

注意：

- 数据库的增、删、改操作会加排它锁；
- 而查询操作并不会加锁，只能通过在 select 语句后显式加 lock in share mode 来加共享锁 或者 for update 来加排它锁。

给表加读锁：

```mysql
mysql> lock table t1 read;
mysql> select * from t1;
+----+---------+-------+
| id | name    | sex   |
+----+---------+-------+
|  1 | ff      | 0     |
+----+---------+-------+

mysql> insert into t1 values(5, 'jj','man');
ERROR 1099 (HY000): Table 't1' was locked with a READ lock and can't be updated
```

给表加写锁：

```mysql
mysql> lock table t1 write;
# 事务B
mysql> beigin;
mysql> select * from t1;//本步骤阻塞，直到t1解锁。

```

### 3 按操作粒度： 行锁Record Lock & 表锁Table Lock

- 行锁只锁定数据行，锁的粒度小，加锁慢，发生锁冲突的概率小，并发度高；(保证可重复读隔离就需要加行锁，保证行更新阻塞)

- 表锁则锁定整个表，锁的粒度大，加锁快，开销小，但是锁冲突的概率大，并发度低；（一般用在整表数据迁移的情况，解决幻读就需要加表锁，保证insert也阻塞）

注意：update 有 where 条件时，若是有索引的条件，则直接锁定指定行；若是没有索引的条件下，就获取所有行，都加上行锁，然后MySQL会再次过滤符合条件的行并释放锁，只有符合条件的行才会继续持有锁；

> MySQL中：
>
> - InnoDB 支持行锁，也支持事务；
>
> - MyISAM 不支持事务，也不支持行锁。

加行锁：

```mysql
SELECT * FROM test WHERE id=1 FOR UPDATE;
# 它会在 id=1 的记录上加上记录锁，以阻止其他事务插入，更新，删除 id=1 这一行。
```

加表锁

```mysql
#手动增加表锁
lock table 表名称 read/write,表名称2 read/write;
#查看表上加过的锁 
show open tables;
#删除表锁 
unlock tables;
```

### 4 间隙锁(范围锁) Gap Locks 

**间隙锁**：封锁索引记录中的间隙，或者第一条索引记录之前的范围，又或者最后一条索引记录之后的范围。

```mysql
SELECT * FROM user WHERE id>3 AND id<5 FOR UPDATE;
# 会在区间 (3，5) 之间加上 Gap Locks。
```

间隙锁就是锁住行记录之间的空隙，从而防止其他事务进行插入。间隙锁常用于可重复读隔离级别 (MySQL 默认级别)，在某些情况下可以解决幻读问题。

> 当采用范围条件查询数据时，InnoDB 会对这个范围内的数据进行加锁。
>
> 比如有 id 为：1、3、5、7 的 4 条数据，我们查找 1-7 范围的数据。
>
> - 那么 1-7 都会被加上锁。
> - 2、4、6 也在 1-7 的范围中，但实际不存在。这些 2、4、6 就被称为间隙。
>
> - 范围查找时，会把整个范围的数据全部锁定，包括所有间隙：比如我要在 1、3、5、7 中插入 2，这个时候 1-7 都被锁定住了，所以根本无法插入 2。在某些场景下会严重影响性能。

### 4 临键锁 Next-Key Locks

Next-Key Locks 是 Gap Locks + Record Locks 形成闭区间锁(加区间锁并给区间内所有行加行锁)；

```mysql
select * from User where id>=3 and id=<5 for update;
# 会在区间 [3,5] 之间加上 Next-Key Locks。
```



## 事务的4种隔离级别

以下四种隔离级别的，隔离性由弱到强，并发处理性能由强到弱。

![在这里插入图片描述](img/pg_transaction_isolation.png)

### 1 读未提交 READ UNCOMMITTED

读未提交，即可以读取其他事务未提交的数据。

最低的隔离级别，事务之间没什么隔离性限制，但是并发性能最高。

存在**脏读**问题。

> **脏读(dirty read)**：
>
> 事务A读到了事务B未提交的内容，这就造成了脏读；
>
> | 事务A                                                | 事务B                     |
> | ---------------------------------------------------- | ------------------------- |
> | Begin;                                               | Begin;                    |
> |                                                      | insert/delete/update tab1 |
> | select * from tab1;  读到了事务B未提交的数据（脏读） |                           |
> |                                                      | Rollback;                 |

### 2 读已提交 READ COMMITTED

读已提交，即只能读到其他事务已经提交数据(未提交的读取不到)。

解决了脏读的问题；存在 **不可重复读（虚读）** 问题。

> **不可重复读/虚读(Non-Repeatable Read)**：(主要update引起)  在一个事务内两次读到的数据是不一样的，即原始读取不可重复。
>
> **解决不可重复读问题只需锁住满足条件的行。**
>
> | 事务A                                                        | 事务B                                  |
> | ------------------------------------------------------------ | -------------------------------------- |
> | Begin;                                                       |                                        |
> | 读取T表的数据(id=1, name='lily')                             | Begin;                                 |
> |                                                              | update t1 set name = 'lyy' where id=1; |
> |                                                              | Commit;                                |
> | 读取T表数据(id=1, name='lyy')，数据和原先读取的已经变了（不可重复读）； |                                        |

### 3 可重复读 REPEATABLE READ 

解决了不可重复读/虚读(通过对数据加行锁，其他事务update行时会阻塞)；存在**幻读**问题：

> **幻读(Phantom Read)**：(insert引起) 事务A对全表update，中间事务B执行insert并提交，结果事务A查询时出现未更新的行。
>
> **解决幻读需要锁表**。
>
> | 事务A                                                   | 事务B                                         |
> | ------------------------------------------------------- | --------------------------------------------- |
> | Begin;                                                  |                                               |
> | update t1 set sex='0';                                  | Begin;                                        |
> |                                                         | insert into t1 values(3, 'yolanda', 'woman'); |
> |                                                         | Commit;                                       |
> | select * from t1; 结果存在sex != ‘0’ 的行数据。（幻读） |                                               |

### 4 串行化 SERIALIZABLE

串行化，即所有的事务串行执行，一个事务完成(commit/rollback)后另一个事务才能执行操作；

最高的隔离级别，并行处理性能极差。

### 数据库的隔离级别设置

#### Oracle提供了2种隔离级别（从低到高）：

- **READ COMMITTED (默认级别)**
- SERIALIZABLE 

#### PostgreSQL提供了4种隔离级别：

- READ UNCOMMITTED
- **READ COMMITTED（默认级别）**
- REPEATABLE READ
- SERIALIZABLE

使用方法：

```sql
BEGIN TRANSACTION ISOLATION LEVEL REPEATABLE READ;
COMMIT/ROLLBACK;
```

#### MySQL提供了4种隔离级别（从低到高）：

- READ UNCOMMITTED
- READ COMMITTED
- **REPEATABLE READ（默认级别）**
- SERIALIZABLE 

MySQL设置数据库隔离级别：

```
#查看隔离级别
SELECT @@tx_isolation;
show variables like 'tx_isolation';
#设置隔离级别
SET [GLOBAL | SESSION] TRANSACTION ISOLATION LEVEL 隔离级别名称；
SET tx_isolation='REPEATABLE-READ';
```

**注意：**

- **设置数据库的隔离级别一定要是在开启事务之前，否则无效！**

- **隔离级别的设置只对当前连接有效。**

> - 对于使用MySQL命令窗口而言，一个窗口就相当于一个连接，当前窗口设置的隔离级别只对当前窗口中的事务有效；
>
> - 对于JDBC操作数据库来说，一个 Connection 对象相当于一个链接，而对于Connection对象设置的隔离级别只对该Connection对象有效，与其他链接Connection对象无关。



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
Connection conn = null;
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



## 事务多版本并发控制(MVCC)

MVCC (Multi-Version Concurrency Control)是事务的多版本并发控制机制，关系型数据库MySQL(InnoDB 引擎支持事务)、PostgreSQL都有实现。

### InnoDB 事务隔离级别实现机制

| MySQL InnoDB 隔离级别 | 实现机制 |      |
| --------------------- | -------- | ---- |
| 读未提交              | 无隔离   |      |
| 读已提交              | MVCC     |      |
| 可重复读              | MVCC     |      |
| 串行化                | 互斥锁   |      |

读已提交/可重复读 是 通过MVCC机制来实现的：对一行数据的读和写两个操作，默认不通过加锁互斥来保证隔离性，避免了频繁加锁互斥；

> InnoDB 日志：
>
> - undo log 回滚日志：事务提交前的原行数据，用于事务Rollback；
> - redo log 重做日志：事务提交后的新行数据，用于事务Commit;
>

### InnoDB 的 MVCC 实现

MySQL InnoDB 行记录中，除了行数据本身，还存了一些额外信息：DATA_TRX_ID，DATA_ROLL_PTR，DB_ROW_ID，DELETE BIT 。

| 列名          | 长度  | 备注                                                         |
| ------------- | ----- | ------------------------------------------------------------ |
| DATA_TRX_ID   | 6byte | 标记了最新更新这条行记录的transaction id，每处理一个事务，值自动+1； |
| DATA_ROLL_PTR | 7byte | 指向当前记录项的rollback segment的undo log记录，找之前版本的数据就是通过这个指针 |
| DB_ROW_ID     | 6byte | InnoDB自动产生聚集索引时，聚集索引包括这一列，否则聚集索引中不包括这个值。 |
| DELETE BIT    |       | 位用于标识该记录是否被删除，这里的不是真正的删除数据，而是标志出来的删除。真正的删除是在commit的时候。 |



MySQL  InnoDB 中，每个事务都有一个自己的事务id，并且是唯一的、递增的 。对于每一个数据行，每次事务更新数据时，都会生成一个新的数据版本，存储在该数据行的 trx_id 中。

![image-20200916120033561](/Users/liuyuanyuan/Library/Application Support/typora-user-images/image-20200916120033561.png)











