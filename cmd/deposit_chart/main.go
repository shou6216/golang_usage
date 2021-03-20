package main

import (
	"encoding/json"
	"fmt"
	"golang_usage/internal/color"
	"golang_usage/internal/db"
	"time"

	"github.com/Equanox/gotron"
	"github.com/guregu/null"
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

	labels := createLabels()
	lineChartDataSets := make([]LineChartDataSet, 0)
	if year2deposits.IsValid() {
		iter := year2deposits.MapRange()
		colorIndex := 0
		colors := color.GetRgbas()
		for iter.Next() {
			date2deposit := make(map[string]int)
			for _, deposit := range iter.Value().Interface().([]db.Deposit) {
				date, _ := time.Parse("2006-01-02", deposit.Date)
				date2deposit[date.Format("01/02")] = deposit.Money
			}

			// label分のデータを作成
			data := make([]null.Int, len(labels))
			for i, label := range labels {
				if money, ok := date2deposit[label]; ok {
					data[i] = null.NewInt(int64(money), true)
				}
			}

			bordercolor := colors[colorIndex]
			if colorIndex < len(colors) {
				colorIndex++
			}

			lineChartDataSets = append(lineChartDataSets, LineChartDataSet{
				Label:           fmt.Sprintf("%d年", iter.Key().Int()),
				Data:            data,
				Bordercolor:     bordercolor,
				Backgroundcolor: "rgba(0,0,0,0)",
				Borderwidth:     1,
				SpanGaps:        true,
			})
		}
	}

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
	for start.Unix() < end.Unix() {
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
	Label           string     `json:"label"`
	Data            []null.Int `json:"data"`
	Backgroundcolor string     `json:"backgroundColor"`
	Bordercolor     string     `json:"borderColor"`
	Borderwidth     int        `json:"borderWidth"`
	SpanGaps        bool       `json:"spanGaps"`
}
