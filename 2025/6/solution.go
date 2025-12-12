package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

	items := parse(file)

	fmt.Println("Part 1:", solve(items))
	fmt.Println("Part 2:", solve2(parse2(file)))
}

func solve(items [][]string) (total int) {
	for _, col := range items {
		op := col[len(col)-1]
		var colTotal int
		if op == "*" {
			colTotal = 1
		}
		for i, n := range col {
			// skip the operator value
			if i == len(col)-1 {
				break
			}
			if op == "*" {
				colTotal *= toi(n)
			} else {
				colTotal += toi(n)
			}
		}
		total += colTotal
	}

	return
}

func solve2(lines []string) (total int) {
	items := transpose2(lines)
	return
}

func parse(file io.Reader) (items [][]string) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		items = append(items, strings.Fields(line))
	}
	return transpose(items)
}

func parse2(file io.Reader) (lines []string) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

func transpose(in [][]string) (out [][]string) {
	if len(in) == 0 {
		return
	}

	cols := len(in[0])
	out = make([][]string, cols)
	for i := range out {
		out[i] = make([]string, len(in))
	}

	for row := range in {
		for col := range in[row] {
			out[col][row] = in[row][col]
		}
	}

	return
}

func transpose2(in []string) (out [][]string) {
	if len(in) == 0 {
		return
	}

	for i := range len(in[0]) {

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
