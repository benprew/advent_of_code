package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	// Part 1
	sum := solve(filename, 'X')
	fmt.Println("Part 1 Sum:", sum)
	// Part 2
	sum = solve(filename, 'A')
	fmt.Println("Part 2 Sum:", sum)
}

func solve(filename string, ch rune) (sum int) {
	charGrid := readCharGrid(filename)

	for y, line := range charGrid {
		for x, char := range line {
			if char == ch {
				if ch == 'X' {
					sum += findXMAS(charGrid, x, y)
				} else {
					sum += findMASX(charGrid, x, y)
				}
			}
		}
	}
	return sum
}

func findXMAS(charGrid []string, x, y int) (sum int) {
	// find XMAS in 8 directions
	yLen := len(charGrid)
	xLen := len(charGrid[y])
	strs := []string{"-1", "-2", "-3", "-4", "-5", "-6", "-7", "-8"}
	if y >= 3 {
		strs[0] = string([]byte{charGrid[y][x], charGrid[y-1][x], charGrid[y-2][x], charGrid[y-3][x]})
	}
	if yLen-y > 3 {
		strs[1] = string([]byte{charGrid[y][x], charGrid[y+1][x], charGrid[y+2][x], charGrid[y+3][x]})
	}
	if x >= 3 {
		strs[2] = string([]byte{charGrid[y][x], charGrid[y][x-1], charGrid[y][x-2], charGrid[y][x-3]})
	}
	if xLen-x > 3 {
		strs[3] = string([]byte{charGrid[y][x], charGrid[y][x+1], charGrid[y][x+2], charGrid[y][x+3]})
	}
	if y >= 3 && x >= 3 {
		strs[4] = string([]byte{charGrid[y][x], charGrid[y-1][x-1], charGrid[y-2][x-2], charGrid[y-3][x-3]})
	}
	if yLen-y > 3 && xLen-x > 3 {
		strs[5] = string([]byte{charGrid[y][x], charGrid[y+1][x+1], charGrid[y+2][x+2], charGrid[y+3][x+3]})
	}
	if y >= 3 && xLen-x > 3 {
		strs[6] = string([]byte{charGrid[y][x], charGrid[y-1][x+1], charGrid[y-2][x+2], charGrid[y-3][x+3]})
	}
	if yLen-y > 3 && x >= 3 {
		strs[7] = string([]byte{charGrid[y][x], charGrid[y+1][x-1], charGrid[y+2][x-2], charGrid[y+3][x-3]})
	}
	for _, str := range strs {
		if str == "XMAS" {
			sum++
		}
	}
	return
}

func findMASX(charGrid []string, x, y int) (sum int) {
	yLen := len(charGrid)
	xLen := len(charGrid[y])

	strs := []string{"-1", "-2"}

	if y > 0 && x > 0 && y < yLen-1 && x < xLen-1 {
		strs[0] = string([]byte{charGrid[y-1][x-1], charGrid[y][x], charGrid[y+1][x+1]})
		strs[1] = string([]byte{charGrid[y-1][x+1], charGrid[y][x], charGrid[y+1][x-1]})
	}

	if (strs[0] == "MAS" || strs[0] == "SAM") && (strs[1] == "MAS" || strs[1] == "SAM") {
		sum++
	}
	return sum
}

func readCharGrid(filename string) (charGrid []string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		charGrid = append(charGrid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return charGrid
}
