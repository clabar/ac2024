package d2

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"

	"bartoli.no/d1"
)

type Safe string

const (
	s       = Safe("Safe")
	u       = Safe("Unsafe")
	maxStep = 3
	minStep = 1
)

func SafetyCheck(lines [][]int) map[int]Safe {
	res := make(map[int]Safe, len(lines))
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		res[i] = isSliceSafe(line)
		if res[i] == u {
			res[i] = safe2(line)
		}
	}
	return res
}

func safe2(line []int) Safe {
	for j := 0; j < len(line); j++ {
		ns := isSliceSafe(remove(line, j))
		if ns == s {
			return s
		}
	}
	return u
}

func remove(slice []int, s int) []int {
	ints := make([]int, 0)
	for i := 0; i < len(slice); i++ {
		if i == s {
			continue
		}
		ints = append(ints, slice[i])
	}
	return ints
}

func isSliceSafe(line []int) (res Safe) {
	res = s
	l := len(line)
	if l == 1 {
		return
	}
	sing := math.Signbit(float64(line[1] - line[0]))
	for i := 1; i < l; i++ {
		diff := d1.IntAbs(line[i] - line[i-1])
		if diff < minStep || diff > maxStep {
			// fmt.Println("Unsafe step for ", line[i], line[i-1])
			return u
		}
		if sing != math.Signbit(float64(line[i]-line[i-1])) {
			// fmt.Println("Unsafe sign for ", line[i], line[i-1])
			return u
		}
	}
	return
}

func ParseInput(file *os.File) [][]int {
	r := bufio.NewReader(file)
	res := make([][]int, 0)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		res = append(res, parseLine(string(line)))
	}
	return res
}

func CountSafe(lines map[int]Safe) int {
	res := 0
	for i := 0; i < len(lines); i++ {
		if lines[i] == s {
			res += 1
		}
	}
	return res
}

func parseLine(line string) (res []int) {
	tokens := strings.Split(line, " ")
	for _, token := range tokens {
		i, err := strconv.Atoi(token)
		if err != nil {
			panic(err)
		}
		res = append(res, i)
	}
	return
}
