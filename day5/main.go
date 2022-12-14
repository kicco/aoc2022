package main

import (
	"fmt"
	"os"
	"strings"
)

type stack map[int][]rune

func readStack() (s stack) {
	return stack{
		1: {'G', 'T', 'R', 'W'},
		2: {'G', 'C', 'H', 'P', 'M', 'S', 'V', 'W'},
		3: {'C', 'L', 'T', 'S', 'G', 'M'},
		4: {'J', 'H', 'D', 'M', 'W', 'R', 'F'},
		5: {'P', 'Q', 'L', 'H', 'S', 'W', 'F', 'J'},
		6: {'P', 'J', 'D', 'N', 'F', 'M', 'S'},
		7: {'Z', 'B', 'D', 'F', 'G', 'C', 'S', 'J'},
		8: {'R', 'T', 'B'},
		9: {'H', 'N', 'W', 'L', 'C'},
	}
	// return stack{
	// 	1: {'Z', 'N'},
	// 	2: {'M', 'C', 'D'},
	// 	3: {'P'},
	// }
}

func parseOrder(order string) (count, from, to int) {
	fmt.Sscanf(order, "move %d from %d to %d", &count, &from, &to)
	return
}

func partOne(s stack, orders []string) {
	for _, order := range orders {
		amount, from, to := parseOrder(order)
		s.Move3000(amount, from, to)
	}
}

func partTwo(s stack, orders []string) {
	for _, order := range orders {
		amount, from, to := parseOrder(order)
		s.Move3001(amount, from, to)
	}
}

func (s stack) Move3000(amount, from, to int) {
	for i := 0; i < amount; i++ {
		n := len(s[from]) - 1
		s[to] = append(s[to], s[from][n])
		s[from] = s[from][:n]
	}
}

func (s stack) Move3001(amount, from, to int) {
	n := len(s[from]) - amount
	s[to] = append(s[to], s[from][n:]...)
	s[from] = s[from][:n]
}

func main() {
	ordersI, _ := os.ReadFile("orders.txt")
	orders := strings.Split(strings.TrimSpace(string(ordersI)), "\n")
	s := readStack()
	// partOne(s, orders)
	partTwo(s, orders)

	// print top container
	var res string
	for idx := 1; idx <= len(s); idx++ {
		res += string(s[idx][len(s[idx])-1])
	}
	fmt.Println("Part 1: top container for every stack:", res)
}
