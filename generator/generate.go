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
	templates []*Parcel
}

func (g *Generator) Generate(parcelsDir, templatesDir string, sizex, sizey int, desiredParcels int) *Level {
	rnd.InitDefault()
	g.level = Level{}
	g.parcels = make([]*Parcel, 0)
	g.templates = make([]*Parcel, 0)

	// init fucking templates
	items, _ := ioutil.ReadDir(templatesDir)
	for _, item := range items {
		if item.IsDir() {

		} else {
			newTemplate := Parcel{}
			newTemplate.UnmarshalFromFile(templatesDir + "/" + item.Name())
			g.templates = append(g.templates, &newTemplate)
		}
	}
	if len(g.templates) == 0 {
		g.level.init(sizex, sizey)
	} else {
		templateForInit := g.templates[rnd.RandInRange(0, len(g.templates)-1)]
		templateForInit.Rotate(rnd.Rand(4))
		g.level.initFromTemplate(templateForInit)
	}
	//init fucking parcels
	items, _ = ioutil.ReadDir(parcelsDir)
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
	prc.Rotate(rnd.Rand(4))
	if rnd.OneChanceFrom(2) {
		prc.MirrorX()
	}
	if rnd.OneChanceFrom(2) {
		prc.MirrorY()
	}
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
	g.level.applyParcelAtCoords(prc, &clearCoords[rnd.Rand(len(clearCoords))])
	return true
}
