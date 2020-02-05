# 数据结构

[TOC]

**参考：**Data Structures and Algorithm Analysis:[in Java (Third Edition)](https://users.cs.fiu.edu/~weiss/#dsaajava3)  (by [Mark Allen Weiss](https://users.cs.fiu.edu/~weiss/))



## 树（tree）

#### 二叉树（binary tree）

二叉树（binary tree）是一棵树，且每个节点最多有2个子节点。

一棵平均二叉树的深度要比节点个数N小得多，这个性质有时很重要。二叉树的平均深度为:
$$
O(\sqrt{N})
$$
二叉树有两个作用：一个是用于查找（二叉查找树），另一个是用于编译器的设计领域（表达式树）。

```
//二叉树节点类
class BinaryNode {
	Object     element;//data in the node
	BinaryNode left;//left child
	BinaryNode right;//right child
}
```



#### 表达式树（expression tree）

二元操作的表达式树：所有的树叶都是**操作数(operand)**，所有的父节点都是**操作符(operator)**。如：

<img src="/Users/liuyuanyuan/github/StrongCode/java/images/expression-tree.png" style="zoom:33%;" />
$$
(a+b*c) + ((d*e+f)*g)
$$


> 注意：一目减运算符（unary minus operator）的表达式树中：一个节点只有一个子节点。如：-1。

中序遍历（inorder traversal）：左子，节点，右子；

后序遍历（inorder traversal）：左子，右子，节点；



#### 二叉查找树（binary search tree）

二叉查找树中，任意结点中的项，大于左子树中任意节点中的项，小于右子树中任意节点中的项。

二叉查找树要求所有节点中的项都能够排序（即可比较的，Java中二叉查找树的类需要实现Comparable接口，使用compareTo方法来进行两项间比较）。

[BinarySearchTree.java](https://users.cs.fiu.edu/~weiss/dsaajava3/code/BinarySearchTree.java)



#### AVL树：带平衡条件（balance condition）的二叉树

平衡条件必须容易保持，并且保证树的深度必须是：O(log N)。

平衡条件：

要求根节点的左、右子树具有相同的高度；（这会出现左子树只有左节点，右子树只有右接节点的情况。）

要求每个节点的具有相同高度的左、右子树。如果空子树的高度定义为-1（通常如此），那么只有具有（2的k次方-1）节点的理想平衡树满足这个条件。因此这种平衡树保证了树的深度小，但是它太严格而难以使用，需要方框条件。

**一棵AVL树是其每个节点的左子树和右子树的高度最多差1的二叉查找树（空树的高度定义为-1）。**实际AVL树的高度只略大于logN。

向AVL树中插入新节点可能会破坏平衡，通过对树进行简单修正来达到平衡条件的要求，这称作旋转（rotation）。

[AvlTree.java](https://users.cs.fiu.edu/~weiss/dsaajava3/code/AvlTree.java)



#### 红黑树（red black tree）：节点带红、黑着色的二叉查找树

历史上AVL树流行的一个变种是红黑树（red black tree）。红黑树的最大事件复杂度为O(log N)。

红黑树是具有下列着色性质的二叉查找树：

- 每个节点要么着黑色，要么着红色；
- 根节点是黑色；
- 如果一个节点是红色，那么其子节点必须是黑色；
- 从任一节点到一个null引用的每一条路径，必须包含相同数目的黑色节点。

着色法则的一个结论是，红黑树的高度最多是2log(N+1) 。



向一个红黑树中插入一个新的节点项，是困难的，**需要颜色的改变和树的旋转**。

[RedBlackTree.java](https://users.cs.fiu.edu/~weiss/dsaajava3/code/RedBlackTree.java)



## 表（list）

> **抽象数据类型**（Abstract data type，ADT）是带有一组操作的一些对象的集合。

大小为0的特殊表称为**空表（empty list）**。

**除空表外，列表中的元素都是依次连续（顺序）存储的。**（有时列表称为序列）

**基本操作**：printList打印所有元素，makeEmpty置空，find返回某一项首次出现的位置，insert插入元素，remove删除元素，findKth返回某位置上的元素。

**表的实现：**

- 数组列表
- 链表

#### 数组列表

数组可以实现列表上的所有操作，列表最简单的实现是通过数组。但是数组是定长的，列表是变长的，因此会涉及数组的扩展以保证可增长性。

```java
//数组扩展的实现
int[] arr = new int[10];
//...
int[] newArr = new int[arr.length*2];
for(int i=0; i<arr.lenght; i++){
  newArr[i] = arr[i];
}
newArr = arr;
```

**操作消耗**：printList消耗线性时间，findKth消耗常数时间；但插入和删除潜藏着昂贵开销（为了保证连续存储，开销取决于操作元素的位置，平均需要移动一半元素，因此仍然需要线性时间）。

**Java中的ArrayList就是可增长的数组实现的。**优点是get和set花费常数时间，缺点是新项删除和现有项删除代价昂贵，除非变动发生在ArrayList末端。

#### 链表

为了避免插入和删除的线性开销，需要保证列表可以不连续存储。

**简单链表**由一系列的节点组成，这些节点不必在内存中相连，每个节点含有该元素和包含该元素后继的节点的链（link），称作next链，最后一个节点的next链引用null。

**操作消耗**：

执行printList或find(x)是从第一个节点开始并通过后继next链来遍历点，这样消耗的时间是线性的，和数组实现时一样；

findKth操作不如数组实现时效率高，花费O(i)时间；

remove方法可以通过修改一个next引用来实现；insert方法需要用new操作从系统中取得一个节点，然后执行两次引用的调整，这两个操作只涉及常数个节点链的改变。



**双向链表（double linked list）**是在简单链表的基础上，每个节点含有：元素，后继节点的链以及前驱节点的链。

**Java中的LinkedList是双向链表。**优点是（在变动位置已知时）新项插入和现有项删除开销均很小，缺点时不容易作索引，因此get调用昂贵，除非调用端点附件的。



## 栈（stack）

栈（stack）是限制插入和删除只能在末端的表，这个位置称作栈的顶（top）。栈有时又称作后进先出（LIFO）表。

**基本操作**：

- 进栈（push），相当于插入；

- 出栈（pop），是删除最后插入的元素。

对空栈进行pop或top一般认为是栈ADT中的一个错误；当运行push时空间用尽是一个实现限制，但不是ADT错误。

**栈的实现：**由于栈是一个表，因此任何实现表的操作都能实现栈。ArrayList和LinkedList都支持栈操作，并在99%情况下是最合理的选择，偶尔设计特殊目的实现可能会更快。

- 栈的链表实现
- 栈的数组实现



## 队列（queue）

像栈一样，队列（queue）也是表。但使用队列时，插入在一端进行，删除则在另一端进行。





## 堆（heap）