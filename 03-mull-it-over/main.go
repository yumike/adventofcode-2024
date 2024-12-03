package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	path := os.Args[1]
	mem, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	result := 0
	do := true
	re := regexp.MustCompile(`mul\((?P<left>[0-9]{1,3}),(?P<right>[0-9]{1,3})\)|do\(\)|don't\(\)`)
	for _, match := range re.FindAllSubmatch(mem, -1) {
		instr := string(match[0])
		if instr == "do()" {
			do = true
		} else if instr == "don't()" {
			do = false
		} else if do {
			left, _ := strconv.Atoi(string(match[1]))
			right, _ := strconv.Atoi(string(match[2]))
			result += left * right
		}
	}
	fmt.Printf("Result: %d\n", result)
}
