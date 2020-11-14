package main

import (
	"github.com/sidav/golibrl/console"
	"parcelcreationtool/generator"
	"strconv"
)

func generateAndRenderSample() {
	gen := generator.Generator{}
	key := ""
	for key != "ESCAPE" {
		console.Clear_console()
		w, h := console.GetConsoleSize()
		lvl := gen.Generate("parcels", "templates", w, h, 100)
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
		// render waypoints
		for routeNum := range lvl.Routes {
			for wpNum := range lvl.Routes[routeNum].Waypoints {
				x := lvl.Routes[routeNum].Waypoints[wpNum].X
				y := lvl.Routes[routeNum].Waypoints[wpNum].Y
				outputSymbol := strconv.Itoa(wpNum)
				if len(outputSymbol) > 1 {
					outputSymbol = string(rune(int('a') + wpNum - 10))
				}
				console.SetFgColor(console.MAGENTA)
				console.PutString(outputSymbol, x, y)
			}
		}
		console.Flush_console()
		
		key = console.ReadKey()
	}
}
