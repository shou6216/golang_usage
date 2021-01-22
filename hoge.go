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

	n6 := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n6), cap(n6), n6)
	n6 = append(n6, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n6), cap(n6), n6)
	n6 = append(n6, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n6), cap(n6), n6)

	n61 := make([]int, 0)
	var n62 []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(n61), cap(n61), n61)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n62), cap(n62), n62)

	n63 := make([]int, 5)
	n64 := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		n63 = append(n63, i)
		n64 = append(n64, i)
		fmt.Println(n63)
		fmt.Println(n64)
	}
	fmt.Println(n63)
	fmt.Println(n64)

	n7 := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(n7)
	fmt.Println(n7["apple"])

	n71, ok1 := n7["apple"]
	fmt.Println(n71, ok1)

	n72, ok2 := n7["nothing"]
	fmt.Println(n72, ok2)

	n8 := []byte{72, 73}
	fmt.Println(n8)
	fmt.Println(string(n8))

	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	f19 := func(x int) {
		fmt.Println("inner func", x)
	}
	f19(1)

	func(x int) {
		fmt.Println("inner func2", x)
	}(2)

	counter20 := incrementGenerator()
	fmt.Println(counter20())
	fmt.Println(counter20())
	fmt.Println(counter20())
	fmt.Println(counter20())

	c201 := circleArea(3.14)
	fmt.Println(c201(2))

	c202 := circleArea(3)
	fmt.Println(c202(2))

	foo(10, 20)
	foo(10, 20, 30)

	s21 := []int{1, 2, 3}
	foo(s21...)
}

func add(x, y int) (int, int) {
	return x + y, x - y
}

func cal(price, item int) (result int) {
	result = price * item
	return
}

func converter(price int) float64 {
	return float64(price)
}

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}
