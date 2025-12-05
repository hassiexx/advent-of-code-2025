package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadInput(day int, test bool) []string {
	f, err := os.Open(filename(day, test))
	if err != nil {
		panic(fmt.Errorf("failed to read input: %w", err))
	}

	defer f.Close()

	lines := make([]string, 0, 1000)

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		panic(fmt.Errorf("failed to read input: %w", err))
	}

	return lines
}

func filename(day int, test bool) string {
	filename := strconv.Itoa(day)
	if day < 10 {
		filename = "0" + filename
	}

	filename += "-input"

	if test {
		filename += "-test"
	}

	filename += ".txt"

	return "../input/" + filename
}
