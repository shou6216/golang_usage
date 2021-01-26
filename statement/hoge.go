package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	// for
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)

	// 無限ループ
	//for {
	//
	//}

	// range
	l26 := []string{"python", "go", "java"}
	for i, v := range l26 {
		fmt.Println(i, v)
	}
	for _, v := range l26 {
		fmt.Println(v)
	}

	m26 := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m26 {
		fmt.Println(k, v)
	}
	for k := range m26 {
		fmt.Println(k)
	}
	for _, v := range m26 {
		fmt.Println(v)
	}

	// switch
	// defaultなくていい
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac!!!")
	case "windows":
		fmt.Println("Windows!!")
	}

	// 評価式なくていい
	t27 := time.Now()
	switch {
	case t27.Hour() < 12:
		fmt.Println("morning")
	case t27.Hour() < 17:
		fmt.Println("afternoon")
	}

	// defer
	// deferを定義した関数が終わった後に実行する
	defer fmt.Println("world")
	fmt.Println("hello")

	// stacking defer 最初に定義したものが最後に実行
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println("hello2")

	file28, _ := os.Open("./hoge.go")
	defer file28.Close()
	data28 := make([]byte, 100)
	file28.Read(data28)
	fmt.Println(string(data28))

	log.Println("logging!")
	log.Printf("%T %v", "test", "test")
	// log.Fatalfを実行するとコードは終了する
	// _, err := os.Open("adniosandiaonia")
	// if err != nil {
	// 	log.Fatalln("Exit", err)
	// }

	save()
	fmt.Println("OK!")

	// Q1 . 以下のスライスから一番小さい数を探して
	// 出力するコードを書いてください。
	l32 := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

	min32 := l32[0]
	for _, v := range l32 {
		if v < min32 {
			min32 = v
		}
	}
	fmt.Println("min=", min32)

	// Q2. 以下の果物の価格の合計を出力するコードを書いてください。

	m32 := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	sum32 := 0
	for _, v := range m32 {
		sum32 += v
	}
	fmt.Println("sum=", sum32)
}

func getOsName() string {
	return "mac"
}

func loggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// 標準出力とログファイルに出力
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

func thirdPartyConnectDB() {
	panic("Unable to connect database!")
}
