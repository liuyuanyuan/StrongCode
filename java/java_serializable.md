# Java 对象的序列化和反序列化

[TOC]

### 序列化与反序列化的概念

Java 提供了对象序列化的机制：一个对象可以转换为一个字节序列（包括该对象的类型、数据及数据的类型等信息）；反序列化的机制：就是将字节序列（对象的类型、数据及数据类型等信息）恢复为对象的过程。

整个过程都是 JVM独立的，也就是说，在一个平台上序列化的对象可以在另一个完全不同的平台上反序列化该对象。

### 对象序列化的2种用途：

1 把对象的字节序列永久地保存到硬盘上，通常存放在一个文件中；
2 在网络上传送对象的字节序列。



### 类的对象要实现序列化的2种方式：

1 实现 Serializable 接口；实现Serializable接口的类，可以采用默认的序列化方式 。

2 实现 Externalizable 接口。实现Externalizable接口的类完全由自身来控制序列化的行为；

   Externalizable接口继承自 Serializable接口；

### 对象序列化的步骤：

1 创建一个对象输出流，它可以包装一个其他类型的目标输出流，如文件输出流；
2 通过对象输出流的writeObject()方法写对象。

ObjectInputStream  类包含序列化对象的方法：

```
//该方法序列化一个对象，并将它发送到输出流。
public final void writeObject(Object x) throws IOException
```

相似的， ObjectInputStream 类包含反序列化一个对象的方法：

```
//该方法从流中取出下一个对象，并将对象反序列化。它的返回值为Object，因此，你需要将它转换成合适的数据类型。
public final Object readObject() throws IOException,                                 ClassNotFoundException
```



### 使用Serializable实现序列化

#### serialVersionUID的取值：

当serialVersionUID显式的定义时，其值是明确不变的；

当serialVersionUID未显式的定义时，其值是变化的不明确的。serialVersionUID的取值是Java运行时环境根据类的内部细节自动生成的。如果对类的源代码作了修改，再重新编译，新生成的类文件的serialVersionUID的取值有可能也会发生变化。类的serialVersionUID的默认值完全依赖于Java编译器的实现，对于同一个类，用不同的Java编译器编译，有可能会导致不同的 serialVersionUID，也有可能相同。

**为了提高serialVersionUID的独立性和确定性，强烈建议在一个可序列化类中显示的定义serialVersionUID，为它赋予明确的值**。

#### 显式地定义serialVersionUID的2种用途：

1 在某些场合，希望类的不同版本对序列化兼容，因此需要确保类的不同版本具有相同的serialVersionUID；
2 在某些场合，不希望类的不同版本对序列化兼容，因此需要确保类的不同版本具有不同的serialVersionUID。

```
public class TestSerial {
	public static void main(String[] args) throws FileNotFoundException, IOException, ClassNotFoundException {
		SerialPerson sp = new SerialPerson("Yolanda", 26);
		serializeObj(sp);
		SerialPerson desp = deserializeObj();
		System.out.println(desp.toString());
	}

	private static void serializeObj(SerialPerson customer) throws FileNotFoundException, IOException {
		// 对象输出流
		ObjectOutputStream oos = new ObjectOutputStream(new FileOutputStream(new File("SerialPerson.txt")));
		oos.writeObject(customer);
		System.out.println("对象序列化成功！");
		oos.close();
	}
	private static SerialPerson deserializeObj() throws FileNotFoundException, IOException, ClassNotFoundException {
		ObjectInputStream ois = new ObjectInputStream(new FileInputStream(new File("SerialPerson.txt")));
		SerialPerson customer = (SerialPerson) ois.readObject();
		System.out.println("对象反序列化成功！");
		return customer;
	}
}

public class TestSerial {
	public static void main(String[] args) throws FileNotFoundException, IOException, ClassNotFoundException {
		SerialPerson sp = new SerialPerson("Yolanda", 26);
		SerializeObj(sp);
		SerialPerson desp = DeserializeObj();
		System.out.println(desp.toString());
	}
	private static void SerializeObj(SerialPerson customer) throws FileNotFoundException, IOException {
		// 对象输出流
		ObjectOutputStream oos = new ObjectOutputStream(new FileOutputStream(new File("SerialPerson.txt")));
		oos.writeObject(customer);
		System.out.println("对象序列化成功！");
		oos.close();
	}
	private static SerialPerson DeserializeObj() throws FileNotFoundException, IOException, ClassNotFoundException {
		ObjectInputStream ois = new ObjectInputStream(new FileInputStream(new File("SerialPerson.txt")));
		SerialPerson customer = (SerialPerson) ois.readObject();
		System.out.println("对象反序列化成功！");
		return customer;
	}
}

public class SerialPerson implements Serializable {
  private static final long serialVersionUID = -8726051873003188533L;
	//private static final long serialVersionUID = 1L;	
	private String name;
	private int age;
	public SerialPerson(String name, int age)
	{
		this.name = name;
		this.age = age;
	}
	public String getName() {
		return name;
	}
	public void setName(String name) {
		this.name = name;
	}
	public int getAge() {
		return age;
	}
	public void setAge(int age) {
		this.age = age;
	}
	@Override
	public String toString() {
		return "name=" + name + ", age=" + age;
	}
}
```

