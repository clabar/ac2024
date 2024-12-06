package d6

import (
	"fmt"
)

type World interface {
	Up() *Point
	Down() *Point
	Left() *Point
	Right() *Point
	Add(p *Point)
	NextPos() *Point
	String() string
	Steps() int
	Obstacles() int
}

func New() World {
	var world = w{
		index:         map[string]*Point{},
		visited:       map[string]*Point{},
		turningPoints: map[string]*Point{},
		obstacles:     map[string]*Point{},
	}
	return &world
}

func NewPoint(x, y int, val int32, w World) *Point {
	if x < 0 || y < 0 {
		panic("Invalid input ")
	}
	p := &Point{
		x:   x,
		y:   y,
		val: val,
	}
	w.Add(p)
	return p
}

type Point struct {
	x, y int
	val  int32
}

func (p *Point) OutOfBound(maxX, maxY, minX, minY int) bool {
	return p.x < minX || p.x >= maxX || p.y < minY || p.y >= maxY
}

func (p *Point) coor() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

type w struct {
	index         map[string]*Point
	visited       map[string]*Point
	currentPos    *Point
	turningPoints map[string]*Point
	obstacles     map[string]*Point
}

func (wrld *w) Add(p *Point) {
	if p.val == startPos {
		if wrld.currentPos != nil {
			panic("too many start pos")
		}
		wrld.currentPos = p
		wrld.visited[p.coor()] = p
	}
	wrld.index[p.coor()] = p
}

func (wrld *w) Up() *Point {
	u := &Point{
		x: wrld.currentPos.x,
		y: wrld.currentPos.y - 1,
	}
	newPos, found := wrld.index[u.coor()]
	if !found {
		return u
	}
	return newPos
}

func (wrld *w) Down() *Point {
	dw := &Point{
		x: wrld.currentPos.x,
		y: wrld.currentPos.y + 1,
	}
	newPos, found := wrld.index[dw.coor()]
	if !found {
		return dw
	}
	return newPos
}

func (wrld *w) Left() *Point {
	l := &Point{
		x: wrld.currentPos.x - 1,
		y: wrld.currentPos.y,
	}
	newPos, found := wrld.index[l.coor()]
	if !found {
		return l
	}
	return newPos
}

func (wrld *w) Right() *Point {
	r := &Point{
		x: wrld.currentPos.x + 1,
		y: wrld.currentPos.y,
	}
	newPos, found := wrld.index[r.coor()]
	if !found {
		return r
	}
	return newPos
}

func (wrld *w) String() string {
	res := ""
	var wInt = make([][]int32, maxY)
	for i, _ := range wInt {
		wInt[i] = make([]int32, maxX)
	}
	for _, point := range wrld.index {
		wInt[point.y][point.x] = point.val
	}
	for _, point := range wrld.obstacles {
		wInt[point.y][point.x] = 'O'
	}
	for _, vals := range wInt {
		for _, val := range vals {
			res = res + string(val)
		}
		res = res + "\n"
	}

	return res
}

func (wrld *w) NextPos() *Point {
	var candidate *Point
	cursor := wrld.currentPos.val
	switch cursor {
	case up:
		candidate = wrld.Up()
	case down:
		candidate = wrld.Down()
	case left:
		candidate = wrld.Left()
	case right:
		candidate = wrld.Right()
	default:
		panic(fmt.Sprintf("Invalid direction %c", wrld.currentPos.val))
	}
	if candidate.OutOfBound(maxX, maxY, 0, 0) {
		// managed to exit the filed EOG
		wrld.index[wrld.currentPos.coor()].val = passed
		return candidate
	}
	if candidate.val == obstacle {
		// rotate 90 deg clockwise
		wrld.rotate()
		return wrld.NextPos()
	}
	// candidate is a valid place to move to
	wrld.updatePos(candidate, cursor)
	wrld.addArtificialObstacle()
	return candidate
}

func (wrld *w) updatePos(candidate *Point, cursor int32) {
	if wrld.isTurningPoint() {
		// is a turn
		wrld.index[wrld.currentPos.coor()].val = turn
	} else {
		wrld.index[wrld.currentPos.coor()].val = passedVal(cursor)
	}

	wrld.index[candidate.coor()].val = cursor
	wrld.currentPos = candidate
	wrld.visited[candidate.coor()] = candidate
}

func passedVal(c int32) int32 {
	switch c {
	case up:
		return vert
	case down:
		return vert
	case left:
		return or
	case right:
		return or
	default:
		panic("Invalid dir")
	}
}

func (wrld *w) Steps() int {
	return len(wrld.visited)
}

func (wrld *w) Obstacles() int {
	return len(wrld.obstacles)
}

func (wrld *w) rotate() {
	wrld.turningPoints[wrld.currentPos.coor()] = wrld.currentPos
	wrld.currentPos.val = rotate(wrld.currentPos.val)
}

func (wrld *w) isTurningPoint() bool {
	_, found := wrld.turningPoints[wrld.currentPos.coor()]
	return found
}

func rotate(curr int32) int32 {
	switch curr {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	default:
		panic("Invalid dir")
	}
}

func (wrld *w) addArtificialObstacle() {
	if !wrld.isTurningPoint() {
		dir := rotate(wrld.currentPos.val)
		if dir == up {
			for i := wrld.currentPos.y; i >= 0; i-- {
				p := Point{x: wrld.currentPos.x, y: i}
				if wrld.index[p.coor()].val == obstacle {
					prec := Point{x: wrld.currentPos.x, y: i + 1}
					_, found := wrld.turningPoints[prec.coor()]
					if found {
						coor := fmt.Sprintf("%d,%d", wrld.currentPos.x-1, wrld.currentPos.y)
						wrld.obstacles[coor] = wrld.index[coor]
						break
					}
				}
			}
		}
		if dir == down {
			for i := wrld.currentPos.y; i < maxY; i++ {
				p := Point{x: wrld.currentPos.x, y: i}
				if wrld.index[p.coor()].val == obstacle {
					prec := Point{x: wrld.currentPos.x, y: i - 1}
					_, found := wrld.turningPoints[prec.coor()]
					if found {
						coor := fmt.Sprintf("%d,%d", wrld.currentPos.x+1, wrld.currentPos.y)
						wrld.obstacles[coor] = wrld.index[coor]
						break
					}
				}
			}
		}
		if dir == left {
			for i := wrld.currentPos.x; i >= 0; i-- {
				p := Point{x: i, y: wrld.currentPos.y}
				if wrld.index[p.coor()].val == obstacle {
					prec := Point{x: i + 1, y: wrld.currentPos.y}
					_, found := wrld.turningPoints[prec.coor()]
					if found {
						coor := fmt.Sprintf("%d,%d", wrld.currentPos.x, wrld.currentPos.y+1)
						wrld.obstacles[coor] = wrld.index[coor]
						break
					}
				}
			}
		}
		if dir == right {
			for i := wrld.currentPos.x; i < maxX; i++ {
				p := Point{x: i, y: wrld.currentPos.y}
				if wrld.index[p.coor()].val == obstacle {
					prec := Point{x: i - 1, y: wrld.currentPos.y}
					_, found := wrld.turningPoints[prec.coor()]
					if found {
						coor := fmt.Sprintf("%d,%d", wrld.currentPos.x, wrld.currentPos.y-1)
						wrld.obstacles[coor] = wrld.index[coor]
						break
					}
				}
			}
		}
	}
}
