package d5

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type ElfPageSorter interface {
	AddRoles(r []Role)
	CheckOrder(s string) bool
	MiddleElem(s string) int
	Sort(s string) string
}

type Role struct {
	Grater, Lesser Page
}
type Page string

func NewSorter() ElfPageSorter {
	return &sorter{grater: map[Page][]Page{}}
}

type sorter struct {
	grater map[Page][]Page
}

func (s *sorter) AddRoles(r []Role) {
	for _, role := range r {
		s.grater[role.Grater] = append(s.grater[role.Grater], role.Lesser)
	}
}

func (s *sorter) CheckOrder(line string) bool {
	return s.Sort(line) == line
}

func (s *sorter) MiddleElem(line string) int {
	pages := strings.Split(line, ",")
	midElem := pages[(len(pages) / 2)]
	res, err := strconv.Atoi(midElem)
	if err != nil {
		panic(err)
	}
	return res
}

func (s *sorter) Sort(line string) string {
	pages := strings.Split(line, ",")

	sort.Slice(pages, func(i, j int) bool {
		pi := Page(pages[i])
		pj := Page(pages[j])
		l := s.grater[pi]
		if l != nil {
			if slices.Contains(l, pj) {
				return true
			}
		}
		l = s.grater[pj]
		if l != nil {
			if slices.Contains(l, pi) {
				return false
			}
		}
		panic(fmt.Sprintf("not found role for (%s, %s)", pi, pj))
	})
	return strings.Join(pages, ",")
}

func D5(input string) (int, int) {
	r, lines := parseInput(input)
	s := NewSorter()
	s.AddRoles(r)
	resA := 0
	resB := 0
	for _, line := range lines {
		if s.CheckOrder(line) {
			resA += s.MiddleElem(line)
		} else {
			sorted := s.Sort(line)
			resB += s.MiddleElem(sorted)
		}
	}
	return resA, resB
}

func parseInput(input string) (r []Role, p []string) {
	split := strings.Split(input, "\n\n")
	roles := strings.Split(split[0], "\n")
	for _, role := range roles {
		rs := strings.Split(role, "|")
		r = append(r, Role{
			Grater: Page(rs[0]),
			Lesser: Page(rs[1]),
		})
	}
	p = strings.Split(split[1], "\n")
	return
}
