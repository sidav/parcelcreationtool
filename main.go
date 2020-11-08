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
	currParcel = Parcel{}
	currParcel.init(10, 10)
	currMode = mode{}
	currMode.name = "Terrain placement"
	currMode.placedTerrain = WALL

	mainLoop()
}

func mainLoop() {
	for running {
		renderScreen()
		control()
	}
}
