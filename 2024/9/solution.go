package main

import (
	"fmt"
	"os"
	"strconv"
)

// part 1
// start 10pm
// finish 11:30pm

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	// if you only have 2 values you can solve it with a binary string
	// in a for loop. IE. to generate all combinations of a length k list
	// you need to iterate 2^k times.
	// You can generate the binary string using Sprintf "%0*b"
	fmt.Println("Part 1:", solve(filename))
	// fmt.Println("Part 2:", solve2(filename))
}

func solve(filename string) int {
	values := parse(filename)

	start := 0
	end := len(values) - 1

	for start < end {
		if values[start] == -1 {
			for values[end] == -1 {
				end--
			}
			if end < start {
				break
			}
			values[start], values[end] = values[end], values[start]
		}
		start++
	}

	// validate
	if !validate(values) {
		// fmt.Println(values)
		fmt.Println("start", start, "end", end)
		panic("invalid")
	}

	checkSum := 0
	for i := range values {
		if values[i] == -1 {
			break
		}
		checkSum += values[i] * i
	}
	return checkSum
}

func validate(values []int) (valid bool) {
	valid = true
	end := false
	for i := range values {
		if end && values[i] != -1 {
			fmt.Println("invalid", i, values[i])
			valid = false
		}
		if values[i] == -1 {
			end = true
		}
	}
	return valid
}

func parse(filename string) (blocks []int) {
	blocksStr, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	for i, c := range string(blocksStr) {
		val, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}

		fileId := -1
		if i%2 == 0 {
			fileId = i / 2
		}

		for range val {
			blocks = append(blocks, fileId)
		}
	}

	return blocks
}
