package main

import (
	"fmt"
	"os"

	"github.com/yumike/adventofcode-2024/internal/inputs"
)

type Point struct {
	X, Y int
}

type Direction struct {
	X, Y int
}

func (d Direction) asFace() byte {
	if d.X == 0 {
		if d.Y == -1 {
			return '^'
		}
		return 'v'
	}
	if d.X == -1 {
		return '<'
	}
	return '>'
}

type Map struct {
	M   [][]byte
	pos Point
	dir Direction

	Cycle bool
}

func NewMap(path string) *Map {
	m := Map{}
	y := 0
	for line := range inputs.ReadLines(path) {
		row := []byte(line)
		for x := 0; x < len(row); x++ {
			if isFace(line[x]) {
				m.dir = NewDirection(row[x])
				m.pos = Point{x, y}
			}
		}
		m.M = append(m.M, row)
		y++
	}
	return &m
}

func (m *Map) Clone() *Map {
	m2 := Map{}
	m2.M = make([][]byte, len(m.M))
	for y, row := range m.M {
		m2.M[y] = make([]byte, len(row))
		copy(m2.M[y], row)
	}
	m2.pos = Point{m.pos.X, m.pos.Y}
	m2.dir = Direction{m.dir.X, m.dir.Y}
	m2.Cycle = false
	return &m2
}

func (m *Map) Step() bool {
	nextPos := Point{m.pos.X + m.dir.X, m.pos.Y + m.dir.Y}
	if nextPos.Y < 0 || nextPos.Y >= len(m.M) || nextPos.X < 0 || nextPos.X >= len(m.M[nextPos.Y]) {
		return false
	}

	if m.M[nextPos.Y][nextPos.X] == '#' {
		m.dir = turnRight(m.dir)
		return true
	}

	if isFace(m.M[nextPos.Y][nextPos.X]) {
		if m.dir.asFace() == m.M[nextPos.Y][nextPos.X] {
			m.Cycle = true
			return false
		}
	}

	m.M[nextPos.Y][nextPos.X] = m.dir.asFace()
	m.pos = nextPos
	return true
}

func (m *Map) Walk() {
	for {
		ok := m.Step()
		if !ok {
			break
		}
	}
}

func (m *Map) Count() int {
	cnt := 0
	for _, row := range m.M {
		for _, cell := range row {
			if isFace(cell) {
				cnt++
			}
		}
	}
	return cnt
}

func NewDirection(face byte) Direction {
	if face == '^' {
		return Direction{0, -1}
	} else if face == 'v' {
		return Direction{0, 1}
	} else if face == '<' {
		return Direction{-1, 0}
	} else if face == '>' {
		return Direction{1, 0}
	}
	panic("Invalid face")
}

func isFace(b byte) bool {
	return b == '^' || b == 'v' || b == '<' || b == '>'
}

func turnRight(dir Direction) Direction {
	if dir.X == 0 {
		return Direction{-dir.Y, 0}
	}
	return Direction{0, dir.X}
}

func part1(path string) int {
	m := NewMap(path)
	m.Walk()
	return m.Count()
}

func part2(path string) int {
	m := NewMap(path)
	o := Point{}
	cnt := 0
	for {
		if o.Y >= len(m.M) {
			break
		}
		if o.X >= len(m.M[o.Y]) {
			o.Y++
			o.X = 0
			continue
		}
		if m.M[o.Y][o.X] != '.' {
			o.X++
			continue
		}
		m2 := m.Clone()
		m2.M[o.Y][o.X] = '#'
		m2.Walk()
		if m2.Cycle {
			cnt++
		}
		o.X++
	}
	return cnt
}

func main() {
	path := os.Args[1]
	fmt.Println(part1(path))
	fmt.Println(part2(path))
}
