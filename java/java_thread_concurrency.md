## 线程安全性与原子性

**线程安全性（thread-safe）**：在并发环境和单线程环境都能保证正确性（行为与其规范完全一致）。

无状态对象（无共享变量）一定是线程安全的。

尽可能使用现有的线程安全对象（如AtomicLong）来管理类的状态。

要保持状态的一致性，就需要在单个原子操作中更新所有相关的状态变量。(x++不是复合操作，不具有原子性，它包含读取-写入-修改三个操作)





## synchronized与volatile 关键字

volatile是轻量级的synchronized，在多处理器（多线程）开发中，保证共享变量的“可见性”。

可见性的意思是，当一个线程修改一个共享变量时，另一个线程能读到这个修改的值。

#### synchronized关键字

synchronized关键字解决多个线程之间访问资源的同步性，被它修饰的 **方法或代码块** 在任意时刻只能有一个线程执行。

**使用方式：**

用于实例方法：作用于*当前对象实例*加锁，进入同步代码前要获得当前对象实例的锁；

用于静态方法：作用于*当前类对象*加锁，进入同步代码前要获得当前类对象的锁；

用于代码块：作用于*synchronized括号里的给定对象*加锁，进入同步代码块前要获得给定对象的锁。

#### Volatile关键字

volatile 关键字解决多线程中 **变量** 的可见性，然后还有一个作用是防止指令重排序。将变量声明为 volatile，这就指示 JVM这个变量是不稳定的，每次使用它都到主存中进行读取。

![image-20200122170241969](/Users/liuyuanyuan/github/StrongCode/java/images/threads-volatile.png)

- volatile关键字是线程同步的**轻量级实现**，所以volatile性能肯定比synchronized关键字要好。（synchronized常被称作重量级锁，但是synchronized关键字在 JavaSE1.6 之后为了减少获得锁和释放锁带来的性能消耗，而引入的偏向锁和轻量级锁及其它各种优化之后，执行效率有了显著提升；实际开发中使用 synchronized 关键字的场景还是更多一些。 ）

- volatile关键字主要解决变量在多个线程之间的可见性，而synchronized关键字解决的是多个线程之间访问资源的同步性。

  volatile关键字只能用于变量，而synchronized关键字可以修饰方法和代码块。

- 多线程访问volatile关键字不会发生阻塞，而synchronized关键字可能会发生阻塞。

- volatile关键字能保证数据的可见性，但不能保证数据的原子性。synchronized关键字两者都能保证。



## Atomic原子类

这里的Atomic是指一个操作是不可中断的。即使在多个线程一起执行时，一个操作一旦开始就不会被其他线程干扰。所以，所谓原子类说简单点就是具有原子操作特征的类。
原子类都存放在 java.util.concurrent.atomic 下，如下图所示。

![image-20200122172430238](/Users/liuyuanyuan/github/StrongCode/java/images/package-atomic.png)

JUC包中的原子类：

```
基本类型：使用原子的方式更新基本类型
AtomicInteger:整形原子类 
AtomicLong:长整型原子类
AtomicBoolean :布尔型原子类

数组类型：使用原子的方式更新数组里的某个元素
AtomicIntegerArray:整形数组原子类 
AtomicLongArray:长整形数组原子类 
AtomicReferenceArray :引用类型数组原子类

引用类型
AtomicReference:引用类型原子类 
AtomicStampedRerence:原子更新引用类型里的字段原子类 
AtomicMarkableReference :原子更新带有标记位的引用类型

对象的属性修改类型 
AtomicIntegerFieldUpdater:原子更新整形字段的更新器
AtomicLongFieldUpdater:原子更新长整形字段的更新器
AtomicStampedReference :原子更新带有版本号的引用类型。该类将整数值与引用关联起来，可用于解决原 子的更新数据和数据的版本号，可以解决使用 CAS 进行原子更新时可能出现的 ABA 问题。
```





#### AQS(AbstractQueuedSynchronizer)

java.util.concurrent.locks.AbstractQueuedSynchronizer



## 并发工具类

#### **Semaphore(信号量): **

允许多个线程同时访问，synchronized 和 ReentrantLock 都是一次只允许一个线程访问某个资源，Semaphore(信号量)可以指定多个线程同时访问某个资源。

#### **CountDownLatch (倒计时器):** 

是一个同步工具类，用来协调多个线程之间的同步。这 个工具通常用来控制线程等待，它可以让某一个线程等待直到倒计时结束，再开始执行。 

#### **CyclicBarrier(循环栅栏): **

CyclicBarrier 和 CountDownLatch 非常类似，它也可以实现线程间的技术等待，但是它的功能比 CountDownLatch 更加复杂和强大。主要应用场景和 CountDownLatch 类似。CyclicBarrier 的字面意思是可循环使用(Cyclic)的屏障(Barrier)。它要做的事情是，让一组线程到达一个屏障(也可以叫同步点)时被阻塞，直到最后一个线程到达屏障时，屏障才会开门，所有被屏障拦截的线程才会继续干活。 CyclicBarrier默认的构造方法是 CyclicBarrier(int parties)，其参数表示屏障拦截的线程数量，每个线程调用 await 方法告诉 CyclicBarrier 我已经到达了屏障，然后当前线程被阻塞。