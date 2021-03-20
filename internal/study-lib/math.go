/*
custom lib
*/
package study_lib

import "fmt"

type Person struct {
	// 変数も大文字にしないとPublicにならない
	Name string
	Age  int
}

func (p *Person) Say() {
	fmt.Println(p.Name)
}

// Average returns the average of a series of numbers
func Average(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}
	return int(total / len(s))
}
