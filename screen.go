package main

import (
	"fmt"
	"github.com/sidav/golibrl/console"
	"strconv"
)

const MapRenderVOffset = 3

func renderScreen() {
	console.Clear_console()
	console.SetFgColor(console.WHITE)
	renderModeData()
	renderParcel()
	renderItems()
	renderWaypoints()
	renderCursor()
	renderData()
	console.Flush_console()
}

func renderModeData() {
	console.PutString("Placement: "+modes[currMode.modeIndex], 0, 0)
	if modes[currMode.modeIndex] == "Terrain" {
		console.PutString(fmt.Sprintf("Placing %s %s", terrainsNames[currMode.placedTerrainIndex],
			string(currMode.getPlacedTerrain())), 0, 1)
	}
	if modes[currMode.modeIndex] == "Routes" {
		waypointsNum := 0
		if len(currParcel.Routes) > currMode.placedRouteIndex {
			waypointsNum = len(currParcel.Routes[currMode.placedRouteIndex].Waypoints)
		}
		console.PutString(fmt.Sprintf("Placing route %d (of %d), %dth waypoint", currMode.placedRouteIndex+1,
			len(currParcel.Routes), waypointsNum), 0, 1)
	}
	if modes[currMode.modeIndex] == "Items" {
		if len(savedItems) > 0 {
			console.PutString(fmt.Sprintf("Placing %s %s",
				savedItems[currMode.placedItemIndex].Name,
				string(savedItems[currMode.placedItemIndex].DisplayedChar),
			), 0, 1)
		} else {
			console.PutString("No items created...", 0, 1)
		}
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
	if modes[currMode.modeIndex] == "Routes" && len(currParcel.Routes) > currMode.placedRouteIndex {
		for i := range currParcel.Routes[currMode.placedRouteIndex].Waypoints {
			x := currParcel.Routes[currMode.placedRouteIndex].Waypoints[i].X
			y := currParcel.Routes[currMode.placedRouteIndex].Waypoints[i].Y
			outputSymbol := strconv.Itoa(i)
			if len(outputSymbol) > 1 {
				outputSymbol = string(rune(int('a') + i - 10))
			}
			console.PutString(outputSymbol, x, y+MapRenderVOffset)
		}
	}
}

func renderItems() {
	console.SetFgColor(console.RED)
	// if modes[currMode.modeIndex] == "Items" {
	for _, i := range currParcel.Items {
		console.PutChar(i.DisplayedChar, i.X, i.Y+MapRenderVOffset)
	}
	// }
}

func renderParcel() {
	for x := range currParcel.Terrain {
		for y := range currParcel.Terrain[x] {
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
	putStringOnRightest(fmt.Sprintf("Whole parcel width: %d, height %d", len(currParcel.Terrain), len(currParcel.Terrain[0])), 0)
	if crs.isRectPlacing {
		w, h := crs.getRectSize()
		putStringOnRightest(fmt.Sprintf("Curr rect width: %d, height %d", w, h), 1)
	}
	putStringOnRightest(crs.lastKeypress, 5)
}

func putStringOnRightest(str string, y int) {
	w, _ := console.GetConsoleSize()
	console.PutString(str, w-len(str), y)
}
