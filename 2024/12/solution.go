package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Point struct {
	x, y int
}

const VISITED = '0'

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
	grid := parse(file)

	fmt.Println("Part 1:", solve(grid))
	fmt.Println("Part 1:", solve2(grid))
}

func solve(grid [][]rune) (sum int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			val := grid[y][x]
			if val == VISITED {
				continue
			}
			p, a := calcPerimeter(grid, Point{x, y})
			fmt.Println("Perimeter:", p, "Area:", a, "Value:", string(val))
			sum += p * a
		}
	}
	return
}

func solve2(grid [][]rune) (sum int) {
	return
}

// look through the grid finding connected values
func calcPerimeter(grid [][]rune, p Point) (perimeter, area int) {

	value := grid[p.y][p.x]
	queue := []Point{p}
	visited := map[Point]bool{}
	i := 0
	for i < len(queue) {
		n := queue[i]
		i++
		val := grid[n.y][n.x]
		visited[n] = true

		if grid[n.y][n.x] == VISITED {
			fmt.Println("Already visited", n)
			continue
		}

		// check the ordinal directions
		// if they match value, add to queue
		// if they don't, increment permimeter
		up := Point{n.x, n.y - 1}
		down := Point{n.x, n.y + 1}
		left := Point{n.x - 1, n.y}
		right := Point{n.x + 1, n.y}

		dirs := []Point{up, down, left, right}

		grid[n.y][n.x] = VISITED
		area++

		for _, dir := range dirs {
			if dir.x < 0 || dir.y < 0 || dir.x >= len(grid[0]) || dir.y >= len(grid) {
				perimeter++
			} else if grid[dir.y][dir.x] == value {
				queue = append(queue, dir)
			} else if !visited[dir] {
				perimeter++
			}
		}
		if val == 'A' {
			fmt.Println("Perimeter:", perimeter, "Area:", area, "Value:", string(val))
		}
	}
	printGrid(grid)
	return
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func parse(file io.Reader) [][]rune {
	grid := [][]rune{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	return grid
}
