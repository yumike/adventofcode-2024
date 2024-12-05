package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/yumike/adventofcode-2024/internal/inputs"
)

func main() {
	path := os.Args[1]
	rules := map[int][]int{}
	updates := [][]int{}
	rulesSection := true
	for line := range inputs.ReadLines(path) {
		if line == "" {
			rulesSection = false
		} else if rulesSection {
			pages := strings.Split(line, "|")
			curr, _ := strconv.Atoi(pages[0])
			next, _ := strconv.Atoi(pages[1])
			rules[curr] = append(rules[curr], next)
		} else {
			pages := strings.Split(line, ",")
			update := []int{}
			for _, page := range pages {
				page, _ := strconv.Atoi(page)
				update = append(update, page)
			}
			updates = append(updates, update)
		}
	}

	correctSum := 0
	incorrectSum := 0
	for _, update := range updates {
		prevs := []int{}
		valid := true
		for _, page := range update {
			for _, prev := range prevs {
				for _, next := range rules[page] {
					if prev == next {
						valid = false
						break
					}
				}
			}
			prevs = append(prevs, page)
		}

		if valid {
			correctSum += update[len(update)/2]
		} else {
			slices.SortFunc(update, func(a, b int) int {
				for _, p := range rules[a] {
					if p == b {
						return 1
					}
				}
				for _, p := range rules[b] {
					if p == a {
						return -1
					}
				}
				return 0
			})
			incorrectSum += update[len(update)/2]
		}
	}
	fmt.Println(correctSum)
	fmt.Println(incorrectSum)
}
