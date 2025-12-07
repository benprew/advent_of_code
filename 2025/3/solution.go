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

	batteries := parse(file)

	fmt.Println("Part 1:", solve(batteries))
	fmt.Println("Part 2:", solve2(batteries))
}

func solve(batteries []string) (total int) {
	for _, battery := range batteries {
		max := 0
		for i, strN := range battery {
			n := toi(string(strN))
			if n*10 > max && i < len(battery)-1 {
				max = n * 10
			} else if max/10*10+n > max {
				max = max/10*10 + n
			}
		}
		total += max
	}
	return total
}

func solve2(batteries []string) (total int) {
	for _, battery := range batteries {
		max := [12]int{}
		for i, strN := range battery {
			n := toi(string(strN))
			for j, maxN := range max {
				// fmt.Println(n, maxN, len(battery)-i, 12-j)
				if n > maxN && len(battery)-i >= 12-j {
					max[j] = n
					// zero out the remainder since it's no longer valid
					for k := j + 1; k < len(max); k++ {
						max[k] = 0
					}
					break
				}
			}
		}
		maxStr := ""
		for _, n := range max {
			maxStr += strconv.Itoa(n)
		}
		fmt.Println(battery, maxStr)
		total += toi(maxStr)
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

func parse(file io.Reader) (lines []string) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return
}
