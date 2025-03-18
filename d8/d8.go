package d8

import (
	"fmt"
	"math"

	"utils"
)

func D8(input string) int {
	plane := utils.LoadInput(input)
	ymax = len(plane)
	xmax = len(plane[0])
	antennas := searchAntennas(plane)
	resonance := map[utils.Point]struct{}{}
	//fmt.Println(fmt.Sprintf("%v", antennas))
	for _, ant := range antennas {
		//fmt.Println(fmt.Sprintf("%v", ant))
		if len(ant) < 2 {
			break
		}
		for i := 0; i < len(ant)-1; i++ {
			points := []utils.Point{}
			for j := i + 1; j < len(ant); j++ {
				//fmt.Println(fmt.Sprintf("checking combo %v, %v", ant[i], ant[j]))
				points = append(points, findResonancePoints(ant[i], ant[j])...)
			}
			for _, point := range points {
				if isInWithinBounderies(point) {
					resonance[point] = struct{}{}
				}
			}
		}
	}
	fmt.Println(fmt.Sprintf("found: %v", resonance))
	return len(resonance)
}

func D8P2(input string) int {
	plane := utils.LoadInput(input)
	ymax = len(plane)
	xmax = len(plane[0])
	antennas := searchAntennas(plane)
	resonance := map[utils.Point]struct{}{}
	//fmt.Println(fmt.Sprintf("%v", antennas))
	for _, ant := range antennas {
		//fmt.Println(fmt.Sprintf("%v", ant))
		if len(ant) < 2 {
			break
		}
		for i := 0; i < len(ant)-1; i++ {
			points := []utils.Point{}
			for j := i + 1; j < len(ant); j++ {
				//fmt.Println(fmt.Sprintf("checking combo %v, %v", ant[i], ant[j]))
				points = append(points, findAllResonancePoints(ant[i], ant[j])...)
			}
			for _, point := range points {
				if isInWithinBounderies(point) {
					resonance[point] = struct{}{}
				}
			}
		}
	}
	//fmt.Println(fmt.Sprintf("found: %v", resonance))
	for point, _ := range resonance {
		if plane[point.Y()][point.X()] == '.' {
			plane[point.Y()][point.X()] = '#'
		}
	}
	utils.Print(plane)
	return len(resonance)
}

func findAllResonancePoints(p1 utils.Point, p2 utils.Point) []utils.Point {
	dx := int(math.Abs(float64(p2.X() - p1.X())))
	dy := int(math.Abs(float64(p2.Y() - p1.Y())))
	var x []int
	var y []int
	for i := 0; i < xmax; i += dx {
		x = append(x, i+p1.X()%dx)
	}
	for i := 0; i < ymax; i += dy {
		y = append(y, i+p1.Y()%dy)
	}
	res := []utils.Point{}
	var l int
	if len(x) >= len(y) {
		l = len(y)
	} else {
		l = len(x)
	}
	for i := 0; i < l; i++ {
		res = append(res, utils.NewPoint(x[i], y[i]))
	}
	return res
}

var (
	xmax = 0
	ymax = 0
)

func isInWithinBounderies(point utils.Point) bool {
	return point.X() > -1 && point.X() < xmax && point.Y() > -1 && point.Y() < ymax
}

func findResonancePoints(p1 utils.Point, p2 utils.Point) []utils.Point {
	x1, x2 := dist(p1.X(), p2.X())
	y1, y2 := dist(p1.Y(), p2.Y())
	res := []utils.Point{utils.NewPoint(x1, y1), utils.NewPoint(x2, y2)}
	fmt.Println(fmt.Sprintf("For combo [%v, %v]: %v", p1, p2, res))
	return res
}

func dist(x1 int, x2 int) (int, int) {
	d := x2 - x1
	return x1 - d, x2 + d
}

func searchAntennas(plane [][]int32) map[int32][]utils.Point {
	res := make(map[int32][]utils.Point)
	for i := 0; i < len(plane); i++ {
		for j := 0; j < len(plane[i]); j++ {
			val := plane[i][j]
			if isAntenna(val) {
				point := utils.NewPoint(j, i)
				res[val] = append(res[val], point)
			}
		}
	}
	return res
}

func isAntenna(val int32) bool {
	return (val >= A && val <= Z) || (val >= a && val <= z) || (val >= zero && val <= nine)
}

const (
	A    = 'A'
	Z    = 'Z'
	a    = 'a'
	z    = 'z'
	zero = '0'
	nine = '9'
)
