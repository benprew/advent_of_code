package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

type Point struct {
	x, y int
}

const START = 'S'
const END = 'E'
const WALL = '#'
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

	grid, start, end := parse(file)

	fmt.Println("Start:", start)
	fmt.Println("End:", end)
	printGrid(grid)

	SAVINGS_MIN := 100

	fmt.Println("Part 1:", solve(grid, start, end, SAVINGS_MIN))
	fmt.Println("Part 2:", solve2(grid, start, end, SAVINGS_MIN))
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func solve(grid [][]rune, start Point, end Point, save_min int) (sum int) {
	// Implement BFS to find the shortest path without cheating
	costs := bfs(grid, start, end)
	fmt.Println("Shortest Path:", costs)

	// Extract keys and values from the costs map
	type kv struct {
		Key   Point
		Value int
	}
	var sortedCosts []kv
	for k, v := range costs {
		sortedCosts = append(sortedCosts, kv{k, v})
	}

	// Sort the keys based on their corresponding values in descending order
	sort.Slice(sortedCosts, func(i, j int) bool {
		return sortedCosts[i].Value > sortedCosts[j].Value
	})

	// Print the sorted keys and values
	fmt.Println("Costs sorted by values (descending):")
	for _, kv := range sortedCosts {
		fmt.Printf("Point: %v, Cost: %d\n", kv.Key, kv.Value)
	}

	// Calculate savings for each possible cheat
	cheats := calculateCheats(grid, start, end, costs)

	// sort the cheats in ascending order
	sort.Ints(cheats)

	fmt.Println("Cheats:", cheats)
	// Count the number of cheats that save at least 100 picoseconds

	for _, cheat := range cheats {
		if cheat >= save_min {
			sum++
		}
	}

	return sum
}

func solve2(grid [][]rune, start Point, end Point, save_min int) (sum int) {
	// Implement BFS to find the shortest path without cheating
	costs := bfs(grid, start, end)
	// Calculate savings for each possible cheat
	cheats := calculateCheats(grid, start, end, costs)
	for _, cheat := range cheats {
		if cheat >= save_min {
			sum++
		}
	}
	return
}

func bfs(grid [][]rune, start Point, end Point) map[Point]int {
	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := []Point{end}
	visited := make(map[Point]bool)
	visited[end] = true
	distance := make(map[Point]int)
	distance[end] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == start {
			return distance
		}

		for _, dir := range directions {
			next := Point{current.x + dir.x, current.y + dir.y}
			if next.x >= 0 && next.x < len(grid[0]) && next.y >= 0 && next.y < len(grid) && !visited[next] && grid[next.y][next.x] != WALL {
				queue = append(queue, next)
				visited[next] = true
				distance[next] = distance[current] + 1
			}
		}
	}
	panic("couldn't find end")
}

func calculateCheats(grid [][]rune, start Point, end Point, costs map[Point]int) []int {
	cheats := []int{}
	// we can cheat up to 2 squares, so look for empty_spaces 2-3 squares away
	// we don't need to look 1 away because that will be an empty_space we could get to
	directions := []Point{{0, 2}, {2, 0}, {0, -2}, {-2, 0}, {0, 3}, {3, 0}, {0, -3}, {-3, 0}}

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] != EMPTY_SPACE {
				continue
			}
			for _, dir := range directions {
				cheatStart := Point{x, y}
				cheatEnd := Point{x + dir.x, y + dir.y}
				if isValidCheat(grid, cheatStart, cheatEnd) {
					savings := costs[cheatStart] - costs[cheatEnd] - 2
					if savings > 0 {
						if savings > 20 {
							fmt.Println("found cheat", cheatStart, cheatEnd, savings)
						}
						cheats = append(cheats, savings)
					}
				}
			}
		}
	}
	return cheats
}

func isValidCheat(grid [][]rune, start, end Point) bool {
	if end.x < 0 || end.x >= len(grid[0]) || end.y < 0 || end.y >= len(grid) || grid[end.y][end.x] == WALL {
		return false
	}

	if start.x != end.x {
		for i := min(start.x, end.x) + 1; i < max(start.x, end.x); i++ {
			if grid[start.y][i] != WALL {
				return false
			}
		}
		return true
	} else {
		for i := min(start.y, end.y) + 1; i < max(start.y, end.y); i++ {
			if grid[i][start.x] != WALL {
				return false
			}
		}
		return true
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parse(file io.Reader) ([][]rune, Point, Point) {
	grid := [][]rune{}
	scanner := bufio.NewScanner(file)
	var start, end Point

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		grid = append(grid, []rune(line))
	}

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == START {
				start = Point{x, y}
			}
			if grid[y][x] == END {
				end = Point{x, y}
			}
		}
	}

	return grid, start, end
}
