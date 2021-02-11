package lib

type Person struct {
	// 変数も大文字にしないとPublicにならない
	Name string
	Age  int
}

func Average(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}
	return int(total / len(s))
}
