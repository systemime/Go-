package main

import "fmt"

func updateSlice(s []int){
	s[0]=100

}
//go语言切片左闭右开
func main() { // Sllce切片是对数组的view视图，本身没有数据
	arr := [...]int{0,1,2,3,4,5,6,7}
	fmt.Println(arr[2:6])
	fmt.Println(arr[:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:])
	fmt.Println("----After updateSlice(s1)-----")
	s1:=arr[2:]
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)
	fmt.Println("----After updateSlice(s2)-----")
	s2:=arr[:]
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice: ",s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 =s2[2:]
	fmt.Println(s2)

	fmt.Println("切片拓展用法：")
	arr1:=[...]int{0,1,2,3,4,5,6,7}
	s1 = arr1[2:6]
	s2 = s1[3:6]// s1中第4位后面的数字数量不足切片数量
	fmt.Println("arr1=",arr1)
	fmt.Println("s1=",s1)
	fmt.Println("s2=",s2)
	// 解释，s1以arr为底层数组切片，我们能够得到并看到的数字为2345，实际上后面还有67，
	// 当s2以s1作为底层数组进行切片，数组内字符数量不够时会从s1的底层arr中获取后面的数值，
	// 若arr再不足，将报错
	//证明
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n",
		s1,len(s1),cap(s1))
	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n",
		s2,len(s2),cap(s2))

}
