# Java反射机制

1：SUN提供的反射机制的类：

java.lang.Class<T>

java.lang.reflect.Constructor<T>

java.lang.reflect.Field

java.lang.reflect.Method

java.lang.reflect.Modifier

2：什么是反射

JAVA反射机制是在运行状态中，对于任意一个类。都能都知道这个类的所有属性和方法，对于任意一个对象，都能够调用它的任意一个方法和属性；这种动态获取的信息以及动态调用对象的方法的功能称之为java语言的反射机制；

3：反射的作用

反编译 .class --à .java

通过反射机制可以访问java对象中的属性，方法，构造方法

4：创建Class对象的三种方式

```java
 public static void main(String[] args) {
             Person p1 = new Person("小明" ,20,'男' );
             Person p2 = new Person("小红" ,23,'女' );
   
             //创建Class对象的方式一：(对象.getClass())，获取person类中的字节码文件
             Class class1 = p1.getClass();
             System. out.println(p1.getClass().getName());
             Class class2 = p2.getClass();
             System. out.println(class1 == class2 );
             
             //创建Class对象的方式二：(类.class:需要输入一个明确的类，任意一个类型都有一个静态的class属性)
             Class class3 = Person.class;
             System. out.println(class1 == class2);
             
             //创建Class对象的方式三：(forName():传入时只需要以字符串的方式传入即可)
             //通过Class类的一个forName（String className)静态方法返回一个Class对象，className必须是全路径名称；
             Class class4 = null;
             try {
                class4 = Class.forName("cn.itcast.Person");
                System. out.println(class4 == class3);
             } catch (ClassNotFoundException e) {
                e.printStackTrace();
            }
}
```

