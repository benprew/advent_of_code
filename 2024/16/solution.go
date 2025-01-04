package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

type Point struct {
	x, y int
}

type Step struct {
	point Point
	dir   Direction
}

type Path struct {
	point Point
	dir   Direction
	cost  int
	steps int
}

type Path2 struct {
	point Point
	dir   Direction
	cost  int
	steps []Point
}

const START = 'S'
const END = 'E'
const WALL = '#'

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

	maze := parse(file)

	fmt.Println("Part 1:", solve(maze))
	fmt.Println("Part 2:", solve2(maze))
}

func solve(maze [][]rune) int {
	visited := make(map[Step]int, 0)
	lowestCost := -1
	start, end := findStartEnd(maze)
	queue := []Path{{start, EAST, 0, 0}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		prevCost, ok := visited[Step{path.point, path.dir}]
		if ok && prevCost < path.cost {
			continue
		}

		visited[Step{path.point, path.dir}] = path.cost

		if path.point == end {
			// fmt.Println("Found end")
			if lowestCost == -1 || path.cost < lowestCost {
				lowestCost = path.cost
			}
		}

		for _, d := range []int{-1, 0, 1} {
			direction := Direction((int(path.dir) + d) % 4)
			if direction == -1 {
				direction = 3
			}
			point := pointAt(path.point, direction)
			if maze[point.y][point.x] == WALL {
				continue
			}
			newPath := Path{point, direction, path.cost + 1, path.steps + 1}
			if d != 0 {
				newPath.cost += 1000
			}
			queue = append(queue, newPath)
		}
	}
	return lowestCost
}

// 122492 shortest path cost
func solve2(maze [][]rune) int {
	visited := make(map[Step]int, 0)
	lowestCost := -1
	solutions := []Path2{}
	start, end := findStartEnd(maze)
	queue := []Path2{{start, EAST, 0, []Point{start}}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		prevCost, ok := visited[Step{path.point, path.dir}]
		if ok && prevCost < path.cost {
			continue
		}

		visited[Step{path.point, path.dir}] = path.cost

		if path.point == end {
			if path.cost == lowestCost {
				fmt.Println("Adding path", path.cost)
				solutions = append(solutions, path)
			} else if lowestCost == -1 || path.cost < lowestCost {
				fmt.Println("Adding path", path.cost)
				solutions = []Path2{path}
				lowestCost = path.cost
			}
		}

		// add new directions
		for _, d := range []int{-1, 0, 1} {
			direction := Direction((int(path.dir) + d) % 4)
			if direction == -1 {
				direction = 3
			}
			point := pointAt(path.point, direction)
			if maze[point.y][point.x] == WALL {
				continue
			}
			newSteps := make([]Point, len(path.steps)+1)
			copy(newSteps, path.steps)
			newSteps[len(path.steps)] = point
			newPath := Path2{point, direction, path.cost + 1, newSteps}

			if d != 0 {
				newPath.cost += 1000
			}
			queue = append(queue, newPath)
		}
	}
	printGridPaths(maze, solutions)
	return countUniquePoints(solutions)
}

func printGridPaths(maze [][]rune, solutions []Path2) {
	for _, path := range solutions {
		for _, s := range path.steps {
			if maze[s.y][s.x] == WALL {
				panic("Invalid path")
			}
			maze[s.y][s.x] = 'O'
		}
	}
	printGrid(maze)
}

func countUniquePoints(paths []Path2) int {
	fmt.Println("paths:", len(paths))
	uniquePoints := make(map[Point]bool)
	for _, path := range paths {
		for _, point := range path.steps {
			uniquePoints[point] = true
		}
	}

	return len(uniquePoints)
}

func findStartEnd(maze [][]rune) (start, end Point) {
	for y := range maze {
		for x := range maze[y] {
			if maze[y][x] == START {
				start = Point{x, y}
			}
			if maze[y][x] == END {
				end = Point{x, y}
			}
		}
	}
	return
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func pointAt(p Point, d Direction) Point {
	switch d {
	case NORTH:
		return Point{p.x, p.y - 1}
	case EAST:
		return Point{p.x + 1, p.y}
	case SOUTH:
		return Point{p.x, p.y + 1}
	case WEST:
		return Point{p.x - 1, p.y}
	}
	panic(fmt.Sprintf("Invalid direction: %d", d))
}

func parse(file io.Reader) [][]rune {
	grid := [][]rune{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		grid = append(grid, []rune(line))
	}

	return grid
}
