package d1

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ComputeDistance(s1, s2 []int) int {
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})
	l1 := len(s1)
	l2 := len(s2)
	if l1 >= l2 {
		return d(s1, s2)
	} else {
		return d(s2, s1)
	}
}

// assumes len(s1 > len s2)
func d(s1 []int, s2 []int) int {
	l1 := len(s1)
	l2 := len(s2)
	if l2 > l1 {
		panic("invalid input: AOB")
	}
	var dist int
	for i := 0; i < l1; i++ {
		if i > l2 {
			dist += s1[i]
		} else {
			res := IntAbs(s1[i] - s2[i])
			dist += res
		}
	}
	return dist
}

func IntAbs(i int) int {
	return int(math.Abs(float64(i)))
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

func parseLine(s string) []int {
	var res []int
	split := strings.Split(s, " ")
	for i := 0; i < len(split); i++ {
		a, err := strconv.Atoi(split[i])
		if err != nil {
			panic(err)
		}
		res = append(res, a)
	}
	return res
}
func FirstCol(both [][]int) (res []int) {
	for i := 0; i < len(both); i++ {
		res = append(res, both[i][0])
	}
	return
}

func SecondCol(both [][]int) (res []int) {
	for i := 0; i < len(both); i++ {
		res = append(res, both[i][1])
	}
	return
}

func Doit() int {
	fi, err := os.Open("../d1/input1.txt")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	input := ParseInput(fi)
	return ComputeDistance(FirstCol(input), SecondCol(input))
}
