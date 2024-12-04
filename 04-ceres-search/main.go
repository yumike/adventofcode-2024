package main

import (
	"fmt"
	"os"

	"github.com/yumike/adventofcode-2024/internal/inputs"
)

func check(ws string, mx [][]byte, x, y, xs, ys int) bool {
	w := []byte(ws)
	n := 0
	xn := x
	yn := y
	for {
		if n >= len(w) {
			return true
		}
		if xn < 0 || xn >= len(mx) || yn < 0 || yn >= len(mx[xn]) {
			return false
		}
		if mx[xn][yn] != w[n] {
			return false
		}
		n++
		xn += xs
		yn += ys
	}
}

func part1(mx [][]byte) int {
	cnt := 0
	for x, row := range mx {
		for y, cell := range row {
			if cell != 'X' {
				continue
			}
			if check("XMAS", mx, x, y, 1, 0) {
				cnt++
			}
			if check("XMAS", mx, x, y, 0, 1) {
				cnt++
			}
			if check("XMAS", mx, x, y, -1, 0) {
				cnt++
			}
			if check("XMAS", mx, x, y, 0, -1) {
				cnt++
			}
			if check("XMAS", mx, x, y, 1, 1) {
				cnt++
			}
			if check("XMAS", mx, x, y, 1, -1) {
				cnt++
			}
			if check("XMAS", mx, x, y, -1, 1) {
				cnt++
			}
			if check("XMAS", mx, x, y, -1, -1) {
				cnt++
			}
		}
	}
	return cnt
}

func part2(mx [][]byte) int {
	cnt := 0
	for x := 1; x < len(mx)-1; x++ {
		for y := 1; y < len(mx[x])-1; y++ {
			if mx[x][y] != 'A' {
				continue
			}
			if (check("MAS", mx, x-1, y-1, 1, 1) || check("MAS", mx, x+1, y+1, -1, -1)) &&
				(check("MAS", mx, x-1, y+1, 1, -1) || check("MAS", mx, x+1, y-1, -1, 1)) {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	path := os.Args[1]
	mx := [][]byte{}
	for line := range inputs.ReadLines(path) {
		mx = append(mx, []byte(line))
	}

	fmt.Println(part1(mx))
	fmt.Println(part2(mx))
}
