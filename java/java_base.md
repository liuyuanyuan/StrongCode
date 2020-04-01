# Java基础知识点

[TOC]

## 1 ==(引用目标比较) 、equals()(值比较)

- == ：判断两个对象的地址是否相等，即判断两个对象是否是同一个对象(基本数据类型==比较的是值，引用数据类型==比较的是内存地址)。
- equals() :  判断两个对象是否相等(一般是判读属性值)。



## 2 字符串的比较

```
    //对于字符串常量，值拼写相同的共享同一个实例；
		String s1 = "abc";
		String s2 = "abc";
		System.out.println("1--" + (s1==s2));//比较引用 true
		System.out.println("2--" + s1.equals(s2));//比较值 true
		
		//使用构造函数创建的字符串，会生成新的实例
		String s3= new String("abc");
		System.out.println("3--" + (s1==s3));//比较引用 false
		System.out.println("4--" + s1.equals(s3));//比较值 true
```



## 3 final关键字

- 用于变量：对于基本数据类型的变量，其数值初始化后便不能更改；对于引用类型的变量，其初始化后便不能指向另一个对象。

- 用于方法：一是为了锁定方法，防止继承类修改它；二是为了效率（早期版本中会将final方法转为内嵌调用，现在已经不需要final方法来优化）。

  类中所有的private方法都隐式地指定为final。

- 用于类：表示这个类不能被继承，final类中的所有成员方法都被隐式指定为final。

  

## 4 String、StringBuilder、StringBuffer

- String：String 类中使用 final 关键字字符数组保存字符串，`private final char value[]` ，所以 String 对象是不可变的。
- StringBuilder：父类AbstractStringBuilder，线程不安全
- StringBuffer：父类AbstractStringBuilder，线程安全

```java
abstract class AbstractStringBuilder implements Appendable, CharSequence {
    //没有使用final，所以对象是可变的
    char[] value; 
    int count;
    AbstractStringBuilder() {
    }
    AbstractStringBuilder(int capacity) {
        value = new char[capacity];
    }
```

对于三者使用的总结:

1 操作少量的数据用：String

2 单线程，操作字符串缓冲区下操作大量数据用 StringBuilder 

3 多线程，操作字符串缓冲区下操作大量数据用StringBuffer



## 5 Object类

Object类 是所有类的父类。

```java
/**
 * Class {@code Object} is the root of the class hierarchy.
 * Every class has {@code Object} as a superclass. All objects,
 * including arrays, implement the methods of this class.
 *
 * @author  unascribed
 * @see     java.lang.Class
 * @since   1.0
 */
public class Object {
    private static native void registerNatives();
    static {
        registerNatives();
    }
  
    /**
     * Constructs a new object.
     */
    @HotSpotIntrinsicCandidate
    public Object() {}

    @HotSpotIntrinsicCandidate
    public final native Class<?> getClass();

    @HotSpotIntrinsicCandidate
    public native int hashCode();

    public boolean equals(Object obj) {
        return (this == obj);
    }

    @HotSpotIntrinsicCandidate
    protected native Object clone() throws CloneNotSupportedException;
  
    public String toString() {
        return getClass().getName() + "@" + Integer.toHexString(hashCode());
    }
    
    @HotSpotIntrinsicCandidate
    public final native void notify();
    
    public final void wait() throws InterruptedException {
        wait(0L);
    }

    public final native void wait(long timeoutMillis) throws InterruptedException;
    
    public final void wait(long timeoutMillis, int nanos) throws InterruptedException {
        if (timeoutMillis < 0) {
            throw new IllegalArgumentException("timeoutMillis value is negative");
        }
        if (nanos < 0 || nanos > 999999) {
            throw new IllegalArgumentException(
                                "nanosecond timeout value out of range");
        }
        if (nanos > 0 && timeoutMillis < Long.MAX_VALUE) {
            timeoutMillis++;
        }
        wait(timeoutMillis);
    }
  
    @Deprecated(since="9")
    protected void finalize() throws Throwable { }
}
```



## 6 获取键盘输入的两种方法

```java
Scanner input = new Scanner(System.in);
String s = input.nextLine();
System.out.println("receive: " + s);
input.close();

BufferedReader input2 = new BufferedReader(new InputStreamReader(System.in));
String s2 = input2.readLine();
System.out.println("2receive: " + s2);
```







