package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/yumike/adventofcode-2024/internal/inputs"
)

func part1(m []int) int {
	b := 0
	disk := []int{}
	file := true
	for _, n := range m {
		var c int
		if file {
			c = b
			b++
		} else {
			c = -1
		}
		for i := 0; i < n; i++ {
			disk = append(disk, c)
		}
		file = !file
	}

	var n int
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] != -1 {
			n = i
			break
		}
	}

	for i := 0; i < n; i++ {
		if disk[i] != -1 {
			continue
		}
		disk[i] = disk[n]
		disk[n] = -1
		for j := n - 1; j >= 0; j-- {
			if disk[j] != -1 {
				n = j
				break
			}
		}
	}

	checksum := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			break
		}
		checksum += disk[i] * i
	}
	return checksum
}

type Space struct {
	C    int
	Size int
}

func part2(m []int) int {
	b := 0
	disk := []Space{}
	file := true
	for _, n := range m {
		var c int
		if file {
			c = b
			b++
		} else {
			c = -1
		}
		disk = append(disk, Space{C: c, Size: n})
		file = !file
	}

	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i].C == -1 {
			continue
		}
		for j := 0; j < i; j++ {
			if disk[j].C != -1 {
				continue
			}
			if disk[j].Size < disk[i].Size {
				continue
			}
			if disk[j].Size == disk[i].Size {
				disk[j].C = disk[i].C
				disk[i].C = -1
				break
			}
			space := Space{C: disk[i].C, Size: disk[i].Size}
			disk[j].Size -= disk[i].Size
			disk[i].C = -1
			disk = append(disk[:j], append([]Space{space}, disk[j:]...)...)
			i++
			break
		}
	}
	fmt.Println(disk)

	checksum := 0
	n := 0
	for i := 0; i < len(disk); i++ {
		for j := 0; j < disk[i].Size; j++ {
			if disk[i].C != -1 {
				checksum += disk[i].C * n
			}
			n++
		}
	}

	return checksum
}

func main() {
	path := os.Args[1]
	m := []int{}
	for line := range inputs.ReadLines(path) {
		for _, c := range line {
			n, _ := strconv.Atoi(string(c))
			m = append(m, n)
		}
	}
	fmt.Println(part1(m))
	fmt.Println(part2(m))
}
