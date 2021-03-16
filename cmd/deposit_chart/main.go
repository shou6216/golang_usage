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
		deposits := db.FindAll()
		window.Send(&RegisterResponse{
			Event:    &gotron.Event{Event: "init"},
			Deposits: deposits,
		})
	})

	window.On(&gotron.Event{Event: "register"}, func(bin []byte) {
		var request RegisterRequest
		if err := json.Unmarshal(bin, &request); err != nil {
			panic(err)
		}

		db.SaveOrUpdate(request.Params.Date, request.Params.Money)

		deposits := db.FindAll()
		window.Send(&RegisterResponse{
			Event:    &gotron.Event{Event: "register"},
			Deposits: deposits,
		})
	})

	<-done
}

type RegisterRequest struct {
	Event  string `json:"event"`
	Params struct {
		Date  string `json:"date"`
		Money int    `json:"money"`
	} `json:"params"`
}

type RegisterResponse struct {
	*gotron.Event `json:"event"`
	Deposits      []db.Deposit `json:"deposits"`
}
