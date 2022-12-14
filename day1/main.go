package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)

	fileScanner.Split(bufio.ScanLines)

	// scan file into these vars
	elvesCals := make([]int, 0, 10)
	elfPtr := 0
	var tmpCount int
	var cals int

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			elvesCals = append(elvesCals, tmpCount)
			elfPtr++
			tmpCount = 0
		} else {
			cals, _ = strconv.Atoi(line)
			tmpCount += cals
		}
	}

	var max int = 0

	for _, cals := range elvesCals {
		if cals > max {
			max = cals
		}
	}

	sort.Ints(elvesCals)

	var total int

	for _, amt := range elvesCals[len(elvesCals)-3:] {
		total += amt
	}

	fmt.Println("Calories of top 3 elves:", total)
}
