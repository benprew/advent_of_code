package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// You can push each button as many times as you like. However, to save on time,
// you will need to determine the fewest total presses required to correctly
// configure all indicator lights for all machines in your list.

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

	input := parse(file)

	fmt.Println("Part 1:", solve(input))
	fmt.Println("Part 2:", solve2(input))
}

// To help the Elves figure out which path is causing the issue, they need you to
// find every path from you to out.
func solve(routes map[string][]string) int {
	memo := map[string]int{}
	return dfs("you", "out", true, true, routes, memo)
}

// find every path from svr (the server rack) to out that visit dac and fft
func solve2(routes map[string][]string) int {
	memo := map[string]int{}
	return dfs("svr", "out", false, false, routes, memo)
}

// path is stored a space-separate string because slices can't be map keys
//
// recursive dfs with memoization
func dfs(path, goal string, hasDAC, hasFFT bool, routes map[string][]string, memo map[string]int) (numPaths int) {
	node := head(path)
	key := memoKey(node, hasDAC, hasFFT)

	cnt, ok := memo[key]
	if ok {
		return cnt
	}

	if strings.Contains(path, "dac") {
		hasDAC = true
	}
	if strings.Contains(path, "fft") {
		hasFFT = true
	}

	if node == goal && hasDAC && hasFFT {
		numPaths++
	} else {
		for _, n := range routes[node] {
			newPath := path + " " + n
			numPaths += dfs(newPath, goal, hasDAC, hasFFT, routes, memo)
		}
	}

	memo[key] = numPaths
	return numPaths
}

func memoKey(node string, hasDAC, hasFFT bool) string {
	var dac, fft string
	if hasDAC {
		dac = "dac"
	}
	if hasFFT {
		fft = "fft"
	}

	return fmt.Sprintf("%s-%s-%s", node, dac, fft)

}

// return the most recently visted node in the path
func head(path string) string {
	pathList := strings.Fields(path)
	return pathList[len(pathList)-1]
}

func parse(file io.Reader) (routes map[string][]string) {
	scanner := bufio.NewScanner(file)
	routes = make(map[string][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		pieces := strings.Split(line, ":")
		head := pieces[0]
		steps := strings.Fields(pieces[1])
		routes[head] = steps
	}
	return
}
