[TOC]

## 栈 Stack - LIFO

栈(stack)，是限制插入和删除只能在末端的线性表，这个位置称作栈顶（top）。又称作**后进先出（LIFO）表**。

### 限制操作：

**基本操作：**

- 进栈(push)：在栈顶插入元素；

- 出栈(pop)：在栈顶删除并返回最后插入的元素。

**操作特性：**

- 对空栈，进行 pop 或 top，一般认为是栈 ADT 错误；

- 当执行 push 时，空间用尽是一个实现限制，但不是 ADT 错误。

<img src="img/4stack_struture.png" alt="image-20200326165442756" style="zoom: 33%;" />

### 栈的实现原理:

由于栈是一个线性表，因此任何实现线性表的操作都能实现栈。ArrayList 和 LinkedList 都支持栈操作，并在99%情况下是最合理的选择，偶尔设计特殊目的实现可能会更快。

- [栈的链表实现](https://www.cs.usfca.edu/~galles/visualization/StackLL.html)

  <img src="img/4stack_linkedlist.png" alt="image-20200904093814209" style="zoom:33%;" />

- [栈的数组实现](https://www.cs.usfca.edu/~galles/visualization/StackArray.html)

  <img src="img/4stack_array.png" alt="image-20200904093625658" style="zoom: 33%;" />

### Java 中的实现：

#### 栈：Stack<T> (扩展数组实现 class Stack<E> extends Vector<E> )

```java
				Stack<Integer> stack = new Stack(); 
        stack.push(2);
        stack.push(1);
        stack.push(3);
        stack.push(0);
        System.out.println(stack.peek());
        while(!stack.isEmpty()){
            System.out.print(stack.pop() + ", ");
        }
        System.out.println();// 0, 3, 1, 2
```





