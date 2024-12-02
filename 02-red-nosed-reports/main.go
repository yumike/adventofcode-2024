package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func isSafe(levels []int) bool {
	if len(levels) == 1 {
		return true
	}

	var ord int
	if levels[0] < levels[1] {
		ord = 1
	} else {
		ord = -1
	}

	for i := 0; i+1 < len(levels); i++ {
		first := levels[i]
		second := levels[i+1]
		diff := (second - first) * ord
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("./02-red-nosed-reports/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reports := [][]int{}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	re := regexp.MustCompile(`\s+`)
	for scanner.Scan() {
		line := scanner.Text()
		levelStrs := re.Split(line, -1)
		fmt.Println(levelStrs)

		levels := []int{}
		for _, levelStr := range levelStrs {
			level, err := strconv.Atoi(levelStr)
			if err != nil {
				panic(err)
			}
			levels = append(levels, level)
		}
		reports = append(reports, levels)
	}

	cnt := 0
	for _, levels := range reports {
		safe := isSafe(levels)
		if safe {
			cnt++
			continue
		}

		for i := 0; i < len(levels); i++ {
			fixedLevels := append([]int{}, levels[:i]...)
			fixedLevels = append(fixedLevels, levels[i+1:]...)
			safe = isSafe(fixedLevels)
			fmt.Println(levels, fixedLevels, safe)
			if safe {
				cnt++
				break
			}
		}
	}
	fmt.Println("Safe: ", cnt)
}
