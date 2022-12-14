package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func buildFS(input []string) *Fs {
	root := Node{Name: "/", IsDir: true}
	fs := Fs{Root: &root}
	fs.Path = append(fs.Path, &root)

	for _, line := range input {
		if strings.HasPrefix(line, "$") {
			cmd := strings.Split(line, " ")
			switch cmd[1] {
			case "cd":
				fs.Cd(cmd[2])
			case "ls":
				// noop
			}
		} else if strings.HasPrefix(line, "dir") {
			name := strings.Split(line, " ")
			fs.Mkdir(name[1])
		} else { // it's a file
			fd := strings.Split(line, " ")
			size, _ := strconv.Atoi(fd[0])
			fs.Touch(fd[1], size)
		}
	}

	return &fs
}

func partOne(fs *Fs) int {
	var t int
	max_size := 1000000
	for _, s := range fs.Root.Search(max_size) {
		t += s.Bytes
	}
	return t
}

func partTwo(fs *Fs) int {
	used := fs.Root.Size()
	total := 70000000
	needed := 30000000
	unused := total - used
	reclaim := needed - unused

	fmt.Println("reclaim", reclaim)

	// just get ALL dirs ds and find the best match. I'm tired
	ds := fs.Root.Dirsizes()
	sort.Ints(ds)
	i := sort.Search(len(ds), func(i int) bool { return ds[i] >= reclaim })
	return ds[i]
}

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSpace(string(data)), "\n")
	fs := buildFS(input)
	fmt.Println("total bytes for dirs smaller than requested size", partOne(fs))
	fmt.Println("bytes total for dirs to be deleted", partTwo(fs))
}
