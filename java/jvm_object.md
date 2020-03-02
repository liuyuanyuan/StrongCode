# Java对象的内存布局

在HotSpot VM中，对象在堆内存中的存储布局可以划分为三个部分：

1 对象头(Header）：占12bytes=96bits

​	mark world（8bytes=64bit）：存储的内容不固定

​    klass pointer（4bytes=32bit） ：

2 实例数据（Instance Data）

3 对齐数据（Padding）



## openjdk-HotSpot doc

http://openjdk.java.net/groups/hotspot/docs/HotSpotGlossary.html

**object header**

Common structure at the beginning of every GC-managed heap object. 

(Every oop points to an object header.) Includes fundamental 

information about the heap object's layout, type, GC state, 

synchronization state, and identity hash code. Consists of two 

words. In arrays it is immediately followed by a length field.

 Note that both Java objects and VM-internal objects have a 

 common object header format.

**mark word**

The first word of every object header. Usually a set of bitfields including synchronization state and identity hash code. May also be a pointer (with characteristic low bit encoding) to synchronization related information. During GC, may contain GC state b 

**klass pointer**

The second word of every object header. Points to another object (a metaobject) which describes the layout and behavior of the original object. For Java objects, the "klass" contains a C++ style "vtable".



## 编码验证

64bit - openjdk12 - JVM中：

```java
public class TestObject {
	boolean flag = false;
	
	public static void main(String[] args) {
		TestObject obj = new TestObject();
		System.out.println(ClassLayout.parseInstance(obj).toPrintable());
	}
}

输出：
thread.TestObject object internals:
 OFFSET SIZE   TYPE DESCRIPTION                VALUE
   0   4      (object header)              05 00 00 00 (00000101 00000000 00000000 00000000) (5)
   4   4      (object header)              00 00 00 00 (00000000 00000000 00000000 00000000) (0)
   8   4      (object header)              08 bc 22 00 (00001000 10111100 00100010 00000000) (2276360)
   12   1  boolean TestObject.flag              false
   13   3      (loss due to the next object alignment)

Instance size: 16 bytes
Space losses: 0 bytes internal + 3 bytes external = 3 bytes total
```



```java
public class TestObject {
	boolean flag = false;
	
	public static void main(String[] args) {
		TestObject obj = new TestObject();
		System.out.println(Integer.toHexString(obj.hashCode()));
		countHash(obj);
		System.out.println(ClassLayout.parseInstance(obj).toPrintable());
	}
  public static void countHash(Object obj) throws NoSuchFieldException, SecurityException, IllegalArgumentException, IllegalAccessException {
		/*手动计算hashcode*/
		Field field = Unsafe.class.getDeclaredField("theUnsafe");
		field.setAccessible(true);
		Unsafe unsafe = (Unsafe)field.get(null);
		long hashCode = 0;
		for(long index = 7; index > 0; index--) {
			hashCode |= (unsafe.getByte(obj, index) & 0xFF) << ((index-1)*8);
		}	
		String code = Long.toHexString(hashCode);
		System.out.println("countHash: 0x" + code);
	}
}

输出：
3764951d
countHash: 0x3764951d
# WARNING: Unable to get Instrumentation. Dynamic Attach failed. You may add this JAR as -javaagent manually, or supply -Djdk.attach.allowAttachSelf
# WARNING: Unable to attach Serviceability Agent. You can try again with escalated privileges. Two options: a) use -Djol.tryWithSudo=true to try with sudo; b) echo 0 | sudo tee /proc/sys/kernel/yama/ptrace_scope
thread.TestObject object internals:
 OFFSET  SIZE      TYPE DESCRIPTION                               VALUE
      0     4           (object header)                           01 1d 95 64 (00000001 00011101 10010101 01100100) (1687493889)
      4     4           (object header)                           37 00 00 00 (00110111 00000000 00000000 00000000) (55)
      8     4           (object header)                           08 bc 22 00 (00001000 10111100 00100010 00000000) (2276360)
     12     1   boolean TestObject.flag                           false
     13     3           (loss due to the next object alignment)
Instance size: 16 bytes
Space losses: 0 bytes internal + 3 bytes external = 3 bytes total
```

