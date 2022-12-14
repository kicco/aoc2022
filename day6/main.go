package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func partOne(input string, markerSize int) int {
	// track window of 4 letters
	var tracked []rune
	// 1st 3 letters into tracked
	for j := 0; j < markerSize; j++ {
		char, _ := utf8.DecodeRuneInString(string(input[j]))
		tracked = append(tracked, char)
	}
	for i := markerSize; i < len(input); i++ {
		tracked = append(tracked, rune(input[i]))
		if len(tracked) > markerSize {
			tracked = tracked[1:]
		}
		letters := make(map[rune]int)
		dupe := false
		for _, c := range tracked {
			letters[c]++
			if letters[c] > 1 {
				dupe = true
			}
		}
		if !dupe {
			return i + 1
		}
	}
	return 0
}

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSpace(string(data)), "\n")
	// 4 non-repeating chars
	fmt.Println(partOne(input[0], 4))
	// 14 non-repeating chars
	fmt.Println(partOne(input[0], 14))
}
