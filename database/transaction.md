# 数据库：事务四大特性&事务的隔离级别

[TOC]

如果说数据库支持事务，就是说数据库具备事务的四大特征。数据库中事务具有四大特性（ACID）：

### 1 原子性（Atomicity）

事务包含的所有操作，要么全部成功，要么全部失败会滚。

### 2 一致性（Consistency）

一致性是指事务必须使数据库从一个一致性状态变换到另一个一致性状态，也就是说一个事务执行之前和执行之后都必须处于一致性状态。（银行存钱、取钱）

### 4 持久性（Durability）

持久性是指已被提交的事务对数据库的修改应该永久保存在数据库中。

即一个事务一旦被提交了，那么对数据库中的数据的改变就是永久性的，即便是在数据库系统遇到故障的情况下也不会丢失提交事务的操作。

> 检查点checkpoint：
>
> 

### 3 隔离性（Isolation）

隔离性是当多个用户并发访问数据库时，比如操作同一张表时，数据库为每一个用户开启的事务，不能被其他事务的操作所干扰，多个并发事务之间要相互隔离。

#### 数据库提供多种隔离级别：

- 脏读
- 不可重复读
- 虚读(幻读)

#### Oracle提供了2种隔离级别（从高到低）：

- Serializable (串行化)

- Read committed (读已提交，默认级别)

### PostgreSQL提供了4种隔离级别：

- SERIALIZABLE (串行化)

- REPEATABLE READ（可重复读）

- READ COMMITTED（读已提交，默认隔离级别）

- READ UNCOMMITTED（读为提交）

![在这里插入图片描述](/Users/liuyuanyuan/github/StrongCode/java/images/pg_transaction_isolation.png)

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

- Serializable (串行化)：可避免脏读、不可重复读、幻读的发生。

- Repeatable read (可重复读，默认级别)：可避免脏读、不可重复读的发生。

- Read committed (读已提交)：可避免脏读的发生。

- Read uncommitted (读未提交)：最低级别，任何情况都无法保证。

  MySQL设置数据库隔离级别：

  ```sql
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

#### JDBC的隔离级别设置：

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







