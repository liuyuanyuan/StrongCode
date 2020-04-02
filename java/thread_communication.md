# Java中线程间通信

[TOC]

分布式系统中说的两种通信机制：共享内存机制和消息通信机制。

管道通信更像消息传递机制，也就是说：通过管道，将一个线程中的消息发送给另一个。其他几种通信方式更像是共享内存机制。

## 一、共享内存机制

### 1 while(volatile x) 轮询变量

```
public class Test {
	static Integer total = 5;
	public static void main(String[] args) {
		Thread t1 = new Thread() {
			@Override
			public void run() {
				System.out.println(Thread.currentThread().getName());
				while (total > 0) {
					total--;
					System.out.println("total=" + total);
				}
			}
		};
		Thread t2 = new Thread() {
			@Override
			public void run() {
				System.out.println(Thread.currentThread().getName());
				while (total > 0) {
					//keep here until break
				}
				System.err.println("total is finished");
			}
		};
		t2.start();
		t1.start();
	}
}

输出：
thread2
thread1
total=4
total=3
total=2
total=1
total=0
total is finished
```

线程t2通过不停执行while语句检测条件(total>0)， 直到 total <= 0 才执行 ，从而实现了线程间的通信。但是这种方式会浪费CPU资源，因为 JVM 调度器将 CPU 交给线程B执行时，它没做啥"有用”的工作，只是在不断地测试某个条件是否成立。

### 2 Synchronized x同步

```java
public class SyncAccount {
	private int balance;
	public SyncAccount(int balance) {
		this.balance = balance;
	}
	public synchronized void draw(int drawAmount) {
		System.out.println(Thread.currentThread().getName() + " before draw " + balance);
		balance = balance - drawAmount;
		System.out.println(Thread.currentThread().getName() + " after draw " + balance);

	}
	public synchronized void deposit(int depositAmount) {
		System.out.println(Thread.currentThread().getName() + " before deposit " + balance);
		balance = balance + depositAmount;
		System.out.println(Thread.currentThread().getName() + " after deposit " + balance);
	}
	
	public static void main(String[] args) {
		SyncAccount account = new SyncAccount(10000);
		Thread t1 = new Thread("t1") {
			@Override
			public void run() {
				System.out.println(Thread.currentThread().getName());
				account.draw(100);
			}
		};
		Thread t2 = new Thread("t2") {
			@Override
			public void run() {
				System.out.println(Thread.currentThread().getName());
				account.deposit(20);
			}
		};
		Thread t3 = new Thread("t3") {
			@Override
			public void run() {
				System.out.println(Thread.currentThread().getName());
				account.draw(40);
				SyncAccount.staticMethod();
			}
		};
		t1.start();
		t2.start();
		t3.start();
	}
}
```

线程t1、t2、t3持有同一个SyncAcount类的实例account，虽然3个线程调用不同的方法，但是它们是同步执行的，即一个线程执行完毕才可以运行另一个。所以：3个线程谁拿到了锁（获得了执行权），谁就可以执行，执行完毕释放锁（让出执行权）。

### 3 Object wait - notify()/notifyAll() 对象的等待-唤醒机制

```java
public class TestWaitNotify {
  public static void main(String[] args) {
		Object lock = new Object();
		Thread t1 = new Thread("thread1") {
			@Override
			public void run() {
				synchronized (lock) {
					System.out.println(Thread.currentThread().getName());
					int total = 5;
					while (total > 0) {
						total--;
						System.out.println("total=" + total);
					}
					lock.notifyAll();
				}
			}
		};
		Thread t2 = new Thread("thread2") {
			@Override
			public void run() {
				synchronized (lock) {
					System.out.println(Thread.currentThread().getName());
					try {
						lock.wait();
					} catch (InterruptedException e) {
						e.printStackTrace();
					}
					System.err.println(Thread.currentThread().getName() + " total is finished");
				}
			}
		};
		Thread t3 = new Thread("thread3") {
			@Override
			public void run() {
				synchronized (lock) {
					System.out.println(Thread.currentThread().getName());
					try {
						lock.wait();
					} catch (InterruptedException e) {
						e.printStackTrace();
					}
					System.err.println(Thread.currentThread().getName() + " total is finished");
				}
			}
		};
		
		t3.start();
		t2.start();
		t1.start();
	}
}
```

线程t2和t3调用 wait() 放弃CPU，并进入阻塞状态。当条件满足时，线程t1调用 notifyAll() 通知线程t2和t3（就是唤醒线程t2和t3，并让它进入可运行状态）。

