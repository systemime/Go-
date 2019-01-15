package main
//day2  常量与枚举
import (
	"fmt"
	"math"
)
const filename1 = "123.txt"//常量也可以在包内定义
func consts(){
	/*关于常量，go语言中常量书写方式一般不大写，因为在go中首字母大写有其他含义
	日后学习中再总结*/
	const(
		filename string = "abc.txt"
		a,b=3,4
	)
	var c int
	c = int(math.Sqrt(a*a+b*b))
	//上面定义未指定a，b类型，可以是int也可以是float
	// 所以在上面不需要显示转换，被认为是float64
	fmt.Println(filename,c)
}
func enums1()  {//go没有特殊枚举关键字，一般用一组const
	const( //const必须给其赋值
		cpp = 0
		java = 1
		python = 3
		golang = 5
	)
	const(//const高级用法
		b = 1<<(10*iota) //移位，下面变量全部使用该公式，结果递增
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp,java,python,golang)
	fmt.Println(b,kb,mb,gb,tb,pb)
}
func enums2()  {
	const( //const自增属性
		cpp = iota //若直接cpp=1，则const内所有变量均为1
		//自增顺序以const组中变量书写顺序为准（0开始）
		java
		python
		golang
	)
	fmt.Println(cpp,java,python,golang)
}
func main() {
	fmt.Println(filename1)//输出全局常量
	consts()
	enums1()
	enums2()
}
