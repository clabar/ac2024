package d4

import "strings"

// SearchForXmas counts the instances of the word XMAS in all directions
func SearchForXmas(input string) int {
	matrix := toIntMatrix(input)
	res := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == firstChar {
				res += checkWord(matrix, i, j)
			}
		}
	}
	return res
}

type direction struct {
	name           string
	deltaX, deltaY int
}

const (
	firstChar = 'X'
	m         = 'M'
	a         = 'A'
	s         = 'S'
)

var (
	N = direction{
		name:   "North",
		deltaX: 0,
		deltaY: 1,
	}
	S = direction{
		name:   "South",
		deltaX: 0,
		deltaY: -1,
	}
	E = direction{
		name:   "East",
		deltaX: -1,
		deltaY: 0,
	}
	W = direction{
		name:   "West",
		deltaX: 1,
		deltaY: 0,
	}
	NW = direction{
		name:   "NorthWest",
		deltaX: 1,
		deltaY: 1,
	}
	NE = direction{
		name:   "NorthEast",
		deltaX: -1,
		deltaY: 1,
	}
	SW = direction{
		name:   "SouthWest",
		deltaX: 1,
		deltaY: -1,
	}
	SE = direction{
		name:   "SouthEast",
		deltaX: -1,
		deltaY: -1,
	}
	directions = []direction{
		N, S, W, E, NW, NE, SW, SE,
	}
)

type directionChecker struct {
	d              direction
	lastRead       int32
	matrix         [][]int32
	istart, jstart int
}

func checkWord(matrix [][]int32, i int, j int) int {
	res := 0
	for _, dir := range directions {
		ns := directionChecker{
			d:        dir,
			lastRead: firstChar,
			matrix:   matrix,
			istart:   i,
			jstart:   j,
		}
		res += ns.checkDir()
	}
	return res
}

func (cd *directionChecker) checkDir() int {
	if cd.istart+cd.d.deltaY < 0 || cd.istart+cd.d.deltaY >= len(cd.matrix) {
		return 0
	}
	row := cd.matrix[cd.jstart]
	if cd.jstart+cd.d.deltaX < 0 || cd.jstart+cd.d.deltaX >= len(row) {
		return 0
	}
	nc := findNextChar(cd.lastRead)
	if cd.matrix[cd.istart+cd.d.deltaY][cd.jstart+cd.d.deltaX] == nc {
		if nc == s {
			// gotcha !!
			return 1
		}
		ns := directionChecker{
			d:        cd.d,
			lastRead: nc,
			matrix:   cd.matrix,
			istart:   cd.istart + cd.d.deltaY,
			jstart:   cd.jstart + cd.d.deltaX,
		}
		return ns.checkDir()
	}
	return 0
}

func findNextChar(read int32) int32 {
	switch read {
	case firstChar:
		return m
	case m:
		return a
	case a:
		return s
	}
	return 0
}

func toIntMatrix(input string) [][]int32 {
	lines := strings.Split(input, "\n")
	res := make([][]int32, len(lines))
	for i, line := range lines {
		res[i] = parseLine(line)
	}
	return res
}

func parseLine(s string) []int32 {
	var res []int32
	for _, c := range s {
		res = append(res, c)
	}
	return res
}
