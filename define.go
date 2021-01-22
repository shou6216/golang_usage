package main

import (
	"fmt"
)

func main() {
	// Q1. 以下の1.11をint型に変換して出力してください。
	f := 1.11
	i := int(f)
	fmt.Printf("f=%v(%T) i=%v(%T)\n", f, f, i, i)

	// Q3. 以下のコードを実行した時に
	// fmt.Printf("%T %v", m, m)
	// 以下のような出力結果となるmを作成してください。
	// map[string]int map[Mike:20 Nancy:24 Messi:30]

	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Printf("%T %v", m, m)
}
