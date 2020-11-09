package generator

import (
	"github.com/sidav/golibrl/random/additive_random"
	"io/ioutil"
	. "parcelcreationtool/parcel"
)

var rnd additive_random.FibRandom
const triesForParcel = 5

type Generator struct {
	level   Level
	parcels []*Parcel
}

func (g *Generator) Generate(parcelsDir string, sizex, sizey int, desiredParcels int) *Level {
	rnd.InitDefault()
	g.level = Level{}
	g.level.init(sizex, sizey)
	g.parcels = make([]*Parcel, 0)

	//init fucking parcels
	items, _ := ioutil.ReadDir(parcelsDir)
	for _, item := range items {
		if item.IsDir() {

		} else {
			newParcel := Parcel{}
			newParcel.UnmarshalFromFile(parcelsDir + "/" + item.Name())
			g.parcels = append(g.parcels, &newParcel)
		}
	}
	if len(g.parcels) == 0 {
		panic("No parcels in folder!")
	}
	// randomly place parcels on the map
	for tries := 0; tries < desiredParcels; tries++ {
		for i :=0; i < triesForParcel; i++ {
			if g.placeRandomParcel() {
				break
			}
		}
	}
	g.level.cleanup()
	return &g.level
}

func (g *Generator) placeRandomParcel() bool {
	prc := g.parcels[rnd.Rand(len(g.parcels))]
	pw, ph := len(prc.Terrain), len(prc.Terrain[0])
	w, h := g.level.getSize()
	clearCoords := make([][]int, 0)
	for x:=0; x<w-pw; x++ {
		for y:=0; y<h-ph;y++{
			if g.level.isRectClearForPlacement(x, y, pw, ph) {
				clearCoords = append(clearCoords, []int{x, y})
			}
		}
	}
	if len(clearCoords) == 0 {
		return false
	}
	g.applyParcelAtCoords(prc, &clearCoords[rnd.Rand(len(clearCoords))])
	return true
}

func (g *Generator) applyParcelAtCoords(prc *Parcel, xy *[]int) {
	x, y := (*xy)[0], (*xy)[1]
	pw, ph := len(prc.Terrain), len(prc.Terrain[0])
	for i := 0; i < pw; i++ {
		for j:=0; j < ph; j++{
			g.level.Terrain[i+x][j+y] = prc.Terrain[i][j]
		}
	}
}
