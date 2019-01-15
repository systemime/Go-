package main

import (
	"fmt"
	"io/ioutil"
)
func iftext(){
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	//go返回值有两个
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
	// 这句话会报错，因为contents2生存期只在if中存在
}
func grade(score int) string {
	g :=""
	switch {
	case score<60:
		g="F"
	case score<80:
		g="C"
	case score<90:
		g="B"
	case score<=100:
		g="A"
	default:
		panic(fmt.Sprintf("Wrong score:%d",score))
	}
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
		)//若输入错误，整个程序中断，
		// 输出default中内容，同时不显示其他参数执行结果
}
