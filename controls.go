package main

import "github.com/sidav/golibrl/console"

type cursor struct {
	x, y int
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
	switch key {
	case "UP": crs.y--
	case "RIGHT": crs.x++
	case "DOWN": crs.y++
	case "LEFT": crs.x--
	case "ENTER": enterKeyForMode()
	case "ESCAPE": running = false 
	}
	crs.normalizeCoords()
}

func enterKeyForMode() {
	currParcel.Terrain[crs.x][crs.y] = currMode.placedTerrain
}
