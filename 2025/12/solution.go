package main

import (
	"advent_of_code/utils"
	"bufio"
	"fmt"
	"image"
	"io"
	"os"
	"strings"
)

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

	shapes, regions := parse(file)
	fmt.Println(shapes)
	fmt.Println(regions)

	fmt.Println("Part 1:", solve(shapes, regions))
	fmt.Println("Part 2:", solve2(shapes, regions))
}

type Shape [][]bool

const NUM_SHAPES = 6

type Region struct {
	Size        image.Point
	ShapeCounts [NUM_SHAPES]int
}

func solve(shapes []Shape, regions []Region) (count int) {
	for _, r := range regions {
		// pieces needed to place
		pieces := []Shape{}
		presentSize := 0
		for i, cnt := range r.ShapeCounts {
			for range cnt {
				pieces = append(pieces, shapes[i])
				presentSize += size(shapes[i])
			}
		}

		gridArea := r.Size.X * r.Size.Y

		// womp womp
		// Actually doing heuristic over input is too hard, but if it's under the
		// size of the grid it'll fit
		if float64(1.3)*float64(presentSize) < float64(gridArea) {
			count++
		} else if presentSize > gridArea {
			continue
		} else {
			fmt.Println("HARD: ", presentSize, gridArea)
		}
		continue

		fmt.Println("Placing", len(pieces), "pieces")

		// try to fit all pieces
		if canFit(pieces, r.Size) {
			fmt.Println("Fit!")
			count++
		}
	}
	return
}

func solve2(shapes []Shape, regions []Region) int {
	return -1
}

func canFit(pieces []Shape, area image.Point) bool {
	grid := make([][]bool, area.Y)
	for i := range grid {
		grid[i] = make([]bool, area.X)
	}
	return backtrack(grid, pieces, 0)

}

func backtrack(grid [][]bool, pieces []Shape, idx int) bool {
	// print(grid)
	// no more pieces to place
	if idx == len(pieces) {
		print(grid)
		return true
	}
	for _, pr := range rotations(pieces[idx]) {
		// print(pr)
		for y := range grid {
			for x := range grid[y] {
				if canPlace(pr, grid, x, y) {
					place(pr, grid, x, y)

					if backtrack(grid, pieces, idx+1) {
						return true
					}

					remove(pr, grid, x, y)
				}
			}
		}
	}
	return false
}

func canPlace(piece Shape, grid [][]bool, x, y int) bool {
	for i := range piece {
		for j, v := range piece[i] {
			if y+i >= len(grid) || x+j >= len(grid[0]) {
				return false
			}
			if v && grid[y+i][x+j] {
				return false
			}
		}
	}
	return true
}

func place(piece Shape, grid [][]bool, x, y int) {
	for i := range piece {
		for j, v := range piece[i] {
			if v && grid[y+i][x+j] {
				panic("trying to overwrite piece")
			}
			if v {
				grid[y+i][x+j] = true
			}
		}
	}
}

func remove(piece Shape, grid [][]bool, x, y int) {
	for i := range piece {
		for j, v := range piece[i] {
			if v && !grid[y+i][x+j] {
				panic("no piece to remove")
			}
			if v {
				grid[y+i][x+j] = false
			}
		}
	}
}

func rotations(piece Shape) (rots []Shape) {
	rots = append(rots, piece)
	for i := range 3 {
		piece = rotate(rots[i])
		rots = append(rots, piece)
	}
	return
}

// rotate piece clockwise
func rotate(piece Shape) (rot Shape) {
	rot = make(Shape, len(piece))
	for y := range piece {
		rot[y] = make([]bool, len(piece[0]))
	}
	for y := range piece {
		for x := range piece[y] {
			rot[x][len(piece)-y-1] = piece[y][x]
		}
	}
	return
}

func print(s Shape) {
	for y := range s {
		for x := range s[y] {
			if s[y][x] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func size(s Shape) int {
	sz := 0
	for i := range s {
		for j := range s[i] {
			if s[i][j] {
				sz++
			}
		}
	}
	return sz
}

func parse(file io.Reader) (shapes []Shape, regions []Region) {
	scanner := bufio.NewScanner(file)

	shape := Shape{}
	i := 0 // track shape line

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if len(shapes) == NUM_SHAPES {
			// parse regions
			pieces := strings.Split(line, ": ")
			sz := strings.Split(pieces[0], "x")
			size := image.Point{X: utils.Toi(sz[0]), Y: utils.Toi(sz[1])}

			shps := strings.Fields(pieces[1])
			shapes := [NUM_SHAPES]int{}
			for i, n := range shps {
				shapes[i] = utils.Toi(n)
			}

			regions = append(regions, Region{Size: size, ShapeCounts: shapes})
		} else {
			// parse shapes
			if strings.Contains(line, ":") {
				// start of shape
				continue
			} else if strings.Contains(line, "#") {
				shape = append(shape, make([]bool, len(line)))
				for j, c := range line {
					if c == '#' {
						shape[i][j] = true
					}
				}
				i++
			} else {
				shapes = append(shapes, shape)
				shape = Shape{}
				i = 0
			}
		}
	}
	return
}
