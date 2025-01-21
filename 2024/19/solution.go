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

	towels, patterns := parse(file)
	fmt.Println("Part 1:", solve(towels, patterns))
	fmt.Println("Part 2:", solve2(towels, patterns))
}

func solve(towels []string, patterns []string) (validDesigns int) {
	return
}

func solve2(towels []string, patterns []string) (validDesigns int) {
	return
}

func parse(file io.Reader) (towels []string, patterns []string) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		towels = append(towels, line)
	}
	for scanner.Scan() {
		line := scanner.Text()
		patterns = append(patterns, line)
	}
	return
}
