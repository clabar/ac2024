package d6

import (
	"fmt"
)

type World interface {
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
		obstacles:     map[string]*Point{},
		turningPoints: map[string]*Point{},
		hashes:        map[string]*Point{},
		v2:            map[string]int32{},
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
	hashes        map[string]*Point
	turningPoints map[string]*Point
	currentPos    *Point
	start         *Point
	pos2          *Point
	v2            map[string]int32
	obstacles     map[string]*Point
}

func (wrld *w) Add(p *Point) {
	if p.val == startPos {
		if wrld.currentPos != nil {
			panic("too many start pos")
		}
		wrld.currentPos = p
		wrld.start = p
		wrld.visited[p.coor()] = p
	}
	if p.val == obstacle {
		wrld.hashes[p.coor()] = p
	}
	wrld.index[p.coor()] = p
}

func (p *Point) Up() *Point {
	u := &Point{
		x:   p.x,
		y:   p.y - 1,
		val: up,
	}
	return u
}

func (p *Point) Down() *Point {
	dw := &Point{
		x:   p.x,
		y:   p.y + 1,
		val: down,
	}
	return dw
}

func (p *Point) Left() *Point {
	l := &Point{
		x:   p.x - 1,
		y:   p.y,
		val: left,
	}
	return l
}

func (p *Point) Right() *Point {
	r := &Point{
		x:   p.x + 1,
		y:   p.y,
		val: right,
	}
	return r
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
	if wrld.pos2 != nil {
		wInt[wrld.pos2.y][wrld.pos2.x] = '*'
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
		candidate = wrld.currentPos.Up()
	case down:
		candidate = wrld.currentPos.Down()
	case left:
		candidate = wrld.currentPos.Left()
	case right:
		candidate = wrld.currentPos.Right()
	default:
		panic(fmt.Sprintf("Invalid direction %c", wrld.currentPos.val))
	}
	if candidate.OutOfBound(maxX, maxY, 0, 0) {
		// managed to exit the filed EOG
		wrld.index[wrld.currentPos.coor()].val = passed
		return candidate
	}
	candidate = wrld.index[candidate.coor()]
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
		p := &Point{
			x:   wrld.currentPos.x,
			y:   wrld.currentPos.y,
			val: rotate(wrld.currentPos.val),
		}
		wrld.v2 = map[string]int32{}
		currentCandidate = *p
		coor := searchLoop(wrld, p)
		if coor != "" && coor != wrld.start.coor() {
			wrld.obstacles[coor] = wrld.index[coor]
		}
	}
}

func (wrld *w) isObstacle(p *Point) bool {
	_, found := wrld.hashes[p.coor()]
	return found
}

var currentCandidate Point

func searchLoop(world *w, current *Point) string {
	var nextPoint *Point
	switch current.val {
	case up:
		nextPoint = current.Up()
	case down:
		nextPoint = current.Down()
	case left:
		nextPoint = current.Left()
	case right:
		nextPoint = current.Right()
	default:
		panic("X" + string(current.val))
	}
	p, found := world.v2[nextPoint.coor()]
	//fmt.Println(fmt.Sprintf("Found %s, dir: %c, %c", nextPoint.coor(), p, nextPoint.val))
	if found && p == nextPoint.val || currentCandidate.coor() == nextPoint.coor() {
		fmt.Println(fmt.Sprintf("found %v", nextPoint.coor()))
		return nextPoint.coor()
	} else {
		if nextPoint.OutOfBound(maxX, maxY, 0, 0) {
			//world.pos2 = current
			//printWorld(world)
			//fmt.Println("~~~~~" + string(current.val) + "~~~~~")
			return ""
		}
		world.v2[nextPoint.coor()] = nextPoint.val
		//fmt.Println(fmt.Sprintf("visited %v", v2))
		if world.isObstacle(nextPoint) {
			current.val = rotate(current.val)
			return searchLoop(world, current)
		}
		return searchLoop(world, nextPoint)
	}
	return ""
}
