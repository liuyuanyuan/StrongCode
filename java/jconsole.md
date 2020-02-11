# JConsole监控分析JVM运行

[Jconsole]()是JDK自带的性能监控分析器，用于JVM中内存、线程、类等监控，并提供大量可视化图表。还可以使用JTop插件。**可以监控本地和远程的JVM，也可以同时监控几个JVM。**

这款工具的好处在于，占用系统资源少，几乎不消耗；而且结合Jstat，可以有效监控到java内存的变动情况，以及引起变动的原因；在Java项目追踪内存泄露问题时，很实用。

可以从命令行（直接输入jconsole）或在 GUI shell （jdk/bin下打开）中运行；它用于连接正在运行的本地或远程的[JVM](http://www.51testing.com/?uid-116228-action-viewspace-itemid-149296)，对运行在Java应用程序的资源消耗和性能进行监控，并画出大量的图表，提供强大的可视化界面。

**示意图**：命令行终端执行jconsole，打开其连接界面：

![image-20200207133732382](/Users/liuyuanyuan/github/StrongCode/java/images/jconsole-connect.png)



