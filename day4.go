package 日常学习
//if与switch
import (
	"fmt"
	"io/ioutil"
)
func iftext(){
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	//go函数返回值有两个
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n",contents)
	}
	//特色写法
	if contents2, err2 := ioutil.ReadFile(filename);err2 !=nil {
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n",contents2)
	}
	//fmt.Println(contents2)
	// 这句话在特色写法中会报错，因为contents2生存期只在if中存在
}
func grade(score int) string {
	g :=""
	switch {
	case score < 0 || score >100:
		panic(fmt.Sprintf("Wrong score:%d",score))
	case score<60:
		g="F"
	case score<80:
		g="C"
	case score<90:
		g="B"
	case score<=100:
		g="A"
	default://第一条存在情况，此句话无效,已包含错误情况
		panic(fmt.Sprintf("Wrong score:%d",score))
	}//default存在出错情况会直接中断整个程序，不输出正确部分
	return g
}
func main() {
	fmt.Println()
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(100),
		grade(101),
		)//若输入错误，整个程序中断，
		// 输出default中内容，同时不显示其他参数执行结果
}
