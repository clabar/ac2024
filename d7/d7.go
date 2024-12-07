package d7

import (
	"math"
	"strconv"
	"strings"
)

func D7(input string) (int, int) {
	lines := strings.Split(input, "\n")
	res1 := 0
	res2 := 0
	for _, line := range lines {
		expectedRes, inputs := parseLine(line)
		res1 += tryAllCombinations(inputs, expectedRes)
		res2 += part2(inputs, expectedRes)
	}
	return res1, res2
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

func part2(inputs []int, res int) int {

	si, mi, ci := tryOps2(inputs[0], inputs[1])
	tmp := []int{si, mi, ci}
	if len(inputs) > 2 {
		for i := 2; i <= len(inputs)-1; i++ {
			var t2 []int
			for _, val := range tmp {
				s, m, c := tryOps2(val, inputs[i])
				//fmt.Println(fmt.Sprintf("(%d op %d) -> s %d, m %d, c %d", inputs[i], val, s, m, c))
				t2 = append(t2, s)
				t2 = append(t2, m)
				t2 = append(t2, c)
			}
			tmp = t2
		}
	}
	//fmt.Println(fmt.Sprintf("%v, expected res %d", tmp, res))
	for _, val := range tmp {
		if res == val {
			return res
		}
	}
	return 0
}

func tryOps2(i, j int) (int, int, int) {
	return sum(i, j), mul(i, j), Conc(i, j)
}
func tryOps(i, j int) (int, int) {
	return sum(i, j), mul(i, j)
}

func Conc(i, j int) int {
	digits := int(math.Log10(float64(j)) + 1)
	//fmt.Println(fmt.Sprintf("digits %d", digits))
	return i*int(math.Pow(10, float64(digits))) + j
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
