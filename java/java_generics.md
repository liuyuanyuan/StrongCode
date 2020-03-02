# Java 泛型（generics）

[TOC]

**参考：**

- [The Java™ Tutorials](https://docs.oracle.com/javase/tutorial/extra/generics/fineprint.html)

- [Java 泛型详解博客](https://www.cnblogs.com/coprince/p/8603492.html)](https://www.cnblogs.com/coprince/p/8603492.html)

## 概念：泛型的本质是“参数化类型”

Java 泛型（generics）是 JDK 5 中引入的一个新特性，泛型提供了编译时类型安全检测机制，该机制允许程序员在编译时检测到非法的类型。

**泛型的本质是“参数化类型”**，也就是说所操作的数据类型被指定为一个参数。

提到参数，最熟悉的就是定义方法时有形参，然后调用此方法时传递实参。
参数化类型，顾名思义，就是将类型由原来的具体的类型参数化，类似于方法中的变量参数，此时类型也定义成参数形式（可以称之为类型形参），然后在使用/调用时传入具体的类型（类型实参）。

泛型的本质是为了参数化类型（在不创建新的类型的情况下，通过泛型指定的不同类型来控制形参具体限制的类型）。也就是说在泛型使用过程中，操作的数据类型被指定为一个参数，这种参数类型可以用在类、接口和方法中，分别被称为泛型类、泛型接口、泛型方法。

## 特性：泛型只在编译阶段有效

Java 中的泛型基本上都是在编译器这个层次来实现的。在生成的 Java 字节代码中是不包含泛型中的类型信息的。也就是说，**Java中的泛型只在编译阶段有效**。在编译过程中，正确检验泛型结果后，会将泛型的相关信息擦除（这个过程称为**类型擦除**），并且在对象进入和离开方法的边界处添加类型检查和类型转换的方法。也就是说，泛型信息不会进入到运行时阶段。

泛型类型在逻辑上可看作多个不同的类型，实际上都是相同的基本类型。

如：在代码中定义的 List<Object>和 List<String>等类型，在编译之后都会变成 List；JVM 看到的只是 List，而由泛型附加的类型信息对 JVM 来说是不可见的。

	public static <T> void main(String[] args) {
			List<String>  list1 = new ArrayList();	
			List<Integer>  list2 = new ArrayList();
			Class list1Class = list1.getClass();
			Class list2Class = list2.getClass();
			System.out.println(list1Class);
			System.out.println(list2Class);
			System.out.println(list1Class.equals(list2Class));		
	}	
	输出结果：
	class java.util.ArrayList
	class java.util.ArrayList
	true
#### 类型擦除的基本过程：

**首先是找到用来替换类型参数的具体类**（**这个具体类一般是 Object**，如果指定了类型参数的上界的话，则使用这个上界）；**然后把代码中的类型参数都替换成具体的类。**

#### 泛型的上下边界：

在使用泛型的时候，可以为传入的泛型类型实参进行上下边界的限制，如：类型实参只准传入某种类型的父类或某种类型的子类。

- 为泛型添加上边界：即传入的类型实参必须是指定类型的子类型；
- 为泛型添加下边界：即传入的类型实参必须是指定类型的父类型。

泛型的上下边界添加，必须与泛型的声明在一起 。

## 泛型的使用

泛型可以用在类、接口和方法中，分别被称为泛型类、泛型接口、泛型方法。

### 泛型类<T>

泛型类型用于类的定义中，被称为泛型类。通过泛型可以完成对一组类的操作对外开放相同的接口。最典型的就是各种容器类，如：List、Set、Map。

```
//泛型类测试
public class Generic<T> {
	  private T key;//key这个成员变量的类型为T,T的类型由外部指定  
    public Generic(T key) { //泛型构造方法形参key的类型也为T，T的类型由外部指定
        this.key = key;
    }
    public T getKey(){ //泛型方法getKey的返回值类型为T，T的类型由外部指定
        return key;
    }
    
    public static void main(String[] args) {
    	Generic<Integer>  genericInteger = new Generic<Integer>(123456);
    	Generic<String>  genericString = new Generic<String>("test_key");
    	System.out.println(genericInteger.getKey());
    	System.out.println(genericString.getKey());   	
	}
}

输出结果：
123456
test_key
```

```java
//ArrayList.class 摘要
public class ArrayList<E>{
		/**
     * Shared empty array instance used for default sized empty instances. We
     * distinguish this from EMPTY_ELEMENTDATA to know how much to inflate when
     * first element is added.
     */
    private static final Object[] DEFAULTCAPACITY_EMPTY_ELEMENTDATA = {};
		/**
     * The array buffer into which the elements of the ArrayList are stored.
     * The capacity of the ArrayList is the length of this array buffer. Any
     * empty ArrayList with elementData == DEFAULTCAPACITY_EMPTY_ELEMENTDATA
     * will be expanded to DEFAULT_CAPACITY when the first element is added.
     */
    transient Object[] elementData; // non-private to simplify nested class access
	  //Constructs an empty list with an initial capacity of ten.
    public ArrayList() {
        this.elementData = DEFAULTCAPACITY_EMPTY_ELEMENTDATA;
    }
    
    public boolean add(E e) {
        modCount++;
        add(e, elementData, size);
        return true;
    }
    public E get(int index) {
        Objects.checkIndex(index, size);
        return elementData(index);
    }
    
    //...
}
```

### 泛型接口<T>

泛型类型用于接口类的定义中，叫泛型接口，它与泛型类的定义及使用基本相同。

泛型接口常被用在各种类的生产器中，例如：

```
//定义一个泛型接口
public interface Generator<T> {
    public T next();
}
//泛型接口未传入实参时
class FruitGenerator<T> implements Generator<T>{
    @Override
    public T next() {
        return null;
    }
}
//泛型接口传入实参时
public class FruitGenerator implements Generator<String> {
    private String[] fruits = new String[]{"Apple", "Banana", "Pear"};
    @Override
    public String next() {
        Random rand = new Random();
        return fruits[rand.nextInt(3)];
    }
}
```

### 类型通配符<?>

类型通配符一般是使用?代替具体的类型参数。注意：此处’？’是类型实参，而不是类型形参 。

再直白点的意思就是：此处的？和Number、String、Integer一样都是一种实际的类型，可以把？看成所有类型的父类，是一种真实的类型；`List<?>`在逻辑上是 `List<Number>, List<String> , List<Integer>` 等所有List<具体类型实参>的父类。

- 类型通配符上限通过形如 `List<? extends Number> ` 来定义，如此定义就是通配符泛型值接受 Number 及其下层子类类型。

- 类型通配符下限通过形如 ` List<? super Number>`来定义，表示类型只能接受Number及其三层父类类型，如 Object 类型的实例

```
public static void main(String[] args) {
		List<String> name = new ArrayList<String>();
		List<Integer> age = new ArrayList<Integer>();
		List<Number> number = new ArrayList<Number>();
		name.add("An");
		age.add(1);
		number.add(2019);

		getData(name);
		getData(age);
		getData(number);
		
		//getUpperNumber(name);//error
		getUpperNumber(age);
		getUpperNumber(number);
		
		//getLowerNumber(name);//error
		//getLowerNumber(age);//error
		getLowerNumber(number);
	}
	public static void getData(List<?> data) {
		System.out.println("data :" + data.get(0));
	}
	public static void getUpperNumber(List<? extends Number> data) {
		System.out.println("data :" + data.get(0));
	}	
	public static void getLowerNumber(List<? super Number> data) {
		System.out.println("data :" + data.get(0));
	}
```



### 泛型方法<E>

一个泛型方法，法在调用时可以接收不同类型的参数。根据传递给泛型方法的参数类型，编译器适当地处理每一个方法调用。

定义泛型方法的规则：

- 所有泛型方法声明都有一个类型参数声明部分（由尖括号分隔），该类型参数声明部分在方法返回类型之前（在下面例子中的<E>）。
- 每一个类型参数声明部分包含一个或多个类型参数，参数间用逗号隔开。一个泛型参数，也被称为一个类型变量，是用于指定一个泛型类型名称的标识符。
- 类型参数能被用来声明返回值类型，并且能作为泛型方法得到的实际参数类型的占位符。
- 泛型方法体的声明和其他方法一样。注意类型参数只能代表引用型类型，不能是原始类型（像int,double,char的等）

> 注意区分：
>
> - 泛型类，是在实例化类的时候指明泛型的具体类型；
>
> - 泛型方法，是在调用方法的时候指明泛型的具体类型 。



#### 泛型类中的泛型方法

当泛型类通过具体类实例化后，其中的泛型方法，也只能接受该类类型的参数。



#### 泛型方法与可变参数T..

再看一个泛型方法和可变参数的例子：

```
public class GenericTest {
	public static <T> void print(T... args) {
		for(T t : args) {
			System.out.println("t is " + t);
		}
	}
	public static void main(String[] args) {
		GenericTest.print(1, 2 , "lily", true);		
	}
}

output:
t is 1
t is 2
t is lily
t is true
```

####  静态方法与泛型

静态方法无法访问泛型类上定义的泛型；如果静态方法操作的引用数据类型不确定的时候，必须要将泛型定义在方法上。

```
public class StaticGenerator<T> {
    /**
     * 即使静态方法不允许使用泛型类中已经声明过的泛型。
     * 如：public static void show(T t){..},此时编译器会提示错误信息：
          "StaticGenerator cannot be refrenced from static context"
     */
    public static <T> void show(T t){
    }
}
```



## 注意：泛型数组

Sun的说明文档，在Java中是**不能创建一个确切的泛型类型的数组**的。

如下的例子是不允许的：

```
List<String>[] ls = new ArrayList<String>[10];  //错误
```

使用通配符创建泛型数组是允许的：

```
List<?>[] ls = new ArrayList<?>[10]; //正确
```

下面也是允许的：

```
List<String>[] ls = new ArrayList[10];//正确
```

下面使用https://docs.oracle.com/javase/tutorial/extra/generics/fineprint.html中的例子来说明该问题：

```java
// Not really allowed.
List<String>[] lsa = new List<String>[10];
Object o = lsa;
Object[] oa = (Object[]) o;
List<Integer> li = new ArrayList<Integer>();
li.add(new Integer(3));
// Unsound, but passes run time store check
oa[1] = li;

// Run-time error: ClassCastException.
String s = lsa[1].get(0);
```

这种情况下，由于JVM泛型的擦除机制，在运行时JVM是不知道泛型信息的，所以可以给oa[1]赋上一个ArrayList，而不会出现异常；但是在取出数据的时候却要做一次类型转换，所以就会出现ClassCastException；如果可以进行泛型数组的声明，上面说的这种情况在编译期将不会出现任何的警告和错误，只有在运行时才会出错。而对泛型数组的声明进行限制，对于这样的情况，可以在编译期提示代码有类型安全问题，比没有任何提示要强很多。

下面采用通配符的方式是被允许的:数组的类型不可以是类型变量，除非是采用通配符的方式，因为对于通配符的方式，最后取出数据是要做显式的类型转换的。

```java
// OK, array of unbounded wildcard type.
List<?>[] lsa = new List<?>[10];
Object o = lsa;
Object[] oa = (Object[]) o;
List<Integer> li = new ArrayList<Integer>();
li.add(new Integer(3));
// Correct.
oa[1] = li;
// Run time error, but cast is explicit.
String s = (String) lsa[1].get(0);
```

在下一个会导致编译时错误的变体中，我们避免创建一个数组对象，该数组对象的元素类型已参数化，但仍使用带有参数化元素类型的数组类型。

```
// Error.
List<String>[] lsa = new List<?>[10];
```

同样，尝试创建元素类型为类型变量的数组对象会导致编译时错误：

```
<T> T[] makeArray(T t) {
    return new T[100]; // Error.
}
```

由于类型变量在运行时不存在，因此无法确定实际的数组类型。

解决此类局限性的方法是将类文字用作运行时类型标记，如文档 [Class Literals as Runtime-Type Tokens](https://docs.oracle.com/javase/tutorial/extra/generics/literals.html) 所述。