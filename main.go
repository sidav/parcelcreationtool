package main

import "github.com/sidav/golibrl/console"

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
	currParcel.init(w, h)
	currMode = mode{}
	currMode.name = "Terrain placement"
}

func mainLoop() {
	for running {
		renderScreen()
		control()
	}
}
