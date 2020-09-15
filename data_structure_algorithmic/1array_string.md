[TOC]

## 1 数组Array、字符串String

### 数据结构特点：

- 元素：可重复、可为null；

- 存储空间：

  - 按顺序连续存储，所以创建简单且空间利用率高；
  - 容量固定，扩容成本高：必须创建一个更大的新数组，然后将原数组的元素逐个拷贝到新数组中。

- 排序顺序：按添加顺序存储，顺序为添加顺序，序号即为索引；

- 操作的时间复杂度：

  - 遍历foreach：按插入顺序遍历，时间复杂度为O(n)；

  - 随机查找get(index)：

    按序号查找的时间复杂度为O(1)，因为本身就是按添加顺序存储；(按元素值查找的时间复杂度为O(n)，因为没有按元素值排序，数组默认也不提供该功能；)

  - 随机修改set(index, element)：

    按序号修改的时间复杂度为O(1)，因为按序号直接可以查找，然后修改元素；

  - 随机插入add(element)/add(index,element)：

    在末尾添加元素的时间复杂度为O(1)；在除末尾外的其他位置，插入元素的时间复杂度为O(n)，因为找到元素后，需要将该元素后面的所有元素整体后移1个；

  - 随机删除remove(index)/remove(element)：

    在末尾添加元素的时间复杂度为O(1)；在除末尾外的其他位置，删除的时间复杂度为O(n)，因为查找到元素后，还需要将该元素后面的所有元素整体前移1个；

> 总结：
>
> - 数组是按顺序、连续存储，容易创建且空间利用率高，是最基础的数据结构；
> - 数组是定长的，如果需要保证可增长性则需要对数组扩容，数组扩容的时间、空间消耗都是线性的；
>
> - 数组的顺序遍历消耗线性时间O(n)，随机查找/修改消耗常数时间O(1)；
>
> - 数组的插入、删除操作的成本取决于操作元素的位置，在末尾的成本很低，在其他位置则成本很高；

### Java中的实现：

##### 纯数组(定长)：Object[]  array = new Object[fixedCap];

##### 可增长的数组：ArrayList<T> arrayList = new ArrayList<T>();

数组扩展过程的源码实现：

```java
//数组扩展的实现的基础原理：
int[] arr = new int[10];
//创建新的数组，将原数组元素逐一复制过去
int oldCap = arr.length;
int newCap = oldCap + oldCap>>1;
int[] newArr = new int[newCap];
for(int i=0; i<arr.lenght; i++){
  newArr[i] = arr[i];
}
newArr = arr;
```

ArrayList (摘自openjdk12) 的扩容实现：

```java
/**
     * Returns a capacity at least as large as the given minimum capacity.
     * Returns the current capacity increased by 50% if that suffices.
     * Will not return a capacity greater than MAX_ARRAY_SIZE unless
     * the given minimum capacity is greater than MAX_ARRAY_SIZE.
     *
     * @param minCapacity the desired minimum capacity
     * @throws OutOfMemoryError if minCapacity is less than zero
     */
    private int newCapacity(int minCapacity) {
        // overflow-conscious code
        int oldCapacity = elementData.length;
        int newCapacity = oldCapacity + (oldCapacity >> 1);
        if (newCapacity - minCapacity <= 0) {
            if (elementData == DEFAULTCAPACITY_EMPTY_ELEMENTDATA)
                return Math.max(DEFAULT_CAPACITY, minCapacity);
            if (minCapacity < 0) // overflow
                throw new OutOfMemoryError();
            return minCapacity;
        }
        return (newCapacity - MAX_ARRAY_SIZE <= 0)
            ? newCapacity
            : hugeCapacity(minCapacity);
    }
```

