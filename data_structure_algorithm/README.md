# 数据结构与算法

数据结构是算法的基石，优秀的算法往往取决于数据结构的选型，所以两者需要融会贯通。本部分主要介绍数据结构与算法理论基础、用途以及在 Java 中的具体实现。

##### 参考

- [usfca Data Structure Visualizations](https://www.cs.usfca.edu/~galles/visualization/Algorithms.html)
- [Data Structures and Algorithm Analysis in Java (Third Edition)](https://users.cs.fiu.edu/~weiss/#dsaajava3)  (by [Mark Allen Weiss](https://users.cs.fiu.edu/~weiss/))
- Java 8+ Source Code

## 数据结构

1. [数组、字符串](ds1array_string.md)-数据类型统一、基于索引查找迅速

2. 线性表-数据的顺序性

   [链表(单向/双向/循环)](ds2linkedlist.md)

   栈 - LIFO

   队列/双端队列 - FIFO

3. 树(二叉树、二叉查找树，前/中/后序)-数据的层级性

4. [哈希表](ds3hashtable) - 条件查找迅速

5. 堆(最大堆/最小堆)

##### 特点分析角度

- 元素：是否唯一/可重复、是否可为Null；
- 存储空间：连续性、扩展性；
- 排序顺序：是否有序性、什么顺序、是否有索引
- 操作特点和时间复杂度：顺序遍历、随机查找、随机修改、随机插入、随机删除(删/查都是先查找，再删/改)；

##### [数据结构的Java实现](data_structure_java.md)

## 算法思维

##### 思维

- 递归
- 分治
- 动态规划

##### 算法

2. [排序算法](a1sort.md)
3. [查找算法](a2search,md)
4. [加密算法](a3encrypt.md)
5. 动态规划

## 解决问题的思路

在开发前，必须做好问题定位：先对问题的复杂度进行分析，再做好技术(数据结构和算法)选型。只有把这个过程做好，才能更好地解决问题。

常用的分析问题的方法有以下 4 种：

- 复杂度分析。估算问题中复杂度的上限和下限。

- 定位问题。根据问题类型，确定采用何种算法思维。

- 数据操作分析。根据增、删、查和数据顺序关系去选择合适的数据结构，利用空间换取时间。

- 编码实现。

## 用途：

- 代码优化
- 性能优化