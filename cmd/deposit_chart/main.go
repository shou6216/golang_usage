package main

import (
	"encoding/json"
	"golang_usage/internal/db"

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
	// window.WindowOptions.WebPreferences.NodeIntegration = false

	done, err := window.Start()
	if err != nil {
		panic(err)
	}

	// 	window.OpenDevTools()

	window.On(&gotron.Event{Event: "init"}, func(bin []byte) {
		sendResponse(window, "init")
	})

	window.On(&gotron.Event{Event: "register"}, func(bin []byte) {
		var request RegisterRequest
		if err := json.Unmarshal(bin, &request); err != nil {
			panic(err)
		}

		db.SaveOrUpdate(request.Params.Date, request.Params.Money)
		sendResponse(window, "register")
	})

	<-done
}

func sendResponse(window *gotron.BrowserWindow, eventName string) {
	deposits := db.FindAll()
	window.Send(&DepositResponse{
		Event:    &gotron.Event{Event: eventName},
		Deposits: deposits,
	})
}

type RegisterRequest struct {
	Event  string `json:"event"`
	Params struct {
		Date  string `json:"date"`
		Money int    `json:"money"`
	} `json:"params"`
}

type DepositResponse struct {
	*gotron.Event
	Deposits []db.Deposit `json:"deposits"`
}

// return {
// 	label: `${year}年`,
// 	data: labels.map(label => Math.floor(Math.random() * Math.floor(100000))),
// 	backgroundColor: `${rgb},0.2)`,
// 	borderColor: `${rgb},1)`,
// 	borderWidth: 1
// }
