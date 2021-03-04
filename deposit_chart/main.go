package main

import (
	"github.com/Equanox/gotron"
)

func main() {
	window, err := gotron.New()
	if err != nil {
		panic(err)
	}

	window.WindowOptions.Width = 1200
	window.WindowOptions.Height = 980
	window.WindowOptions.Title = "Gotron"

	done, err := window.Start()
	if err != nil {
		panic(err)
	}

	<-done
}
