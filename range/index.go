package main

import (
	"fmt"
)

func main() {
	list := []string{"a", "b", "c", "d"}
	for k, v := range list {
		fmt.Printf("第%v位是：%v \n", k, v)
	}
}
