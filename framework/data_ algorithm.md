# 数据结构与算法

[TOC]

## 数据结构：

- 数组

  Object[]

  ArrayList

  

- 链表  

  ```java
  LinkedList<T> list = new LinkedList();
  ```

  

- 哈希表  

  HashMap<K, V>

  LinkedHashMap<K,V>（按插入/访问排序）

- 队列 ：先进先出

  ```java
  Queue<T> queue= new LinkedList<T>();
  // 按入队顺序排序
  queue.offer(ele); // 入队
  queue.poll(); // 出队，队列为空时返回null不报异常
  queue.peek(); //队列的第一个元素
  ```

- 栈：后进先出

  ```java
  Stack<T>  stack = new Stack(); 
  stack.push(ele); // 入栈
  stack.pop();  // 出栈，队列为空时报异常
  stack.peek(); // 获取栈顶元素
  ```

- 堆：堆序性（最大堆/最小堆）

  ```java
  PriorityQueue<T> minHeap = new  PriorityQueue(); 
  
  minHeap.offer(ele); // 入队
  minHeap.poll(); // 出队
  queue.peek(); // 获取队列第1个元素，即堆中最小的元素
  ```



## 算法思维

- 二分查找
- 

- 动态规划

  

![image-20200915165735239](/Users/liuyuanyuan/Library/Application Support/typora-user-images/image-20200915165735239.png)



## 用途：

- 性能优化
- 代码优化



## 解决思路：

在开发前，一定要对问题的复杂度进行分析，做好技术(数据结构和算法)选型。这就是定位问题的过程。只有把这个过程做好，才能更好地解决问题。

常用的分析问题的方法有以下 4 种：

- 复杂度分析。估算问题中复杂度的上限和下限。

- 定位问题。根据问题类型，确定采用何种算法思维。

- 数据操作分析。根据增、删、查和数据顺序关系去选择合适的数据结构，利用空间换取时间。

- 编码实现。

