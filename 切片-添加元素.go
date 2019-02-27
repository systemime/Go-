package main

import "fmt"

func printSlice(s []int){
	fmt.Printf("value = %v,len = %d,cap= %d \n",s,len(s),cap(s))
}// cap字节长

func main() {
	arr := [...]int{0,1,2,3,4,5,6,7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println("arr = ",arr)
	fmt.Println("s1 = ",s1)
	fmt.Println("s2 = ",s2)

	s3 := append(s2,10) //必须接受append返回值
	s4 := append(s3,11)
	s5 := append(s4,12)
	fmt.Println("s3 = ",s3)
	fmt.Println("s4 = ",s4)
	fmt.Println("s5 = ",s5)
	fmt.Println("After arr = ",arr)
	// 因为在slice的cap指针影响下，s2可以影响到arr的第六位
	// 第六位即7，在此情况下，s2最后一位是10，
	// 由于指针的关系，arr第六位变成10
	// 视觉上s2是增加了一个10，实际上是改变了cap指针指向的值
	// 添加元素时如果超越cap，系统会重新分配更大的底层数组

	fmt.Println("---------------------------------------------")
	fmt.Println("Creat Slice: ")
	var s []int // Zero value for slice is nil
				// s为0值
	for i:=0;i<100;i++{
		printSlice(s)
		s=append(s,2*i+1)
	}
	fmt.Println(s)

	s1 = []int{2,4,6,8}
	printSlice(s1)
	s2 = make([]int,16) //建立一个长度为16的slice
	s3 = make([]int,10,32) //10个位，但是长度申请32
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice: ")
	copy(s2,s1) // 拷贝后s1内容在s2前面
	printSlice(s2)

	fmt.Println("Deleting elements from slice: ")
	s2 = append(s2[:3],s2[4:]...)
	//append第二位置是可变参数
	// 想放一个数组需要后面加...
	printSlice(s2)

	fmt.Println("Popping from front: ")//删除头尾
	front := s2[0]
	s2 = s2[1:]
	printSlice(s2)      
	fmt.Println("Popping from back: ")
	tail :=  s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(front,tail)
	printSlice(s2)
}
