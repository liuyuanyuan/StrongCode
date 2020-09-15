[TOC]

## 链表 Linked List

### 分类：

##### 单向链表(单链表single linked list)：只有自身和后继元素(next)；

##### 双向链表(双链表double linked list)：有自身、前驱元素(pre)、后继元素(next)；

##### 循环链表(circular linked list)：首尾相接的链表

![image-20200915235511597](img/2linkedlist.png)

### 数据结构特点：

- 元素：可重复、可为null；

- 存储空间：

  - 不连续存储，所以空间利用率低；
  - 容量可变，没有扩容成本：直接增加指向新元素的链即可。

- 排序顺序：存储不连续，也没有按照元素排序，所以没有索引；

- 操作的时间复杂度：

  - 顺序遍历foreach：按插入顺序遍历、时间复杂度为O(n)；

  - 随机查找get(index)：

    按序号/元素查找的时间复杂度为O(index)，因为没有索引所以只能挨个遍历；

  - 随机修改set(index, element)：

    按序号修改的时间复杂度为O(index)，因为没有索引所以只能挨个遍历查找，然后修改元素；

  - 随机插入add(element)/add(index,element)：

    在头/尾添加元素的时间复杂度为O(1)；在头/尾外的其他位置，插入元素的时间复杂度为O(index)，因为虽然修改操作本身无成本，需要逐个遍历来查找元素；

  - 随机删除remove(index)/remove(element)：

    在头/尾添加元素的时间复杂度为O(1)；在头/尾外的其他位置，删除的时间复杂度为O(n)，因为虽然删除操作本身无成本，但是需要逐个遍历来查找元素，；

> 总结：
>
> - 链表是不连续存储，按插入顺序指向链，空间利用率低；
> - 链表是变长的，扩容没有成本；
>
> - 链表的顺序遍历消耗线性时间O(n)，随机查找消耗线性时间O(n)；
>
> - 链表的修改、插入、删除操作的成本取决于操作元素的位置，在头/尾消耗常数时间O(1)，在其他位置则需要消耗线性时间O(n)。(因为查找过程需要O(n))

### Java中的实现：

##### 双向链表(实现了List和Deque)：LinkedList<T> linkedList = new LinkedList<T>();

LinkedList (摘自openjdk12)的源码解析：

```java
public class LinkedList<E>
    extends AbstractSequentialList<E>
    implements List<E>, Deque<E>, Cloneable, java.io.Serializable
{
    transient int size = 0;
    // Pointer to first node.
    transient Node<E> first;
    // Pointer to last node.
    transient Node<E> last;
  
    /**
    Appends the specified element to the end of this list.
    <p>This method is equivalent to {@link #addLast}.
    */
    public boolean add(E e) {
        linkLast(e);
        return true;
    }
    public void add(int index, E element) {
        checkPositionIndex(index);
        if (index == size)
            linkLast(element);
        else
            linkBefore(element, node(index));
    }
    public void addFirst(E e) {
        linkFirst(e);
    }
    public void addLast(E e) {
        linkLast(e);
    }
  
    public E set(int index, E element) {
        checkElementIndex(index);
        Node<E> x = node(index);
        E oldVal = x.item;
        x.item = element;
        return oldVal;
    }
  
    public E get(int index) {
        checkElementIndex(index);
        return node(index).item;
    }
    public E getFirst() {
        final Node<E> f = first;
        if (f == null)
            throw new NoSuchElementException();
        return f.item;
    }
    public E getLast() {
        final Node<E> l = last;
        if (l == null)
            throw new NoSuchElementException();
        return l.item;
    }
  
    //由于实现了Deque 所以还支持以下操作方式
    public void push(E e) {
        addFirst(e);
    }
    public E pop() {
        return removeFirst();
    }
  
    public boolean offer(E e) {
        return add(e);
    }
    public boolean offerFirst(E e) {
        addFirst(e);
        return true;
    }
    public boolean offerLast(E e) {
        addLast(e);
        return true;
    }
    public E poll() {
        final Node<E> f = first;
        return (f == null) ? null : unlinkFirst(f);
    }
    public E pollFirst() {
        final Node<E> f = first;
        return (f == null) ? null : unlinkFirst(f);
    }
    public E pollLast() {
        final Node<E> l = last;
        return (l == null) ? null : unlinkLast(l);
    }
    public E peek() {
        final Node<E> f = first;
        return (f == null) ? null : f.item;
    }
    public E peekFirst() {
        final Node<E> f = first;
        return (f == null) ? null : f.item;
    }
    public E peekLast() {
        final Node<E> l = last;
        return (l == null) ? null : l.item;
    }   
}
```

