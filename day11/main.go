package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operation struct {
	operand  string
	operator int
}

type monkey struct {
	id           int
	items        []int
	operation    operation
	divisibility int
	ifTrue       int
	ifFalse      int
	inspections  int
}

func (m *monkey) throw(item int, dest *monkey) {
	// fmt.Println("monkey", m.id, "throws", item, "to", dest.id)
	dest.catch(item)
	m.items = m.items[1:]
}

func (m *monkey) catch(item int) {
	if item == 0 {
		return
	}
	m.items = append(m.items, item)
}

func parseBlock(block []string) *monkey {
	m := monkey{}
	var monkeyId int
	var itemsString string

	fmt.Sscanf(block[0], "%s %d:", &itemsString, &monkeyId)
	m.id = monkeyId

	itemsLine := strings.Split(strings.TrimSpace(block[1]), ":")
	itemsStr := strings.Split(strings.TrimSpace(itemsLine[1]), ", ")
	for _, itemInt := range itemsStr {
		intItem, _ := strconv.Atoi(itemInt)
		m.items = append(m.items, intItem)
	}

	var op operation
	fmt.Sscanf(block[2], "  Operation: new = old %s %d", &op.operand, &op.operator)
	m.operation = op

	fmt.Sscanf(block[3], "  Test: divisible by %d", &m.divisibility)

	fmt.Sscanf(block[4], "    If true: throw to monkey %d", &m.ifTrue)
	fmt.Sscanf(block[5], "    If false: throw to monkey %d", &m.ifFalse)

	return &m
}

func partOne(ms []*monkey) {
	for i := 0; i < 10000; i++ {
		for _, m := range ms {
			for _, it := range m.items {
				m.inspections++
				op := m.operation.operator
				var wl int
				switch m.operation.operand {
				case "+":
					wl = it + op
				case "-":
					wl = it - op
				case "*":
					if m.operation.operator == 0 {
						wl = it * it
					} else {
						wl = it * op
					}
				case "/":
					wl = it / op
				}
				// wl = wl / 3
				if wl%m.divisibility == 0 {
					m.throw(wl, ms[m.ifTrue])
				} else {
					m.throw(wl, ms[m.ifFalse])
				}
			}
		}
		fmt.Println("End of Round", i+1)
		for _, m := range ms {
			fmt.Printf("monkey id: %d inspections: %d\n", m.id, m.inspections)
		}
	}
}

// low 14413975108
// high 16858797605

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	var monkeys []*monkey
	idx := 0
	for {
		monkeys = append(monkeys, parseBlock(lines[idx:idx+6]))
		idx += 7
		if idx > len(lines) {
			break
		}
	}
	partOne(monkeys)
}
