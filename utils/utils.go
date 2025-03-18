package utils

import (
	"fmt"
	"strings"
)

func LoadInput(input string) [][]int32 {
	lines := strings.Split(input, "\n")
	res := [][]int32{}
	for y, line := range lines {
		maxX := len(line)
		res = append(res, make([]int32, maxX))
		for x, c := range line {
			res[y][x] = c
		}
	}
	return res
}

func Print(world [][]int32) {
	var w string
	for i := 0; i < len(world); i++ {
		for j := 0; j < len(world[i]); j++ {
			w = fmt.Sprintf("%s%c", w, world[i][j])
		}
		w = fmt.Sprintf("%s\n", w)
	}
	fmt.Println(w)
}

func NewPoint(x, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

type Point struct {
	x, y int
}

func (p *Point) X() int {
	return p.x
}

func (p *Point) Y() int {
	return p.y
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}
