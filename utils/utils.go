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
