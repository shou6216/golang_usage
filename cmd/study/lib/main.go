package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"time"
)

const (
	c1 = iota
	c2
	c3
)

const (
	// _は使わないという意味。値は0
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func main() {
	t67 := time.Now()
	fmt.Println(t67)
	fmt.Println(t67.Format(time.RFC3339))
	fmt.Println(t67.Year(), t67.Month(), t67.Day(),
		t67.Hour(), t67.Minute(), t67.Second())

	match68, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match68)
	r68 := regexp.MustCompile("a([a-z]+)e")
	ms68 := r68.MatchString(("apple"))
	fmt.Println(ms68)

	r268 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs68 := r268.FindString("/view/test")
	fmt.Println(fs68)
	fss68 := r268.FindStringSubmatch("/view/test")
	fmt.Println(fss68, fss68[0], fss68[1], fss68[2])
	fss168 := r268.FindStringSubmatch("/edit/test")
	fmt.Println(fss168, fss168[0], fss168[1], fss168[2])
	fss268 := r268.FindStringSubmatch("/save/test")
	fmt.Println(fss268, fss268[0], fss268[1], fss268[2])

	i69 := []int{5, 4, 2, 8, 7}
	s69 := []string{"d", "a", "f"}
	p69 := []struct {
		Name string
		Age  int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 30},
		{"Bob", 50},
	}
	fmt.Println(i69, s69, p69)
	sort.Ints(i69)
	sort.Strings(s69)
	sort.Slice(p69, func(i, j int) bool {
		return p69[i].Name < p69[j].Name
	})
	fmt.Println(i69, s69, p69)

	// iota
	fmt.Println(c1, c2, c3)
	fmt.Println(KB, MB, GB)

	// context
	ch71 := make(chan string)
	// goroutineが長くて終わらない場合終了するときにcontext
	ctx71 := context.Background()
	// タイムアウトはしないけどとりあえずcontext渡して動かす
	//ctx71 := context.TODO()
	ctx71, cancel := context.WithTimeout(ctx71, 3*time.Second)
	defer cancel()
	go longProcess(ctx71, ch71)

CTXLOOP:
	for {
		select {
		case <-ctx71.Done():
			fmt.Println(ctx71.Err())
			break CTXLOOP
		case <-ch71:
			fmt.Println("success")
			break CTXLOOP
		}
	}
	fmt.Println("#############################")

	content72, err72 := ioutil.ReadFile(("README.md"))
	if err72 != nil {
		log.Fatalln(err72)
	}
	fmt.Println(string(content72))
}
