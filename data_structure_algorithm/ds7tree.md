[TOC]

## 树 (Tree)

树结构：

- 只有结点和连线，没有环；

- 结点与结点之间是一对多的关系；

  

## 二叉树(binary tree)

二叉树（binary tree）是每个节点最多有2个子节点的树。

一棵平均二叉树的深度要比节点个数N小得多，这个性质有时很重要。二叉树的平均深度为:
$$
O(\sqrt{N})
$$
### 二叉树的层次与深度/高度

树中结点的最大层次数，就是这棵树的深度(也称为高度)。

<img src="img/5binarytree_height.png" alt="img" style="zoom: 20%;" />

### 二叉树的分类与存储

#### 二叉树的分类

- 满二叉树(除了叶子结点外，所有结点都有两个子结点)
- 完全二叉树(存储空间利用率的100%)
- 非完全二叉树

<img src="img/7binarytree_classify.png" alt="image" style="zoom:25%;" />

#### 二叉树的2种存储方式

- 基于指针的链式存储法(通用方法)

  ```java
  //二叉树节点类定义
  class BinaryNode {
  	Object     value;//data in the node
  	BinaryNode left;//left child
  	BinaryNode right;//right child
  }
  ```

- 基于数组的顺序存储法

链式存储法，也就是像链表一样，每个结点有3个字段，一个存储数据，另外两个分别存放指向左右子结点的指针，如下图所示：

<img src="https://s0.lgstatic.com/i/image/M00/1F/E1/CgqCHl7nVhKAJVYKAABbMx2OS5o954.png" alt="image" style="zoom: 20%;" />

顺序存储法，就是按照规律把结点存放在数组里，如下图所示，为了方便计算，我们会约定把根结点放在下标为 1 的位置。随后，B 结点存放在下标为 2 的位置，C 结点存放在下标为 3 的位置，依次类推。

<img src="https://s0.lgstatic.com/i/image/M00/1F/E1/CgqCHl7nVhyAF-yqAAFEIfF2-z4697.png" alt="img" style="zoom: 20%;" />

根据这种存储方法，我们可以发现如果结点 X 的下标为 i，那么 X 的左子结点总是存放在 2 * i 的位置，X 的右子结点总是存放在 2 * i + 1 的位置。

之所以称为完全二叉树，是从存储空间利用效率的视角来看的。对于一棵完全二叉树而言，仅仅浪费了下标为 0 的存储位置。而如果是一棵非完全二叉树，则会浪费大量的存储空间。

<img src="https://s0.lgstatic.com/i/image/M00/1F/D5/Ciqc1F7nVi2AVfUZAAFA7ZImLgI310.png" alt="image" style="zoom:20%;" />

### 结点遍历方式(父结点)-使用递归遍历

- 先序遍历：**节点自身**，左子树，右子树；

  对树中的任意结点，先打印结点自身，然后前序遍历它的左子树，最后前序遍历它的右子树。

- 中序遍历：左子树，**节点自身**，右子树；

  对树中的任意结点，先中序遍历它的左子树，然后打印结点自身，最后中序遍历它的右子树。

- 后序遍历：左子树，右子树，**父节点**；

  对树中的任意结点，先中序遍历它的左子树，然后打印结点自身，最后中序遍历它的右子树。

<img src="img/7binarytree_order_traverse.png" alt="image" style="zoom:25%;" />

```java
// 先序遍历
public static void preOrderTraverse(Node node) {
    if (node == null){
    	return;
    } 
    System.out.print(node.data + " ");
    preOrderTraverse(node.left);
    preOrderTraverse(node.right);
}
// 中序遍历
public static void inOrderTraverse(Node node) {
    if (node == null){
    	return;
    }
    inOrderTraverse(node.left);
    System.out.print(node.data + " ");
    inOrderTraverse(node.right);
}
// 后序遍历
public static void postOrderTraverse(Node node) {
    if (node == null){
    	return;
    }
    postOrderTraverse(node.left);
    postOrderTraverse(node.right);
    System.out.print(node.data + " ");
}
//计算树结点总数
public static int countNode(Node root){
  if (node == null){
    	return 0;
  }
  return 1 + countNode(root.left) + countNode(root.right);
}
//根据层次遍历树结点
public static void levelTraverse(Node root) {
    if (root == null) {
        return;
    }
    LinkedList<Node> queue = new LinkedList<Node>();
    Node current = null;
    queue.offer(root); // 根节点入队
    while (!queue.isEmpty()) { // 只要队列中有元素，就可以一直执行，非常巧妙地利用了队列的特性
        current = queue.poll(); // 出队队头元素
        System.out.print("-->" + current.data);
        // 左子树不为空，入队
        if (current.leftChild != null){
           queue.offer(current.leftChild);
        } 
        // 右子树不为空，入队
        if (current.rightChild != null){
           queue.offer(current.rightChild);
        }
    }
}

```



