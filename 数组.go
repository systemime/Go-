package main

import "fmt"
// 数组是值类型，证明：
func printArray(arr [5]int){
	arr[0] = 100
	for i,v :=range arr{
		fmt.Println(i,v)
	}
}
//go语言一般不直接使用数组，使用切片
// 改变数组内容的方法——指针
// 此处指针的方式可以改写成切片 func printArray2(arr []int) 参数长度也不限制，函数调用改为arr1[:]或其他切片形式
func printArray2(arr *[5]int){
	arr[0] = 100
	for i,v :=range arr{
		fmt.Println(i,v)
	}
}
func main() {
	var arr1 [5]int
	arr2 := [3]int{1,3,5}
	arr3 := [...]int{2,4,6,8,10}
	fmt.Println(arr1,arr2,arr3)

	var grid [4][5]int
	fmt.Println(grid)
	// 方案一
	for i:=0;i<len(arr3);i++{
		fmt.Println(arr3[i])
	}
	// 方案二
	for i:=range arr3{
		fmt.Println(arr3[i])
	} // i为下标
	// 方案三
	for i,v :=range arr3{ // for _,v :=range arr3{
		fmt.Println(i,v) // fmt.Println(v)
	}// i=下标、v=值

	printArray(arr1)
	printArray(arr3)
	fmt.Println(arr1,arr3) //输出并没有变，证明数组是值类型
	// printArray(arr2) //函数中定义了参数长度为5，arr2为3
	// go语言中[10]int和[20]int是不同类型
	// 没有C语言中数组名就是头指针一说

	//现在指针是改变后的数组
	printArray2(&arr1)
	printArray2(&arr3)
	fmt.Println(arr1,arr3)
}
