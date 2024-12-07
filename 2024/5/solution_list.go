package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

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

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	fmt.Println("Part 1", solve(filename, false))
	fmt.Println("Part 2", solve(filename, true))
}

func solve(filename string, reorder bool) (sum int) {
	rules, updates := parseInput(filename)

	for i := range updates {
		var count int
		var update []int
		update = append(update, updates[i]...)
		x, y, hasErr := findError(update, rules)
		for hasErr && reorder && count < 2000 {
			fmt.Println("Reordering", x, y, hasErr, update)
			update = moveBack(update, x, y)
			x, y, hasErr = findError(update, rules)
			count++
		}
		if !hasErr {
			_, _, hasErr = findError(update, rules)
			if hasErr {
				panic(update)
			}

			sum += update[len(update)/2]
		} else {
			fmt.Println("======> Could not fix", update)
		}
	}
	return
}

func findError(update []int, rules map[int][]int) (int, int, bool) {
	// fmt.Println("FE", rules)
	// fmt.Println("FE", update)
	for i := len(update) - 1; i >= 0; i-- {
		num := update[i]
		// return the location of the first number that causes the update to be invalid
		// and the location of the number that is causing the invalidation
		y := findErrorForN(rules[num], update[:i])
		if y != -1 {
			// fmt.Println("FE Found", i, y, y+i+1)
			return i, y, true
		}
	}
	return -1, -1, false
}

func findErrorForN(rules, list []int) (y int) {
	set := make(map[int]struct{})

	// fmt.Println("FindErrorForN", rules, list)
	for _, num := range rules {
		set[num] = struct{}{}
	}

	// fmt.Println("Rules", rules, set)

	for i := len(list) - 1; i >= 0; i-- {
		num := list[i]
		if _, found := set[num]; found {
			return i
		}
	}
	return -1
}

// move x behind y in update
// x and y are indexes in update
func moveBack(update []int, y, x int) []int {
	// fmt.Println("MoveBack", x, y, update)
	var list []int
	list = append(list, update[0:x]...)
	list = append(list, update[x+1:y+1]...)
	list = append(list, update[x])
	list = append(list, update[y+1:]...)
	// fmt.Println("MoveBack", list)
	return list

}

// Prompt:
// how do I scan a file looking for 2 numbers separated by |? Or N numbers separated by commas
// now take the numbers separated by | and add them to a map, where the 2nd number is the key and the 1st number is appended to the value (int array).
// For the numbers separate by commas, convert them to a list of integers and append that to the updates list.
func parseInput(filename string) (map[int][]int, [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer file.Close()

	rules := make(map[int][]int)
	var updates [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			if len(parts) == 2 {
				num1, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
				num2, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err1 == nil && err2 == nil {
					rules[num1] = append(rules[num1], num2)
				}
			}
		} else if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			var nums []int
			for _, part := range parts {
				num, err := strconv.Atoi(strings.TrimSpace(part))
				if err == nil {
					nums = append(nums, num)
				}
			}
			if len(nums) > 0 {
				updates = append(updates, nums)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return rules, updates
}
