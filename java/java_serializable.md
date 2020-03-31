# Java 对象的序列化和反序列化

[TOC]

### 序列化与反序列化的概念

- 序列化：把内存中的对象 变成可传输（如网络传输）、可存储（如存储到硬盘）的字节序列的过程；
- 反序列化：把字节序列 恢复到 内存中的对象的过程；

字节序列：包括该对象的类型信息、数据及数据的类型信息。

整个过程都是 JVM独立的，也就是说，在一个平台上序列化的对象可以在另一个完全不同的平台上反序列化该对象。



### 对象序列化的2种用途：

- 把对象的字节序列永久地保存到硬盘上，通常存放在一个文件中；（用于磁盘File IO）
- 在网络上传送对象的字节序列；（用于网络IO）



### 序列化的条件

类的对象序列化需要满足两个条件

- 该类必须实现 java.io.Serializable 对象。
- 该类的所有属性必须是可序列化的。如果有一个属性不是可序列化的，则该属性必须注明是短暂的（通过 **transient** 关键字声明）。

想知道一个 Java 标准类是否是可序列化的，请查看该类的文档。检验一个类的实例是否能序列化十分简单， 只需要查看该类有没有实现 java.io.Serializable 接口。



### 类的对象要实现序列化的2种方式：

1 实现 Serializable 接口；实现Serializable接口的类，可以采用默认的序列化方式 。

2 实现 Externalizable 接口。实现Externalizable接口的类完全由自身来控制序列化的行为；

   Externalizable接口继承自 Serializable接口；



### 对象序列化的步骤：

1 创建一个对象输出流，它可以包装一个其他类型的目标输出流，如文件输出流；
2 通过对象输出流的writeObject()方法写对象。

ObjectInputStream  类包含序列化对象的方法;

类似的， ObjectInputStream 类包含反序列化一个对象的方法;

```java
public class ObjectSerialTest {
	
	public static void main(String[] args) 
			throws IOException, ClassNotFoundException {
		String filePath = "/Users/liuyuanyuan/user1";
		User user = new User(1, "yuan", 18);
		ObjectSerialTest.writeObject(user, filePath);
		User restoreUser = (User) ObjectSerialTest.readObject(filePath);
	}

	//序列化
	public static final void writeObject(Object obj, String filePath) throws IOException {
		FileOutputStream fos = null;
		ObjectOutputStream oos = null;
		try {		
			fos = new FileOutputStream(filePath);
			oos = new ObjectOutputStream(fos);
			oos.writeObject(obj);
			System.out.println("Serialized data is saved in " + filePath);
		} catch (Exception e) {
			e.printStackTrace();
		} finally {
			if (oos != null) {
				oos.close();
			}
			if (fos != null) {
				fos.close();
			}
		}
	}

	//反序列化
	public static final Object readObject(String filePath) 
			throws IOException, ClassNotFoundException {
		User user = null;
		FileInputStream fis = null;
		ObjectInputStream ois = null;
		try {
			fis = new FileInputStream(filePath);
			ois = new ObjectInputStream(fis);
			user = (User) ois.readObject();
			System.out.println("Deserialize User: ");
			System.out.println("id : " + user.getId());
			System.out.println("name : " + user.getName());
			System.out.println("age : " + user.getAge());
		} catch (Exception e) {
			e.printStackTrace();
		} finally {
			if (ois != null) {
				ois.close();
			}
			if (fis != null) {
				fis.close();
			}
		}
		return user;
	}
}
//序列化的对象
public class User implements Serializable{
	//default serial
	//private static final long serialVersionUID = 1L;
	//generated serial
	private static final long serialVersionUID = -7884044395190826641L;
	
	private int id;
	private String name;
	private int age;
	public User(int id,String name, int age) {
		this.id = id;
		this.name = name ;
		this.age = age;
	}
  //此处getter、setter方法省略...
}
```



### 使用Serializable实现序列化

#### serialVersionUID的取值：

当serialVersionUID显式的定义时，其值是明确不变的；

当serialVersionUID未显式的定义时，其值是变化的不明确的。serialVersionUID的取值是Java运行时环境根据类的内部细节自动生成的。如果对类的源代码作了修改，再重新编译，新生成的类文件的serialVersionUID的取值有可能也会发生变化。类的serialVersionUID的默认值完全依赖于Java编译器的实现，对于同一个类，用不同的Java编译器编译，有可能会导致不同的 serialVersionUID，也有可能相同。

**为了提高serialVersionUID的独立性和确定性，强烈建议在一个可序列化类中显示的定义serialVersionUID，为它赋予明确的值**。

#### 显式地定义serialVersionUID的2种用途：

1 当希望类的不同版本对序列化兼容，因此需要确保类的不同版本，具有相同的serialVersionUID；
2 当不希望类的不同版本对序列化兼容，因此需要确保类的不同版本，具有不同的serialVersionUID。

```javascript
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
  private static final long serialVersionUID = -8726051873003188533L;//generated value
	//private static final long serialVersionUID = 1L;	//default value
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

