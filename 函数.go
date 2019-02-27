package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)
// 函数中两个返回值指名对调用者无意义，不建议使用
//仅仅用于简单函数
func eval(a,b int , op string) (int, error) {
	switch op{
	case "+":
		return a+b,nil //每一个nil对应error返回值
	case "-":
		return a-b,nil
	case "*":
		return a*b,nil
	case "/":
		q,_ :=div(a,b)
		return q,nil
	default:
		return 0, fmt.Errorf(
			"unsupported open: %s", op)
	}
}
//go语言函数式编程
// func div(a,b int)(int int){
// 		q = a/b
//		r = a%b
//`		return
// }
func div(a,b int)(q,r int){
	return a/b,a%b
}
func apply(op func(int, int) int,a,b int) int{
	p := reflect.ValueOf(op).Pointer() //reflect获得函数真实的值，然后.Pointer()获取函数真实指针
	opName := runtime.FuncForPC(p).Name() // 获取函数名称
	fmt.Printf("Calling Function %s with args "+
		"(%d, %d)", opName,a, b)
	return op(a,b)
}
func pow(a,b int) int{ //复写math.Pow，但是Go中没有函数重载
	return int(math.Pow(float64(a),float64(b)))
}
//...表示可变参数数量
func sun(numbers ...int) int{ //...随便参数数量
	s:=0
	for i:=range numbers{
		s+=numbers[i]
	}
	return s
}

func swap(a,b *int){
	*b,*a=*a,*b
}

func swap2(a,b int)(int, int){
	return b,a
}

func main() {
	fmt.Println(eval(4,2,"/"))
	if result, err := eval(3,4,"x");err!=nil{
		fmt.Println("Error:",err)
	} else {
		fmt.Println(result)
	}

	q,r := div(13,3)
	fmt.Println(q,r)

	fmt.Println(apply(pow,3,4))
	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(
				float64(a),float64(b)))
		},3,4))//匿名函数调用内部函数为参数对应apply函数中的op函数

	// go语言可变参数列表，没有默认参数，函数重载等花哨方法
	fmt.Println(sun(1,2,3,4,5))

	// go语言指针不能运算
	// 只有值类型传递，不能引用传递
	// swap2更好
	a,b:=3,4
	swap(&a,&b)
	fmt.Println(a,b)
	fmt.Println(swap2(a,b))
	}
