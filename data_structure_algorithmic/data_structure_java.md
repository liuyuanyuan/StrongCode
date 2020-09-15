# 数据结构、算法与Java

[TOC]

## 数据结构的 Java 实现

- 数组

  - 纯定长数组：Object[] 

  - 可增长的数组：ArrayList<T> arrayList = new ArrayList<T>();

- 链表  

  ```java
  LinkedList<T> linkedList = new LinkedList<T>();
  ```

- 哈希表  

  HashMap<K, V> 

  LinkedHashMap<K,V>（按插入/访问排序）

- 队列(先进先出)：

  ```java
  Queue<T> queue= new LinkedList<T>();
  // 按入队顺序排序
  queue.offer(ele); // 入队
  queue.poll(); // 出队，队列为空时返回null不报异常
  queue.peek(); //队列的第一个元素
  DQueue<T> dequeue= new LinkedList<T>();
  ```

- 栈(后进先出)：

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



- 

