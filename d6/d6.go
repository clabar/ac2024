package d6

import (
	"fmt"
	"strings"
	"time"
)

func parseInput(input string) World {
	lines := strings.Split(input, "\n")
	maxY = len(lines)
	world := New()
	for y, line := range lines {
		maxX = len(line)
		for x, c := range line {
			NewPoint(x, y, c, world)
		}
	}
	return world
}

var (
	maxX, maxY int
)

func D6(input string) int {
	world := parseInput(input)
	steps := 0
	defer printWorld(world)
	defer func() {
		fmt.Println("Obstacles found:", world.Obstacles())
	}()
	for {
		nextPos := world.NextPos()
		steps++
		if steps%100 == 0 {
			printWorld(world)
		}
		// printWorld(world)
		if nextPos.OutOfBound(maxX, maxY, 0, 0) {
			return world.Steps()
		}
		if steps > 10000 {
			panic("too many steps aborting")
		}
	}
}

func printWorld(w World) {
	fmt.Println(w)
	time.Sleep(33 * time.Millisecond)
}

const (
	startPos = '^'
	up       = '^'
	down     = 'v'
	left     = '<'
	right    = '>'
	obstacle = '#'
	passed   = 'X'
	vert     = '|'
	or       = '-'
	turn     = '+'
)
