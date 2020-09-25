[TOC]

## 堆Heap

堆，就是一棵被完全填满的二叉树，只有底层例外，底层上的元素从左往右填入；这样的二叉树称为**完全二叉树**。

堆就是用**数组**实现的完全二叉树，所以它没有使用父指针或者子指针。**堆根据“堆属性”来排序**，“堆属性”决定了树中节点的位置。

**堆属性**

堆分为两种：*最大堆* 和*最小堆*，两者的差别在于节点的排序方式。

- 在最大堆中：父节点的值比每一个子节点的值都要大。

- 在最小堆中：父节点的值比每一个子节点的值都要小。

这就是所谓的“堆属性”，并且这个属性对堆中的每一个节点都成立。

**堆的性质**

- 结构性
- 堆序性（heap-order property）

对堆的一次操作可能破坏这堆的两个性质中的一个，因此，对堆的操作必须到堆所有性质都被满足才能终止。

**应用场景**

- 构建优先队列
- 快速找出一个集合中的最小值(最小堆)或者最大值(最大堆)；
- 支持堆排序

### Java中的实现：

#### 优先队列：PriorityQueue<T>（基于Object[]数组实现）

```java
public static void testHeap(){
        //heap
        PriorityQueue<Integer> minHeap = new PriorityQueue<>();
        minHeap.offer(2);
        minHeap.offer(1);
        minHeap.offer(3);
        minHeap.offer(0);
        System.out.println(minHeap.peek());
        while(!minHeap.isEmpty()){
            System.out.print(minHeap.poll() + ", ");
        }
        System.out.println();
        System.out.println("----heap end----");
}
```

