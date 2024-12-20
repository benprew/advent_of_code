package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	blocksStr, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	parsedData := parseGrid(string(blocksStr))
	// fmt.Println("Parsed Data:", parsedData)

	fmt.Println("Part 1:", solve(parsedData))
	fmt.Println("Part 2:", solve2(parsedData))
}

// parseGrid parses the input string into a list of lists of integers
func parseGrid(input string) [][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := make([][]int, len(lines))

	for i, line := range lines {
		result[i] = make([]int, len(line))
		for j, char := range line {
			if char == '.' {
				result[i][j] = -1 // or any other value to represent empty spaces
			} else {
				num, err := strconv.Atoi(string(char))
				if err != nil {
					panic(err)
				}
				result[i][j] = num
			}
		}
	}

	return result
}

type Point struct {
	x, y int
}

// solve and solve2 functions should be updated to accept [][]int as input
func solve(data [][]int) (sum int) {
	sums := pathSets(data)
	for _, set := range sums {
		unique := map[Point]bool{}
		for _, v := range set {
			unique[v] = true
		}
		// fmt.Println(len(unique), unique)
		sum += len(unique)
	}

	return sum
}

func solve2(data [][]int) (sum int) {
	sums := pathSets(data)
	for _, set := range sums {
		sum += len(set)
	}
	return
}

func pathSets(data [][]int) (sums [][]Point) {
	// var sums [][]Point
	for y, row := range data {
		for x := range row {
			if data[y][x] == 0 {
				sums = append(sums, countPaths(data, Point{x, y}, 0))
			}
		}
	}
	return sums
}

func countPaths(data [][]int, p Point, curr int) (sum []Point) {
	sum = []Point{}
	// look at the 4 ordinal directions and call countPaths on them
	up := Point{p.x, p.y - 1}
	down := Point{p.x, p.y + 1}
	left := Point{p.x - 1, p.y}
	right := Point{p.x + 1, p.y}

	directions := []Point{up, down, left, right}

	// fmt.Println("Current Point:", p, "Current Value:", curr)

	for _, d := range directions {
		if d.x < 0 || d.y < 0 || d.x >= len(data[0]) || d.y >= len(data) {
			continue
		}

		if data[d.y][d.x] == -1 {
			continue
		}

		if data[d.y][d.x] == 9 && curr == 8 {
			sum = append(sum, Point{d.x, d.y})
		}

		if data[d.y][d.x] == curr+1 {
			sum = append(sum, countPaths(data, d, curr+1)...)
		}
	}
	return sum
}
