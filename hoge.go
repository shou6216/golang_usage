package main

import (
	"fmt"
	"os/user"
	"strconv"
	"strings"
	"time"
)

const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
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

	fmt.Println(Pi, Username, Password)

	var (
		u8  uint8     = 255
		i8  int8      = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i
	)
	fmt.Println(u8, i8, f32, c64)
	fmt.Printf("type=%T value=%v", u8, u8)

	fmt.Println(string("Hello World"[0]))

	var ss string = "Hello World"
	fmt.Println(strings.Replace(ss, "H", "X", 1))
	fmt.Println(ss)

	fmt.Println("\"")
	fmt.Println(`"`)

	var s3 string = "14"
	// _ 利用しなくてもエラーにならないようにする表記
	i3, _ := strconv.Atoi(s3)
	fmt.Printf("%T %v", i3, i3)

	var a4 [2]int
	a4[0] = 100
	a4[1] = 200
	fmt.Println(a4)
	var b4 [2]int = [2]int{300, 400}
	fmt.Println(b4)

	n5 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n5)
	fmt.Println(n5[2])
	fmt.Println(n5[2:4])
	fmt.Println(n5[:2])
	fmt.Println(n5[2:])
	fmt.Println(n5[:])

	n5 = append(n5, 1000, 2000, 3000, 4000)
	fmt.Println(n5)
}
