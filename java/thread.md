# 线程的概念

[TOC]

## 线程的状态

![image-20200205193457003](/Users/liuyuanyuan/github/StrongCode/java/images/thread-states.png)



1. 初始(NEW)：新创建了一个线程对象，但还没有调用start()方法。
2. 运行(RUNNABLE)：Java线程中将就绪（ready）和运行中（running）两种状态笼统的称为“运行”。
   线程对象创建后，其他线程(比如main线程）调用了该对象的start()方法。该状态的线程位于可运行线程池中，等待被线程调度选中，获取CPU的使用权，此时处于就绪状态（ready）。就绪状态的线程在获得CPU时间片后变为运行中状态（running）。
3. 阻塞(BLOCKED)：表示线程阻塞于锁。
4. 等待(WAITING)：进入该状态的线程需要等待其他线程做出一些特定动作（通知或中断）。
5. 超时等待(TIMED_WAITING)：该状态不同于WAITING，它可以在指定的时间后自行返回。
6. 终止(TERMINATED)：表示该线程已经执行完毕。



## Thread.java类详解：

几种构造方法：

<img src="images/thread-constructor.png" alt="image-20200205190341297" style="zoom:50%;" />

```java
   // Thread.java 代码段
   public Thread(Runnable target) {
        this(null, target, "Thread-" + nextThreadNum(), 0);
    }
     /* @param  group
     *         the thread group. If {@code null} and there is a security
     *         manager, the group is determined by {@linkplain
     *         SecurityManager#getThreadGroup SecurityManager.getThreadGroup()}.
     *         If there is not a security manager or {@code
     *         SecurityManager.getThreadGroup()} returns {@code null}, the group
     *         is set to the current thread's thread group.
     *
     * @param  target
     *         the object whose {@code run} method is invoked when this thread
     *         is started. If {@code null}, this thread's run method is invoked.
     *
     * @param  name
     *         the name of the new thread
     *
     * @param  stackSize
     *         the desired stack size for the new thread, or zero to indicate
     *         that this parameter is to be ignored
     *
     * @param  inheritThreadLocals
     *         if {@code true}, inherit initial values for inheritable
     *         thread-locals from the constructing thread, otherwise no initial
     *         values are inherited
     *
     * @throws  SecurityException
     *          if the current thread cannot create a thread in the specified
     *          thread group
     *
     * @since 9
     */
    public Thread(ThreadGroup group, Runnable target, String name,
                  long stackSize, boolean inheritThreadLocals) {
        this(group, target, name, stackSize, null, inheritThreadLocals);
    }

    @Override
    public void run() {
        if (target != null) {
            target.run();
        }
    }
```

### 1 start() 启动线程 & run() 运行线程体

- start() 是启动线程
- run()  是运行线程体中的具体代码

当一个线程类同时实现Runnable和继承Thread时，只有重写的Thread的run方法会被执行。

（当Thread的run方法没有重写时，才会执行Runnable的run方法。）

```java
 Thread t = new Thread(new Runnable() {
		@Override
		public void run() {
			System.out.println(" do runnable run..."); //not work
		}	
	}) {
		@Override
		public void run() {
			System.out.println(" do thread run..."); //this work
		}
	};
	
  t.start();
```

### 2 setDaemon() 守护线程

将线程设置为守护线程setDaemon(true)后，当主线程结束后守护线程也会结束。

> 守护线程：也称服务线程，是后台线程，为用户线程提供公共服务，在没有用户线程可服务时会自动离开。
>
> 优先级：守护线程的优先级比较低，用于为系统中的其它对象和线程提供服务。
>
> 设置方法：在线程对象启动之前，用 setDaemon(true) 设置。在 Daemon 线程中产生的新线程也是 Daemon 的。

```
public static void main(String[] args) {
		MyThread th1 = new MyThread("ThreadImpl1");
		MyThread th2 = new MyThread("ThreadImpl2");
		
		/* The JVM exits when the only threads running are all daemon threads.
		 * after setDaemon(true), th1 and th2 will terminate when main thread terminate
	   */
		th1.setDaemon(true);
		th2.setDaemon(true);
		
		th1.start();
		th2.start();
}
```

### 3 interpreted()  终止线程

线程的终止使用interpreted（stop()已弃用，因为它只是让线程无限等待并未真正释放所有资源）。

```java
public class MyThread extends Thread{
	@Override
	public void run() {
    //通过线程终止状态拍决定是否执行
		while(!interrupted()){
			System.out.println(getName()+ " is running");
		}
	}

	public static void main(String[] args) {
		MyThread th1 = new MyThread("ThreadImpl1");
		MyThread th2 = new MyThread("ThreadImpl2");

		th1.start();
		th2.start();
		
		//th1.stop();//Deprecated, 只是无限期等待并未释放所有资源，不是真正的停止
		th1.interrupt();//停止当前线程，并将终止状态interrupted置为true
	}
```



## 常见线程方法：

### Thread.sleep(long millis)

一定是当前线程调用此方法，当前线程进入TIMED_WAITING状态，但不释放对象锁，millis后线程自动苏醒进入就绪状态。

作用：给其它线程执行机会的最佳方式。

>区分 Thread.sleep(interval) 与 Object .wait()
>
>- sleep()属于Thread类，wait()属于Object类；
>
>- 调用sleep()方法时，线程不会释放对象锁；
>
>  sleep()方法使程序暂停执行（指定时长），让出cpu给其他线程，但它的监控器状态依然保持，当指定时间到了，又会自动恢复运行状态。
>
>- 调用wait()方法时，线程会释放对象锁，
>
>  进入该<对象的等待锁定池>，只有针对此对象调用notify()方法后，本线程才进入<对象的锁定池>，准备获取对象锁进入运行状态。

### Thread.yield()

一定是当前线程调用此方法，当前线程放弃获取的CPU时间片，但不释放锁资源，由运行状态变为就绪状态，让OS再次选择线程。

作用：让相同优先级的线程轮流执行，但并不保证一定会轮流执行。

实际中无法保证yield()达到让步目的，因为让步的线程还有可能被线程调度程序再次选中。Thread.yield()不会导致阻塞。该方法与sleep()类似，只是不能由用户指定暂停多长时间。

### object wait - notify / notifyAll 机制

在Object.java中，定义了wait(), notify()和notifyAll()等接口。

wait()的作用是让当前线程释放它所持有的对象锁，让当前线程进入等待队列/等待状态。需要依靠notify()/notifyAll() 唤醒，或者使用 wait(long timeout)时 timeout时间到自动唤醒。

notify()和notifyAll()的作用则是唤醒当前对象上的等待的线程；notify()是唤醒单个线程，选择是随机的；而notifyAll()是唤醒在此对象监视器上等待的所有的线程。

Object类中关于等待/唤醒的API详细信息如下：

- **notify()**：唤醒在此对象监视器上等待的单个线程。
- **notifyAll()**：唤醒在此对象监视器上等待的所有线程。
- **wait()**：让当前线程处于“等待(阻塞)状态”，“直到其他线程调用此对象的 notify() 方法或 notifyAll() 方法”，当前线程被唤醒(进入“就绪状态”)。
- **wait(long timeout)**：让当前线程处于“等待(阻塞)状态”，“直到其他线程调用此对象的 notify() 方法或 notifyAll() 方法，或者超过指定的时间量”，当前线程被唤醒(进入“就绪状态”)。
- **wait(long timeout, int nanos)**：让当前线程处于“等待(阻塞)状态”，“直到其他线程调用此对象的 notify() 方法或 notifyAll() 方法，或者其他某个线程中断当前线程，或者已超过某个实际时间量”，当前线程被唤醒(进入“就绪状态”)。

### LockSupport.park() / LockSupport.parkNanos(long nanos) 

 LockSupport.parkUntil(long deadlines)，当前线程进入 WAITING/TIMED_WAITING 状态。对比 wait 方法，不需要获得锁就可以让线程进入 WAITING/TIMED_WAITING 状态，需要通过 LockSupport.unpark(Thread thread) 唤醒。

### thread.join() / thread.join(long millis)

让父线程等待子线程结束之后才能继续运行。

当前线程调用子线程的join()，使当前线程进入WAITING/TIMED_WAITING状态，当前线程不会释放已经持有的对象锁。线程执行完毕或者millis时间到，当前线程一般会进入RUNNABLE状态，也可能进入BLOCKED状态（因为join是基于wait实现的）。

join()方法定义在Thread.java中:

```java
 //join()等效于join(0),  永久等待直到该子线程结束
 public final void join() throws InterruptedException{
		join(0);
 }
 //最多等待 millis毫秒，超时直接回到主线程，主线程与子线程并发
 public final synchronized void join(final long millis)
    throws InterruptedException {
    ...
 
//最多等待 milli毫秒+nanos纳秒，超时直接回到主线程，主线程与子线程并发
 public final synchronized void join(long millis, int nanos)
    throws InterruptedException {
    ...
 }
```

用法1：

```
Thread t1 = new Thread();
Thread t2 = new Thread();
Thread t3 = new Thread();
//t1>main>t2>main>t3>main
t1.start();
t1.join();
t2.start();
t2.join();
t3.start();
t3.join();
```

<img src="/Users/liuyuanyuan/github/StrongCode/java/images/thread_join.png" alt="image-20200309140703617" style="zoom:50%;" />

用法2:

```java
public static void main(String[] args) throws InterruptedException {
		Thread t = new Thread(new Runnable() {
			@Override
			public void run() {
				try {
					Thread.sleep(5);
					System.out.println("subthread....5");
					Thread.sleep(15);
					System.out.println("subthread....20");
					Thread.sleep(10);
					System.out.println("subthread....30");
				} catch (InterruptedException e) {
					e.printStackTrace();
				}
			}
		});
		t.start();
		t.join(10);
		System.out.println("main...10");
		Thread.sleep(10);
		System.out.println("main....20");
	}
```







