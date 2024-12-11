package main

import (
	"fmt"
	"os"

	"github.com/yumike/adventofcode-2024/internal/inputs"
)

type Point struct {
	x, y int
}

type Antenna struct {
	letter byte
	x, y   int
}

type Map struct {
	Antennas  [][]byte
	Antinodes [][]byte
}

func (m *Map) addAntinode(p Point) {
	if p.y < 0 || p.y >= len(m.Antinodes) {
		return
	}
	if p.x < 0 || p.x >= len(m.Antinodes[p.y]) {
		return
	}
	m.Antinodes[p.y][p.x] = '#'
}

func (m *Map) getAninodes(ai, aj Antenna) []Point {
	if ai.letter != aj.letter {
		panic("Antennas must have the same letter")
	}
	ps := []Point{}
	dx := aj.x - ai.x
	dy := aj.y - ai.y
	x := ai.x
	y := ai.y
	for {
		if y < 0 || y >= len(m.Antennas) {
			break
		}
		if x < 0 || x >= len(m.Antennas[y]) {
			break
		}
		ps = append(ps, Point{x, y})

		x -= dx
		y -= dy
	}
	return ps
}

func main() {
	path := os.Args[1]
	m := Map{}
	for line := range inputs.ReadLines(path) {
		m.Antennas = append(m.Antennas, []byte(line))

		antinodes := make([]byte, len(line))
		for i := range antinodes {
			antinodes[i] = '.'
		}
		m.Antinodes = append(m.Antinodes, antinodes)
	}

	as := []Antenna{}
	for y, row := range m.Antennas {
		for x, cell := range row {
			if cell != '.' {
				as = append(as, Antenna{cell, x, y})
			}
		}
	}

	for i, ai := range as {
		for j := i + 1; j < len(as); j++ {
			aj := as[j]
			if ai.letter == aj.letter {
				for _, p := range m.getAninodes(ai, aj) {
					m.addAntinode(p)
				}
				for _, p := range m.getAninodes(aj, ai) {
					m.addAntinode(p)
				}
			}
		}
	}
	cnt := 0
	for _, row := range m.Antinodes {
		fmt.Println(string(row))
		for _, cell := range row {
			if cell == '#' {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
