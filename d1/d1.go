package d1

import (
	"math"
	"sort"
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
