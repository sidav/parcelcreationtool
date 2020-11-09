package parcel

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
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

func (p *Parcel) GetSize() (int, int) {
	return len(p.Terrain), len(p.Terrain[0])
}

func (p *Parcel) MarshalToFile(filename string) {
	folderName := strings.Split(filename, "/")[0]

	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		os.Mkdir(folderName, 0777)
	}
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
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

func (p *Parcel) Rotate(times int) {
	for t := 0; t < times; t++ {
		w, h := p.GetSize()
		ter := make([][]rune, h)
		for i := range ter {
			ter[i] = make([]rune, w)
		}
		for x := range p.Terrain {
			for y := range p.Terrain[0] {
				ter[y][w-x-1] = p.Terrain[x][y]
			}
		}
		p.Terrain = ter
		for i := range p.Routes {
			for j := range p.Routes[i].Waypoints {
				oldx := p.Routes[i].Waypoints[j].X
				oldy := p.Routes[i].Waypoints[j].Y
				p.Routes[i].Waypoints[j].X = oldy
				p.Routes[i].Waypoints[j].Y = w-oldx-1
			}
		}
	}
}
