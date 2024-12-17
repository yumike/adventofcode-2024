package main

import (
	// "bufio"
	"fmt"
	"os"

	"github.com/yumike/adventofcode-2024/internal/inputs"
)

type Point struct {
	X, Y int
}

type Move struct {
	X, Y int
}

type Warehouse struct {
	Map   [][]byte
	Robot Point
}

func NewMove(b byte) Move {
	switch b {
	case '<':
		return Move{-1, 0}
	case '>':
		return Move{1, 0}
	case '^':
		return Move{0, -1}
	case 'v':
		return Move{0, 1}
	}
	panic(fmt.Sprintf("invalid move: %c", b))
}

func (w *Warehouse) Move(p Point, m Move, apply bool) bool {
	o := w.Map[p.Y][p.X]
	if o == '#' {
		return false
	}
	if o == '.' {
		return true
	}
	if o == '@' || o == '[' && m.Y == 0 || o == ']' && m.Y == 0 {
		ok := w.Move(Point{p.X + m.X, p.Y + m.Y}, m, apply && true)
		if !ok {
			return false
		}

		if apply {
			w.Robot = Point{p.X + m.X, p.Y + m.Y}
			w.Map[p.Y+m.Y][p.X+m.X] = o
			w.Map[p.Y][p.X] = '.'
		}
		return true
	}
	if o == '[' && m.X == 0 {
		ok1 := w.Move(Point{p.X, p.Y + m.Y}, m, false)
		ok2 := w.Move(Point{p.X + 1, p.Y + m.Y}, m, false)
		if !ok1 || !ok2 {
			return false
		}

		if apply {
			w.Move(Point{p.X, p.Y + m.Y}, m, true)
			w.Move(Point{p.X + 1, p.Y + m.Y}, m, true)
			w.Map[p.Y+m.Y][p.X] = '['
			w.Map[p.Y+m.Y][p.X+1] = ']'
			w.Map[p.Y][p.X] = '.'
			w.Map[p.Y][p.X+1] = '.'
		}
		return true
	}
	if o == ']' && m.X == 0 {
		ok1 := w.Move(Point{p.X, p.Y + m.Y}, m, false)
		ok2 := w.Move(Point{p.X - 1, p.Y + m.Y}, m, false)
		if !ok1 || !ok2 {
			return false
		}

		if apply {
			w.Move(Point{p.X, p.Y + m.Y}, m, true)
			w.Move(Point{p.X - 1, p.Y + m.Y}, m, true)
			w.Map[p.Y+m.Y][p.X-1] = '['
			w.Map[p.Y+m.Y][p.X] = ']'
			w.Map[p.Y][p.X] = '.'
			w.Map[p.Y][p.X-1] = '.'
		}
		return true
	}
	panic(fmt.Sprintf("invalid object: %c", o))
}

func (w *Warehouse) MoveRobot(m Move) bool {
	return w.Move(w.Robot, m, true)
}

func (w Warehouse) String() string {
	s := ""
	for y := 0; y < len(w.Map); y++ {
		s += string(w.Map[y]) + "\n"
	}
	return s
}

func main() {
	warehouse := Warehouse{}
	moves := []byte{}

	path := os.Args[1]
	part := "map"
	for line := range inputs.ReadLines(path) {
		if part == "map" {
			if len(line) == 0 {
				part = "moves"
				continue
			}
			row := []byte{}
			for _, c := range line {
				if c == '#' {
					row = append(row, '#', '#')
				} else if c == 'O' {
					row = append(row, '[', ']')
				} else if c == '.' {
					row = append(row, '.', '.')
				} else if c == '@' {
					row = append(row, '@', '.')
				} else {
					panic(fmt.Sprintf("invalid object: %c", c))
				}
			}
			warehouse.Map = append(warehouse.Map, row)
		} else if part == "moves" {
			moves = append(moves, []byte(line)...)
		}
	}

	for y := 0; y < len(warehouse.Map); y++ {
		for x := 0; x < len(warehouse.Map[y]); x++ {
			if warehouse.Map[y][x] == '@' {
				warehouse.Robot = Point{x, y}
			}
		}
	}

	// fmt.Println(warehouse)
	// bufio.NewReader(os.Stdin).ReadBytes('\n')

	for _, m := range moves {
		warehouse.MoveRobot(NewMove(m))
		// fmt.Println(warehouse)
		// bufio.NewReader(os.Stdin).ReadBytes('\n')
	}

	result := 0
	for y := 0; y < len(warehouse.Map); y++ {
		for x := 0; x < len(warehouse.Map[y]); x++ {
			if warehouse.Map[y][x] == '[' {
				result += 100*y + x
			}
		}
	}
	fmt.Println(result)
}
