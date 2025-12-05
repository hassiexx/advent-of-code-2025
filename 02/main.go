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

	idRanges := strings.SplitSeq(input[0], ",")

	for idRange := range idRanges {
		spl := strings.Split(idRange, "-")
		start, _ := strconv.Atoi(spl[0])
		end, _ := strconv.Atoi(spl[1])
		end += 1

	outer:
		for i := start; i < end; i++ {
			s := strconv.Itoa(i)
			slen := len(s)

			if s[0:slen/2] == s[slen/2:] {
				total += i
				continue
			}

			if atLeastTwice {
				term := slen/2 + 1
				for j := 1; j < term; j++ {
					sr := strings.ReplaceAll(s, s[:j], "")
					if sr == "" {
						total += i
						continue outer
					}
				}
			}
		}
	}

	return total
}
