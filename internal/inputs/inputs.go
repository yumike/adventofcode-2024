package inputs

import (
	"bufio"
	"iter"
	"os"
)

func ReadLines(file string) iter.Seq[string] {
	return func(yield func(string) bool) {
		file, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				break
			}
		}
	}
}