这种方式的一个好处就是CPU的利用率提高了（相比于while轮询）。

缺点是：线程t2或t3如果在t1执行完 notifyAll() 才执行wait()，那这个线程永远就不可能被唤醒了。因为，线程 t1 已经发了通知了，以后不再发通知了。这说明：**通知过早，会打乱程序的执行逻辑。**

### 4 Lock - Condition 机制

```java
public class Account {
	private Lock lock = new ReentrantLock();
	
	private long balance;//余额
	public Account(int initBalance){
		this.balance = initBalance;
	}
	//取款
	public void draw(int drawAmount) {
		lock.lock();
		System.out.println(Thread.currentThread().getName() + " before draw " + balance);
		balance = balance - drawAmount;
		System.out.println(Thread.currentThread().getName() + " after draw " + balance);
		lock.unlock();
	}
	//存款
	public void deposit(int depositAmount) {
		lock.lock();
		System.out.println(Thread.currentThread().getName() + " before deposit " + balance);
		balance = balance + depositAmount;
		System.out.println(Thread.currentThread().getName() + " after deposit " + balance);
		lock.unlock();
	}
	
	public static void main(String[] args) {
		Account account = new Account(10000);
		Thread t1 = new Thread("t1") {
			@Override
			public void run() {
				System.out.println(Thread.currentThread().getName());
				account.draw(100);
			}
		};
		Thread t2 = new Thread("t2") {
			@Override
			public void run() {
				System.out.println(Thread.currentThread().getName());
				account.deposit(20);
			}
		};
		Thread t3 = new Thread("t3") {
			@Override
			public void run() {
				System.out.println(Thread.currentThread().getName());
				account.draw(40);
			}
		};
		t1.start();
		t2.start();
		t3.start();
	}
}
```



## 二、5 管道(PipedxxStream)通信机制

管道流是 Java 中线程通讯的常用方式之一，基本流程如下：

1）创建管道输出流: PipedOutputStream pos 和管道输入流 PipedInputStream pis ;

2）将pos和pis匹配: pos.connect(pis);

3）将pos赋给信息输入线程，pis赋给信息获取线程，就可以实现线程间的通讯了

管道流虽然使用起来方便，但是也有一些缺点

1）管道流只能在两个线程之间传递数据

线程consumer1和consumer2同时从pis中read数据，当线程producer往管道流中写入一段数据后，每一个时刻只有一个线程能获取到数据，并不是两个线程都能获取到producer发送来的数据，因此一个管道流只能用于两个线程间的通讯。

不仅仅是管道流，其他 IO 方式都是一对一传输。

2）管道流只能实现单向发送，如果要两个线程之间互通讯，则需要两个管道流

可以看到上面的例子中，线程producer通过管道流向线程consumer发送数据，如果线程consumer想给线程producer发送数据，则需要新建另一个管道流pos1和pis1，将pos1赋给consumer1，将pis1赋给producer。

```java
public class TestPipeConnect {
	public static void main(String[] args) {
		  //创建管道输出流
        PipedOutputStream pos = new PipedOutputStream();
        //创建管道输入流
        PipedInputStream pis = new PipedInputStream();
        try {
            //将管道输入流与输出流连接,此过程也可通过重载的构造函数来实现
            pos.connect(pis);
        } catch (IOException e) {
            e.printStackTrace();
        }
        //创建生产者线程
        Producer p = new Producer(pos);
        //创建消费者线程
        Consumer c1 = new Consumer(pis);

        //启动线程
        p.start();
        c1.start();
	}
}
public class Producer extends Thread {
	private PipedOutputStream pos;
	public Producer(PipedOutputStream pos) {
		this.pos = pos;
	}
	public void run() {
		int i = 0;
		try {
			while (true) {
				this.sleep(3000);
				pos.write(i);
				i++;
			}
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
}
public class Consumer extends Thread {
	private PipedInputStream pis;
	public Consumer(PipedInputStream pis) {
        this.pis = pis;
  }
	public void run() {
		try {
			while (true) {
				System.out.println(Thread.currentThread().getName() + "consumer:" + pis.read());
			}
		} catch (IOException e) {
			e.printStackTrace();
		}
	}
}

程序启动后，就可以看到producer线程往consumer1线程发送数据:
Thread-1consumer:0
Thread-1consumer:1
Thread-1consumer:2
Thread-1consumer:3
...
```

