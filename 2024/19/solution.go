package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var cache = make(map[string]bool)

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
	fmt.Println("towels:", towels)

	towelMap := make(map[string]bool)
	maxLen := 0

	for _, t := range towels {
		towelMap[t] = true
		if len(t) > maxLen {
			maxLen = len(t)
		}
	}

	for _, p := range patterns {
		fmt.Println("pattern:", p)
		if isValidDesign(towelMap, maxLen, p) {
			validDesigns++
		} else {
			fmt.Println("invalid:", p)
		}
	}
	return
}

func solve2(towels []string, patterns []string) (validDesigns int) {
	return
}

func isValidDesign(towelMap map[string]bool, maxTowelLen int, pattern string) bool {
	if pattern == "" {
		return true
	}
	ret, ok := cache[pattern]
	if ok {
		return ret
	}

	l := min(len(pattern), maxTowelLen)

	for i := range l + 1 {
		t := pattern[:i]
		if towelMap[t] && isValidDesign(towelMap, maxTowelLen, pattern[i:]) {
			cache[pattern] = true
			return true
		}
	}
	cache[pattern] = false
	return false
}

func parse(file io.Reader) (towels []string, patterns []string) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		towels = strings.Split(line, ", ")
	}
	for scanner.Scan() {
		line := scanner.Text()
		patterns = append(patterns, line)
	}
	return
}
