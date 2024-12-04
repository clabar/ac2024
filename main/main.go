package main

import (
	"fmt"
	"os"

	"bartoli.no/d1"
	"bartoli.no/d2"
	"bartoli.no/d3"
	"bartoli.no/d4"
)

func main() {
	a := []int{3, 4, 2, 1, 3, 3}
	b := []int{3, 4, 5, 3, 9, 3}
	fmt.Println(d1.ComputeDistance(a, b))
	fmt.Println(d1.Doit())

	fi, err := os.Open("../d2/input.txt")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	fmt.Println(d2.CountSafe(d2.SafetyCheck(d2.ParseInput(fi))))

	//fmt.Println(d3.ParseAll("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"))
	fmt.Println(d3.ParseAll(d3.Input))

	//fmt.Println(d4.SearchForXmas(testInput4))
	fmt.Println(d4.SearchForXmas2(d4.Input4))
	fmt.Println(d4.SearchForXmas(d4.Input4))
}

var testInput4 = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
