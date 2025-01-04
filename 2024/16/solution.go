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
	// fmt.Println("Part 1:", solve2(grid, instructions))
}

func solve(maze [][]rune) int {
	visited := make(map[Step]int, 0)
	lowestCost := -1
	depth := 0
	start, end := findStartEnd(maze)
	queue := []Path{{start, EAST, 0, 0}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		if path.steps > depth {
			depth = path.steps
			fmt.Println("Depth:", depth)
		}

		prevCost, ok := visited[Step{path.point, path.dir}]
		if ok && prevCost < path.cost {
			continue
		}

		visited[Step{path.point, path.dir}] = path.cost

		if path.point == end {
			fmt.Println("Found end")
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
func solve2(maze [][]rune) (sum int) {
	return
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
