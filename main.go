package main

import (
	"fmt"
	"golang_usage/syntax"
)

func main() {
	fmt.Println("##### Define #####")
	syntax.Define()
	fmt.Println("##### Statement #####")
	syntax.Statement()
	fmt.Println("##### Struct #####")
	syntax.Struct()
	fmt.Println("##### Pointer #####")
	syntax.Pointer()
	fmt.Println("##### Goroutine #####")
	syntax.Goroutine()
}
