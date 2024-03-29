package main

import "github.com/sidav/golibrl/console"
import . "github.com/sidav/parcelcreationtool/parcel"

var (
	crs cursor
	running bool
	currParcel Parcel
	currMode mode
)

func main() {
	console.Init_console("Parcel creation tool", console.TCellRenderer)
	defer console.Close_console()

	running = true
	initVars(10, 10)

	mainLoop()
}

func initVars(w, h int) {
	currParcel = Parcel{}
	currParcel.Init(w, h)
	currMode = mode{}
}

func mainLoop() {
	for running {
		renderScreen()
		control()
	}
}
