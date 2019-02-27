package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string{
	result := ""
	if n == 0 {
		result = "0"
		return result
	}
	for ;n >0;n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	} // strconv.Itoa(lsb)转字符型
	return result
}
func printFile(filename string)  {
	file,err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){ //go语言没有while
		fmt.Println(scanner.Text())
	}
}
func forever(){
	for {
		fmt.Println("abc")
	}
}
func main() {
	fmt.Println(convertToBin(5), //101
	convertToBin(13), //  1011
	convertToBin(72387885),
	convertToBin(0),
	)
	printFile("abc.txt")
	forever()
}
