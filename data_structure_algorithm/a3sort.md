# 8大排序算法

[TOC]

## 什么是排序问题

排序，就是让一组无序数据变成有序的过程。 一般默认的有序都是从小到大的排列顺序。

衡量一个排序算法的优劣，主要从3 个角度分析：

- **时间复杂度**：具体包括，最好时间复杂度、最坏时间复杂度、平均时间复杂度。
- **空间复杂度**：如果空间复杂度为 1，也叫作原地排序。
- **稳定性**：排序的稳定性是指相等的数据对象，在排序之后，顺序是否能保证不变。



## 排序算法的选择

各种排序算法之间没有绝对的好坏，都是也有利弊的，所以要根据场景选择。

- **如果对数据规模比较小的数据进行排序，可以选择时间复杂度为 O(n\*n) 的排序算法**。因为当数据规模小的时候，时间复杂度 O(nlogn) 和 O(n*n) 的区别很小，它们之间仅仅相差几十毫秒，因此对实际的性能影响并不大。

- **但对数据规模比较大的数据进行排序，就需要选择时间复杂度为 O(nlogn) 的排序算法了**
  - 归并排序的空间复杂度为 O(n)，也就意味着当排序 100M 的数据，就需要 200M 的空间，所以对空间资源消耗会很多。
  - 快速排序在平均时间复杂度为 O(nlogn)，但是如果分区点选择不好的话，最坏的时间复杂度也有可能逼近 O(n*n)。而且快速排序不具备稳定性，这也需要看你所面对的问题是否有稳定性的需求。



## 常见的排序算法及其思想

![img](img/a1sort_algorithm.png)

### 冒泡排序：

基本思想：持续比较相邻的两个数字，如果左边比右边大，就交换他们两个，否则保持不变。通过多轮迭代，直到没有交换操作为止。（即每当两相邻的数比较后，如果它们的排序与要求相反时，就互换它们。）

#### 性能

- **冒泡排序最好的时间复杂度为O(n)：**也就是当输入数组刚好是顺序的时候，只需要挨个比较一遍就行了，不需要做交换操作，所以时间复杂度为 O(n)；

  **冒泡排序的最坏时间复杂度为O(n^2)**：也就是说当数组刚好是完全逆序的时候，每轮排序都需要挨个比较 n 次，并且重复 n 次，所以时间复杂度为 O(n*n)。

  **因此冒泡排序总的平均时间复杂度为O(n^2)。**

- **冒泡排序不需要额外的空间，所以空间复杂度是 O(1)。**

- 冒泡排序过程中，当元素相同时不做交换，所以**冒泡排序是稳定的排序算法。**

- 算法适用于少量数据的排序。

  <img src="img/a1sort_bublbe.png" alt="img" style="zoom:25%;" />

```java
public static void bubbleSort(int[] arr){
  System.out.println("原始数据: " + Arrays.toString(arr));
	for (int i = 1; i < arr.length; i++) {
     // 每次冒泡后，最后一个值都是本次的最大值。所以一次冒泡的比较次数是(arr.length - i)
		for (int j = 0; j < arr.length - i; j++) {
			if (arr[j] > arr[j + 1]) {
				int temp = arr[j];
				arr[j] = arr[j + 1];
				arr[j + 1] = temp;
			} 
		}
	}
	System.out.println("冒泡排序: " + Arrays.toString(arr));
}
```



### 插入排序

选取未排序的元素，插入到已排序区间的合适位置，直到未排序区间为空。插入排序顾名思义，就是从左到右维护一个已经排好序的序列。直到所有的待排数据全都完成插入动作。

#### 性能

- **插入排序最好时间复杂度是 O(n)**：即当数组刚好是完全顺序时，每次只用比较一次就能找到正确的位置。这个过程重复 n 次，就可以清空未排序区间。

  **插入排序最坏时间复杂度则需要 O(n^2)**：即当数组刚好是完全逆序时，每次都要比较 n 次才能找到正确位置。这个过程重复 n 次，就可以清空未排序区间，所以最坏时间复杂度为 O(n^2)。

  **插入排序的平均时间复杂度是 O(n^2)**：这是因为往数组中插入一个元素的平均时间复杂度为 O(n)，而插入排序可以理解为重复 n 次的数组插入操作，所以平均时间复杂度为 O(n*n)。

- 插入排序不需要开辟额外的空间，**所以空间复杂度是 O(1)**。

- **插入排序是稳定的排序算法**。

<img src="img/a1sort_insert.png" alt="动画2.gif" style="zoom:25%;" />

```java
public static void insertSort(int[] arr){
  System.out.println("原始数据: " + Arrays.toString(arr));
	for (int i = 1; i < arr.length; i++) {
		int temp = arr[i];
    // 用第i（属于[1, length-1]）个元素，与排好序的[0, i<1]数组从后往前比较。
    int j = i - 1;
		for (; j >= 0; j--) {
			if (arr[j] <= temp) {
				break;
			} else {
				arr[j + 1] = arr[j]; // 将大于的元素后移一位；
			}
		}
		arr[j + 1] = temp; // 直到所有大于的元素移动完毕，插入该元素；
	}
	System.out.println("插入排序: " + Arrays.toString(arr));	
}
```





### 归并排序

归并排序的原理就是分治法。

它首先将数组不断地二分，直到最后每个部分只包含 1 个数据。然后再对每个部分分别进行排序，最后将排序好的相邻的两部分合并在一起，这样整个数组就有序了。

