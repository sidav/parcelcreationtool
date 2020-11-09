package parcel

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	FLOOR = '.'
	WALL = '#'
	DOOR = '+'
)


type Parcel struct {
	Terrain [][]rune
	Routes[] Route
}

func (p *Parcel) Init(w, h int) {
	p.Terrain = make([][]rune, w)
	for i := range p.Terrain {
		p.Terrain[i] = make([]rune, h)
		for j := range p.Terrain[i] {
			p.Terrain[i][j] = FLOOR
		}
	}
	p.Routes = make([]Route, 0)
}

func (p *Parcel) MarshalToFile(filename string) {
	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(b)
	if err != nil {
		panic(err)
	}
}

func (p *Parcel) UnmarshalFromFile(filename string) {
	jsn, err := ioutil.ReadFile(filename)
	if err == nil {
		json.Unmarshal(jsn, p)
	}
}

type Waypoint struct {
	X, Y int
	Props string
}

type Route struct {
	Waypoints []Waypoint
}

func (r *Route) AddWaypoint(w *Waypoint) {
	r.Waypoints = append(r.Waypoints, *w)
}
