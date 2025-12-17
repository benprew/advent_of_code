package main

import (
	"advent_of_code/utils"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	items := parse(file)
	file.Seek(0, 0)

	fmt.Println("Part 1:", solve(items))
	fmt.Println("Part 2:", solve(parse2(file)))
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
				colTotal *= utils.Toi(n)
			} else {
				colTotal += utils.Toi(n)
			}
		}
		total += colTotal
	}

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

func parse2(file io.Reader) (items [][]string) {
	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return transpose2(lines)
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

	var nums []string

	// fmt.Println(in)

	length := len(in[0])

	for i := range length {
		x := length - i - 1
		var sb strings.Builder

		for y := range in {
			sb.WriteString(string(in[y][x]))
		}
		str := sb.String()
		fmt.Println(str, nums)
		last := string(str[len(str)-1])
		if last == "*" || last == "+" {
			nums = append(nums, str[:len(str)-1])
			nums = append(nums, last)
			out = append(out, strings.Fields(strings.Join(nums, " ")))
			nums = []string{}
		} else {
			nums = append(nums, str)
		}
	}
	fmt.Println(out)
	return out
}
