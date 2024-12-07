package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/yumike/adventofcode-2024/internal/inputs"
)

type eq struct {
	test    int
	numbers []int
}

func parse(line string) eq {
	parts := strings.Split(line, ": ")
	test, _ := strconv.Atoi(parts[0])
	numbers := []int{}
	for _, nstr := range strings.Split(parts[1], " ") {
		n, _ := strconv.Atoi(nstr)
		numbers = append(numbers, n)
	}
	return eq{test: test, numbers: numbers}
}

func ends(test int, n int) bool {
	return strings.HasSuffix(strconv.Itoa(test), strconv.Itoa(n))
}

func trim(test int, n int) int {
	s := strings.TrimSuffix(strconv.Itoa(test), strconv.Itoa(n))
	if s == "" {
		return 0
	}
	i, _ := strconv.Atoi(s)
	return i
}

func check(numbers []int, test int) bool {
	if len(numbers) == 0 {
		return test == 0
	}
	if len(numbers) == 1 {
		return numbers[0] == test
	}
	n := numbers[(len(numbers) - 1)]
	rest := numbers[:len(numbers)-1]
	if test%n == 0 && check(rest, test/n) {
		return true
	}
	if test-n >= 0 && check(rest, test-n) {
		return true
	}
	if ends(test, n) && check(rest, trim(test, n)) {
		return true
	}
	return false
}

func main() {
	path := os.Args[1]
	sum := 0
	for line := range inputs.ReadLines(path) {
		eq := parse(line)
		if check(eq.numbers, eq.test) {
			fmt.Printf("%v\n", eq)
			sum += eq.test
		}
	}
	fmt.Println(sum)
}
