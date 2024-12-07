package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Part 1
// start 9pm
// finish 9:18pm
// Part 2
// start 10:00pm
// stop 11:42pm
// Part 2 9453 -- too high
// Part 2 9206 -- too high still
// no more invalid lists
// start 5 minutes
// start 7:41pm
// end 10:40pm
// I realized I needed to only add the rows that were reordered, not all of them :( :( :(

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	part1, part2 := solve(filename)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func solve(filename string) (part1, part2 int) {
	rules, updates := parseInput(filename)
	// fmt.Println(rules)

	cmp := func(i, j int) int {
		if rules[i][j] == 1 {
			return -1
		} else if rules[j][i] == 1 {
			return 1
		}
		return 0
	}

	for _, update := range updates {
		orig := append([]int{}, update...)
		slices.SortFunc(update, cmp)
		fmt.Println(update)

		middle := update[len(update)/2]

		if slices.IsSortedFunc(orig, cmp) {
			part1 += middle
		} else {
			part2 += middle
		}
	}

	return
}

// Prompt:
// how do I scan a file looking for 2 numbers separated by |? Or N numbers separated by commas
// now take the numbers separated by | and add them to a map, where the 2nd number is the key and the 1st number is appended to the value (int array).
// For the numbers separate by commas, convert them to a list of integers and append that to the updates list.
func parseInput(filename string) (map[int]map[int]int, [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer file.Close()

	rules := make(map[int]map[int]int)
	var updates [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "|") {
			var before, after int
			fmt.Sscanf(line, "%d|%d", &before, &after)
			if _, exists := rules[before]; !exists {
				rules[before] = make(map[int]int)
			}
			rules[before][after] = 1
		} else if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			var nums []int
			for _, part := range parts {
				num, _ := strconv.Atoi(strings.TrimSpace(part))
				nums = append(nums, num)
			}
			updates = append(updates, nums)
		} else if line == "" {
			continue
		} else {
			fmt.Println(line)
			panic("unknown line")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return rules, updates
}
