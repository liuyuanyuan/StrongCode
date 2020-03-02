# Java IO &  NIO

[TOC]

**参考**： [Java SE 12 Doc - Base -  IO/NIO ](https://docs.oracle.com/en/java/javase/12/docs/api/java.base/java/io/package-summary.html)

## I/O模型

### 预备知识：同步/异步、阻塞/非阻塞

**同步和异步**

同步：如果有多个任务或者事件发生，这些任务或者事件必须逐个地进行，一个事件或者任务的执行会导致整个流程的暂时等待，这些事件没有办法并发地执行；（串行）

异步：如果有多个任务或者事件发生，这些事件可以并发地执行，一个事件或者任务的执行不会导致整个流程的暂时等待。（并发）

举例说明：假如有一个任务包括两个子任务A和B，对于同步，当A在执行的过程中，B只有等待，直至A执行完毕，B才能执行；而对于异步，就是A和B可以并发地执行，B不必等待A执行完毕之后再执行，这样就不会由于A的执行导致整个任务的暂时等待。

**阻塞与非阻塞**

阻塞就是：当某个事件或者任务在执行过程中，它发出一个请求操作，但是由于该请求操作需要的条件不满足，那么就会一直在那等待，直至条件满足；

非阻塞就是：当某个事件或者任务在执行过程中，它发出一个请求操作，如果该请求操作需要的条件不满足，会立即返回一个标志信息告知条件不满足，不会一直在那等待。

举例说明：假如我要读取一个文件中的内容，如果此时文件中没有内容可读，对于同步来说就是会一直在那等待，直至文件中有内容可读；而对于非阻塞来说，就会直接返回一个标志信息告知文件中暂时无内容可读。



### 1 阻塞IO模型（Blocking-IO）

是最传统的一种IO模型，即在读、写数据时发生阻塞现象。

当用户线程发出IO请求之后，内核会去查看数据是否就绪，如果没有就绪就会等待数据就绪，而用户线程就会处于阻塞状态，用户线程交出 CPU。当数据就绪之后，内核会将数据拷贝到用户线程，并返回结果给用户线程，用户线程才解除 block 状态。

典型的阻塞 IO 模型的例子为:

```java
data = socket.read(); //一直阻塞在 read 方法，直到返回请求的数据。
```

### 2 非阻塞模型（NoneBlocking-IO ）

当用户线程发起一个 read 操作后，并不需要等待，而是马上就得到了一个结果。如果结果是一个 error 时，它就知道数据还没有准备好，于是它可以再次发送 read 操作。一旦内核中的数据准备好了，并且又再次收到了用户线程的请求，那么它马上就将数据拷贝到了用户线程，然后返回。 所以事实上，在非阻塞 IO 模型中，用户线程需要不断地询问内核数据是否就绪，也就说非阻塞 IO 不会交出 CPU，而会一直占用 CPU。

典型的非阻塞 IO 模型一般如下:

```java
while(true){
	data = socket.read(); // 返回请求的数据或者error标记信息
	if(data!= error){ 
		//处理数据
		break;
	}
}
```

注意：对于非阻塞 IO 有一个严重的问题，在 while 循环中需要不断地询问内核数据是否就绪，这样会导致 CPU 占用率非常高，因此一般很少使用 while 循环这种方式来读取数据。

>**一个完整的IO读请求操作包括两个阶段：**
>
>- 1查看数据是否就绪
>- 2 进行数据拷贝（内核将数据拷贝到用户线程）
>
>那么阻塞（blocking IO）和非阻塞（non-blocking IO）的区别就在于第一个阶段，如果数据没有就绪，在查看数据是否就绪的过程中是一直等待，还是直接返回一个标志信息。

### 3 多路复用IO模型

是目前使用比较多的模型。Java NIO 实际上就是多路复用 IO。

在多路复用 IO 模型中，会有一个线程不断去轮询多个 socket 的状态，只有当 socket 真正有读写事件时，才真正调用实际的 IO 读写操作。因为在多路复用 IO 模型中，只需要使用一个线程就可以管理多个 socket，系统不需要建立新的进程或者线程，也不必维护这些线程和进程，并且仅在真正有 socket 读写事件进行时，才会使用 IO 资源，所以它大大减少了资源占用。在 Java NIO 中，是通 过selector.select()去查询每个通道是否有到达事件，如果没有事件，则一直阻塞在那里，因此这 种方式会导致用户线程的阻塞。多路复用 IO 模式，通过一个线程就可以管理多个 socket，只有当 socket 真正有读写事件发生才会占用资源来进行实际的读写操作。因此，多路复用 IO 比较适合连 接数比较多的情况。	

> 多路复用IO 比 非阻塞IO 的效率高的原因：
>
> 在非阻塞IO 中，不断询问 socket 状态是通过用户线程去进行的；而在多路复用 IO 中，轮询每个 socket 状态是在内核进行的，这个效率要比用户线程要高的多。 								

注意：多路复用 IO 模型是通过轮询的方式来检测是否有事件到达，并且对到达的事件逐一进行响应。因此对于多路复用 IO 模型来说，一旦事件响应体很大，那么就会导致后续的事件 迟迟得不到处理，并且会影响新的事件轮询。

### 4 信号驱动IO模型（SIG-IO）

在信号驱动 IO 模型中，当用户线程发起一个 IO 请求操作，会给对应的 socket 注册一个信号函数，然后用户线程会继续执行，当内核数据就绪时会发送一个信号给用户线程，用户线程接收到信号之后，便在信号函数中调用 IO 读写操作来进行实际的 IO 请求操作。

### 5 异步IO模型（Async-IO）

**异步 IO 模型才是最理想的 IO 模型。**

在异步 IO 模型中，当用户线程发起 read 操作之后，立刻就可以开始去做其它的事。而另一方面，从内核的角度，当它收到一个 asynchronous read 之后， 它会立刻返回，说明 read 请求已经成功发起了，因此不会对用户线程产生任何 block。然后，内核会等待数据准备完成，然后将数据拷贝到用户线程，当这一切都完成之后，内核会给用户线程 发送一个信号，告诉它 read 操作完成了。也就说用户线程完全不需要知道实际的整个 IO 操作是如何进行的，只需要先发起一个请求，当接收内核返回的成功信号时表示 IO 操作已经完成，可以直接去使用数据了。

也就说在异步 IO 模型中，IO 操作的两个阶段都不会阻塞用户线程，这两个阶段都是由内核自动完成，然后发送一个信号告知用户线程操作已完成。用户线程中不需要再次调用 IO 函数进行具体的读写。这点是和信号驱动模型有所不同的，在信号驱动模型中，当用户线程接收到信号表示数据已经就绪，然后需要用户线程调用 IO 函数进行实际的读写操作；而在异步 IO 模型中，收到信号表示 IO 操作已经完成，不需要再在用户线程中调用 IO 函数进行实际的读写操作。

注意：异步 IO 是需要操作系统的底层支持，在 Java 7 中，提供了 Asynchronous IO。参考：http://www.importnew.com/19816.html 。



## Java IO

IO 基于字节流和字符流进行操作。

<img src="images/io-class.png" style="zoom:67%;" />



## Java New IO

NIO的三个核心部分：**通道 Channel、缓冲区 Buffer、选择区 Selector**。

NIO 基于Channel 和 Buffer 进行操作，数据总是从Channel读取到Buffer中，或者从Buffer写入到Channel中去。Selector用于监听多个通道的事件(比如：连接打开，数据到达)，因此，单个线程可以监听多个数据通道。

> NIO 与 IO的最大区别是：IO是面向Stream（byte-stream, character-stream），NIO是面向Buffer的。

java nio的核心抽象：

- [*Buffers*](https://docs.oracle.com/en/java/javase/12/docs/api/java.base/java/nio/package-summary.html#buffers), which are containers for data, and provides an overview of the other NIO packages;
- [*Charsets*](https://docs.oracle.com/en/java/javase/12/docs/api/java.base/java/nio/charset/package-summary.html) and their associated *decoders* and *encoders*, which translate between bytes and Unicode characters;
- [*Channels*](https://docs.oracle.com/en/java/javase/12/docs/api/java.base/java/nio/channels/package-summary.html) of various types, which represent connections to entities capable of performing I/O operations; and
- *Selectors* and *selection keys*, which together with *selectable channels* define a [multiplexed, non-blocking I/O](https://docs.oracle.com/en/java/javase/12/docs/api/java.base/java/nio/channels/package-summary.html#multiplex) facility.

<img src="images/nio_class.png" style="zoom: 67%;" />

### [Channel](https://docs.oracle.com/en/java/javase/12/docs/api/java.base/java/nio/channels/Channel.html)

Channel 常翻译为“通道”，和 IO 中的 Stream差不多一个等级。但是，Stream 是单向的，如InputStream/OutputStream 和 Reader/Writer 分别用来读和写；而 Channel 是双向的，既可以用来进行读操作，又可以用来进行写操作。

NIO 中 Channel 是抽象类，其实现类有:

AbstractInterruptibleChannel, AbstractSelectableChannel, AsynchronousFileChannel, AsynchronousServerSocketChannel, AsynchronousSocketChannel, **DatagramChannel**, **FileChannel**, Pipe.SinkChannel, Pipe.SourceChannel, SctpChannel, SctpMultiChannel, SctpServerChannel, SelectableChannel, **ServerSocketChannel, SocketChannel**

- **FileChannel**：面向文件IO
- **DatagramChannel**：面向UDP
- **ServerSocketChannel**：面向TCP Server
- **SocketChannel**：面向TCP Client



### [Buffer](https://docs.oracle.com/en/java/javase/12/docs/api/java.base/java/nio/Buffer.html)

Buffer 字面意为缓冲区，实际上是一个容器，是一个连续数组。Channel 提供从文件、 网络读取数据的渠道，但是读取或写入的数据都必须经由 Buffer。

NIO中Buffer是顶层父类、抽象类，Buffer 及其子类如下所示：

| Buffer 及其子类                                              | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [`Buffer`](https://docs.oracle.com/en/java/javase/12/docs/api/java.base/java/nio/Buffer.html) | Position, limit, and capacity; clear, flip, rewind, and mark/reset |
| [`ByteBuffer`](https://docs.oracle.com/en/java/javase/12/docs/api/java.base/java/nio/ByteBuffer.html) | Get/put, compact, views; allocate, wrap                      |
| [`MappedByteBuffer`](https://docs.oracle.com/en/java/javase/12/docs/api/java.base/java/nio/MappedByteBuffer.html) | A byte buffer mapped to a file                               |
| [`CharBuffer`](https://docs.oracle.com/en/java/javase/12/docs/api/java.base/java/nio/CharBuffer.html) | Get/put, compact; allocate, wrap                             |
| [`DoubleBuffer`](https://docs.oracle.com/en/java/javase/12/docs/api/java.base/java/nio/DoubleBuffer.html) | Get/put, compact; allocate, wrap                             |
| [`FloatBuffer`](https://docs.oracle.com/en/java/javase/12/docs/api/java.base/java/nio/FloatBuffer.html) | Get/put, compact; allocate, wrap                             |
| [`IntBuffer`](https://docs.oracle.com/en/java/javase/12/docs/api/java.base/java/nio/IntBuffer.html) | Get/put, compact; allocate, wrap                             |
| [`LongBuffer`](https://docs.oracle.com/en/java/javase/12/docs/api/java.base/java/nio/LongBuffer.html) | Get/put, compact; allocate, wrap                             |
| [`ShortBuffer`](https://docs.oracle.com/en/java/javase/12/docs/api/java.base/java/nio/ShortBuffer.html) | Get/put, compact; allocate, wrap                             |
| [`ByteOrder`](https://docs.oracle.com/en/java/javase/12/docs/api/java.base/java/nio/ByteOrder.html) | Typesafe enumeration for byte orders                         |

<img src="images/nio-channel.png" alt="image-20200228123035691" style="zoom:50%;" />



### [Selector](https://docs.oracle.com/en/java/javase/12/docs/api/java.base/java/nio/channels/Selector.html)

Selector 选择区，Selector 能够检测多个注册的通道上是否有事件发生，如果有事件发生，便获取事件然后针对每个事件进行相应的响应处理。这样，只是用一个单线程就可以管理多个通道，也就是管理多个连接。这样使得只有在连接真正有读写事件发生时，才会调用函数来进行读写，就大大地减少了系统开销，并且不必为每个连接都创建一个线程，不用去维护多个线程，并且避免了多线程之间的上下文切换导致的开销。

 <img src="images/nio-selector.png" alt="image-20200228123212999" style="zoom:50%;" />

所有的Channel都归Selector管理，这些channel中只要有至少一个有IO动作，就可以通过Selector.select方法检测到，并且使用selectedKeys得到这些有IO的channel，然后对它们调用相应的IO操作。







