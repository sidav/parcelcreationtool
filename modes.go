package main

import "github.com/sidav/golibrl/console"
import . "parcelcreationtool/parcel"

var(
	modes = [...]string{"Terrain", "Routes", "Items"}
	terrains      = [...]rune{'#', '.', '+', '\'', '?'}
	terrainsNames = [...]string{"Wall", "Floor", "Door", "Window", "Place for parcel"}
	terrainsColors = [...]int{console.RED, console.WHITE, console.YELLOW, console.CYAN, console.BLACK}
	savedItems = []Item{}
)

type mode struct {
	modeIndex int 
	
	placedTerrainIndex int

	placedRouteIndex int
	currWaypointIndex int

	placedItemIndex int
}

func (m *mode) switchMode() {
	m.modeIndex++
	if m.modeIndex >= len(modes) {
		m.modeIndex = 0
	}
	if len(currParcel.Routes) == 0 || len(currParcel.Routes) <= currMode.placedRouteIndex {
		currParcel.Routes = append(currParcel.Routes, Route{})
	}
}

func (m *mode) getPlacedTerrain() rune {
	return terrains[m.placedTerrainIndex]
}

func (m *mode) switchTerrain() {
	m.placedTerrainIndex++
	if m.placedTerrainIndex >= len(terrains) {
		m.placedTerrainIndex = 0
	}
}

func (m *mode) switchItem() {
	m.placedItemIndex++
	if m.placedItemIndex >= len(savedItems) {
		m.placedItemIndex = 0
	}
}

