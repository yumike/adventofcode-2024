package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/yumike/adventofcode-2024/internal/inputs"
)

type Point struct {
	X, Y int
}

type Velocity struct {
	X, Y int
}

type Robot struct {
	p Point
	v Velocity
}

type Space struct {
	width  int
	height int
	robots []Robot
}

func (s *Space) tick() {
	for i := range s.robots {
		r := &s.robots[i]
		r.p.X = (r.p.X + r.v.X + s.width) % s.width
		r.p.Y = (r.p.Y + r.v.Y + s.height) % s.height
	}
}

func (s *Space) quadrants() []int {
	qs := make([]int, 4)
	for _, r := range s.robots {
		if r.p.X < s.width/2 && r.p.Y < s.height/2 {
			qs[0]++
		} else if r.p.X > s.width/2 && r.p.Y < s.height/2 {
			qs[1]++
		} else if r.p.X < s.width/2 && r.p.Y > s.height/2 {
			qs[2]++
		} else if r.p.X > s.width/2 && r.p.Y > s.height/2 {
			qs[3]++
		}
	}
	return qs
}

func (s Space) String() string {
	m := make([][]byte, 0, s.height)
	for i := 0; i < s.height; i++ {
		m = append(m, make([]byte, s.width))
		for j := 0; j < s.width; j++ {
			m[i][j] = '.'
		}
	}
	for _, r := range s.robots {
		m[r.p.Y][r.p.X] = '#'
	}

	str := ""
	for i := 0; i < s.height; i++ {
		str += string(m[i]) + "\n"
	}
	return str
}

func main() {
	s := Space{width: 101, height: 103}

	path := os.Args[1]
	re := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
	for line := range inputs.ReadLines(path) {
		matches := re.FindStringSubmatch(line)
		if len(matches) != 5 {
			panic("invalid line")
		}
		robot := Robot{}
		robot.p.X, _ = strconv.Atoi(matches[1])
		robot.p.Y, _ = strconv.Atoi(matches[2])
		robot.v.X, _ = strconv.Atoi(matches[3])
		robot.v.Y, _ = strconv.Atoi(matches[4])
		s.robots = append(s.robots, robot)
	}

	step := 0
	for {
		if step%101 == 46 {
			sep := "----------------------------------------"
			fmt.Println(sep)
			fmt.Println("Step", step)
			fmt.Println(sep)

			fmt.Println(s)
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
		s.tick()
		step++
	}
}
