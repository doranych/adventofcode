package main

import (
	"fmt"
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
	intervals := strings.Split(input, ",")
	res := 0
	for _, r := range intervals {
		rng := strings.Split(r, "-")
		low, _ := strconv.Atoi(rng[0])
		high, _ := strconv.Atoi(rng[1])
		for i := low; i <= high; i++ {
			val := fmt.Sprintf("%d", i)
			if !part2 {
				if val[0:len(val)/2] == val[len(val)/2:] {
					res += i
				}
				continue
			}

			for j := 0; j <= len(val)/2; j++ {
				p := val[0 : j+1]
				if len(val)%len(p) != 0 {
					continue
				}
				count := 1
				for c := len(p); c < len(val); c += len(p) {
					if val[c:c+len(p)] != p {
						break
					}
					count++
				}
				if count != 1 && count == len(val)/len(p) {
					res += i
					break
				}
			}
		}
	}
	return res
}
