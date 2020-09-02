# Java 8 Lambda 表达式 编程



### 例子

```java
package com.lyy.arithmetic;

/**
 Lambda 表达式，也可称为闭包，它是推动 Java 8 发布的最重要新特性。
 Lambda 允许把函数作为一个方法的参数（函数作为参数传递进方法中）。
 使用 Lambda 表达式可以使代码变的更加简洁紧凑。
 *
 * @author liuyuanyuan
 * @version 1.0.0
 * @create 2020/8/26
 */
public class LambdaTest {

    interface MathOperation {
        int operation(int a, int b);
    }

    private int operate(MathOperation mathOperation, int a, int b){
        return mathOperation.operation(a, b);
    }


    interface GreetingService {
        void sayMessage(String message);
    }

    public static void main(String[] args) {
        MathOperation add = (int a, int b) -> a + b;
        MathOperation plus = (int a, int b) -> a * b;
        MathOperation xxx = (int a, int b) -> {
            int x = (a+b)*100;
            return x;
        };

        LambdaTest lt = new LambdaTest();
        System.out.println(lt.operate(add, 1, 2));
        System.out.println(lt.operate(plus, 1, 2));
        System.out.println(lt.operate(xxx, 2, 3));


        // 不用括号
        GreetingService greetService1 = message ->
        System.out.println("Hello " + message);
        // 用括号
        GreetingService greetService2 = (message) ->
        System.out.println("Hi " + message);

        greetService1.sayMessage("World");
        greetService2.sayMessage("Google");
    }

}

```