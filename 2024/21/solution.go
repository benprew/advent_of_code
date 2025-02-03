package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Point struct {
	x, y int
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

	door_codes := parse(file)

	fmt.Println("Part 1:", solve(door_codes))
	// fmt.Println("Part 2:", solve2(door_codes))

}

func solve(door_codes []string) (sum int) {
	door_grid := [][]rune{
		{'7', '8', '9'},
		{'4', '5', '6'},
		{'1', '2', '3'},
		{'Z', '0', 'A'},
	}
	door_start := Point{2, 3}
	door_cache := make(map[CacheKey][]QItem)
	robot_grid := [][]rune{
		{'Z', '^', 'A'},
		{'<', 'v', '>'},
	}
	robot_start := Point{2, 0}
	robot_cache := make(map[CacheKey][]QItem)
	for _, code := range door_codes {
		moves := getMoves([]string{code}, door_grid, door_start, door_cache)
		// fmt.Printf("moves_door: %s\n", moves)
		moves = getMoves(moves, robot_grid, robot_start, robot_cache)
		// fmt.Printf("moves_bot1: %s\n", moves)
		moves = getMoves(moves, robot_grid, robot_start, robot_cache)
		// fmt.Printf("moves: %s\n", moves)

		// fmt.Printf("%s: %s\n", code, moves)

		num, err := strconv.Atoi(code[:3])
		if err != nil {
			panic(err)
		}

		fmt.Printf("num: %d, moves len: %d\n", num, len(moves[0]))

		sum += len(moves[0]) * num

	}
	return
}

func getMoves(codes []string, grid [][]rune, start Point, cache map[CacheKey][]QItem) (moves []string) {
	curr := start
	moves_final := []string{}
	for _, code := range codes {
		moves = []string{""}
		// fmt.Println(code)
		for _, c := range code {
			// fmt.Printf("CODE => %c\n", c)
			ends := bfs(grid, curr, c, cache)
			// these are all the same ending position
			curr = ends[0].pos
			moves2 := []string{}
			for i := range moves {
				for item := range ends {
					moves2 = append(moves2, moves[i]+ends[item].path)
				}
			}
			moves = moves2
			// fmt.Println("moves: ", moves)
			// fmt.Println(i)
		}
		if len(moves_final) == 0 || len(moves[0]) < len(moves_final[0]) {
			moves_final = moves
		} else if len(moves[0]) == len(moves_final[0]) {
			moves_final = append(moves_final, moves...)
		}
	}

	min := -1
	var shortest_moves []string
	for _, m := range moves_final {
		if min == -1 || len(m) == min {
			shortest_moves = append(shortest_moves, m)
		}
		if len(m) < min {
			shortest_moves = []string{m}
		}
	}

	// fmt.Println("shortest_moves: ", shortest_moves)

	return shortest_moves
}

type QItem struct {
	path string
	pos  Point
}

type CacheKey struct {
	start Point
	end   rune
}

func bfs(grid [][]rune, start Point, end rune, cache map[CacheKey][]QItem) (ends []QItem) {
	directions := []Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	queue := []QItem{{"", start}}
	visited := make(map[Point]int)
	visited[start] = 0

	shortest := -1

	// fmt.Printf("starting: %v end: %c\n", start, end)

	if val, ok := cache[CacheKey{start, end}]; ok {
		return val
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current.pos] > 0 && visited[current.pos] < len(current.path) {
			continue
		}

		visited[current.pos] = len(current.path)

		if gridChr(grid, current.pos) == end {
			current.path += "A"
			// fmt.Println("found end: ", current.path)
			if shortest == -1 || len(current.path) == shortest {
				ends = append(ends, current)
				cache[CacheKey{start, end}] = ends
				shortest = len(current.path)
			} else if len(current.path) < shortest {
				ends = []QItem{current}
				cache[CacheKey{start, end}] = ends
				shortest = len(current.path)
			}
		}

		for _, dir := range directions {
			next := Point{current.pos.x + dir.x, current.pos.y + dir.y}
			if next.x >= 0 && next.x < len(grid[0]) && next.y >= 0 && next.y < len(grid) && grid[next.y][next.x] != 'Z' {
				queue = append(queue, QItem{current.path + dirString(dir), next})
			}
		}
	}
	// fmt.Println("ends: ", ends)
	return ends
}

func gridChr(grid [][]rune, pos Point) rune {
	return grid[pos.y][pos.x]
}

func dirString(dir Point) string {
	switch dir {
	case Point{0, 1}:
		return "v"
	case Point{1, 0}:
		return ">"
	case Point{0, -1}:
		return "^"
	case Point{-1, 0}:
		return "<"
	}
	panic("invalid direction")
}

func parse(file io.Reader) (door_codes []string) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		door_codes = append(door_codes, line)
	}
	return
}
