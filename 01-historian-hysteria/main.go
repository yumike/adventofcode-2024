package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"

	"github.com/yumike/adventofcode-2024/internal/inputs"
)

func main() {
	path := os.Args[1]

	lefts := []int{}
	rights := []int{}

	re := regexp.MustCompile(`\s+`)
	for line := range inputs.ReadLines(path) {
		ns := re.Split(line, -1)

		left, err := strconv.Atoi(ns[0])
		if err != nil {
			panic(err)
		}
		lefts = append(lefts, left)

		right, err := strconv.Atoi(ns[1])
		if err != nil {
			panic(err)
		}
		rights = append(rights, right)
	}
	slices.Sort(lefts)
	slices.Sort(rights)

	distances := []int{}
	for i := 0; i < len(lefts); i++ {
		var d int
		if lefts[i] > rights[i] {
			d = lefts[i] - rights[i]
		} else {
			d = rights[i] - lefts[i]
		}
		distances = append(distances, d)
	}

	sum := 0
	for _, d := range distances {
		sum += d
	}
	fmt.Println("Sum:", sum)

	cnts := map[int]int{}
	for _, n := range rights {
		cnts[n]++
	}

	score := 0
	for _, n := range lefts {
		if cnt, ok := cnts[n]; ok {
			score += cnt * n
		}
	}
	fmt.Println("Score:", score)
}
