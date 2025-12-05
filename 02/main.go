package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hassiexx/advent-of-code-2025/common"
)

func main() {
	part1()
	part2()
}

func part1() {
	input := common.ReadInput(2, false)
	fmt.Println("Day 2 Part 1:", solve(input, false))
}

func part2() {
	input := common.ReadInput(2, false)
	fmt.Println("Day 2 Part 2:", solve(input, true))
}

func solve(input []string, atLeastTwice bool) int {
	var total int

	idRanges := strings.Split(input[0], ",")

	for _, idRange := range idRanges {
		spl := strings.Split(idRange, "-")
		start, _ := strconv.Atoi(spl[0])
		end, _ := strconv.Atoi(spl[1])

	outer:
		for i := start; i < end+1; i++ {
			s := strconv.Itoa(i)

			if atLeastTwice {
				for j := 1; j < len(s)/2+1; j++ {
					sr := strings.ReplaceAll(s, s[:j], "")
					if sr == "" {
						total += i
						continue outer
					}
				}
			}

			if len(s)%2 != 0 {
				continue
			}

			if s[0:len(s)/2] != s[len(s)/2:] {
				continue
			}

			total += i
		}
	}

	return total
}
