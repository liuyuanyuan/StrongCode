# 设计模式

[TOC]

**参考：**

- 书籍：Design Pattern  GoF （鼻祖）
- 书籍：Head First Design Patterns （生动形象）

设计结构对系统性能的影响要远高于代码优化。熟悉一些典型的设计模式和方法，有助于设计高性能软件。

## 概念和分类

设计模式（Design pattern）代表了最佳的实践，是软件开发人员在软件开发过程中面临的一般问题的解决方案。

设计模式是一套被反复使用的、多数人知晓的、经过分类编目的、代码设计经验的总结。使用设计模式是为了重用代码、让代码更容易被他人理解、保证代码可靠性。设计模式使代码编制真正工程化，设计模式是软件工程的基石。

根据设计模式的参考书 **Design Patterns - Elements of Reusable Object-Oriented Software（中文译名：设计模式 - 可复用的面向对象软件元素）** 中所提到的，**总共有 23 种设计模式**。这些模式可以分为**3大类**：**创建型模式（Creational Patterns）、结构型模式（Structural Patterns）、行为型模式（Behavioral Patterns）。**

- 创建型模式(5种，绿色)：

工厂方法模式、抽象工厂模式、单例模式、建造者模式、原型模式。

- 结构型模式(7种，橙色)：

适配器模式、装饰器模式、代理模式、外观模式、桥接模式、组合模式、享元模式。

- 行为型模式(11种，粉色)：

策略模式、模板方法模式、观察者模式、迭代子模式、责任链模式、命令模式、备忘录模式、状态模式、访问者模式、中介者模式、解释器模式。


以下是网上经典的设计模式关系图：

<img src="images/design_pattern_relationship.png" alt="img" style="zoom:75%;" />



## 单例模式（Singleton Design Pattern）

单例模式是最简单的设计模式之一。单例模式是一种**创建型模式**。

提供了创建对象实例的最佳方式：确保系统中一个类只产生一个实例。

**意图：**

1 对频繁使用的对象，可以节约创建与销毁对象所花费的时间 ，尤其是对重量级对象；

2 new操作的次数减少，因此对系统内存的使用频率也降低，这将减轻GC压力，缩短GC停顿时间。

**核心：**保证系统中一个类只有一个实例，并提供一个访问它的全局访问点（构造函数必须私有）。

**主要参与者**：单例类和使用者；

**使用场景：** 

1 要求生产唯一序列号；

2 web中的计数器，不必每次刷新都在数据库里加一次，先用单例缓存起来

3 创建的对象需要消耗的资源过多，比如I/O与数据库的连接等。

4 配置文件的读取

```java
//最简单的单实例实现，不能做到延迟加载实例。(饿汉式，可保证线程安全)
public class Singleton {
  //私有化构造方法
  private Singleton(){
      //todo
  }
  private static Singleton instance = new Singleton();
	public static Singleton getInstance() {
		return instance;
	}
}

/*
* 使用内部类来实现单实例，实例是在类加载时完成创建。
* 既做到了延迟加载，又对多线程友好（不必使用synchronized关键字），是一种比较完善的实现。
*/
public class StaticSingleton {
  //私有化构造方法
  private StaticSingleton(){
      //todo
  }
  
  private static class InstanceHolder {
		private static StaticSingleton instance = new StaticSingleton();//在类加载时创建对象实例
	}
	public static StaticSingleton getInstance() {
		return InstanceHolder.instance;
	}
}
```



## 代理模式（Proxy Design Pattern）

**定义：**

使用代理对象完成用户请求，屏蔽用户对真实对象的访问。（如同现实生活中的代理人一样）属于**结构型模式**。

**主要参与者：**

| 角色         | 作用                                                         |
| ------------ | ------------------------------------------------------------ |
| 主体接口     | 定义代理类和真实主题的公共对外方法，也是代理类代理真实主题的方法 |
| 真实主题     | 真正实现业务逻辑的类                                         |
| 代理类       | 用来代理和封装真实主题                                       |
| 客户端(Main) | 使用代理类和主题接口完成一些工作                             |

**应用场景：**

用于远程调用的网络代理、考虑安全因素的安全代理、延迟加载。

代理模式用于延迟加载，可以有效提升系统的启动速度，对改善用户体验有很大帮助。



### 静态代理

代理类在程序运行前就已经存在，这种代理方式被成为**静态代理**。

