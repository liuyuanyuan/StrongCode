# Java线程的创建

[TOC]

# 线程的创建

Thread 的构造方法如下：

<img src="images/thread-constructor.png" alt="image-20200205190341297" style="zoom:50%;" />

### 1 继承Thread父类：

作为Thread具体子类可直接执行。直接重写父类Thread的run方法。

### 2 实现Runnnable接口：

仅作为线程任务传给Thread执行。将任务与线程分离，以达到解耦；

### 3 通过Thread/Runnable匿名内部类创建线程

通过Runnable或Thread都可以创建匿名内部类；

```java
  //匿名内部类创建线程（同时实现Runnable并继承Thread时，只有Thread的run方法会执行）
	Thread t = new Thread(new Runnable() {
		@Override
		public void run() {
			System.out.println(" do runnable run..."); //not work
		}	
	}){
		@Override
		public void run() {
			System.out.println(" do thread run..."); //this work， this override run of runnable
		}
	};
	
	t.start();
```


### 4 创建带返回值的线程：通过Callable和FutureTask

**实现Callable<V>（参数V是返回值的类型）接口，并使用FutureTask<V>**：作为带返回值并可以抛出异常的线程任务，传入Thread执行；

```java
/**
 * @Desc   有返回值且能抛出异常的线程
 * @Author github.com/liuyuanyuan
 */
public class MyCallable implements Callable<Integer> {
	@Override
	public Integer call() throws Exception {
		System.out.println("computing...");
		return 1;
	}

	public static void main(String[] args) {
		MyCallable ct = new MyCallable();
		FutureTask<Integer> ftask = new FutureTask<>(ct);//Runnable的实现，是线程任务
		Thread t = new Thread(ftask);
		t.start();
		System.out.println("can do other things here...");
		try {
			Integer result = ftask.get();//await until done and return
			System.out.println("result = " + result);
		} catch (InterruptedException|ExecutionException e) {
			System.err.println(e.getMessage());
			e.printStackTrace();
	    } 
	}
}
```



### 5 定时器(Timer)创建多线程

-  **java.util.Timer**

定期执行定时任务，适用于简单的处理。

```
public static void main(String[] args) {
		Timer timer = new Timer();
		timer.schedule(new TimerTask() {
			@Override
			public void run() {
				System.out.println("run timer task ");
			}
		}, 0, 1000);
}
```

- **Quartz**

- **Spring**

  

### 6 线程池(以空间换时间)创建多线程

#### 线程池的创建方法

《阿里巴巴Java开发手册》中强制线程池不允许使用 Executors/ExecutorService去创建，而是通过 ThreadPoolExecutor的方式，这样更加明确线程池的运行规则，规避资源耗尽的风险。

> Executors/ExecutorService返回线程池对象的弊端如下:
>
> - FixedThreadPool 和 SingleThreadExecutor : 允许请求的队列长度为 Integer.MAX_VALUE,可能堆积 大量的请求，从而导致OOM。
> - CachedThreadPool 和 ScheduledThreadPool : 允许创建的线程数量为 Integer.MAX_VALUE ，可能 会创建大量线程，从而导致OOM。

#### 通过线程池工厂类Executors创建线程池：

(绿框中的方法)内部是通过ThreadPoolExecutor实现的。

![image-20200122172101493](images/Executors-thead.png)

- newFixedThreadPool(int) 固定数量线程池
- newCachedThreadPool() 缓存线程池
- newSingleThreadPool() 单个线程的线程池
- newScheduledThreadPool() 可定时执行线程任务的线程池




#### 通过线程池类ThreadPoolExecutor创建线程池（阿里推荐）：

**线程池的构造方法ThreadPoolExecutor**

> public ThreadPoolExecutor(int corePoolSize,
>                               int maximumPoolSize,
>                               long keepAliveTime,
>                               TimeUnit unit,
>                               BlockingQueue<Runnable> workQueue,
>                               ThreadFactory threadFactory,
>                               RejectedExecutionHandler handler)；

**传入参数:**

**corePoolSize** 保留在池中的线程数，即使它们处于空闲状态也被保留，除非设置了`allowCoreThreadTimeOut`。

**maximumPoolSize** 池中允许的最大线程数； 

