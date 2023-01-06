package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type inst struct {
	name   string
	cycles int
	qty    int
}

type instructionSet map[string]*inst

type machine struct {
	x        int
	ip       int
	isa      instructionSet
	strenght int
	tmp      int
	target   int
	screen   [6][40]rune
}

func (m *machine) ExecOne(instruction inst) {
	for i := 0; i < instruction.cycles; i++ {
		m.ip++
		if m.ip == m.target {
			m.tmp = m.tmp + (m.ip * m.x)
			fmt.Printf("---- cycle %d x: %d strength: %d\n", m.ip, m.x, m.ip*m.x)
			m.target += 40
		}
	}
	m.x = m.x + instruction.qty
	m.strenght = m.x * m.ip
}

func (m *machine) Print() {
	// fmt.Println(len(m.monitor[0]))
	for i := 0; i < len(m.screen); i++ {
		for j := 0; j < len(m.screen[0]); j++ {
			if m.screen[i][j] == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%c", m.screen[i][j])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (m *machine) ExecTwo(instruction inst) {
	var row int
	for i := 0; i < instruction.cycles; i++ {
		m.ip++
		row = m.ip / 40
		m.screen[row][m.x] = '.'
	}
	m.x = m.x + instruction.qty
	m.strenght = (m.x * m.ip) / 60
	// col := m.x
	// fmt.Println(row, col)
	// fmt.Println("would put", row, col)
}

func partOne(input []string, m *machine) {
	for _, line := range input {
		iData := strings.Split(line, " ")

		instr := m.isa[iData[0]]
		cycles := instr.cycles
		var qty int
		if len(iData) > 1 {
			qty, _ = strconv.Atoi(iData[1])
		}

		m.ExecTwo(inst{name: instr.name, cycles: cycles, qty: qty})
	}
	m.screen[0][2] = '#'
	m.Print()
	fmt.Println("total", m.tmp)
}

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	isaSet := make(instructionSet, 2)
	isaSet["noop"] = &inst{name: "noop", cycles: 1}
	isaSet["addx"] = &inst{name: "addx", cycles: 2}

	// m := machine{x: 1, ip: 0, isa: isaSet, target: 20, screen: [6][40]rune{}}
	m := machine{x: 1, ip: 0, isa: isaSet, target: 40, screen: [6][40]rune{}}
	partOne(input, &m)
}
