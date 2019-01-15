package main
import (
	"fmt"
	"math"
	"math/cmplx"
)
var aa = 3
var ss="kkk" //没有全局变量说法，只是包变量
//bb:=true //不能再全局用:=，必须用var
var (//省去多次书写var的麻烦
	aaa=3
	sss="kkk"
	bb=true
)
/*go语言必须使用申明的变量*/
func variable()  {
	var a int
	var s string
	fmt.Printf("%d %q\n",a,s)
}
func variableInitialValue(){
	var a,b int = 3,4
	var s string = "abc"
	fmt.Printf("第一次：%d %q %d\n",a,s,b)
	fmt.Println(a,s,b)
}
func variableType(){
	var a,b,c,d = 3,4,true,"def"
	fmt.Println(a,b,c,d)
}
func variableShort(){
	a,b,c,d := 3,4,true,"def"
	//e:=5 //再使用会报错,只能第一次使用
	fmt.Println(a,b,c,d)
}
func euler(){
	c:=3+4i //复数类型不能用4*i，会误以为i是变量
	fmt.Println(cmplx.Abs(c),)
	fmt.Printf("%.3f\n",cmplx.Exp(1i * math.Pi)+1) //1i为了让系统知道i是个复数，不是变量
}
func triangle(){
	a,b:=3,4
	var c int //go语言只能强制类型转换，没有隐式的，所以int（）之类不能省
	c = int(math.Sqrt(float64(a*a+b*b)))
	fmt.Println(c)
	//由于float有时候不准，所以这条式子有可能是4
	//该问题存在于任何编程语言中
}
func main() {
	fmt.Println("hello world")
	variable()
	variableInitialValue()
	variableType()
	variableShort()
	fmt.Println(aa,ss,aaa,sss,bb)
	euler()
	triangle()
}
