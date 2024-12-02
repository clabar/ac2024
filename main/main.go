package main

import (
	"fmt"
	"os"

	"bartoli.no/d1"
	"bartoli.no/d2"
)

func main() {
	a := []int{3, 4, 2, 1, 3, 3}
	b := []int{3, 4, 5, 3, 9, 3}
	fmt.Println(d1.ComputeDistance(a, b))

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
}
