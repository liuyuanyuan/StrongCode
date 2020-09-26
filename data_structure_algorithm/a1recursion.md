## Recursion 递归 - 函数自我调用以缩小问题规模

- [Factorial](https://www.cs.usfca.edu/~galles/visualization/RecFact.html)
- [Reversing a String](https://www.cs.usfca.edu/~galles/visualization/RecReverse.html)
- [N-Queens Problem](https://www.cs.usfca.edu/~galles/visualization/RecQueens.html)



## 概念-

递归的概念：函数定义中自己调用自己；

递归的两部分：

- 递归主体：
- 终止条件(否则会导致无限递归)

递归的数学模型：数学归纳法；

递归的思路：

- 写出递归公式：将大规模问题分解为解决思路相同的子问题；
- 找出终止条件；



## 递归解决汉诺塔问题

```java
public static void main(String[] args) {
    String x = "x";
    String y = "y";
    String z = "z";
    hanio(3, x, y, z);
}
public void hanio(int n, String x, String y, String z) {
    if (n < 1) {
        System.out.println("汉诺塔的层数不能小于1");
    } else if (n == 1) {
        System.out.println("移动: " + x + " -> " + z);
        return;
    } else {
        hanio(n - 1, x, z, y);
        System.out.println("移动: " + x + " -> " + z);
        hanio(n - 1, y, x, z);
    }
}
```

