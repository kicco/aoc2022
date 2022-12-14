package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getSpaces(couple []string) ([]string, []string) {
	return strings.Split(couple[0], "-"), strings.Split(couple[1], "-")
}

func toInt(space []string) (start int, end int) {
	start, _ = strconv.Atoi(space[0])
	end, _ = strconv.Atoi(space[1])
	return
}

func partOne(input []string) (overlapping int) {
	for _, couple := range input {
		leftR, rightR := getSpaces(strings.Split(couple, ","))

		leftStart, leftEnd := toInt(leftR)
		rightStart, rightEnd := toInt(rightR)

		if rightStart >= leftStart && rightEnd <= leftEnd {
			overlapping++
		} else if leftStart >= rightStart && leftEnd <= rightEnd {
			overlapping++
		}
	}
	return
}

func partTwo(input []string) (overlapping int) {
	for _, couple := range input {
		leftR, rightR := getSpaces(strings.Split(couple, ","))

		leftStart, leftEnd := toInt(leftR)
		rightStart, rightEnd := toInt(rightR)

		if rightStart >= leftStart && rightStart <= leftEnd {
			overlapping++
		} else if leftStart >= rightStart && leftStart <= rightEnd {
			overlapping++
		}

	}
	return
}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	fmt.Println("PartOne overlapping spaces:", partOne(lines))
	fmt.Println("PartTwo overlapping spaces:", partTwo(lines))
}
