package main

import (
	"fmt"
	"os/user"
	"time"
)

// 最初に自動で呼ばれる関数
func init() {
	fmt.Println("init")
}

/*
これでコメントもいける
*/

func bazz() {
	fmt.Println("bazz world")
}

func main() {
	bazz()
	//カンマ区切りで複数文字列表示
	fmt.Println("Hello world", "hoge hoge", time.Now())
	fmt.Println(user.Current())

	var i int = 1
	var f64 float64 = 1.2
	var s string = "aaaa"
	var t bool = true
	var f bool = false
	var tt, ff = true, false
	fmt.Println(i, f64, s, t, f, tt, ff)

	// 関数内でしかショートの定義はできない
	xi := 1
	xi = 2
	xf64 := 1.2
	xs := "aaaa"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T", xf64)
}