```java
//主体接口
public interface Shop {
	public abstract void sell();
}
//真实主体
public class FruitShop implements Shop {
	@Override
	public void sell() {
		System.out.println("sell fruit");
	}
}
//代理类
public class ShopProxy implements Shop{
	private Shop shop;
	ShopProxy(Shop shop){
		this.shop = shop;
	}
	@Override
    public void sell() {
		System.out.println("proxy begin sell");
		shop.sell();
		System.out.println("proxy finish sell");
	}
}
//客户端使用代理
public class Client {
	public static void main(String[] args) {
		ShopProxy sp = new ShopProxy(new FruitShop());
		sp.sell();
	}
}
输出：
  proxy begin sell
  sell fruit
  proxy finish sell
```



### 动态代理

代理类在运行时动态生成并加载的，称为动态代理。即：代理类的字节码将在运行时生成并载入当前的 ClassLoader。

**生成动态代理类的方法：**

- JDK自带的动态代理：内置在JDK中，无需引入第三方jar，使用简单；
- CGLib和Javassist：都是高级的字节码生成库，总体性能比JDK自带的动态代理好，且功能十分强大；
- ASM库：是低级的字节码生成工具，使用ASM近乎使用Java bytecode编码，对开发人员要求高，当然也是性能最好。但是ASM程序可维护性差，如非性能要求苛刻，还是建议使用CGLIB和Javassist。

**JDK动态代理、CGLib动态代理的适用场景：**

- JDK 动态代理，使用Java反射技术生成代理类，只能代理实现了接口的类，没有实现接口的类不能实现动态代理；

- CGLib 会在运行时动态的生成一个被代理类的子类，子类重写了被代理类中所有非final的方法，在子类中采用方法拦截的技术拦截所有父类方法的调用，不需要被代理类对象实现接口，从而 CGLIB 动态代理效率比 Jdk 动态代理反射技术效率要高。

#### JDK动态代理

-  java.lang.reflect.Proxy：生成动态代理类和对象；
-  java.lang.reflect.InvocationHandler（处理器接口）：可以通过invoke方法实现对真实目标的代理访问。

```java
//主体接口Shop和真实主体Fruit（与静态代理例子中相同）
public class Client {
	public static void main(String[] args) {
		Shop fs = new FruitShop();
    //静态代理
		//ShopProxy sproxy = new ShopProxy(fs);
		//sproxy.sell();
    
    //生成JDK动态代理
		Shop sproxy = (Shop) Proxy.newProxyInstance(
				fs.getClass().getClassLoader()
				, fs.getClass().getInterfaces()
				, new InvocationHandler() {
					@Override
					public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
						System.out.println("jdk proxy start sell");
						Object invoke = method.invoke(fs, args);
						System.out.println("jdk proxy end sell");
						return invoke;
					}
		});
		sproxy.sell();
	}
}
```



#### CGLib动态代理

目标类不能为 final，目标对象的方法如果为 final / static，那么就不会被拦截，即不会执行目标对象额外的业务方法。

**1 引入 CGLib 依赖**

Spring环境下不需要，因为Spring-Core里已经引入了。以下Maven引入依赖：

```xml
<dependency>
    <groupId>cglib</groupId>
    <artifactId>cglib</artifactId>
    <version>3.2.12</version>
</dependency>
```

**2 创建一个目标类**

```java
public class FruitShop {
	@Override
	public void sell() {
		System.out.println("sell fruit");
	}
}
```

**3 创建 CGLibProxy 代理类**

```java
public class CGLibProxy implements MethodInterceptor{

	private FruitShop fshop;
	CGLibProxy(FruitShop fshop){
		this.fshop = fshop;
	}
	
	@Override
	public Object intercept(Object obj, Method method, Object[] objects, MethodProxy methodProxy) throws Throwable {
		System.out.println("CGLic proxy start sell");
		Object invoke = method.invoke(fshop, objects);
		System.out.println("CGLic proxy end sell");
		return invoke;
	}

	FruitShop proxy() {
		Enhancer enhancer = new Enhancer();
		// 设置代理的目标类
		enhancer.setSuperclass(FruitShop.class);
		// 设置回调方法, this代表是当前类, 因为当前类实现了CallBack
		enhancer.setCallback(this);
		return (FruitShop) enhancer.create();
	}
}

```

**客户端测试类：**

```java
public class Test{
		public static void main(String[] args) {
				FruitShop fshop = new FruitShop();
				FruitShop proxy = new CGLibProxy(fshop).proxy();
        proxy.sell();
    }
}
```



## 工厂模式（Factory Design Pattern）

参考：https://www.jianshu.com/p/d951ac56136e

工厂模式，属于**创建型设计模式**。细分为：

- 简单工厂模式（静态工厂模式）：具体工厂、**抽象产品**、具体产品；
- 工厂方法模式（多态工厂模式）：**抽象工厂**、具体工厂、**抽象产品**、具体产品；
- 抽象工厂模式：**抽象工厂**、具体工厂、**抽象产品族、抽象产品**、具体产品；

### JDK 中的工厂模式实例：

