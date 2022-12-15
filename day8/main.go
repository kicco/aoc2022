package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func buildSchema(input []string) [][]int {
	var schema [][]int
	for _, r := range input {
		var row []int
		for _, c := range strings.Split(r, "") {
			num, _ := strconv.Atoi(c)
			row = append(row, num)
		}
		schema = append(schema, row)
	}
	return schema
}

func printSchema(schema [][]int) {
	for i := 0; i <= len(schema)-1; i++ {
		for j := 0; j <= len(schema[0])-1; j++ {
			fmt.Printf("%d", schema[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func look(direction []int, treeHeight int, invert bool) (bool, int) {
	d := make([]int, len(direction))
	copy(d, direction) // can't fuck with the original
	if invert {
		for i, j := 0, len(d)-1; i < j; i, j = i+1, j-1 {
			d[i], d[j] = d[j], d[i]
		}
	}
	tallerThan := 1
	for _, v := range d {
		if treeHeight <= v {
			return false, tallerThan
		}
		tallerThan++
	}
	return true, tallerThan - 1
}

func treeInfo(i int, j int, grid [][]int) (bool, int) {
	h := grid[i][j]
	col := make([]int, len(grid[0]))
	for c := range grid {
		col[c] = grid[c][j]
	}
	l, ls := look(grid[i][:j], h, true)    // left
	r, rs := look(grid[i][j+1:], h, false) // right
	t, ts := look(col[:i], h, true)        // top
	b, bs := look(col[i+1:], h, false)     // bottom
	return l || r || t || b, ls * rs * ts * bs
}

func partOne(input []string) {
	schema := buildSchema(input)
	visibleCount := (len(schema)*2 + len(schema[0])*2) - 4 // perimeter
	var score int
	for i := 1; i <= len(schema)-2; i++ {
		for j := 1; j <= len(schema[0])-2; j++ {
			visible, scenic := treeInfo(i, j, schema)
			if scenic > score {
				score = scenic
			}
			if visible {
				visibleCount++
			}
		}
	}
	fmt.Println("there's", visibleCount, "visible trees")
	fmt.Println("max scenic score", score)
}

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	partOne(lines)
}
