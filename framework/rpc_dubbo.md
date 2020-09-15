[TOC]

## Java RPC 框架 - Dubbo 

### 概述：

Apache Dubbo 是一款高性能、轻量级的开源 Java RPC 框架，提供了三大核心能力: 

- 面向接口的远程方法调用；RPC>Netty>Java NIO 网络通信>基于TCP协议传输
- 智能容错(重试机制)和负载均衡(集群部署、均衡消费、高可用)；
- 服务自动注册和发现；

简单来说 Dubbo 是一个分布式服务框架，致力于提供高性能和透明化的 RPC 远程服务调用方案，以及 SOA 服务治理方案。

决定RPC效率的两大因素：

- 通信效率
- 所传输对象的序列化/反序列化效率：数据流>JSON>XML

整体架构：

- 服务注册、配置中心
- 服务提供端
- 服务消费端

### 注册中心

- zookeeper



[nacos](https://nacos.io/zh-cn/docs/use-nacos-with-dubbo.html)

推荐使用Spring外部化配置 

pom.xml 

```xml
<dependencies>
    <!-- Dubbo dependency -->
    <dependency>
        <groupId>com.alibaba</groupId>
        <artifactId>dubbo</artifactId>
        <version>[latest version]</version>
    </dependency>
    
    <!-- 使用Spring装配方式时可选: -->
    <dependency>
        <groupId>com.alibaba.spring</groupId>
        <artifactId>spring-context-support</artifactId>
        <version>[latest version]</version>
    </dependency>
</dependencies>
```

application.properties

```properties
## application
dubbo.application.name = your-dubbo-application
## Zookeeper registry address
dubbo.registry.address = zookeeper://10.20.153.10:2181
```









## 分布式微服务框架比较：

最大的区别是：

- Dubbo底层是使用Netty这样的NIO框架，是基于TCP协议传输的，配合以 Hession 序列化完成 RPC 通信。

- 而SpringCloud是基于Http协议+RESTFul接口调用远程过程的通信，

- 相对而言，Http请求会有更大的报文，占的带宽也会更多。但是 REST 相比 RPC 更为灵活，服务提供方和调用方的依赖只依靠一纸契约，不存在代码级别的强依赖，这在强调快速演化的微服务环境下，显得更为合适，至于注重通信速度还是方便灵活性，具体情况具体考虑。

|              | Dubbo                                  | SpringCloud                  |
| ------------ | -------------------------------------- | ---------------------------- |
| 通信协议     | TCP                                    | tp（基于TCP）                |
| 服务调用方式 | RPC接口（基于Netty>Java NIO 网络通信） | RESTFul API接口              |
|              | 长连接                                 | 短链接                       |
| 服务注册中心 | Zookeeper/Nacos                        | Spring Cloud Netfix Eureka   |
| 服务监控     | Dubbo-monitor                          | Spring Boot Admin            |
| 熔断器       | 不完善                                 | Spring Cloud Netflix Hystrix |
| 服务网关     | 无                                     | Spring Cloud Netflix Zuul    |
| 分布式配置   | 无                                     | Spring Cloud Config          |
| 服务跟踪     | 无                                     | Spring Cloud Sleuth          |
| 数据流       | 无                                     | Spring Cloud Stream          |
| 批量任务     | 无                                     | Spring Cloud Task            |
| 信息总线     | 无                                     | Spring Cloud Bus             |

 

 

