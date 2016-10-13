package main

import (
	"fmt"
)

func main() {
	//常量声明
	const NUM int = 10
	// 平行声明变量并赋值
	num1, num2, num3 := 1, 2, 3
	//变量重新赋值
	num3 = 4
	fmt.Print(num1, num2, num3, NUM)
}
