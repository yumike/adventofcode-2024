package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/yumike/adventofcode-2024/internal/inputs"
)

type Device struct {
	A, B, C int
	Pointer int
	Program []int
	Out     []int
}

func (d *Device) Exec() bool {
	if d.Pointer >= len(d.Program) {
		return false
	}
	instruction := d.Program[d.Pointer]
	operand := d.Program[d.Pointer+1]
	switch instruction {
	case 0:
		d.adv(operand)
	case 1:
		d.bxl(operand)
	case 2:
		d.bst(operand)
	case 3:
		d.jnz(operand)
	case 4:
		d.bxc(operand)
	case 5:
		d.out(operand)
	case 6:
		d.bdv(operand)
	case 7:
		d.cdv(operand)
	default:
		panic(fmt.Sprintf("Unknown instruction %d", instruction))
	}
	return true
}

func (d *Device) ExecAll() {
	for {
		ok := d.Exec()
		if !ok {
			break
		}
	}
}

func (d *Device) String() string {
	outstr := []string{}
	for _, n := range d.Out {
		outstr = append(outstr, strconv.Itoa(n))
	}
	return strings.Join(outstr, ",")
}

func (d *Device) Suffix(l int) bool {
	if len(d.Out) != len(d.Program) {
		return false
	}
	for i := len(d.Out) - l; i < len(d.Out); i++ {
		if d.Out[i] != d.Program[i] {
			return false
		}
	}
	return true
}

func (d *Device) combo(operand int) int {
	if operand >= 0 && operand <= 3 {
		return operand
	}
	if operand == 4 {
		return d.A
	}
	if operand == 5 {
		return d.B
	}
	if operand == 6 {
		return d.C
	}
	panic(fmt.Sprintf("Unknown operand %d", operand))
}

func (d *Device) adv(operand int) {
	operand = d.combo(operand)
	d.A = int(math.Trunc(1.0 * float64(d.A) / math.Pow(2, float64(operand))))
	d.Pointer += 2
}

func (d *Device) bxl(operand int) {
	d.B = d.B ^ operand
	d.Pointer += 2
}

func (d *Device) bst(operand int) {
	d.B = d.combo(operand) % 8
	d.Pointer += 2
}

func (d *Device) jnz(operand int) {
	if d.A != 0 {
		d.Pointer = operand
	} else {
		d.Pointer += 2
	}
}

func (d *Device) bxc(_ int) {
	d.B = d.B ^ d.C
	d.Pointer += 2
}

func (d *Device) out(operand int) {
	d.Out = append(d.Out, d.combo(operand)%8)
	d.Pointer += 2
}

func (d *Device) bdv(operand int) {
	operand = d.combo(operand)
	d.B = int(math.Trunc(1.0 * float64(d.A) / math.Pow(2, float64(operand))))
	d.Pointer += 2
}

func (d *Device) cdv(operand int) {
	operand = d.combo(operand)
	d.C = int(math.Trunc(1.0 * float64(d.A) / math.Pow(2, float64(operand))))
	d.Pointer += 2
}

func find(base Device, a, n int) []int {
	result := []int{}
	for {
		d := Device{
			A:       a,
			B:       base.B,
			C:       base.C,
			Pointer: 0,
			Program: base.Program,
			Out:     []int{},
		}
		d.ExecAll()
		if len(d.Out) != len(d.Program) {
			break
		}
		if !d.Suffix(n - 1) {
			break
		}
		if d.Suffix(n) {
			result = append(result, a)
		}
		a += 1 << (3 * (len(d.Program) - n))
	}
	return result
}

func main() {
	device := Device{}

	path := os.Args[1]
	t := "registers"
	re := regexp.MustCompile(`Register (\w): (\d+)`)
	for line := range inputs.ReadLines(path) {
		if t == "registers" {
			if line == "" {
				t = "program"
				continue
			}
			matches := re.FindStringSubmatch(line)
			switch matches[1] {
			case "A":
				device.A, _ = strconv.Atoi(matches[2])
			case "B":
				device.B, _ = strconv.Atoi(matches[2])
			case "C":
				device.C, _ = strconv.Atoi(matches[2])
			}
		} else {
			pstr := strings.Split(line, " ")[1]
			for _, nstr := range strings.Split(pstr, ",") {
				n, _ := strconv.Atoi(nstr)
				device.Program = append(device.Program, n)
			}
		}
	}

	as := []int{1 << (3 * 15)}
	n := 1
	for {
		newas := []int{}
		for _, a := range as {
			newas = append(newas, find(device, a, n)...)
		}
		as = newas
		if n == len(device.Program) {
			break
		}
		n++
	}

	result := as[0]
	for _, a := range as {
		if a < result {
			result = a
		}
	}
	fmt.Printf("Result: %d\n", result)
}
