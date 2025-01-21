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

type Path struct {
	point Point
	cost  int
}

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
	points := parse(file)

	GRID_SIZE := 71

	// printGrid(points[:1024], GRID_SIZE)

	fmt.Println("Part 1:", solve(points[:1024], GRID_SIZE))
	fmt.Println("Part 2:", solve2(points, GRID_SIZE))
	// fmt.Println("Part 1:", solve2(grid, instructions))
}

func printGrid(points []Point, gridSize int) {
	wallMap := makeWallMap(points)
	for y := range gridSize {
		for x := range gridSize {
			if contains(wallMap, Point{x, y}) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func contains(wallMap map[Point]bool, p Point) bool {
	v, ok := wallMap[p]
	return ok && v
}

func solve(walls []Point, gridSize int) (sum int) {
	visited := make(map[Point]int, 0)
	lowestCost := -1
	start, end := Point{0, 0}, Point{gridSize - 1, gridSize - 1}
	queue := []Path{{start, 0}}
	wallMap := makeWallMap(walls)

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		c, ok := visited[path.point]
		if ok && c <= path.cost {
			continue
		}

		visited[path.point] = path.cost

		if path.point.x == end.x && path.point.y == end.y {
			if lowestCost == -1 || path.cost < lowestCost {
				lowestCost = path.cost
			}
		}

		dirs := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

		for _, d := range dirs {
			point := pointAdd(path.point, d)
			if contains(wallMap, point) || isOutside(point, gridSize) {
				continue
			}
			newPath := Path{point, path.cost + 1}
			queue = append(queue, newPath)
		}
	}
	// visits := []Point{}
	// for k := range visited {
	// 	visits = append(visits, k)
	// }
	// printGrid(visits, gridSize)
	return lowestCost
}

func makeWallMap(walls []Point) (wallMap map[Point]bool) {
	wallMap = make(map[Point]bool, 0)
	for _, wall := range walls {
		wallMap[wall] = true
	}
	return
}

func solve2(walls []Point, gridSize int) (sum int) {
	for i := range len(walls) - 1024 {
		sum = solve(walls[:1024+i], gridSize)
		if sum == -1 {
			return i + 1024
		}
	}
	return -1
}

func pointAdd(p1, p2 Point) Point {
	return Point{p1.x + p2.x, p1.y + p2.y}
}

func isOutside(p Point, gridSize int) bool {
	return p.x < 0 || p.x >= gridSize || p.y < 0 || p.y >= gridSize
}

func parse(file io.Reader) (points []Point) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var x, y int
		_, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		points = append(points, Point{x, y})
	}
	return
}