- java.util.Calendar, ResourceBundle 和 NumberFormat 的 getInstance()；
- Boolean、Integer 的 valueOf() 方法。

### 简单工厂模式（静态工厂模式）

**定义：**是由一个工厂对象决定创建出哪一种产品类的实例。实质是由一个工厂类根据传入的参数，动态决定应该创建哪一个产品类（这些产品类继承自一个父类或接口）的实例。

**主要角色**：
**具体工厂：**负责实现创建所有实例的内部逻辑，并提供一个外界调用的方法，创建所需的产品对象。
**抽象产品：**负责描述产品的公共接口
**具体产品：**描述生产的具体产品。

**代码实现：**一个抽象产品类，可以派生出多个具体产品类；一个具体工厂类，通过**往此工厂的 static 方法中传入不同参数，产出不同的具体产品类实例。**

**优点：**将创建使用工作分开，不必关心类对象如何创建，实现了解耦；
**缺点**：违背开闭原则，一旦添加新产品就不得不修改工厂类的逻辑，这样就会造成工厂逻辑过于复杂。

```java
/** 
 * Factory.java
 * 工厂：负责实现创建所有实例的内部逻辑，并提供一个外界调用的方法，创建所需的产品对象。
 */
public class Factory {
    /**
     * 供外界调用的方法（可以看成是对外提供的三种按钮）
     * @param type 产品类型
     * @return 产品实例
     */
    public static Product getProduct(String type) {
        switch (type) {
            case "A":
                return new AProduct();
            case "B":
                return new BProduct();
            default:
                return null;
        }
    }
}

/** Product.java
 *  抽象产品： 描述产品的公共接口
 */
abstract  class Product {
    abstract void intro(); //产品介绍
}

/** AProduct.java
 *  具体产品A（可以看成是一种饮料：可乐）
 */
public class AProduct extends Product{
    @Override
    void intro() {
        System.out.println("可乐");
    }
}

/** BProduct.java
 *  具体产品B（可以看成是一种饮料：奶茶）
 */
public class BProduct extends Product{
    @Override
    void intro() {
        System.out.println("奶茶");
    }
}

//测试
public class Test {
    public static void main(String[] args) {
        //创建具体的工厂
        Factory factory = new Factory();
        //根据传入的参数生产不同的产品实例
        //(按下不同的按钮，获取饮料)
        Product A = factory.getProduct("A");
        A.intro();
        Product B = factory.getProduct("B");
        B.intro();
    }
}
输出：
  可乐
  奶茶
```



### 工厂方法模式（多态工厂模式）

**定义：**工厂方法模式，又称工厂模式、多态工厂模式、虚拟构造器模式。通过**定义工厂父类负责定义创建对象的公共接口，而子类则负责生成具体的对象**。此模式的核心精神是封装类中不变的部分。

**作用：**将类的实例化（具体产品的创建）延迟到工厂类的子类（具体工厂）中完成，即由子类来决定应该实例化（创建）哪一个类。

**主要角色**：
**抽象工厂：**描述工厂的公共接口
**具体工厂：**描述具体工厂，创建产品的实例，供外界调用
**抽象产品：**描述产品的公共接口
**具体产品：**描述生产的具体产品

**实现：**一个抽象产品类，可以派生出多个具体产品类；一个抽象工厂类，可以派生出多个具体工厂类；每个具体工厂类只能创建一个具体产品类的实例。

**优点：**

1符合开-闭原则：新增一种产品时，只需要增加相应的具体产品类和相应的工厂子类即可。

2符合单一职责原则：每个具体工厂类只负责创建对应的产品

**缺点：**

1增加了系统的复杂度：类的个数将成对增加；

2增加了系统的抽象性和理解难度；

3一个具体工厂只能创建一种具体产品；

```java
//抽象产品
public abstract class Drink {
	public abstract void intro();
}
//具体产品：可乐
public class DrinkCola extends Drink {
	@Override
	public void intro() {
		System.out.println("可乐");
	}
}

//抽象工厂
public abstract class Factory {
	//生产产品
    public abstract Drink getDrink();
}
//具体工厂-可乐工厂
public class ColaFactory extends Factory {
	@Override
	public Drink getDrink() {
		return new DrinkCola();
	}
}

//测试
public class Test {
	public static void main(String[] args) {
		ColaFactory cf = new ColaFactory();
		cf.getDrink().intro();
		SpriteFactory sf = new SpriteFactory();
		sf.getDrink().intro();
	}
}
输出：
可乐
雪碧
```



### 抽象工厂模式【重点】

**定义：**提供一个创建一系列相关或相互依赖对象的接口，而无须指定它们具体的类；具体的工厂负责实现具体的产品实例。

