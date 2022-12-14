package main

import (
	"fmt"
	"os"
	"strings"
)

func charToInt(c rune) int {
	if int(c) > 90 {
		// lower case
		return int(c) - 96
	} else {
		// upper case
		return int(c) - 38
	}
}

func partOne(rucksacks []string) (sum int) {

	for _, rs := range rucksacks {
		left := rs[:len(rs)/2]
		right := rs[len(rs)/2:]
		for _, c := range left {
			if strings.Contains(right, string(c)) {
				sum += charToInt(c)
				break
			}
		}
	}
	return
}

func partTwo(rucksack []string) (sum int) {
	m := make(map[rune][3]int)
	partTwoTotal := 0
	groupPointer := 0
	for _, s := range rucksack {
		for _, c := range s {
			if val, ok := m[c]; ok {
				val[groupPointer] = 1
				m[c] = val
				if val[0] == 1 && val[1] == 1 && val[2] == 1 {
					partTwoTotal += charToInt(c)
					break
				}
			} else {
				m[c] = [3]int{0, 0, 0}
				val := m[c]
				val[groupPointer] = 1
				m[c] = val
			}
		}
		groupPointer += 1
		// reset every 3 strings
		if groupPointer == 3 {
			groupPointer = 0
			m = make(map[rune][3]int)
		}
	}
	return partTwoTotal
} // 2644

func main() {
	input, _ := os.ReadFile("input.txt")
	rucksacks := strings.Split(strings.TrimSpace(string(input)), "\n")
	fmt.Println("Part 1 total", partOne(rucksacks))
	fmt.Println("Part 2 total: ", partTwo(rucksacks))
}
