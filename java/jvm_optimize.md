# JVM 调优

[TOC]

## 性能监控与测试

### 性能衡量指标

- 吞吐量
  - QPS：每秒查询数
  - TPS：每秒事务数
  - HPS：每秒HTTP请求数

- 响应时间
  - AVG：平均响应时间
  - TOP Percent：百分位数

- 并发量

- 秒开率
- 正确性



### 性能监控工具

- nmon：Linux操作系统性能数据；

- jvisualvm（通过插件可以获得更多数据的监控）：JVM性能数据；

- jmc： Java 应用详细性能数据；

- arthas：Java请求的调用链耗时；

- wrk： web接口的性能数据；

  

### 基准测试(Benchmark)

- 基准测试工具JMH



![image-20200911135008806](/Users/liuyuanyuan/Library/Application Support/typora-user-images/image-20200911135008806.png)

![image-20200911141342729](/Users/liuyuanyuan/Library/Application Support/typora-user-images/image-20200911141342729.png)



## 性能优化的7类手段

- 业务优化
  - 复用优化
  - 结果集优化
  - 采用高效的业务实现
  - 算法优化
- 技术优化
  - 计算优化
  - 资源冲突优化
  - JVM优化



#### 选择恰当的设计模式编写代码，来减少内存占用、提高运行性能；

#### 通过多线程并发(并行计算)提高处理速度-使用线程池

- I/O密集型任务

  常见的互联网服务，大多是属于I/O密集型的，比如等待数据库的I/O，等待网络的I/O等；

- 计算密集型任务

  计算密集型的的任务较少，比如一些耗时的算法逻辑；

  CPU要想达到最高的利用率，提高吞吐量，最好的方式就是尽量少地在任务之间切换，此时线程数等于CPU数量，是效率最高的。

![image-20200911210403942](/Users/liuyuanyuan/Library/Application Support/typora-user-images/image-20200911210403942.png)

![image-20200911210706619](/Users/liuyuanyuan/Library/Application Support/typora-user-images/image-20200911210706619.png)

<img src="/Users/liuyuanyuan/Library/Application Support/typora-user-images/image-20200911125411690.png" alt="image-20200911125411690" style="zoom:50%;" />







