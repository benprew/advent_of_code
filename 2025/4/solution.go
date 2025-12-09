package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

	grid := parse(file)

	fmt.Println("Part 1:", solve(grid))
	total2 := 0
	for {
		t := solve2(grid)
		if t == 0 {
			break
		}
		total2 += t
	}
	fmt.Println("Part 2:", total2)
}

func solve(grid []string) (total int) {
	for y, line := range grid {
		for x, chr := range line {
			if chr != '@' {
				continue
			}
			Xrange := []int{-1, 0, 1}
			Yrange := []int{-1, 0, 1}
			bales := 0
			for _, xr := range Xrange {
				for _, yr := range Yrange {
					if y+yr >= 0 && x+xr >= 0 && y+yr < len(grid) && x+xr < len(line) && grid[y+yr][x+xr] == '@' {
						bales++
					}
				}
			}
			if bales <= 4 {
				fmt.Println(x, y, bales)
				total++
			}
		}
	}
	return total
}

func solve2(grid []string) (total int) {
	for y, line := range grid {
		for x, chr := range line {
			if chr != '@' {
				continue
			}
			Xrange := []int{-1, 0, 1}
			Yrange := []int{-1, 0, 1}
			bales := 0
			for _, xr := range Xrange {
				for _, yr := range Yrange {
					if y+yr >= 0 && x+xr >= 0 && y+yr < len(grid) && x+xr < len(line) && grid[y+yr][x+xr] == '@' {
						bales++
					}
				}
			}
			if bales <= 4 {
				s := grid[y]
				grid[y] = s[:x] + "x" + s[x+1:]
				total++
			}
		}
	}
	return
}

func parse(file io.Reader) (lines []string) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return
}
