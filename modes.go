package main

import "github.com/sidav/golibrl/console"

var(
	modes = [...]string{"Terrain", "Routes"}
	terrains      = [...]rune{'#', '.', '+', '\''}
	terrainsNames = [...]string{"Wall", "Floor", "Door", "Window"}
	terrainsColors = [...]int{console.RED, console.WHITE, console.YELLOW, console.CYAN}
)

type mode struct {
	modeIndex int 
	
	placedTerrainIndex int

	placedRouteIndex int
	currWaypointIndex int
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
