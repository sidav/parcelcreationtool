package main

import (
	"fmt"
	"github.com/sidav/golibrl/console"
)

const MapRenderVOffset = 3

func renderScreen() {
	console.Clear_console()
	console.SetFgColor(console.WHITE)
	console.PutString("Placement: " + modes[currMode.modeIndex], 0, 0)
	console.PutString(fmt.Sprintf("Placing %s %s", terrainsNames[currMode.placedTerrainIndex],
		string(currMode.getPlacedTerrain())), 0, 1)
	renderParcel()
	renderCursor()
	renderData()
	console.Flush_console()
}

func renderCursor() {
	console.SetFgColor(console.YELLOW)
	console.PutChar('X', crs.x, crs.y+MapRenderVOffset)
	if crs.isRectPlacing {
		console.SetFgColor(console.GREEN)
		console.PutChar('X', crs.origx, crs.origy+MapRenderVOffset)
	}
}

func renderParcel() {
	for x :=range(currParcel.Terrain) {
		for y := range (currParcel.Terrain[x]) {
			for i := range terrains {
				if terrains[i] == currParcel.Terrain[x][y] {
					console.SetFgColor(terrainsColors[i])
					break
				}
			}
			console.PutChar(currParcel.Terrain[x][y], x, y+MapRenderVOffset)
		}
	}
}

func renderData() {
	putStringOnRightest(fmt.Sprintf("Width: %d, height %d", len(currParcel.Terrain), len(currParcel.Terrain[0])), 0)
	putStringOnRightest(crs.lastKeypress, 5)
}

func putStringOnRightest(str string, y int) {
	w, _ := console.GetConsoleSize()
	console.PutString(str, w-len(str), y)
}
