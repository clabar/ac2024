package d7

import (
	"strconv"
	"strings"
)

func D7(input string) int {
	lines := strings.Split(input, "\n")
	res := 0
	for _, line := range lines {
		expectedRes, inputs := parseLine(line)
		res += tryAllCombinations(inputs, expectedRes)
	}
	return res
}

func tryAllCombinations(inputs []int, res int) int {

	si, mi := tryOps(inputs[0], inputs[1])
	tmp := []int{si, mi}
	if len(inputs) > 2 {
		for i := 2; i <= len(inputs)-1; i++ {
			var t2 []int
			for _, val := range tmp {
				s, m := tryOps(inputs[i], val)
				t2 = append(t2, s)
				t2 = append(t2, m)
			}
			tmp = t2
		}
	}
	for _, val := range tmp {
		if res == val {
			return res
		}
	}
	return 0
}

func tryOps(i, j int) (int, int) {
	return sum(i, j), mul(i, j)
}

func sum(i, j int) int {
	return i + j
}

func mul(i, j int) int {
	m := i * j
	return m
}

func parseLine(line string) (int, []int) {
	split := strings.Split(line, ":")
	res, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}
	var fres []int
	factors := strings.Split(split[1], " ")
	for _, factor := range factors {
		if factor != "" {
			num, err := strconv.Atoi(factor)
			if err != nil {
				panic(err)
			}
			fres = append(fres, num)
		}
	}
	return res, fres
}
