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

const BOX = 'O'
const WALL = '#'
const ROBOT = '@'
const EMPTY_SPACE = '.'

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
	var robot Point
	grid, instructions := parse(file)
	for i, row := range grid {
		for j, cell := range row {
			if cell == ROBOT {
				robot = Point{i, j}
			}
		}
	}

	fmt.Println(instructions)
	printGrid(grid)

	fmt.Println("Part 1:", solve(grid, robot, instructions))
	// fmt.Println("Part 1:", solve2(grid, instructions))
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func solve(grid [][]rune, robot Point, instructions string) (sum int) {
	for _, instruction := range instructions {
		grid, robot = move(grid, robot, instruction)
	}

	return scoreGrid(grid)
}

func move(grid [][]rune, robot Point, instruction rune) ([][]rune, Point) {
	var newLoc Point
	switch instruction {
	case '^':
		newLoc = Point{robot.x, robot.y - 1}
	case 'v':
		newLoc = Point{robot.x, robot.y + 1}
	case '<':
		newLoc = Point{robot.x - 1, robot.y}
	case '>':
		newLoc = Point{robot.x + 1, robot.y}
	}

	fmt.Println("Move:", string(instruction))
	fmt.Println("Robot:", robot)
	fmt.Println("NewLoc:", newLoc)
	if canMove(grid, robot, newLoc) {
		grid = moveObject(grid, robot, newLoc, ROBOT)
		grid[robot.y][robot.x] = EMPTY_SPACE
		// grid[newLoc.y][newLoc.x] = ROBOT
		robot = newLoc
	}
	if grid[newLoc.y][newLoc.x] == WALL {
		printGrid(grid)
		return grid, robot
	}
	printGrid(grid)

	return grid, robot
}

func canMove(grid [][]rune, robot Point, newLoc Point) bool {
	if grid[newLoc.y][newLoc.x] == WALL {
		return false
	}
	if grid[newLoc.y][newLoc.x] == EMPTY_SPACE {
		return true
	}
	return canMove(grid, newLoc, Point{newLoc.x + (newLoc.x - robot.x), newLoc.y + (newLoc.y - robot.y)})
}

func moveObject(grid [][]rune, robot Point, newLoc Point, object rune) [][]rune {
	if grid[newLoc.y][newLoc.x] == EMPTY_SPACE {
		grid[newLoc.y][newLoc.x] = object
		return grid
	}
	grid[newLoc.y][newLoc.x] = object
	if grid[newLoc.y][newLoc.x] == WALL {
		panic("Can't move object to wall")
	}
	return moveObject(
		grid, newLoc, Point{newLoc.x + (newLoc.x - robot.x), newLoc.y + (newLoc.y - robot.y)}, BOX,
	)
}

func scoreGrid(grid [][]rune) (sum int) {
	for i, row := range grid {
		for j, cell := range row {
			if cell == BOX {
				sum += 100*i + j
			}
		}
	}
	return sum
}

func solve2(grid [][]rune) (sum int) {
	return
}

func parse(file io.Reader) ([][]rune, string) {
	grid := [][]rune{}
	scanner := bufio.NewScanner(file)
	var instructions string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		grid = append(grid, []rune(line))
	}
	for scanner.Scan() {
		instructions += scanner.Text()
	}

	return grid, instructions
}