### 操作的时间复杂度

- 遍历的时间复杂度：O(n)
- 查找的时间复杂度：O(n)
- 增、删处理的时间复杂度：查找O(n)+操作O(1)

### 二叉树的2种用途：

- 用于编译器的设计领域(表达式树）

- 用于查找(二叉查找树)；

  

## 表达式树(expression tree)

二元操作的表达式树：所有的树叶都是**操作数(operand)**，所有的父节点都是**操作符(operator)**。如：

<img src="img/7expression-tree.png" alt="image-20200304173305792" style="zoom: 25%;" />
$$
(a+b*c) + ((d*e+f)*g)
$$


> 注意：一目减运算符（unary minus operator）的表达式树中：一个节点只有一个子节点。如：-1。





## 二叉查找树(binary search tree)

二叉查找树(也称作二叉搜索树)具备以下特性：

- 任一结点中的值，大于其左子树中任一节点的值，小于其右子树中任一节点的值。

- 二叉查找树要求所有节点中的值都能够排序（即可比较的，Java中二叉查找树的类需要实现Comparable接口，使用compareTo方法来进行两项间比较，如：TreeMap和TreeSet）。[BinarySearchTree.java](https://users.cs.fiu.edu/~weiss/dsaajava3/code/BinarySearchTree.java)

- 在二叉查找树中，会尽可能规避两个结点数值相等的情况。

- 对二叉查找树进行中序遍历，就可以的到一个从小到大的有序数据队列。

  如下图所示，中序遍历的结果就是 10、13、15、16、20、21、22、26。

  <img src="img/7binarysearchtree_inorder_traverse.png" alt="image" style="zoom:25%;" />

### 操作的时间复杂度

- 遍历的时间复杂度：O(n)
- 查找的时间复杂度(因为有序，所以是二分查找)：**O(log n)**
- 增、删处理的时间复杂度：查找O(n)+操作O(1)



### AVL树：带平衡条件的二叉查找树

平衡条件必须容易保持，并且保证树的深度必须是：O(log N)。

平衡条件：

- 要求根节点的左、右子树具有相同的高度；（这会出现左子树只有左节点，右子树只有右接节点的情况。）

- 要求每个节点的具有相同高度的左、右子树。如果空子树的高度定义为-1（通常如此），那么只有具有（2的k次方-1）节点的理想平衡树满足这个条件。因此这种平衡树保证了树的深度小，但是它太严格而难以使用，需要放宽条件。

**一棵AVL树是其每个节点的左子树和右子树的高度最多差1的二叉查找树（空树的高度定义为-1）。**实际AVL树的高度只略大于logN。

**操作消耗**：向AVL树中插入新节点可能会破坏平衡，通过对树进行简单修正来达到平衡条件的要求，这称作旋转（rotation）。

[AvlTree.java](https://users.cs.fiu.edu/~weiss/dsaajava3/code/AvlTree.java)



### 红黑树(red black tree)：节点带红、黑着色的二叉查找树

历史上AVL树流行的一个变种是红黑树（red black tree）。红黑树的最大时间复杂度为O(log N)。

**红黑树是具有下列着色性质的二叉查找树：**

- 每个节点要么着黑色，要么着红色；
- 根节点是黑色；
- 如果一个节点是红色，那么其子节点必须是黑色；
- 从任一节点到一个null引用的每一条路径，必须包含相同数目的黑色节点。

着色法则的一个结论是，红黑树的高度最多是2log(N+1) 。

**操作消耗**：向一个红黑树中插入一个新的节点项，是困难的，**需要颜色的改变和树的旋转**。

[RedBlackTree.java](https://users.cs.fiu.edu/~weiss/dsaajava3/code/RedBlackTree.java)



### B-Tree(平衡多路查找树）

B-tree 又叫平衡多叉查找树。一棵 m 阶的 B-tree (m 叉树)的特性如下(其中 ceil(x)是一个取上限的函数):

> 1. 树中每个结点至多有 m 个孩子;
>
> 2. 除根结点和叶子结点外，其它每个结点至少有有 ceil(m / 2)个孩子;
>
> 3. 若根结点不是叶子结点，则至少有 2 个孩子(特殊情况:没有孩子的根结点，即根结点为叶子
>
>    结点，整棵树只有一个根节点);
>
> 4. 所有叶子结点都出现在同一层，叶子结点不包含任何关键字信息(可以看做是外部结点或查询
>
>    失败的结点，实际上这些结点不存在，指向这些结点的指针都为 null);
>
> 5. 每个非终端结点中包含有 n 个关键字信息: (n，P0，K1，P1，K2，P2，......，Kn，Pn)。其中:
>
>   - Ki (i=1...n)为关键字，且关键字按顺序排序 K(i-1)< Ki。
>    - Pi 为指向子树根的接点，且指针 P(i-1)指向子树种所有结点的关键字均小于 Ki，但都大于 K(i- c) 关键字的个数 n 必须满足: ceil(m / 2)-1 <= n <= m-1。
>



**B-Tree 是为了文件系统（磁盘或其它外存设备）而设计的一种多叉平衡查找树（相对于二叉，B树每个内结点有多个分支，即多叉）。**

> 系统从磁盘读取数据到内存时是以磁盘块（block）为基本单位的，位于同一个磁盘块中的数据会被一次性读取出来，而不是需要什么取什么。
>
> MySQL的InnoDB存储引擎中有页（Page）的概念，页是其磁盘管理的最小单位。InnoDB存储引擎中默认每个页的大小为16KB，可通过参数innodb_page_size将页的大小设置为4K、8K、16K，在[MySQL](http://lib.csdn.net/base/mysql)中可通过如下命令查看页的大小：
>
> ```
> mysql> show variables like 'innodb_page_size';
> - 1
> - 1
> ```
>
> 而系统一个磁盘块的存储空间往往没有这么大，因此InnoDB每次申请磁盘空间时都会是若干地址连续磁盘块来达到页的大小16KB。InnoDB在把磁盘数据读入到磁盘时会以页为基本单位，在查询数据时如果一个页中的每条数据都能有助于定位数据记录的位置，这将会减少磁盘I/O次数，提高查询效率。

**B-Tree结构的数据可以让系统高效的找到数据所在的磁盘块。**为了描述B-Tree，首先定义一条记录为一个二元组[key, data] ，key为记录的键值，对应表中的主键值，data为一行记录中除主键外的数据。对于不同的记录，key值互不相同。

一棵m阶的B-Tree有如下特性： 

1. 每个节点最多有m个孩子。 

2. 除了根节点和叶子节点外，其它每个节点至少有Ceil(m/2)个孩子。 

3. 若根节点不是叶子节点，则至少有2个孩子 

4. 所有叶子节点都在同一层，且不包含其它关键字信息 

5. 每个非终端节点包含n个关键字信息（P0,P1,…Pn, k1,…kn） 

6. 关键字的个数n满足：ceil(m/2)-1 <= n <= m-1 

7. ki(i=1,…n)为关键字，且关键字升序排序。 

8. Pi(i=1,…n)为指向子树根节点的指针。P(i-1)指向的子树的所有节点关键字均小于ki，但都大于k(i-1)

B-Tree中的每个节点根据实际情况可以包含大量的关键字信息和分支，如下图所示为一个3阶的B-Tree： 

![索引](img/7b-tree.png)

B-Tree结构的数据可以让系统高效的找到数据所在的磁盘块。为了描述B-Tree，首先定义一条记录为一个二元组[key, data] ，key为记录的键值，对应表中的主键值，data为一行记录中除主键外的数据。对于不同的记录，key值互不相同。



### B+Tree(B-Tree的优化)

B+Tree是在B-Tree基础上的一种优化，使其更适合实现数据库索引（外存储索引结构）。

> MySQL的InnoDB存储引擎就是用B+Tree实现其索引结构。

从上面B-Tree结构图中可以看到每个节点中不仅包含数据的key值，还有data值。而每一个页的存储空间是有限的，如果data数据较大时将会导致每个节点（即一个页）能存储的key的数量很小，当存储的数据量很大时同样会导致B-Tree的深度较大，增大查询时的磁盘I/O次数，进而影响查询效率。在B+Tree中，所有数据记录节点都是按照键值大小顺序存放在同一层的叶子节点上，而非叶子节点上只存储key值信息，这样可以大大加大每个节点存储的key值数量，降低B+Tree的高度。

**B+Tree相对于B-Tree有几点不同：**

1. **非叶子节点只存储键值信息。**
2. **所有叶子节点之间都有一个链指针。**
3. **数据记录都存放在叶子节点中。**

将上一节中的B-Tree优化，由于B+Tree的非叶子节点只存储键值信息，假设每个磁盘块能存储4个键值及指针信息，则变成B+Tree后其结构如下图所示： 
![索引](img/7b+tree.png)

通常在B+Tree上有两个头指针，一个指向根节点，另一个指向关键字最小的叶子节点，而且所有叶子节点（即数据节点）之间是一种链式环结构。因此可以对B+Tree进行两种查找运算：一种是对于主键的范围查找和分页查找，另一种是从根节点开始，进行随机查找。

可能上面例子中只有22条数据记录，看不出B+Tree的优点，下面做一个推算：

InnoDB存储引擎中页的大小为16KB，一般表的主键类型为INT（占用4个字节）或BIGINT（占用8个字节），指针类型也一般为4或8个字节，也就是说一个页（B+Tree中的一个节点）中大概存储16KB/(8B+8B)=1K个键值（因为是估值，为方便计算，这里的K取值为〖10〗^3）。也就是说一个深度为3的B+Tree索引可以维护10^3 * 10^3 * 10^3 = 10亿 条记录。

实际情况中每个节点可能不能填充满，因此在数据库中，B+Tree的高度一般都在2~4层。[mysql](http://lib.csdn.net/base/mysql)的InnoDB存储引擎在设计时是将根节点常驻内存的，也就是说查找某一键值的行记录时最多只需要1~3次磁盘I/O操作。



> **数据库中的B+Tree索引：可以分为聚集索引（clustered index）和辅助索引（secondary index）。**
>
> 上面B+Tree示例图，在数据库中的实现即为聚集索引，聚集索引的B+Tree中的叶子节点存放的是整张表的行记录数据。
>
> 辅助索引与聚集索引的区别在于辅助索引的叶子节点并不包含行记录的全部数据，而是存储相应行数据的聚集索引键，即主键。
>
> 当通过辅助索引来查询数据时，MySQL 的 InnoDB存储引擎会遍历辅助索引找到主键，然后再通过主键在聚集索引中找到完整的行记录数据。
>
> **为什么说B+树比B树更适合数据库索引？**
>
> 1B+Tree的磁盘读写代价更低：B+树的内部节点并没有指向关键字具体信息的指针，因此其内部节点相对B树更小，如果把所有同一内部节点的关键字存放在同一盘块中，那么盘块所能容纳的关键字数量也越多，一次性读入内存的需要查找的关键字也就越多，相对IO读写次数就降低了。
>
> 2 B+Tree的查询效率更加稳定：由于非终结点并不是最终指向文件内容的结点，而只是叶子结点中关键字的索引。所以任何关键字的查找必须走一条从根结点到叶子结点的路。所有关键字查询的路径长度相同，导致每一个数据的查询效率相当。
>
> 3 由于B+Tree的数据都存储在叶子结点中，分支结点均为索引，方便扫库，只需要扫一遍叶子结点即可，但是B树因为其分支结点同样存储着数据，我们要找到具体的数据，需要进行一次中序遍历按序来扫，所以B+树更加适合在区间查询的情况，所以通常B+树用于数据库索引。


