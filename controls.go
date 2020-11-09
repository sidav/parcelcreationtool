package main

import (
	"fmt"
	"github.com/sidav/golibrl/console"
	"os"
	. "parcelcreationtool/parcel"
)

type cursor struct {
	x, y int
	origx, origy int
	isRectPlacing bool
	lastKeypress string
}

func (c *cursor) normalizeCoords() {
	if c.x < 0 {
		c.x = len(currParcel.Terrain)-1
	}
	if c.y < 0 {
		c.y = len(currParcel.Terrain[0])-1
	}
	if c.x >= len(currParcel.Terrain) {
		c.x = 0
	}
	if c.y >= len(currParcel.Terrain[0]) {
		c.y = 0
	}

}

func control() {
	key := console.ReadKey()
	crs.lastKeypress = key
	switch key {
	case "UP": crs.y--
	case "RIGHT": crs.x++
	case "DOWN": crs.y++
	case "LEFT": crs.x--
	case "ENTER": enterKeyForMode()
	case " ": currMode.switchTerrain()
	case "N": reinitNewParcel()
	case "O": openExistingParcel()
	case "S": saveParcelToFile()
	case "g": generateAndRenderSample()

	case "TAB": currMode.switchMode()
	case "ESCAPE": running = false
	}
	crs.normalizeCoords()
}

func reinitNewParcel() {
	w := inputIntValue("Input new parcel width")
	h := inputIntValue("Input new parcel height")
	inputIntValue(fmt.Sprintf("You inputed %d %d", w, h))
	if w == 0 || h == 0 {
		return
	}
	initVars(w, h)
}

func openExistingParcel() {
	name := inputStringValue("Enter file name: ")
	if name == "" {
		return
	}
	currParcel.UnmarshalFromFile("parcels/"+name+".json")
}

func enterKeyForMode() {
	// terrain mode; draw rect
	if modes[currMode.modeIndex] == "Terrain" {
		if crs.isRectPlacing {
			xfrom := crs.origx
			xto := crs.x
			if xfrom > xto {
				xfrom = crs.x
				xto = crs.origx
			}
			yfrom := crs.origy
			yto := crs.y
			if yfrom > yto {
				yfrom = crs.y
				yto = crs.origy
			}
			for x := xfrom; x <= xto; x++ {
				for y := yfrom; y <= yto; y++ {
					currParcel.Terrain[x][y] = currMode.getPlacedTerrain()
				}
			}
			crs.isRectPlacing = false
		} else {
			crs.origx = crs.x
			crs.origy = crs.y
			crs.isRectPlacing = true
		}
	}
	if modes[currMode.modeIndex] == "Routes" {
		currParcel.Routes[currMode.placedRouteIndex].AddWaypoint(&Waypoint{X: crs.x, Y: crs.y})
	}
}

func saveParcelToFile() {
	name := inputStringValue("Enter file name (blank for auto name): ")
	if name == "" {
		i := 0
		for {
			name = fmt.Sprintf("parcel_%d", i)
			_, err := os.Stat("parcels/" + name + ".json")
			if os.IsNotExist(err) {
				break
			}
			i++
		}
	}
	currParcel.MarshalToFile("parcels/" + name + ".json")
	inputStringValue("Saved as " + "parcels/" + name + ".json")
}

