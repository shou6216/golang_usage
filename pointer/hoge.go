package main

import "fmt"

func main() {
	// ポインタ
	n34 := 100
	fmt.Println(n34)
	fmt.Println(&n34)

	var p34 *int = &n34
	fmt.Println(p34)
	fmt.Println(*p34)

	// 参照渡し（デリファレンス）
	one(&n34)
	fmt.Println(n34)

	// newとmake
	// 宣言時にポインタを返すかどうかの違い
	// 領域確保済み
	var p35 *int = new(int)
	fmt.Println(p35)
	// 領域未確保
	var m35 *int
	fmt.Println(m35)

	// struct
	// 変数名小文字だとpublicにアクセスできない
	v36 := Vertext{X: 1, Y: 2}
	fmt.Println(v36)
	fmt.Println(v36.X, v36.Y)
	v36.X = 100
	fmt.Println(v36.X, v36.Y)
}

func one(x *int) {
	*x = 1
}

type Vertext struct {
	X int
	Y int
}
