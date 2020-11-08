package main

import (
	"fmt"
	"github.com/sidav/golibrl/console"
)

const MapRenderVOffset = 2

func renderScreen() {
	console.Clear_console()
	console.SetFgColor(console.WHITE)
	console.PutString(currMode.name, 0, 0)
	console.PutString(fmt.Sprintf("Placing %s %s", terrains_names[currMode.placedTerrainIndex],
		string(currMode.getPlacedTerrain())), 0, 1)
	renderParcel()
	renderCursor()
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
			console.PutChar(currParcel.Terrain[x][y], x, y+MapRenderVOffset)
		}
	}
}
