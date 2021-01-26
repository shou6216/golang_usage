package main

import (
	"fmt"
)

func main() {
	// for
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)

	// 無限ループ
	//for {
	//
	//}
}
