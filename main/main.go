package main

import (
	"fmt"
	"os"

	"d7"
)

func main() {
	// a := []int{3, 4, 2, 1, 3, 3}
	// b := []int{3, 4, 5, 3, 9, 3}
	// fmt.Println(d1.ComputeDistance(a, b))
	// fmt.Println(d1.Doit())

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
	// fmt.Println(d2.CountSafe(d2.SafetyCheck(d2.ParseInput(fi))))

	//fmt.Println(d3.ParseAll("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"))
	// fmt.Println(d3.ParseAll(d3.Input))

	//fmt.Println(d4.SearchForXmas(testInput4))
	// fmt.Println(d4.SearchForXmas2(d4.Input4))
	// fmt.Println(d4.SearchForXmas(d4.Input4))

	//fmt.Println(d5.D5(d5.Input)) // 6949
	//fmt.Println(d6.D6(inputD6))
	//fmt.Println(d6.D6(d6.Input))

	fmt.Println(d7.D7(inputD7))
	fmt.Println(d7.D7(d7.Input))
	//fmt.Println(d7.Conc(48, 6))
}

var inputD7 = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

var inputD6 = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

var tInput5 = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

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
