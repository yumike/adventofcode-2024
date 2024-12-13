package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/yumike/adventofcode-2024/internal/inputs"
)

type Point struct {
	X, Y int64
}

type Move struct {
	X, Y int64
}

type Machine struct {
	A, B  Move
	Prize Point
}

type Solution struct {
	I, J int64
}

func solve(m Machine) (Solution, bool) {
	il := m.B.Y*m.A.X - m.A.Y*m.B.X
	ir := m.Prize.Y*m.A.X - m.A.Y*m.Prize.X
	i := ir / il

	jl := m.A.X*m.B.Y - m.B.X*m.A.Y
	jr := m.Prize.X*m.B.Y - m.B.X*m.Prize.Y
	j := jr / jl

	if m.A.X*j+m.B.X*i == m.Prize.X && m.A.Y*j+m.B.Y*i == m.Prize.Y {
		return Solution{I: i, J: j}, true
	}
	return Solution{}, false
}

func main() {
	path := os.Args[1]
	machines := []Machine{}

	are := regexp.MustCompile(`Button A: X\+([0-9]+), Y\+([0-9]+)`)
	bre := regexp.MustCompile(`Button B: X\+([0-9]+), Y\+([0-9]+)`)
	prizere := regexp.MustCompile(`Prize: X=([0-9]+), Y=([0-9]+)`)

	t := "a"
	m := Machine{}
	for line := range inputs.ReadLines(path) {
		if t == "a" {
			match := are.FindStringSubmatch(line)
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			m.A = Move{X: int64(x), Y: int64(y)}
			t = "b"
		} else if t == "b" {
			match := bre.FindStringSubmatch(line)
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			m.B = Move{X: int64(x), Y: int64(y)}
			t = "prize"
		} else if t == "prize" {
			match := prizere.FindStringSubmatch(line)
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			m.Prize = Point{X: int64(x) + 10000000000000, Y: int64(y) + 10000000000000}
			machines = append(machines, m)
			m = Machine{}
			t = "blank"
		} else if t == "blank" {
			t = "a"
		}
	}

	var cnt int64 = 0
	for _, m := range machines {
		s, ok := solve(m)
		if ok {
			fmt.Printf("%+v\n", m)
			fmt.Println(s)
			cnt += s.J*3 + s.I
		}
	}
	fmt.Println(cnt)
}
