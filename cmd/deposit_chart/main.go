package main

import (
	"encoding/json"
	"fmt"
	"golang_usage/internal/db"
	"time"

	"github.com/Equanox/gotron"
	"github.com/wesovilabs/koazee"
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
	// 年別に集計
	year2deposits, _ := koazee.StreamOf(deposits).GroupBy(func(deposit db.Deposit) int {
		date, _ := time.Parse("2006-01-02", deposit.Date)
		return date.Year()
	})

	lineChartDataSets := make([]LineChartDataSet, 0)
	if year2deposits.IsValid() {
		iter := year2deposits.MapRange()
		for iter.Next() {
			lineChartDataSets = append(lineChartDataSets, LineChartDataSet{
				Label:           fmt.Sprintf("%d年", iter.Key().Int()),
				Data:            []int{1, 2, 3, 4, 5},
				Backgroundcolor: []int{1, 2, 3, 4, 5},
				Bordercolor:     []int{1, 2, 3, 4, 5},
				Borderwidth:     1,
			})
		}
	}

	labels := createLabels()
	window.Send(&DepositResponse{
		Event: &gotron.Event{Event: eventName},
		LineChartData: struct {
			Labels   []string           `json:"labels"`
			DataSets []LineChartDataSet `json:"datasets"`
		}{
			labels,
			lineChartDataSets,
		},
	})
}

func createLabels() []string {
	// 閏年のある年で1年間の日付を作成
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	end := start.AddDate(1, 0, 0)

	labels := make([]string, 0)
	for start.Unix() <= end.Unix() {
		labels = append(labels, start.Format("01/02"))
		start = start.AddDate(0, 0, 1)
	}

	return labels
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
	LineChartData struct {
		Labels   []string           `json:"labels"`
		DataSets []LineChartDataSet `json:"datasets"`
	} `json:"lineChartData"`
}

type LineChartDataSet struct {
	Label           string `json:"label"`
	Data            []int  `json:"data"`
	Backgroundcolor []int  `json:"backgroundColor"`
	Bordercolor     []int  `json:"borderColor"`
	Borderwidth     int    `json:"borderWidth"`
}
