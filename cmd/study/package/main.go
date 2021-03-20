package main

import (
	"fmt"
	lib "golang_usage/internal/study-lib"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
	// talib "github.com/markcheno/go-talib" import名を定義できる
)

func main() {
	s60 := []int{1, 2, 3, 4, 5}
	v60 := lib.Average(s60)
	fmt.Println(v60)

	person61 := lib.Person{Name: "ほげ", Age: 30}
	fmt.Println(person61)

	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
}
