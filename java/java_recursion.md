# Java的递归实现

#### 分析

递归好处：代码更简洁清晰，可读性更好；
递归坏处：由于递归需要系统堆栈，所以空间消耗要比非递归代码要大很多。而且，如果递归深度太大，可能系统撑不住。 
总结：**递归要慎用，非必要时不使用递归。**

**典型例子：**求阶乘，求文件夹大小，删除文件夹等等。
参考：https://www.jianshu.com/p/edfc4e35f383

#### 例1：求任意整数的阶乘

x!=x*(x-1)! 或者 x!=x*(x-1)*...3*2*1
实现：

```java
// 递归方法实现
public static BigInteger getFactorial(int num)
{
    if (num == 1)
    {
				return BigInteger.ONE;
    }
    return BigInteger.valueOf(num).multiply(getFactorial(num - 1));
}	
// 普通方法实现
public static BigInteger getFactorial2(int x)
{
    BigInteger y = BigInteger.ONE;
    for (int i = x; i >= 1; i--)
    {
        y = y.multiply(BigInteger.valueOf(i));
    }
    return y;
}

public static void main(String[] args)
{
   System.out.println("10的阶乘是：" + getFactorial(10));   
   System.out.println("10的阶乘是：" + getFactorial2(10)); 
}
```



#### 例2：打印99乘法表

嵌套for循环 和  用递归实现 的比较:
栈主要是用来存放栈帧的，每执行一个方法就会出现压栈操作，所以采用递归的时候产生的栈帧比较多，递归就会影响到内存，非常消耗内存；
而使用for循环就执行了一个方法，压入栈帧一次，只存在一个栈帧，所以比较节省内存。

```java
public class MultiTable
{
    final int x=9;	
    public static void main(String[] args)
    {		
				MultiTable mt = new MultiTable();
        mt.normal();//普通方法实现
				mt.recursive2(1);//递归方法实现
    }
  
		public void normal()
		{    	
			for(int i=1; i<=x; i++)
			{
    		for (int j = 1; j <= x; j++)
    		{
					System.out.print(i + "*" + j + "=" + i * j + "  ");
    		}
    		System.out.println();
			}    	
		}
     
		public void recursive(int i)
		{
 			if (i == 10)
 			{
     		return;
 			}
	
 			for (int j = 1; j <= x; j++)
 			{
     		System.out.print(i + "*" + j + "=" + (i * j) + " ");
			}
 			System.out.println();
	
 			recursive(i+1);
		}

		public String recursive2(int i)
		{
				String s = "";
				if (i == 10)
				{
   				 return s;
				}
	
				for (int j = 1; j <= x; j++)
				{
     			s = s + i + "*" + j + "=" + (i * j) + " ";
				}
				s = s + "\n";
				return s + recursive2(i+1);
		}
}
```