**解决的问题：**每个工厂只能创建一类产品（工厂方法模式）

抽象工厂模式与工厂方法模式最大的区别：抽象工厂中每个工厂可以创建多种类的产品；而工厂方法每个工厂只能创建一类

**主要对象**
**抽象工厂**：描述具体工厂的公共接口
**具体工厂**：描述具体工厂，创建产品的实例，供外界调用
**抽象产品族**：描述抽象产品的公共接口
**抽象产品**：描述具体产品的公共接口
**具体产品**：具体产品

**实现：** 多个抽象产品类，每个抽象产品类可以派生出多个具体产品类。一个抽象工厂类，可以派生出多个具体工厂类。 每个具体工厂类可以创建多个具体产品类的实例。

**优点：**降低耦合，符合开-闭原则，符合单一职责原则，不使用静态工厂方法，可以形成基于继承的等级结构。

**缺点：**难以扩展新种类产品

>**例子：**假设有各种食品工厂，主要生产饮料和零食两类食品；A食品工厂生产的饮料是果汁，零食是薯片；B食品工厂生产的饮料是可乐，零食是虾条。

```java

//抽象产品族
public abstract class Product {
	public abstract void intro();
}
//抽象产品：饮料
public abstract class Drink extends Product {
	@Override
	public abstract void intro();
}
//抽象产品：零食
public abstract class Snakes extends Product{
	@Override
	public abstract void intro();
}
//具体产品：果汁
public class DrinkJuice extends Drink {
	@Override
	public void intro() {
		System.out.println("果汁");
	}
}
//具体产品：薯片
public class SnakesChips extends Snakes {
	@Override
	public void intro() {
		System.out.println("薯片");
	}
}

//抽象工厂
public abstract class Factory {
	  //生产饮料、零食
    public abstract Drink getDrink();
    public abstract Snakes getSnakes();
}
//具体工厂
public class AFactory extends Factory {
	@Override
	public Drink getDrink() {
		return new DrinkJuice();//生产果汁
	}
	@Override
	public Snakes getSnakes() {
		return new SnakesChips();//生产薯片
	}
}

//测试
public class Test {
	public static void main(String[] args) {
		AFactory cf = new AFactory();
		cf.getDrink().intro();
    cf.getSnakes().intro();
	}
}
输出：
	果汁
	薯片
```



##  观察者模式（Observer Pattern）

观察者模式（Observer Pattern）用于一个对象的行为依赖另一个对象的状态的情况。比如，当一个对象被修改时，则会自动通知它的依赖对象。它属于**行为型模式**。

**意义：**在单线程中，使得某一对象，及时得知自身所依赖对象的状态变化（要考虑到易用和低耦合，保证高度的协作）。

**主要解决：**一个对象状态改变时给其他对象通知的问题，而且要考虑到易用和低耦合，保证高度的协作。

**使用场景：**事件监听、通知发布等，确保观察者在不使用轮询监控的情况下，及时收到相关消息和事件。

**注意：** 1 JDK已经实现了一套观察者模式，可以直接复用相关代码。 2 避免循环引用。 3 如果顺序执行，某一观察者错误会导致系统卡壳，一般采用异步方式。

**观察者模式最典型的应用便是：Swing框架的JButton实现。**当按钮被单击时，通过AbstractButton的fireActionPerformed()方法，回调ActionListener.actionPerformed()方法实现。应用开发中，只需实现ActionListener接口（Observer），就可将其添加到按钮（Subject）的观察者列表中，当点击事件发生，就自动触发监听器的业务处理函数。



## 生产者消费者模式

**生产者消费者问题描述：**

1 生产者和消费者在同一时间段内共用同一个存储空间(内存缓冲区)；
2 生产者往存储空间中添加产品，消费者从存储空间中取走产品；
3 当存储空间为空时，消费者阻塞；当存储空间满时，生产者阻塞。

**解决该问题的3个方案**

- 1 sync，wait，notify

- 2 lock，await，singal

- **3 阻塞队列实现**(简明利落)

​    参考：https://www.jianshu.com/p/e87c8248d64d





## 值对象模式（Value Object Pattern）

当一组属性需要同时访问时，分别访问将会繁琐难维护；**值对象模式将远程调用的传递数据封装在一个串行化的对象中进行传输**。用封装后的对象的在网络中传递，从而使系统拥有更好的交互模式；并减少网络通信数据，从而提高系统性能。

- 对象序列化
- 对象流
- 网络传输（同步阻塞Socket、同步非阻塞Channel、异步阻塞AsynchronizeChannel）



## 享元模式（Flyweight Pattern）

主要用于减少创建对象的数量，以减少内存占用和提高性能。它属于**结构型模式**。它提供了减少对象数量从而改善应用所需的对象结构的方式。