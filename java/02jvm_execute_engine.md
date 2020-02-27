## 垃圾回收GC

### 判断对象可以被回收

- 引用计数器法：引用+1；不再引用-1；不再有引用时回收。

​						缺点是循环引用（两个对象互相引用）时无法回收。

- 可达性分析算法：

### 判断常量可以被回收：

类似于引用计数法。

### 判断类是废弃的：

需要满足三个条件：



### 垃圾回收算法：

标记-清除算法：



标记-复制算法：



标记-整理算法：



分代收集算法（最常用的）



### 垃圾收集器：

Serial 收集器：最基本、历史最悠久的垃圾收集器，单线程；

ParNew收集器：Serial 收集器的多线程版本；

Parallel Scavenge收集器（jdk1.8使用的）：类似于ParNew收集器；

Serial Old 收集器：

Parallel Old 收集器：

CMS 收集器：



G1（Garbage-first）收集器：



### 如何选择垃圾收集器：

1 优先调整堆大小

2 

3

4

5





## 调优

### JVM调优的两个主要指标

- 停顿时间：

- 吞吐量：



### GC调优

1 打印GC日志；

javr -jar	-XX:+PrintGCDetails -XX:+PrintGCTimeStamps -XX:PrintGCDateStamps -Xloggc:./gc.log -XXMeraspaceSize=128M  app.jar

2 分析GC日志（通过工具，如GCeasy）得到当前停顿时间和吞吐量；

3 分析GC原因，调优JVM参数；







