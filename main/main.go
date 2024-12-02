package main

import (
	"fmt"

	"bartoli.no/d1"
)

func main() {
	a := []int{3, 4, 2, 1, 3, 3}
	b := []int{3, 4, 5, 3, 9, 3}
	fmt.Println(d1.ComputeDistance(a, b))
}
