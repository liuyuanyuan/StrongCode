# Java 复制

[TOC]

将一个对象的引用复制给另外一个对象，有3种方式：1是直接赋值，2是浅拷贝，3是深拷贝。这3种实际上都是为了拷贝对象。

### 直接赋值复制（复制对象的引用）

**直接赋值实际上是复制的引用**。

在 Java 中如：`Person student = lily;`，也就是说 student 和 lily 指向的是同一个对象。因此，当 lily 变化的时候，student 里面的成员变量也会跟着变化。



### 浅复制(复制对象本身的属性，及引用的对象的引用）

浅复制：仅复制对象本身的属性值，对包含的引用对象只复制引用。



### 深复制(复制对象的属性，及其引用对象的属性)

深复制：不仅复制对象本身的属性值，对包含的引用对象也复制所有的属性值。

要复制对象的属性值，而不是复制引用，有以下方法：

##### 1 重写java.lang.Object类中的方法clone() （实现 Cloneable接口，但仅用于标记该类是可以复制的）；

##### 2 通过**序列化实现对象的复制**。

 Java 语言里深复制一个对象可以通过序列化实现：先使对象实现 Serializable 接口，然后把对象(实际上只是对象的所有数据)写到一个流里，再从流里读出来，便可以重建对象（新对象和原对象拥有相同的属性数据，但不指向同一个对象）。

##### 3 将A对象的值分别通过set赋值到B对象中；

##### 4 通过 org.apache.commons 中的工具类 BeanUtils 和 PropertyUtils 进行对象复制；

```
//重写clone实现深复制
public class CloneTest {
	public static void main(String[] args) {
		User lily = new User(1, "lily", new Car(1, "smart"));
		User lily2 = new User();
		lily2.setId(lily.getId());
		lily2.setName(lily.getName());
		System.out.println("lily2=" + lily2.toString());
		lily.setId(2);
		System.out.println("lily2=" + lily2.toString());
		
		User lily3 = lily.clone();
		System.out.println("lily3=" + lily3.toString());
		lily.setName("lily jone");
		lily.getCar().setBrand("lamb");
		System.out.println("lily3=" + lily3.toString());	
	}
}

public class User implements Cloneable{
	int id;
	String name;
	Car car;

	public User() {	
	}
	public User(int id, String name, Car car) {
		this.id = id;
		this.name = name;
		this.car = car;
	}
	
	@Override
	public String toString() {
		return "person:" + id + " | " + name + "|" + (car == null ? "null car" : car.toString());
	}
	
	@Override  
	public User clone() {
		User u = null;
		try {
			u = (User) super.clone();
		  u.car = car.clone();//有这句是深复制，没这句是浅复制
		} catch (CloneNotSupportedException e) {
			System.err.println(e.getMessage());
			e.printStackTrace();
		}
		return u;
	}

	public int getId() {
		return id;
	}
	public void setId(int id) {
		this.id = id;
	}
	public String getName() {
		return name;
	}
	public void setName(String name) {
		this.name = name;
	}
	public Car getCar() {
		return car;
	}
	public void setCar(Car car) {
		this.car = car;
	}
}

public class Car implements Cloneable{
	private int id;
	private String brand;
	public Car() {	
	}
	public Car(int id, String brand) {
		this.id = id;
		this.brand = brand;
	}
	
	@Override  
	public Car clone() {
		Car c = null;
		try {
			c = (Car) super.clone();
		} catch (CloneNotSupportedException e) {
			System.err.println(e.getMessage());
			e.printStackTrace();
		}
		return c;
	}

	@Override
	public String toString() {
		return "car:" + id + " | " + brand;
	}

	public int getId() {
		return id;
	}
	public void setId(int id) {
		this.id = id;
	}
	public String getBrand() {
		return brand;
	}
	public void setBrand(String brand) {
		this.brand = brand;
	}
}

```