**keepAliveTime** 当线程数大于**corePoolSize**时，多余的空闲线程在终止之前，等待新任务的最大时间；

**unit**  `keepAliveTime` 参数的时间单位；

**workQueue** 用来保存等待被执行的任务的阻塞队列.。此队列将仅保留`execute`方法提交的`Runnable`任务。

> JDK中提供了如下阻塞队列的实现类：
>
> (1) ArrayBlockingQueue：基于数组结构的有界阻塞队列，按FIFO排序任务； 
> (2) LinkedBlockingQuene：基于链表结构的阻塞队列，按FIFO排序任务，吞吐量通常要高于ArrayBlockingQuene； 
> (3) SynchronousQuene：一个不存储元素的阻塞队列，每个插入操作必须等到另一个线程调用移除操作，否则插入操作一直处于阻塞状态，吞吐量通常要高于LinkedBlockingQuene； 
> (4) PriorityBlockingQuene：具有优先级的无界阻塞队列（无界队列的风险：容易内存溢出）；
>
> LinkedBlockingQueue比ArrayBlockingQueue：在插入删除节点性能方面更优，但是二者在put(), take()任务的时均需要加锁；
>
> SynchronousQueue使用无锁算法，根据节点的状态判断执行，而不需要用到锁，其核心是Transfer.transfer()。

**threadFactory** 创建新线程时使用的工厂；

**handler** 因达到线程边界和队列容量，而阻塞执行时，要使用的处理程序handler。

**异常抛出**:

[IllegalArgumentException](eclipse-javadoc:☂=test/\/Library\/Java\/JavaVirtualMachines\/openjdk-12.0.2.jdk\/Contents\/Home\/lib\/jrt-fs.jar`java.base;~Ljava.util.concurrent.ThreadFactory;~Ljava.util.concurrent.RejectedExecutionHandler;☂IllegalArgumentException) - 当出现以下情况时。
	`corePoolSize < 0`
	`keepAliveTime < 0`
	`maximumPoolSize <= 0`
	`maximumPoolSize < corePoolSize`

[NullPointerException](eclipse-javadoc:☂=test/\/Library\/Java\/JavaVirtualMachines\/openjdk-12.0.2.jdk\/Contents\/Home\/lib\/jrt-fs.jar`java.base;~Ljava.util.concurrent.ThreadFactory;~Ljava.util.concurrent.RejectedExecutionHandler;☂NullPointerException) - 当 `workQueue` 、 `threadFactory` 、 `handler` 为 null 时。

 

**线程池的关闭方法**：

- shutdown() 将线程池里的线程状态设置成SHUTDOWN状态, 然后中断所有没有正在执行任务的线程；
- shutdownNow() 将线程池里的线程状态设置成STOP状态, 然后停止所有正在执行或暂停任务的线程；
  只要调用这两个关闭方法中的任意一个, isShutDown() 返回true. 
  当所有任务都成功关闭了, isTerminated()返回true.

关闭原理是：遍历线程池中的所有线程，然后逐个调用线程的`interrupt`方法来中断线程.



### 线程池的实现原理

**当一个线程任务提交至线程池之后，** 

1 线程池先看：当前运行的线程数量<corePoolSize ？是，则创建一个新的工作线程来执行任务；否，则表示线程池已满且都在工作中，则进入2；

2 判断线程等待队列 workQueue 是否已经满了？没满，则将线程放入 workQueue；满了，进入3；

3 如果新建一个工作线程，使当前运行的线程数量>maximumPoolSize，则交给RejectedExecutionHandler来处理任务。

当 ThreadPoolExecutor 创建新线程时，通过 CAS 来更新线程池的状态ctl。



### 线程池的饱和策略

当池中线程已满且都在工作中，会将新提交的任务先加入工作队列中，等到有空闲线程时再从工作队列中获取。

工作队列有两种实现策略：

- 无界队列（不可取）：不存在饱和的问题，但是其问题是当请求持续高负载的话，任务会无脑的加入工作队列，那么很可能导致内存等资源溢出或者耗尽。
- 有界队列（java线程池采用的）：不会带来高负载导致的内存耗尽的问题，但会出现工作队列已满时，新提交的任务如何管理的难题，这就是线程池工作队列饱和策略要解决的问题。

