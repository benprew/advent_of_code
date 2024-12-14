package main

import (
	"bufio"
	"fmt"
	"os"
)

// Part 1
// start 9:35pm
// finish 10:10pm

// Part 2
// start 9:15pm
// finish 9:25pm

type Point struct {
	x, y int
}

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	// if you only have 2 values you can solve it with a binary string
	// in a for loop. IE. to generate all combinations of a length k list
	// you need to iterate 2^k times.
	// You can generate the binary string using Sprintf "%0*b"
	fmt.Println("Part 1:", solve(filename))
	fmt.Println("Part 2:", solve2(filename))
}

// return the number of antinodes
func solve(filename string) (numAntinodes int) {
	grid, antennas := parse(filename)
	// iterate through the antennas, calculate the distance between them and then
	// check if that distance is again is still inside the grid
	// if it is, we've found an antinode
	for k, v := range antennas {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				rise, run := pointsDist(v[i], v[j])

				anode := Point{v[i].x + run, v[i].y + rise}
				if anode.y >= 0 && anode.y < len(grid) && anode.x >= 0 && anode.x < len(grid[0]) {
					fmt.Println(string(k), "antinode at", v[i].y+rise, v[i].x+run)
					if grid[anode.y][anode.x] != '#' {
						numAntinodes++
					}
					grid[anode.y][anode.x] = '#'

				}
				anode = Point{v[j].x - run, v[j].y - rise}
				if anode.y >= 0 && anode.y < len(grid) && anode.x >= 0 && anode.x < len(grid[0]) {
					fmt.Println(string(k), "antinode at", v[i].y+rise, v[i].x+run)
					if grid[anode.y][anode.x] != '#' {
						numAntinodes++
					}
					grid[anode.y][anode.x] = '#'
				}
			}
		}
	}
	fmt.Println(antennas)
	printGrid(grid)
	return
}

func solve2(filename string) (numAntinodes int) {
	grid, antennas := parse(filename)
	// iterate through the antennas, calculate the distance between them and then
	// check if that distance is again is still inside the grid
	// if it is, we've found an antinode
	for k, v := range antennas {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				rise, run := pointsDist(v[i], v[j])

				anode := Point{v[i].x + run, v[i].y + rise}
				for anode.y >= 0 && anode.y < len(grid) && anode.x >= 0 && anode.x < len(grid[0]) {
					fmt.Println(string(k), "antinode at", v[i].y+rise, v[i].x+run)
					if grid[anode.y][anode.x] != '#' {
						numAntinodes++
					}
					grid[anode.y][anode.x] = '#'
					anode.x += run
					anode.y += rise

				}
				anode = Point{v[j].x - run, v[j].y - rise}
				for anode.y >= 0 && anode.y < len(grid) && anode.x >= 0 && anode.x < len(grid[0]) {
					fmt.Println(string(k), "antinode at", v[i].y+rise, v[i].x+run)
					if grid[anode.y][anode.x] != '#' {
						numAntinodes++
					}
					grid[anode.y][anode.x] = '#'
					anode.x -= run
					anode.y -= rise
				}
			}
		}
	}
	// also include all non '#' and non '.' squares
	for _, row := range grid {
		for _, c := range row {
			if c != '#' && c != '.' {
				numAntinodes++
			}
		}
	}
	fmt.Println(antennas)
	printGrid(grid)
	return
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func pointsDist(a, b Point) (rise int, run int) {
	return a.y - b.y, a.x - b.x
}

// antennas is a map of rune to an array of points
// the rune is the antenna character and the array of points
// is the location of the antenna on the grid
func parse(filename string) (grid [][]rune, antennas map[rune][]Point) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading file", err)
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	antennas = make(map[rune][]Point)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
		for x, c := range line {
			if c != '.' {
				antennas[c] = append(antennas[c], Point{x, len(grid) - 1})
			}
		}
	}

	return grid, antennas
}
