package main

import "github.com/sidav/golibrl/console"

var(
	terrains      = [...]rune{'#', '.', '+', '\''}
	terrainsNames = [...]string{"Wall", "Floor", "Door", "Window"}
	terrainsColors = [...]int{console.RED, console.WHITE, console.YELLOW, console.CYAN}
)

type mode struct {
	name string
	placedTerrainIndex int
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
