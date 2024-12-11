package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/yumike/adventofcode-2024/internal/inputs"
)

type Result map[int]map[int]int

type State struct {
	m [][]int
	x int
	y int
}

func (s State) Height() int {
	return s.m[s.y][s.x]
}

func (s State) Move(x, y int) (State, bool) {
	if y < 0 || y >= len(s.m) || x < 0 || x >= len(s.m[y]) {
		return s, false
	}
	if s.m[y][x] != s.m[s.y][s.x]+1 {
		return s, false
	}
	return State{
		m: s.m,
		x: x,
		y: y,
	}, true
}

func merge(r1, r2 Result) Result {
	r := Result{}
	for y, yv := range r1 {
		for x := range yv {
			if _, ok := r[y]; !ok {
				r[y] = map[int]int{}
			}
			r[y][x] = r1[y][x]
		}
	}
	for y, yv := range r2 {
		for x := range yv {
			if _, ok := r[y]; !ok {
				r[y] = map[int]int{}
			}
			r[y][x] += r2[y][x]
		}
	}
	return r
}

func score(r Result) int {
	s := 0
	for _, yv := range r {
		for _, xv := range yv {
			s += xv
		}
	}
	return s
}

func calc(s State) Result {
	result := Result{}
	if s.Height() == 9 {
		result[s.y] = map[int]int{s.x: 1}
		return result
	}
	if ss, ok := s.Move(s.x+1, s.y); ok {
		result = merge(result, calc(ss))
	}
	if ss, ok := s.Move(s.x-1, s.y); ok {
		result = merge(result, calc(ss))
	}
	if ss, ok := s.Move(s.x, s.y+1); ok {
		result = merge(result, calc(ss))
	}
	if ss, ok := s.Move(s.x, s.y-1); ok {
		result = merge(result, calc(ss))
	}
	return result
}

func calcScore(s State) int {
	r := calc(s)
	return score(r)
}

func main() {
	path := os.Args[1]
	m := [][]int{}
	for line := range inputs.ReadLines(path) {
		row := []int{}
		for _, c := range line {
			n, _ := strconv.Atoi(string(c))
			row = append(row, n)
		}
		m = append(m, row)
	}

	sum := 0
	for y, row := range m {
		for x, v := range row {
			if v == 0 {
				s := State{
					m: m,
					x: x,
					y: y,
				}
				sum += calcScore(s)
			}
		}
	}
	fmt.Println(sum)
}