1 AbortPolicy：默认策略，新任务提交时直接抛出异常RejectedExecutionException；

2 CallerRunsPolicy：既不抛弃任务也不抛出异常，而是将拒绝的任务在调用execute()方法的线程中运行；

3 DiscardPolicy：以静默(没有任何其他反应)方式抛弃提交的任务；

4 DiscardOldestPolicy：抛弃等待队列中最老的请求任务，然后重新尝试execute，直到executor被关闭（ 这时任务也被抛弃）；（不适合工作队列为优先队列场景）

```java
threadPoolExecutor.setRejectedExecutionHandler(new ThreadPoolExecutor.AbortPolicy());
threadPoolExecutor.setRejectedExecutionHandler(new ThreadPoolExecutor.CallerRunsPolicy()));
threadPoolExecutor.setRejectedExecutionHandler(new ThreadPoolExecutor.DiscardPolicy());
threadPoolExecutor.setRejectedExecutionHandler(new ThreadPoolExecutor.DiscardOldestPolicy());
```

实验：

```
public static void main(String[] args) {
		//ExecutorService pool = Executors.newFixedThreadPool(3);//固定缓存线程数
		//ExecutorService pool = Executors.newCachedThreadPool();//智能缓存线程数
        ThreadPoolExecutor pool = 
        		new ThreadPoolExecutor(2, 3, 0l, TimeUnit.MILLISECONDS, new LinkedBlockingDeque<>(5));
        //pool.setRejectedExecutionHandler(new ThreadPoolExecutor.AbortPolicy());//default
        pool.setRejectedExecutionHandler(new ThreadPoolExecutor.CallerRunsPolicy());
        //pool.setRejectedExecutionHandler(new ThreadPoolExecutor.DiscardOldestPolicy());
        //pool.setRejectedExecutionHandler(new ThreadPoolExecutor.DiscardPolicy());
		for(int i=1; i<=10; i++)
		{
			pool.execute(new Runnable(){
				@Override
				public void run() {
					System.out.println(Thread.currentThread().getName() + " is running..");
				}
			});
		}
		//pool.shutdown();//shutdown all threads of the pool
	}
```




### 7 Lambda表达式(since Java8)实现多线程

Lambda 表达式，也可称为闭包，它是推动 Java 8 发布的最重要新特性。

Lambda 允许把函数作为一个方法的参数（函数作为参数传递进方法中）。使用 Lambda 表达式可以使代码变的更加简洁紧凑。

```
public class LambdaThread {
	
	public static void main(String[] args) {
		List<Integer> values = Arrays.asList(10,20,30,40);
		int res = new LambdaThread().add(values);
		System.out.println("result=" + res);
	}

	public int add(List<Integer> values){
		//values.parallelStream().forEach(System.out::println);//并行打印，结果无序
		//values.parallelStream().forEachOrdered(System.out::println);//有序打印
		return values.parallelStream().mapToInt(i -> i).sum();	
	}
}
```



### 8 Spring Framework实现多线程

Eclipse创建Maven简单项目（preference-maven-installation&user config-create maven project - single project），进行Spring Framework（现在是SprintBoot2.2.4）的初始配置，参考官网 [Project - Spring Framework - Quick Start](https://start.spring.io)   和 [Developing Your First Spring Boot Application](https://docs.spring.io/spring-boot/docs/2.2.4.RELEASE/reference/html/getting-started.html#getting-started-first-application) 。

```java
@SpringBootApplication
@EnableAsync
public class AppConfig {
	
	//just for test
	public static void main(String[] args) {
		//SpringApplication.run(AppConfig.class, args);
		AnnotationConfigApplicationContext ac = 
				new AnnotationConfigApplicationContext(AppConfig.class);
		DemoService ds = ac.getBean(DemoService.class);
		ds.a();
		ds.b();
	}
}

@Service
public class DemoService {
	@Async
	public void a() {
		while(true) {
			System.out.println("a---");
			try {
				Thread.sleep(1000);
			} catch (InterruptedException e) {
				e.printStackTrace();
			}
		}
	}
	
	@Async
	public void b() {
		while(true) {
			System.out.println("b---");
			try {
				Thread.sleep(1000);
			} catch (InterruptedException e) {
				e.printStackTrace();
			}
		}
	}
}

```



