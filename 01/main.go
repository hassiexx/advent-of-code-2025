package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/hassiexx/advent-of-code-2025/common"
)

func main() {
	part1()
}

func part1() {
	input := common.ReadInput(1, false)
	fmt.Println("Day 1 Part 1:", solve(input))
}

func solve(input []string) int {
	var password int
	dialPos := 50

	for _, ln := range input {
		direction := string(ln[0])
		n, _ := strconv.Atoi(ln[1:])

		switch direction {
		case "R":
			dialPos += n
			if dialPos > 99 {
				dialPos = dialPos % 100
			}
		case "L":
			dialPos -= n
			if dialPos < 0 {
				rem := int(math.Abs(float64(dialPos))) % 100
				if rem == 0 {
					dialPos = 0
				} else {
					dialPos = 100 - rem
				}
			}
		default:
			log.Fatalln("unknown direction for turning dial")
		}

		if dialPos == 0 {
			password++
		}
	}

	return password
}
