package main

import (
	"fmt"
	"os"

	"github.com/yumike/adventofcode-2024/internal/inputs"
)

type Point struct {
	x, y int
}

func scan(m [][]byte) (Point, bool) {
	for y, row := range m {
		for x, cell := range row {
			if cell != '.' {
				return Point{x: x, y: y}, true
			}
		}
	}
	return Point{}, false
}

func same(m [][]byte, p Point, letter byte) bool {
	if p.y < 0 || p.y >= len(m) || p.x < 0 || p.x >= len(m[p.y]) {
		return false
	}
	return m[p.y][p.x] == letter
}

func sides(p Point) []Point {
	return []Point{
		{x: p.x - 1, y: p.y},
		{x: p.x, y: p.y - 1},
		{x: p.x + 1, y: p.y},
		{x: p.x, y: p.y + 1},
	}
}

func corner(m [][]byte, p Point, c Point) bool {
	if same(m, Point{x: p.x + c.x, y: p.y + c.y}, m[p.y][p.x]) {
		if !same(m, Point{x: p.x + c.x, y: p.y}, m[p.y][p.x]) &&
			!same(m, Point{x: p.x, y: p.y + c.y}, m[p.y][p.x]) {
			return true
		}
		return false
	}
	if same(m, Point{x: p.x + c.x, y: p.y}, m[p.y][p.x]) &&
		same(m, Point{x: p.x, y: p.y + c.y}, m[p.y][p.x]) {
		return true
	}
	if !same(m, Point{x: p.x + c.x, y: p.y}, m[p.y][p.x]) &&
		!same(m, Point{x: p.x, y: p.y + c.y}, m[p.y][p.x]) {
		return true
	}
	return false
}

func calc(m [][]byte, p Point) int {
	area := 0
	perimeter := 0

	letter := m[p.y][p.x]
	queue := []Point{p}
	visited := map[Point]struct{}{}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if _, ok := visited[p]; ok {
			continue
		}
		visited[p] = struct{}{}

		if p.y < 0 || p.y >= len(m) || p.x < 0 || p.x >= len(m[p.y]) {
			continue
		}

		for _, side := range sides(p) {
			if same(m, side, letter) {
				queue = append(queue, side)
			}
		}

		area++
		if corner(m, p, Point{x: -1, y: -1}) {
			perimeter++
		}
		if corner(m, p, Point{x: 1, y: -1}) {
			perimeter++
		}
		if corner(m, p, Point{x: -1, y: 1}) {
			perimeter++
		}
		if corner(m, p, Point{x: 1, y: 1}) {
			perimeter++
		}
	}

	for p := range visited {
		m[p.y][p.x] = '.'
	}
	fmt.Printf("area: %d, perimeter: %d\n", area, perimeter)
	return area * perimeter
}

func main() {
	m := [][]byte{}
	path := os.Args[1]
	for line := range inputs.ReadLines(path) {
		m = append(m, []byte(line))
	}

	price := 0
	for {
		p, ok := scan(m)
		if !ok {
			break
		}
		price += calc(m, p)
	}
	fmt.Println(price)
}
