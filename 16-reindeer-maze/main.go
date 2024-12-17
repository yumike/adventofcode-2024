package main

import (
	"fmt"
	"os"

	"github.com/yumike/adventofcode-2024/internal/inputs"
)

var rotateScore = map[byte]map[byte]int{
	'>': {
		'>': 0,
		'^': 1000,
		'<': 2000,
		'v': 1000,
	},
	'^': {
		'^': 0,
		'<': 1000,
		'v': 2000,
		'>': 1000,
	},
	'<': {
		'<': 0,
		'v': 1000,
		'>': 2000,
		'^': 1000,
	},
	'v': {
		'v': 0,
		'>': 1000,
		'^': 2000,
		'<': 1000,
	},
}

type Point struct {
	X, Y int
}

type Visit struct {
	Direction byte
	Score     int
}

type Cell struct {
	Point Point
	Visit Visit
}

type QueueItem struct {
	Value Cell
	Next  *QueueItem
}

type Queue struct {
	Head *QueueItem
	Tail *QueueItem
}

func (q *Queue) Push(value Cell) {
	item := &QueueItem{Value: value}
	if q.Head == nil {
		q.Head = item
		q.Tail = item
		return
	}
	q.Tail.Next = item
	q.Tail = item
}

func (q *Queue) Pop() Cell {
	if q.Head == nil {
		return Cell{}
	}
	value := q.Head.Value
	q.Head = q.Head.Next
	return value
}

func (q *Queue) IsEmpty() bool {
	return q.Head == nil
}

func NewQueue() *Queue {
	return &Queue{}
}

type Reindeer struct {
	X         int
	Y         int
	Direction byte
	Score     int
}

func NewReindeer(x, y int) Reindeer {
	return Reindeer{
		X:         x,
		Y:         y,
		Direction: '>',
		Score:     0,
	}
}

type Step struct {
	Point     Point
	Direction byte
	Score     int
	Prev      *Step
}

func (s *Step) Cycle() bool {
	prev := s.Prev
	for prev != nil {
		if prev.Point == s.Point {
			return true
		}
		prev = prev.Prev
	}
	return false
}

type StackItem struct {
	Value Step
	Next  *StackItem
}

type Stack struct {
	Head *StackItem
}

func (s *Stack) Push(value Step) {
	item := &StackItem{Value: value}
	if s.Head == nil {
		s.Head = item
		return
	}
	item.Next = s.Head
	s.Head = item
}

func (s *Stack) Pop() (Step, bool) {
	if s.Head == nil {
		return Step{}, false
	}
	value := s.Head.Value
	s.Head = s.Head.Next
	return value, true
}

func (s *Stack) Len() int {
	cnt := 0
	item := s.Head
	for item != nil {
		cnt++
		item = item.Next
	}
	return cnt
}

func calcScore(m [][]byte, reindeer Point) map[Point]Visit {
	visits := map[Point]Visit{}
	queue := NewQueue()
	queue.Push(Cell{
		Point: Point{X: reindeer.X, Y: reindeer.Y},
		Visit: Visit{Direction: '>', Score: 0},
	})
	for !queue.IsEmpty() {
		c := queue.Pop()
		if m[c.Point.Y][c.Point.X] == '#' {
			continue
		}
		if visit, ok := visits[c.Point]; ok {
			if visit.Score <= c.Visit.Score {
				continue
			}
		}
		visits[c.Point] = c.Visit
		queue.Push(Cell{
			Point: Point{X: c.Point.X + 1, Y: c.Point.Y},
			Visit: Visit{Direction: '>', Score: c.Visit.Score + rotateScore[c.Visit.Direction]['>'] + 1},
		})
		queue.Push(Cell{
			Point: Point{X: c.Point.X, Y: c.Point.Y - 1},
			Visit: Visit{Direction: '^', Score: c.Visit.Score + rotateScore[c.Visit.Direction]['^'] + 1},
		})
		queue.Push(Cell{
			Point: Point{X: c.Point.X - 1, Y: c.Point.Y},
			Visit: Visit{Direction: '<', Score: c.Visit.Score + rotateScore[c.Visit.Direction]['<'] + 1},
		})
		queue.Push(Cell{
			Point: Point{X: c.Point.X, Y: c.Point.Y + 1},
			Visit: Visit{Direction: 'v', Score: c.Visit.Score + rotateScore[c.Visit.Direction]['v'] + 1},
		})
	}

	return visits
}

func main() {
	m := [][]byte{}

	path := os.Args[1]
	for line := range inputs.ReadLines(path) {
		m = append(m, []byte(line))
	}

	var reindeer Point
	var end Point
	for y, row := range m {
		for x, cell := range row {
			if cell == 'S' {
				reindeer = Point{X: x, Y: y}
			} else if cell == 'E' {
				end = Point{X: x, Y: y}
			}
		}
	}

	visits := calcScore(m, reindeer)
	score := visits[end].Score
	fmt.Printf("Score: %d\n", score)

	step := Step{
		Point:     reindeer,
		Direction: '>',
		Score:     0,
	}
	paths := []Step{}
	stack := Stack{}
	stack.Push(step)
	for {
		step, ok := stack.Pop()
		if !ok {
			break
		}

		if m[step.Point.Y][step.Point.X] == '#' {
			continue
		}
		if step.Score > score {
			continue
		}
		if step.Cycle() {
			continue
		}
		if step.Point == end && step.Score == score {
			paths = append(paths, step)
			continue
		}

		if visit, ok := visits[step.Point]; ok {
			if visit.Score < step.Score && visit.Direction == step.Direction {
				continue
			}
		}

		stack.Push(Step{
			Point:     Point{X: step.Point.X + 1, Y: step.Point.Y},
			Direction: '>',
			Score:     step.Score + rotateScore[step.Direction]['>'] + 1,
			Prev:      &step,
		})
		stack.Push(Step{
			Point:     Point{X: step.Point.X, Y: step.Point.Y - 1},
			Direction: '^',
			Score:     step.Score + rotateScore[step.Direction]['^'] + 1,
			Prev:      &step,
		})
		stack.Push(Step{
			Point:     Point{X: step.Point.X - 1, Y: step.Point.Y},
			Direction: '<',
			Score:     step.Score + rotateScore[step.Direction]['<'] + 1,
			Prev:      &step,
		})
		stack.Push(Step{
			Point:     Point{X: step.Point.X, Y: step.Point.Y + 1},
			Direction: 'v',
			Score:     step.Score + rotateScore[step.Direction]['v'] + 1,
			Prev:      &step,
		})
	}

	for _, step := range paths {
		s := &step
		for s != nil {
			m[s.Point.Y][s.Point.X] = 'O'
			s = s.Prev
		}
	}

	cnt := 0
	for _, row := range m {
		for _, cell := range row {
			if cell == 'O' {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