#### 性能

- 对于归并排序，它采用了二分的迭代方式，时间复杂度是O(logn)。

  每次的迭代，需要对两个有序数组进行合并，这样的动作在 O(n) 的时间复杂度下就可以完成。因此，归并排序的复杂度就是二者的乘积 O(nlogn)。同时，它的执行频次与输入序列无关，因此，**归并排序最好、最坏、平均时间复杂度都是 O(nlogn)**。

- 空间复杂度方面，由于每次合并的操作都需要开辟基于数组的临时内存空间，**所以空间复杂度为 O(n)。**
- 归并排序合并的时候，相同元素的前后顺序不变，所以**归并是稳定的排序算法**。

### 实现

<img src="img/a1sort_merge.png" alt="动画3.gif" style="zoom:25%;" />

```java
public static void main(String[] args) {
    int[] arr = { 49, 38, 65, 97, 76, 13, 27, 50 };
    int[] tmp = new int[arr.length];
    System.out.println("原始数据: " + Arrays.toString(arr));
    customMergeSort(arr, tmp, 0, arr.length - 1);
    System.out.println("归并排序: " + Arrays.toString(arr));
}
public static void customMergeSort(int[] a, int[] tmp, int start, int end) {
    if (start < end) {
        int mid = (start + end) / 2;
        // 对左侧子序列进行递归排序
        customMergeSort(a, tmp, start, mid);
        // 对右侧子序列进行递归排序
        customMergeSort(a, tmp,mid + 1, end);
        // 合并
        customDoubleMerge(a, tmp, start, mid, end);
    }
}
public static void customDoubleMerge(int[] a, int[] tmp, int left, int mid, int right) {
    int p1 = left, p2 = mid + 1, k = left;
    while (p1 <= mid && p2 <= right) {
        if (a[p1] <= a[p2])
            tmp[k++] = a[p1++];
        else
            tmp[k++] = a[p2++];
    }
    while (p1 <= mid)
        tmp[k++] = a[p1++];
    while (p2 <= right)
        tmp[k++] = a[p2++];
    // 复制回原素组
    for (int i = left; i <= right; i++)
        a[i] = tmp[i];
}
```





### 快速排序

快速排序法的原理也是分治法。

它的每轮迭代，会选取一个**基准元素**（通常选择第一个元素或者最后一个元素），通过一趟扫描，将小于它的元素放在其左侧，大于它的放在其右侧。再利用分治思想，继续分别对左右两侧进行同样的操作，直至每个区间缩小为 1，则完成排序。

#### 性能

- **在快排的最好时间复杂度下**，如果每次选取分区点时，都能选中中位数，把数组等分成两个，那么**此时的时间复杂度和归并一样，都是 O(n\*logn)**。
- **而在最坏的时间复杂度下**，也就是如果每次分区都选中了最小值或最大值，得到不均等的两组。那么就需要 n 次的分区操作，每次分区平均扫描 n / 2 个元素，**此时时间复杂度就退化为 O(n\*n) 了**。
- 快速排序法在大部分情况下，统计上是很难选到极端情况的。**因此它平均的时间复杂度是 O(n\*logn)**。

- **快速排序法的空间方面，使用了交换法，因此空间复杂度为 O(1)**。

- 很显然，快速排序的分区过程涉及交换操作，所以**快排是不稳定的排序算法**。

### 实现

<img src="img/a1sort_quick.png" alt="动画4.gif" style="zoom:35%;" />

```java
public static void main(String[] args) {
	int[] arr = { 6, 1, 2, 7, 9, 11, 4, 5, 10, 8 };
	System.out.println("原始数据: " + Arrays.toString(arr));
	customQuickSort(arr, 0, arr.length - 1);
	System.out.println("快速排序: " + Arrays.toString(arr));
}

public static void customQuickSort(int[] arr, int low, int high) {
        if (low >= high) {
            return;
        }
        int one = arr[low]; //基准元素
        int i = low;
        int j = high;
        // 以基准元素为准，从两端开始比较，将小的左移，大的右移，最终基准元素左侧
        while (i < j) {
            // 先看右边，依次往左比较
            while (one <= arr[j] && i < j) {
                j--;
            }
            // 再看左边，依次往右比较
            while (one >= arr[i] && i < j) {
                i++;
            }
            // 将违背顺序的两个元素交换位置
            int temp = arr[j];
            arr[j] = arr[i];
            arr[i] = temp;
        }
        // 将违背顺序的元素与基准元素互换
        arr[low] = arr[i];
        arr[i] = one;
        // 递归调用左半数组
        customQuickSort(arr, low, j - 1);
        // 递归调用右半数组
        customQuickSort(arr, j + 1, high);
}
```



### 选择顺序

基本思想：每一次从待排序的数据元素中选出最小（或最大）的一个元素，存放在序列的起始位置，直到全部待排序的数据元素排完。

- **时间复杂度 O(n^2)。**
- 选择排序**是不稳定的排序算法。**

```java
public static void selectSort(int[] array){
        for(int i=0; i<array.length-1; i++){
            int min = array[i];
            int minindex = i;
            for(int j = i; j<array.length; j++){
                if(array[j]<min){
                    //选择当前最小的数
                    min = array[j];
                    minindex = j;
                }
            }
            if(i != minindex){
                //若i不是当前元素最小的，则和找到的那个元素交换
                array[minindex] = array[i];
                array[i] = min;
            }
        }
}
```


