package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	file, err := os.Open("./01-historian-hysteria/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lefts := []int{}
	rights := []int{}

	// Read file line by line
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	re := regexp.MustCompile(`\s+`)
	for scanner.Scan() {
		line := scanner.Text()
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
