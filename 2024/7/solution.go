package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Part 1
// start 9:30pm
// finish 9:52pm
//
// Part 2
// start 9:52pm
// finish 10:30pm

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	// if you only have 2 values you can solve it with a binary string
	// in a for loop. IE. to generate all combinations of a length k list
	// you need to iterate 2^k times.
	// You can generate the binary string using Sprintf "%0*b"
	fmt.Println("Part 1:", solve(filename, []int{0, 1}))
	fmt.Println("Part 2:", solve(filename, []int{0, 1, 2}))
}

func solve(filename string, set []int) (calibrations int64) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading file", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var numbers []int
		var answer int

		line := scanner.Text()
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		answer, err = strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			fmt.Println("Error parsing number:", err)
			continue
		}

		numberStrings := strings.Split(parts[1], " ")
		for _, numStr := range numberStrings {
			num, err := strconv.Atoi(strings.TrimSpace(numStr))
			if err != nil {
				fmt.Println("Error parsing number:", err)
				continue
			}
			numbers = append(numbers, num)
		}

		fmt.Println("Numbers", numbers)

		bitSize := len(numbers) - 1
		for _, perm := range generateCombinations(set, bitSize) {
			result := numbers[0]

			for j, bit := range perm {
				if bit == 0 {
					result += numbers[j+1]
				} else if bit == 1 {
					result *= numbers[j+1]
				} else {
					combined := fmt.Sprintf("%d%d", result, numbers[j+1])
					n, _ := strconv.Atoi(strings.TrimSpace(combined))
					result = n
				}
			}
			if result == answer {
				fmt.Println("Result:", result, "Answer:", answer, "Binary:", perm)
				calibrations += int64(answer)
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return
}

func generateCombinations(S []int, size int) [][]int {
	var result [][]int
	comb := make([]int, size)
	generate(S, comb, 0, size, &result)
	return result
}

func generate(S []int, comb []int, index int, size int, result *[][]int) {
	if index == size {
		combCopy := make([]int, size)
		copy(combCopy, comb)
		*result = append(*result, combCopy)
		return
	}

	for _, val := range S {
		comb[index] = val
		generate(S, comb, index+1, size, result)
	}
}
