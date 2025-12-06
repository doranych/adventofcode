package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	inputSlc := strings.Split(input, "\n")
	resOver0 := 0
	res := 0
	curV := 50
	for _, act := range inputSlc {
		n, _ := strconv.Atoi(act[1:])
		if n > 100 {
			resOver0 += n / 100
			n = n % 100
		}

		if act[0] == 'L' {
			if curV != 0 && curV-n < 0 {
				resOver0++
			}
			curV = (curV + 100 - n) % 100
		} else {
			if curV+n > 100 {
				resOver0++
			}
			curV = (curV + n) % 100
		}
		if curV == 0 {
			res++
		}
	}
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return resOver0 + res
	}
	return res
}
