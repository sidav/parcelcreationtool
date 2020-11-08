package main

var(
	terrains = [...]rune{'#', '.', '+', '\''}
	terrains_names = [...]string{"Wall", "Floor", "Door", "Window"}
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
