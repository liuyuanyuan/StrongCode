# 类加载机制

[TOC]

### 类的生命周期

（ class> ）**加载 > 连接 [验证 > 准备 > 解析] > 初始化 > 使用 >  卸载**



加载： 将类的.class文件从磁盘读到内存

连接：

​		验证：验证字节码文件的正确性；

​		准备：给类的静态变量分配内存，并赋予默认值（此处是虚拟机默认的初始值）；

​		解析：类装载器装入类所引用的其他所有类；

初始化：为类的静态变量赋予正确的初始值（此处是程序员为变量分配的真正的初始值），执行静态代码块；

使用：使用

卸载：销毁



### 类加载器种类

- Bootstrap  ClassLoader  启动类加载器

​	负责加载jre核心类库，如jre目录下的rt.jar, charsets.jar等;

- Extension ClassLoader 扩展类加载器（ExtClassLoader）

​	负责加载jre扩展目录ext中的jar类包；

- Application ClassLoader 系统类加载器（AppClassLoader ）

​    负责加载ClassPath路径下的类包

-  User  ClassLoader  用户自定义类加载器（ **extends** ClassLoader）

​    负责加载用户自定义路径下的类包

实战验证：

```java
//Java中用c/c++实现的类，其类加载器打印为null）
public class Test {
	public static void main(String[] args) {
		Map map = new HashMap();
		System.out.println(Test.class.getClassLoader());
		System.out.println(map.getClass().getClassLoader());
		}
}

output:
jdk.internal.loader.ClassLoaders$AppClassLoader@10f87f48
null
null
```



### 类加载机制

- 全盘负责委托机制

  最简单的，当前类中引用的类，依然使用当前类的加载器；

- 双亲委派机制

  复杂的，

  优势：避免类的重复加载，防止系统类的篡改；

