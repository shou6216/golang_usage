package main

import (
	"bytes"
	"fmt"

	"github.com/Equanox/gotron"
)

func main() {
	window, err := gotron.New("webui")
	if err != nil {
		panic(err)
	}

	window.WindowOptions.Width = 1200
	window.WindowOptions.Height = 980
	window.WindowOptions.Title = "Gotron"
	// ElectronのフロントJavaScriptでrequireしているためhtmlでloadすると競合する
	// ここでフロント側のrequireを無効にする
	// ここでfalseにするとjavascript側のglobalが定義されないのでfalseはなし
	//window.WindowOptions.WebPreferences.NodeIntegration = false

	done, err := window.Start()
	if err != nil {
		panic(err)
	}

	window.OpenDevTools()

	window.On(&gotron.Event{Event: "register"}, func(bin []byte) {
		// ここに処理を書いていく
		// fmt.Println(bin)
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)

		window.Send(&CustomEvent{
			Event:           &gotron.Event{Event: "event-name"},
			CustomAttribute: "Hello World!",
		})
	})

	<-done
}

type CustomEvent struct {
	*gotron.Event
	CustomAttribute string `json:"AtrNameInFrontend"`
}
