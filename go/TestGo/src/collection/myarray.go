package collection

import (
	"fmt"
)

func MyArray(){
	//int
    //容量为5的数组
	array := [5]int{10, 20, 30, 40, 50}
	//由初始化值的数量决定数组容量
	array2 := [...]int{10, 20, 30, 40, 50}
	array3 := [5]int{1: 10, 3: 30}
	array3[0] = 10

	//初始化数组第0和1数组元素的值为指向整数的指针
	array4 := [5]*int{0: new(int), 1: new(int)}
	*array4[0] = 1
	*array4[1] = 20

	fmt.Println(array)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array4)

     strArray()
}

func strArray(){
	 //string
	 //数组元素指向字符串数组
	var array11 [3]string
	array12 := [3]string{"red", "blue", "green"}
	array11 = array12
	fmt.Println(array11)
	fmt.Println(array12)
	
	//数组元素指向字符串指针
	var array21 [3]*string
	array22 := [3]*string{new(string), new(string), new(string)}
	*array22[0] = "Red"
	*array22[1] = "Blue"
	*array22[2] = "Green"
	
	array21 = array22
    fmt.Println(array21)
    fmt.Println(array22)	
}

func MultiArray(){
	var array1 [3][2]int
	fmt.Println(array1)
	
	array2 := [3][2]int{{1,11}, {2,22}, {3,33}}
	fmt.Println(array2)
	
	array3 := [3][2]int{1:{1,11}, 2:{3,33}}
	fmt.Println(array3)
	
	array4 := [3][2]int{1:{0:1}, 2:{1:33}}
	fmt.Println(array4)
}

func PassArray(){
	var array [2]int
	array[1] = 1
	passArray(&array)
	
	array2 := [2]int{1, 2}
	collection.PassArray(array2)
}
func passArray1(array *[2]int){
	fmt.Println( array)
}
func passArray2(array [2]int){
	fmt.Println( array)
}









