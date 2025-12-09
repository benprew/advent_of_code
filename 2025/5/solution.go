package main

import (
	"bufio"
	"cmp"
	"fmt"
	"io"
	"os"
	"slices"
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

	ranges, items := parse(file)

	fmt.Println("Part 1:", solve(ranges, items))
	fmt.Println("Part 2:", solve2(ranges))
}

func solve(ranges, items []string) (fresh int) {
	for _, i := range items {
		item := toi(i)
		for _, r := range ranges {
			pieces := strings.Split(r, "-")
			min := toi(pieces[0])
			max := toi(pieces[1])
			if item >= min && item <= max {
				fmt.Println(item)
				fresh++
				break
			}
		}
	}

	return
}

// Algo:
// sort ranges ascending
// store previous range
// iterate over ranges
//   if current range is within previous range, find new range end
//   if not, add the range to list of ranges and set prevR to current range

// answer: 350780324308385
func solve2(ranges []string) (total int) {
	var iRanges [][2]int
	for _, r := range ranges {
		pieces := strings.Split(r, "-")
		min := toi(pieces[0])
		max := toi(pieces[1])
		iRange := [2]int{min, max}
		iRanges = append(iRanges, iRange)
	}

	slices.SortFunc(iRanges, func(i, j [2]int) int {
		return cmp.Compare(i[0], j[0])
	})

	var combRanges [][2]int
	var prevR [2]int

	for i := range iRanges {
		if i == 0 {
			prevR = iRanges[0]
			continue
		}
		// use index so slice prevR doesn't get overwritten when iterator changes
		r := iRanges[i]
		fmt.Println(prevR, r)

		if r[0] <= prevR[1] {
			iRange := [2]int{prevR[0], max(r[1], prevR[1])}
			prevR = iRange
		} else {
			combRanges = append(combRanges, prevR)
			prevR = r
		}
	}
	combRanges = append(combRanges, prevR)
	fmt.Println(combRanges)

	for _, r := range combRanges {
		total += 1 + r[1] - r[0]
	}

	return total
}

func max(i, j int) int {
	if i < j {
		return j
	} else {
		return i
	}
}

func parse(file io.Reader) (ranges, items []string) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		ranges = append(ranges, line)
	}
	for scanner.Scan() {
		line := scanner.Text()
		items = append(items, line)
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
