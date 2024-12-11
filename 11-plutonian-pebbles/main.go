package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/yumike/adventofcode-2024/internal/inputs"
)

func main() {
	path := os.Args[1]
	m := map[int]int{}
	for line := range inputs.ReadLines(path) {
		parts := strings.Split(line, " ")
		for i := len(parts) - 1; i >= 0; i-- {
			n, _ := strconv.Atoi(parts[i])
			if _, ok := m[n]; !ok {
				m[n] = 0
			}
			m[n]++
		}
		break
	}

	for i := 0; i < 75; i++ {
		mnext := map[int]int{}
		for n, c := range m {
			if n == 0 {
				if _, ok := mnext[1]; !ok {
					mnext[1] = 0
				}
				mnext[1] += c
			} else if nstr := strconv.Itoa(n); len(nstr)%2 == 0 {
				nstr1 := nstr[:len(nstr)/2]
				nstr2 := nstr[len(nstr)/2:]
				n1, _ := strconv.Atoi(nstr1)
				n2, _ := strconv.Atoi(nstr2)
				if _, ok := mnext[n1]; !ok {
					mnext[n1] = 0
				}
				mnext[n1] += c
				if _, ok := mnext[n2]; !ok {
					mnext[n2] = 0
				}
				mnext[n2] += c
			} else {
				n *= 2024
				if _, ok := mnext[n]; !ok {
					mnext[n] = 0
				}
				mnext[n] += c
			}
		}
		m = mnext
	}

	cnt := 0
	for _, c := range m {
		cnt += c
	}
	fmt.Println(cnt)
}
