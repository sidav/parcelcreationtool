package main

import (
	"fmt"
	"github.com/sidav/golibrl/console"
)

const MapRenderVOffset = 3

func renderScreen() {
	console.Clear_console()
	console.SetFgColor(console.WHITE)
	renderModeData()
	renderParcel()
	renderWaypoints()
	renderCursor()
	renderData()
	console.Flush_console()
}

func renderModeData() {
	console.PutString("Placement: " + modes[currMode.modeIndex], 0, 0)
	if modes[currMode.modeIndex] == "Terrain" {
		console.PutString(fmt.Sprintf("Placing %s %s", terrainsNames[currMode.placedTerrainIndex],
			string(currMode.getPlacedTerrain())), 0, 1)
	}
	if modes[currMode.modeIndex] == "Routes" {
		console.PutString(fmt.Sprintf("Placing %dth route, %dth waypoint", len(currParcel.Routes),
			len(currParcel.Routes[currMode.placedRouteIndex].Waypoints)), 0, 1)
	}
}

func renderCursor() {
	console.SetFgColor(console.YELLOW)
	console.PutChar('X', crs.x, crs.y+MapRenderVOffset)
	if crs.isRectPlacing {
		console.SetFgColor(console.GREEN)
		console.PutChar('X', crs.origx, crs.origy+MapRenderVOffset)
	}
}

func renderWaypoints() {
	console.SetFgColor(console.YELLOW)
	if modes[currMode.modeIndex] == "Routes" {
		for i := range currParcel.Routes[currMode.placedRouteIndex].Waypoints {
			x := currParcel.Routes[currMode.placedRouteIndex].Waypoints[i].X
			y := currParcel.Routes[currMode.placedRouteIndex].Waypoints[i].Y
			console.PutString(fmt.Sprintf("%d", i), x, y+MapRenderVOffset)
		}
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
