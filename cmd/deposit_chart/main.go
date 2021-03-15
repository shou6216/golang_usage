package main

import (
	"encoding/json"
	"fmt"

	"github.com/Equanox/gotron"
)

func main() {
	window, err := gotron.New("web")
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

	// 	window.OpenDevTools()

	window.On(&gotron.Event{Event: "register"}, func(bin []byte) {
		var event RegisterEvent
		if err := json.Unmarshal(bin, &event); err != nil {
			panic(err)
		}
		fmt.Println(event)

		// window.Send(&CustomEvent{
		// 	Event:           &gotron.Event{Event: "event-name"},
		// 	CustomAttribute: "Hello World!",
		// })
	})

	<-done
}

type RegisterEvent struct {
	Event  string `json:"event"`
	Params struct {
		Date  string `json:"date"`
		Money int    `json:"money"`
	} `json:"params"`
}
