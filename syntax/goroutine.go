package syntax

import (
	"fmt"
	"sync"
	"time"
)

func goroutine(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(100 + time.Millisecond)
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 + time.Millisecond)
		fmt.Println(s)
	}
}

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func goroutine3(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func goroutine4(ch chan string) {
	for {
		ch <- "packet from 1"
		time.Sleep(3 * time.Second)
	}
}

func goroutine5(ch chan string) {
	for {
		ch <- "packet from 2"
		time.Sleep(1 * time.Second)
	}
}

func goroutine6(s []string, c chan string) {
	sum := ""
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func producer(ch chan int, i int) {
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("process", i*1000)
		wg.Done()
	}
	fmt.Println("########################")
}

func producer1(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

// 引数の段階で矢印で送受信を指定可能
func multi2(first <-chan int, second chan<- int) {
	defer close(second)
	for i := range first {
		second <- i * 2
	}
}

func multi4(second <-chan int, third chan<- int) {
	defer close(third)
	for i := range second {
		third <- i * 4
	}
}

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func Goroutine() {
	var wg sync.WaitGroup
	// 1つのgroutineがあることを設定する
	wg.Add(1)

	// mainスレッドが終了すると、groutineも処理されない
	// WaitGroupのポイントを渡す
	go goroutine("world", &wg)
	normal("hello")

	// wg.Done()が呼ばれるのを待つ
	// wg.Done()が呼ばれないとエラーになる
	wg.Wait()

	// チャネルはgoroutineでデータの受け渡しをするもの
	s50 := []int{1, 2, 3, 4, 5}
	c50 := make(chan int)
	go goroutine1(s50, c50)
	go goroutine2(s50, c50)
	x50 := <-c50
	fmt.Println(x50)
	y50 := <-c50
	fmt.Println(y50)

	// Buffered channels
	ch51 := make(chan int, 2)
	ch51 <- 100
	fmt.Println(len(ch51))
	ch51 <- 200
	fmt.Println(len(ch51))
	// 長さが2なので以下を実行しても追加されない
	//ch52 <- 300
	//fmt.Println(len(ch52))
	close(ch51)

	// forで単純に回すと3つめでエラーになる
	// 事前にcloseする必要がある
	for c := range ch51 {
		fmt.Println(c)
	}

	s52 := []int{1, 2, 3, 4, 5}
	c52 := make(chan int, len(s52))
	go goroutine3(s52, c52)
	for i := range c52 {
		fmt.Println(i)
	}

	var wg53 sync.WaitGroup
	ch53 := make(chan int)

	// Producer
	for i := 0; i < 10; i++ {
		// Doneをする数をaddする
		wg53.Add(1)
		go producer(ch53, i)
	}

	// Consumer
	go consumer(ch53, &wg53)
	// 10回Doneされるのを待つ
	wg53.Wait()
	// consumer内でchannelがrangeで取得しようとするため
	// ここでcloseしてあげる
	close(ch53)
	time.Sleep(2 * time.Second)
	fmt.Println("Done")

	first54 := make(chan int)
	second54 := make(chan int)
	third54 := make(chan int)

	go producer1(first54)
	go multi2(first54, second54)
	go multi4(second54, third54)
	for result := range third54 {
		fmt.Println(result)
	}

	// ブロッキングしないでチャネルからデータ取得
	// selectを使う
	c155 := make(chan string)
	c255 := make(chan string)
	go goroutine4(c155)
	go goroutine5(c255)
	for {
		select {
		case msg1 := <-c155:
			fmt.Println(msg1)
		case msg2 := <-c255:
			fmt.Println(msg2)
		}
		break
	}

	tick56 := time.Tick(100 * time.Microsecond)
	boom56 := time.After(500 * time.Microsecond)
OuterLoop:
	for {
		select {
		case t := <-tick56:
			fmt.Println("tick.", t)
		case <-boom56:
			fmt.Println("BOOM!")
			// ここで普通にbreakしても処理は終わらない
			// loopを指定する
			break OuterLoop

		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Microsecond)
		}
	}
	fmt.Println("##########################")

	// Mutex
	c57 := Counter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 10; i++ {
			c57.Inc("key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c57.Inc("key")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c57.Value("key"))

	words58 := []string{"test1!", "test2!", "test3!", "test4!"}
	c58 := make(chan string)
	go goroutine6(words58, c58)
	for w := range c58 {
		fmt.Println(w)
	}
}
