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
		return a+b,nil
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
func div(a,b int)(q,r int){
	return a/b,a%b
}
func apply(op func(int, int) int,a,b int) int{
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling Function %s with args "+
		"(%d, %d)", opName,a, b)
	return op(a,b)
}
func pow(a,b int) int{ //复写math.Pow
	return int(math.Pow(float64(a),float64(b)))
}
//...表示可变参数数量
func sun(numbers ...int) int{
	s:=0
	for i:=range numbers{
		s+=numbers[i]
	}
	return s
}

func main() {
	fmt.Println(eval(4,2,"/"))
	if result, err := eval(3,4,"x");err!=nil{
		fmt.Println("error:",err)
	} else {
		fmt.Println(result)
	}

	q,r := div(13,3)
	fmt.Println(q,r)

	fmt.Println(apply(pow,3,4))
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a),float64(b)))
	},3,4))//匿名函数调用
	// go语言可变参数列表，没有默认参数，函数重载等花哨方法
	fmt.Println(sun(1,2,3,4,5))

	}
