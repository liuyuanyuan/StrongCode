# 8大排序算法

参考：

- https://zhuanlan.zhihu.com/p/121122555

- https://zhuanlan.zhihu.com/p/34168443

![img](img/a1sort_algorithm.png)

### 1 冒泡排序：

基本思想：持续比较相邻的两数字，如果左边比右边大，就交换他们两个。直到没有任何一对数字需要比较。（即：每当两相邻的数比较后，发现它们的排序与排序要求相反时，就将它们互换。）

- 冒泡排序最好的时间复杂度为O(n)。冒泡排序的最坏时间复杂度为O(n^2)。因此冒泡排序总的平均时间复杂度为O(n^2)。
- 算法适用于少量数据的排序，是稳定的排序方法。

```java
public static void bubbleSort(int[] array)
{
            int tmp;
            boolean flag = false;//设置是否发生交换的标志
            for(int i=array.length-1; i>=0; i--)
            {
                for(int j=0; j<i; j++)
                {//每一轮都找到一个最大的数放在右边
                    if(array[j]>array[j+1])
                    {
                        tmp = array[j];
                        array[j] = array[j+1];
                        array[j+1] = tmp;
                        flag = true;//发生了交换
                    }
                }
                if(!flag)  
                  break;//这一轮循环没有发生交换，说明排序已经完成，退出循环
           }
}
```

### 2 快速排序

选择一个**基准元素**（通常选择第一个元素或者最后一个元素），通过一趟扫描，将待排序列分成两部分：一部分比基准元素小,一部分大于等于基准元素；此时基准元素在其排好序后的正确位置，然后再用同样的方法递归地排序划分的两部分。



### 3 选择顺序

- 基本思想：每一次从待排序的数据元素中选出最小（或最大）的一个元素，存放在序列的起始位置，直到全部待排序的数据元素排完。
- 选择排序是不稳定的排序方法。时间复杂度 O(n^2)。

```java
public static void selectSort(int[] array)
{
        for(int i=0; i<array.length-1; i++)
        {
            int min = array[i];
            int minindex = i;
            for(int j = i; j<array.length; j++)
            {
                if(array[j]<min)
                {//选择当前最小的数
                    min = array[j];
                    minindex = j;
                }
            }
            if(i != minindex) 
            {//若i不是当前元素最小的，则和找到的那个元素交换
                array[minindex] = array[i];
                array[i] = min;
            }
        }
}
```







