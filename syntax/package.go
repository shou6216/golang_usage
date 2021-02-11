package syntax

import (
	"fmt"
	"golang_usage/lib"
)

func Package() {
	s60 := []int{1, 2, 3, 4, 5}
	v60 := lib.Average(s60)
	fmt.Println(v60)

	person61 := lib.Person{Name: "ほげ", Age: 30}
	fmt.Println(person61)
}
