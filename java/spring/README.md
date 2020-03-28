# Spring全家桶



## 概念

**概念：**Spring是一个全面的、企业应用开发一站式的解决方案，贯穿表现层、业务层、持久层。但是 Spring仍然可以和其他的框架无缝整合。

**特点**：是一个**轻量级**的**控制反转（IoC）和面向切面（AOP）**的**容器**，是基于Java的 开发框架的集合。

##### 区分 Spring 家族成员：

- **Spring** 是专业开发web项目的开源框架的集合。

- **SpringMVC** 是 Spring内部的一个模块(module)，同样采取 MVC 设计模式。 所以在使用 Spring 开发web项目时，作为核心环节的 MVC 可以使用struts1 / struts2 / SpringMVC。

- **Spring Boot** 是 Build Anything；实现自动配置，降低项目搭建的复杂度。https://spring.io/quickstart 这里可以自动生成初始项目，可导入IDE直接使用。

- **Spring Cloud** 是 Coordinate Anything；

##### 区分 jpa 和 hibernate： 

- JPA 即 Java Persistence API，是Java EE 5的标准ORM（Object Relation Map）接口，也是ejb3规范的一部分。
- Hibernate 是当今很流行的ORM框架，是JPA的一个实现，但其功能是JPA的超集。
- JPA 和 Hibernate之间的关系，可以简单的理解为 JPA 是标准接口，Hibernate是实现。那么Hibernate 与 JPA 的这种关系主要是通过三个组件来实现的：hibernate-annotation、hibernate-entitymanager、hibernate-core。



Spring 特点：



Spring Bean的作用域：







## AOP

AOP思想的实现一般都是基于**代理模式** ，在Java中一般采用JDK动态代理模式，但我们知道，JDK动态代理模式只能代理接口而不能代理类。因此，Spring AOP 会根据情况进行切换，因为Spring AOP 同时支持 CGLIB、 ASPECTJ、JDK动态代理：

如果目标对象的实现类实现了接口，Spring AOP 将会采用 JDK 动态代理来生成 AOP 代理类；如果目标对象的实现类没有实现接口，Spring AOP 将会采用 CGLIB 来生成 AOP 代理类，不过这个选择过程对开发者完全透明、开发者也无需关心。

这部分内容可以查看下面这几篇文章:

https://www.jianshu.com/p/fe8d1e8bd63e 

http://www.cnblogs.com/puyangsky/p/6218925.html

https://juejin.im/post/5a55af9e518825734d14813f

> 静态代理：
>
> 动态代理：



## IOC

Spring IOC的初始化过程:

> XML =读取> Resource  =解析> BeanDefination  =注册> BeanFactory

IOC源码阅读：https://javadoop.com/post/spring-ioc



