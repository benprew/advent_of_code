package main

import (
	"bufio"
	"fmt"
	"image"
	"io"
	"os"
	"strings"

	"advent_of_code/utils"
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

	tiles := parse(file)

	fmt.Println("Part 1:", solve(tiles))
	fmt.Println("Part 2:", solve2(tiles))
}

// brute force solution:
// for each point make trianges with all previous points and store
// the max seen.
//
// 4 corners solution:
// start at each corner, walk inward until you find a point
// after you have all corners, see which 2 points makes the largest rectangle
func solve(tiles []image.Point) int {
	max := image.Rectangle{}
	reds := []image.Point{}
	fmt.Println(area(image.Rectangle{image.Point{7, 1}, image.Point{11, 1}}))
	for _, p := range tiles {
		for _, red := range reds {
			r := image.Rectangle{p, red}
			if cmp(r, max) == 1 {
				fmt.Println(r)
				max = r
			}
		}
		reds = append(reds, p)
	}
	fmt.Println("max:", max)
	return area(max)
}

func solve2(lines []image.Point) int {
	return 0
}

func cmp(r1, r2 image.Rectangle) int {
	r1a := area(r1)
	r2a := area(r2)
	if r1a > r2a {
		return 1
	} else if r1a < r2a {
		return -1
	}
	return 0
}

func area(r image.Rectangle) int {
	return (abs(r.Size().X) + 1) * (abs(r.Size().Y) + 1)
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func parse(file io.Reader) (tiles []image.Point) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		pieces := strings.Split(line, ",")
		tiles = append(tiles, image.Point{utils.Toi(pieces[0]), utils.Toi(pieces[1])})
	}
	return
}
