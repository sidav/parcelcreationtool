package main

import (
	"github.com/sidav/golibrl/console"
	"parcelcreationtool/generator"
)

func generateAndRenderSample() {
	gen := generator.Generator{}
	key := ""
	for key != "ESCAPE" {
		w, h := console.GetConsoleSize()
		lvl := gen.Generate("parcels", w, h, 100)
		// render level 
		for x := 0; x < len(lvl.Terrain); x++ {
			for y := 0; y < len(lvl.Terrain[0]); y++ {
				for i := range terrains {
					if terrains[i] == lvl.Terrain[x][y] {
						console.SetFgColor(terrainsColors[i])
						break
					}
				}
				console.PutChar(lvl.Terrain[x][y], x, y)
			}
		}
		console.Flush_console()
		
		key = console.ReadKey()
	}
}
