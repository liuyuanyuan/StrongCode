# springboot+mybatis

mybatis是Java语言的数据库访问框架，对数据库的JDBC访问进行了封装，提供了一系列API，方便更快实现项目业务；

mybatis-plus是mybatis的增强版框架，提供了更丰富的API。



## Spring系列框架的事务处理

### Springboot中数据库事务处理有两种方式:

- 编程式事务管理：使用 TransactionTemplate (推荐) 和使用 PlatformTransactionManager

  >  Spring 团队推荐使用 TransactionTemplate。

- 声明式事务管理：基于AOP原理的注解 @Transactional 和 @EnableTransactionManagement

  > 在以前的spring版本中，程序启动器还需要加@EnableTransactionManagement
  >
  > 而在Springboot中事务管理是默认添加的。



### 常见坑点

- 遇到检测异常，默认不回滚

Spring的默认的事务规则是遇到运行异常（RuntimeException及其子类）和程序错误（Error）才会进行事务回滚，显然SQLException并不属于这个范围。如果想针对检测异常进行事务回滚，可以在 @Transactional 注解里使用
 rollbackFor 属性明确指定异常。

```java
    @Transactional(rollbackFor = Exception.class)
    public void addMoney() throws Exception {
        //先增加余额
        accountMapper.addMoney();
        //然后遇到故障
        throw new SQLException("发生异常了..");
    }
```

- 在业务层捕捉异常后，发现事务不生效。

在业务层手工捕捉并处理了异常，你都把异常“吃”掉了，Spring自然不知道这里有错，更不会主动去回滚数据。例如：

```java
    @Transactional
    public void addMoney() throws Exception {
        //先增加余额
        accountMapper.addMoney();
        //谨慎：尽量不要在业务层捕捉异常并处理
        try {
            throw new SQLException("发生异常了..");
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
```

推荐做法：若非实际业务要求，则在业务层统一抛出异常，然后在控制层统一处理。

```java
    @Transactional
    public void addMoney() throws Exception {
        //先增加余额
        accountMapper.addMoney();
        //推荐：在业务层将异常抛出
        throw new RuntimeException("发生异常了..");
    }
```

