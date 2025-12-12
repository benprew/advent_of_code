package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	fmt.Println(filename)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := parse(file)

	fmt.Println("Part 1:", solve(lines))
	fmt.Println("Part 2:", solve2(lines))
}

// beam indexes
// read each row, encounter a ^
//  - if is in beam indexes, split++ and add 2 new indexes and remove current index
//  - else do nothing

func solve(lines []string) (splits int) {
	beamIndexes := make([]bool, len(lines[0]))

	for _, l := range lines {
		for i, n := range l {
			if n == 'S' {
				beamIndexes[i] = true

			}
			if n == '^' && beamIndexes[i] {
				beamIndexes[i] = false
				if i > 0 {
					beamIndexes[i-1] = true
				}
				if i < len(l) {
					beamIndexes[i+1] = true
				}
				splits++
			}
		}
	}

	return
}

func solve2(ranges []string) (total int) {
	return total
}

func max(i, j int) int {
	if i < j {
		return j
	} else {
		return i
	}
}

func parse(file io.Reader) (lines []string) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

func toi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
