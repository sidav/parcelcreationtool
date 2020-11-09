package generator

type Level struct {
	Terrain [][]rune
}

func (l *Level) init(w, h int) {
	l.Terrain = make([][]rune, w)
	for i := range l.Terrain {
		l.Terrain[i] = make([]rune, h)
		for y := 0; y < h; y++ {
			l.Terrain[i][y] = '?'
		}
	}
}

func (l *Level) getSize() (int,int) {
	return len(l.Terrain), len(l.Terrain[0])
}

func (l *Level) isRectClearForPlacement(x, y, w, h int) bool {
	for i:=x;i<x+w;i++{
		for j:=y;j<y+h;j++{
			if l.Terrain[i][j] != '?' {
				return false
			}
		}
	}
	return true
}

func (l *Level) cleanup() {
	w, h := l.getSize()
	for x:=0; x<w; x++ {
		for y:=0; y < h; y++ {
			if l.Terrain[x][y] == '?' {
				l.Terrain[x][y] = '.'
			}
		}
	}
}
