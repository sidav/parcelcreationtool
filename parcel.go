package main

type Parcel struct {
	Terrain [][]rune
	Routes[] Route
}

func (p *Parcel) init(w, h int) {
	p.Terrain = make([][]rune, w)
	for i := range p.Terrain {
		p.Terrain[i] = make([]rune, h)
		for j := range p.Terrain[i] {
			p.Terrain[i][j] = FLOOR
		}
	}
	p.Routes = make([]Route, 0)
}

type Waypoint struct {
	X, Y int
	Props string
}

type Route struct {
	Waypoints []Waypoint
}

func (r *Route) addWaypoint(w *Waypoint) {
	r.Waypoints = append(r.Waypoints, *w)
}
