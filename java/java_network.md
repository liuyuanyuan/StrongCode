# Java 网络编程

[TOC]

网络编程是指编写运行在多个设备（计算机）的程序，这些设备都通过网络连接起来。

java.net 包中 Java SE 的 API 包含有类和接口，它们提供低层次的通信细节。你可以直接使用这些类和接口，来专注于解决问题而不用关注通信细节。

java.net 包中提供了两种常见的网络协议的支持：

- **TCP**（传输控制协议），它保障了两个应用程序之间的可靠通信。通常用于互联网协议，被称 TCP / IP。
- **UDP**（用户数据报协议），一个无连接的协议。提供了应用程序之间要发送的数据的数据包。

主要知识点：

- **URL 处理**：更详细地了解在 [Java 语言中的 URL 处理](https://www.runoob.com/java/java-url-processing.html)。
- **Socket 编程**：是使用最广泛的网络概念。



## URL 处理

URL（Uniform Resource Locator，统一资源定位符）有时也被俗称为网页地址。表示为互联网上的资源，如网页或者FTP地址。

### URL组成和格式

URL组成如下：

```
protocol://host:port/path?query#fragment
```

- protocol(协议)：可以是 HTTP、HTTPS、FTP 、nntp File；

- port (端口号)：端口号标识一个主机上进行通信的不同的应用程序。

  - HTTP 协议代理服务器常用端口号：80/8080/3128/8081/9098

  - SOCKS 代理协议服务器常用端口号：1080

  - FTP(文件传输) 协议代理服务器常用端口号：21

  - Telnet(远程登录) 协议代理服务器常用端口号：23

- path为文件路径及文件名。

> 举例说明：
>
> ```
> http://www.runoob.com/index.html?language=cn#j2se
> ```
>
> URL 解析：
>
> - **协议**：http
> - **主机为(host:port)**：www.runoob.com
> - **端口号为(port):** 80 ，以上URL实例并未指定端口，因为 HTTP 协议默认的端口号为 80。
> - **文件路径为(path)：**/index.html
> - **请求参数(query)**：language=cn
> - **定位位置(fragment)：**j2se，定位到网页中 id 属性为 j2se 的 HTML 元素位置 。



### java.net.URL类

可以通过构造函数创建URL，可以通过提供的get方法解析url的组成部分。

```java
import java.net.*;
import java.io.*;
 
public class URLDemo
{
   public static void main(String [] args)
   {
      try
      {
         URL url = new URL("http://www.runoob.com/index.html?language=cn#j2se");
         System.out.println("URL 为：" + url.toString());
         System.out.println("协议为：" + url.getProtocol());
         System.out.println("验证信息：" + url.getAuthority());
         System.out.println("文件名及请求参数：" + url.getFile());
         System.out.println("主机名：" + url.getHost());
         System.out.println("端口：" + url.getPort());
         System.out.println("默认端口：" + url.getDefaultPort());
         System.out.println("路径：" + url.getPath());
         System.out.println("请求参数：" + url.getQuery());
         System.out.println("定位位置：" + url.getRef());
      }catch(IOException e)
      {
         e.printStackTrace();
      }
   }
}
输出：
URL 为：http://www.runoob.com/index.html?language=cn#j2se
协议为：http
验证信息：www.runoob.com
文件名及请求参数：/index.html?language=cn
主机名：www.runoob.com
端口：-1
默认端口：80
路径：/index.html
请求参数：language=cn
定位位置：j2se
```



### java.net.URLConnections 类

openConnection() 返回一个 java.net.URLConnection。

例如：

- 如果你连接HTTP协议的URL，openConnection() 方法返回 HttpURLConnection 对象。
- 如果你连接的URL为一个 JAR 文件， openConnection() 方法将返回 JarURLConnection 对象。
- 等等...

```java
public class UrlConnTest {
	public static void main(String[] args) {
		try {
			URL url = new URL("http://www.baidu.com");
			URLConnection urlConnection = url.openConnection();
			HttpURLConnection connection = null;
			if (urlConnection instanceof HttpURLConnection) {
				connection = (HttpURLConnection) urlConnection;
			} else {
				System.err.println("请输入 HTTP URL 地址");
				return;
			}
			BufferedReader in = new BufferedReader(new InputStreamReader(connection.getInputStream()));
			String urlString = "";
			String current;
			while ((current = in.readLine()) != null) {
				urlString += current + "\n";
			}
			System.out.println(urlString);
		} catch (IOException e) {
			e.printStackTrace();
		}
	}
}
```



## Socket 编程

### 概念基础

#### 1 同步和异步：

同步和异步是针对应用程序和内核的交互而言的:

**同步**是指用户进程触发 IO 操作并等待或者轮询的去查看IO 操作是否就绪；

**异步**是指用户进程触发 IO 操作以后便开始做自己的事情，而当 IO 操作已经完成的时候会得到 IO 完成的通知。

> 以银行取款为例：
>
> **同步** ： 自己亲自出马持银行卡到银行取钱（使用同步 IO 时，Java 自己处理IO 读写）；
>
> **异步** ： 委托一小弟拿银行卡到银行取钱，然后给你（使用异步IO 时，Java 将 IO 读写委托给OS 处理，需要将数据缓冲区地址和大小传给OS(银行卡和密码)，OS 需要支持异步IO操作API）；

#### 2 阻塞和非阻塞：

阻塞和非阻塞是针对于进程在访问数据的时候，根据IO操作的就绪状态来采取的不同方式，说白了是一种读取或者写入操作方法的实现方式。

阻塞方式下，读取或者写入函数将一直等待；

非阻塞方式下，读取或者写入方法会立即返回一个状态值。

>以银行取款为例：
>
>**阻塞** ： ATM排队取款，你只能等待（使用阻塞IO时，Java调用会一直阻塞到读写完成才返回）；
>
>**非阻塞** ： 柜台取款，取个号，然后坐在椅子上做其它事，等号广播会通知你办理，没到号你就不能去，你可以不断问大堂经理排到了没有，大堂经理如果说还没到你就不能去（使用非阻塞IO时，如果不能读写Java调用会马上返回，当IO事件分发器通知可读写时再继续进行读写，不断循环直到读写完成）



### 1 BIO(Blocking IO) 同步阻塞编程

BIO编程方式通常是在JDK1.4版本之前常用的编程方式。编程实现过程为：首先在服务端启动一个ServerSocket来监听网络请求，客户端启动Socket发起网络请求，默认情况下ServerSocket回建立一个线程来处理此请求，如果服务端没有线程可用，客户端则会阻塞等待或遭到拒绝。

且建立好的连接，在通讯过程中，是同步的。在并发处理效率上比较低。大致结构如下：

同步并阻塞，服务器实现模式为一个连接一个线程，即客户端有连接请求时服务器端就需要启动一个线程进行处理，如果这个连接不做任何事情会造成不必要的线程开销，当然可以通过线程池机制改善。

BIO方式适用于连接数目比较小且固定的架构，这种方式对服务器资源要求比较高，并发局限于应用中，JDK1.4以前的唯一选择，但程序直观简单易理解。

使用线程池机制改善后的BIO模型图如下:



### 2 NIO( New IO)：同步非阻塞编程

NIO本身是基于事件驱动思想来完成的，其主要想解决的是BIO的大并发问题，NIO基于Reactor，当socket有流可读或可写入socket时，操作系统会相应的通知引用程序进行处理，应用再将流读取到缓冲区或写入操作系统。也就是说，这个时候，已经不是一个连接就要对应一个处理线程了，而是有效的请求，对应一个线程，当连接没有数据时，是没有工作线程来处理的。

NIO的最重要的地方是当一个连接创建后，不需要对应一个线程，这个连接会被注册到多路复用器上面，所以所有的连接只需要一个线程就可以搞定，当这个线程中的多路复用器进行轮询的时候，发现连接上有请求的话，才开启一个线程进行处理，也就是一个请求一个线程模式。

在NIO的处理方式中，当一个请求来的话，开启线程进行处理，可能会等待后端应用的资源(JDBC连接等)，其实这个线程就被阻塞了，当并发上来的话，还是会有BIO一样的问题



### 3 AIO(Asynchronous IO) 异步非阻塞编程

与NIO不同，当进行读写操作时，只须直接调用API的read或write方法即可。这两种方法均为异步的，对于读操作而言，当有流可读取时，操作系统会将可读的流传入read方法的缓冲区，并通知应用程序；对于写操作而言，当操作系统将write方法传递的流写入完毕时，操作系统主动通知应用程序。即可以理解为，read/write方法都是异步的，完成后会主动调用回调函数。

在 JDK1.7 中，这部分内容被称作NIO.2，主要在java.nio.channels包下增加了下面四个异步通道：

- AsynchronousSocketChannel

- AsynchronousServerSocketChannel

- AsynchronousFileChannel

- AsynchronousDatagramChannel