package main

import (
	"bufio"
	"fmt"
	"os"
)

// start 11:30pm
// stop 12:10am
// 40 minutes

// start 12:30pm
// end 1:40pm

// Ended up doing a brute force solution for part 2,
// putting obstacles only where the guard had been previously and checking if the
// new obstacle resulted in a loop

// I started by adding directions to the grid, previously it was just a single value
// but ended up not needing in... premature optimization strikes again.

// grid type
const (
	Exit     = -2
	Obstacle = -1
	Safe     = 0
)

type Dir int

const (
	Up    = 1
	Right = 2
	Down  = 4
	Left  = 8
)

type GuardPos struct {
	X, Y int
	Dir  Dir
}

const MAX_STEPS = 20000

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	grid, guard := parseInput(filename)

	blankGrid := deepCopyGrid(grid)
	startingGuard := guard

	part1, _ := solve(guard, grid)
	fmt.Println("Part 1:", part1)
	part2 := solveWithObstacles(grid, blankGrid, startingGuard)
	fmt.Println("Part 2:", part2)
	// printGrid(blankGrid, guard)
}

func deepCopyGrid(grid [][]int) [][]int {
	newGrid := make([][]int, len(grid))
	for i := range grid {
		newGrid[i] = make([]int, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	return newGrid
}

func parseInput(filename string) ([][]int, GuardPos) {
	file, _ := os.Open(filename)
	defer file.Close()

	var grid [][]int
	var guard GuardPos
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		line := scanner.Text()
		var row []int
		for _, c := range line {
			if c == '.' {
				row = append(row, Safe)
			} else if c == '^' {
				row = append(row, int(Up)) // guard direction
				guard = GuardPos{X: len(row) - 1, Y: len(grid), Dir: Up}
			} else {
				row = append(row, Obstacle)
			}
		}
		grid = append(grid, row)

	}
	return grid, guard
}

func solve(guard GuardPos, grid [][]int) (int, int) {
	count := 0

	for count < MAX_STEPS {
		// printGrid(grid, guard)
		// fmt.Println()
		nextX, nextY := nextPos(guard)
		next := gridAt(grid, nextX, nextY)
		if next == Exit {
			grid[guard.Y][guard.X] = grid[guard.Y][guard.X] | int(guard.Dir)
			break
		} else if next == Obstacle {
			// mark as patroled in a direction
			grid[guard.Y][guard.X] = grid[guard.Y][guard.X] | int(guard.Dir)
			guard.Dir = (guard.Dir * 2) % 15
		} else {
			// mark as patroled in a direction
			grid[guard.Y][guard.X] = grid[guard.Y][guard.X] | int(guard.Dir)
			guard.X, guard.Y = nextX, nextY
		}
		count++
	}
	fmt.Println("count:", count)
	return countPatrol(grid), count
}

func solveWithObstacles(grid, blankGrid [][]int, guard GuardPos) int {
	count := 0
	for y, row := range grid {
		for x, cell := range row {
			if cell > 0 {
				testGrid := deepCopyGrid(blankGrid)
				testGrid[y][x] = Obstacle
				newGuard := guard
				_, steps := solve(newGuard, testGrid)
				if steps == MAX_STEPS {
					count++
				}
			}
		}
	}
	return count
}

func printGrid(grid [][]int, guard GuardPos) {
	for y, row := range grid {
		for x, cell := range row {
			if y == guard.Y && x == guard.X {
				if guard.Dir == Up {
					fmt.Print("^")
				} else if guard.Dir == Right {
					fmt.Print(">")
				} else if guard.Dir == Down {
					fmt.Print("v")
				} else if guard.Dir == Left {
					fmt.Print("<")
				}
			} else if cell == Safe {
				fmt.Print(".")
			} else if cell == Obstacle {
				fmt.Print("#")
			} else if cell == Up || cell == Down {
				fmt.Print("|")
			} else if cell == Left || cell == Right {
				fmt.Print("-")
			} else {
				fmt.Print(cell)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func gridAt(grid [][]int, x, y int) int {
	if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[0]) {
		return Exit
	}
	return grid[y][x]
}

func nextPos(guard GuardPos) (X, Y int) {
	switch guard.Dir {
	case Up:
		return guard.X, guard.Y - 1
	case Right:
		return guard.X + 1, guard.Y
	case Down:
		return guard.X, guard.Y + 1
	case Left:
		return guard.X - 1, guard.Y
	}
	return
}

func countPatrol(grid [][]int) int {
	count := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell > 0 {
				count++
			}
		}
	}
	return count
}
