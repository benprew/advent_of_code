package main

import (
	"bufio"
	"fmt"
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

	lines := parse(file)

	fmt.Println("Part 1:", solve(lines))
	fmt.Println("Part 2:", solve2(lines))
}

// beam indexes
// read each row, encounter a ^
//   - if is in beam indexes, split++ and add 2 new indexes and remove current index
//   - else do nothing
func solve(lines []string) (splits int) {
	beamIndexes := make([]bool, len(lines[0]))

	for _, l := range lines {
		for i, n := range l {
			if n == 'S' {
				beamIndexes[i] = true
			}
			if n == '^' && beamIndexes[i] {
				beamIndexes[i] = false
				if i > 0 {
					beamIndexes[i-1] = true
				}
				if i < len(l) {
					beamIndexes[i+1] = true
				}
				splits++
			}
		}
	}

	return
}

func solve2(lines []string) (total int) {
	start, graph := mkGraph(lines)
	memo := make(map[string]int, 0)
	return dfs(start, graph, memo)

}

// build a graph from the input so we can run dfs on it
func mkGraph(lines []string) (string, map[string][]string) {
	graph := make(map[string][]string, 0)
	beamIndexes := make([]int, len(lines[0]))
	start := ""

	for y, l := range lines {
		for x, n := range l {
			key := mkKey(x, y)
			if n == 'S' {
				beamIndexes[x] = 1
				graph[key] = []string{}
				start = key
				fmt.Println("start", x, beamIndexes, graph)
			}
			if n == '^' && beamIndexes[x] > 0 {
				beamIndexes[x] = 0
				graph[key] = []string{}

				// a beam can have up to 2 par
				par := parents(x, y, graph)
				if x > 0 {
					beamIndexes[x-1] = y
				}
				if x < len(l) {
					beamIndexes[x+1] = y
				}

				// should only occur for start
				if len(par) == 0 {
					fmt.Println(key)
				}

				for _, p := range par {
					graph[p] = append(graph[p], key)
				}
			}
		}
	}

	fmt.Println(start)
	fmt.Println(graph)

	return start, graph
}

// path is stored a space-separate string because slices can't be map keys
//
// recursive dfs with memoization
//
// I don't think memoization is necessary since we don't have overlapping paths, but it's the same dfs function I used in day 11
func dfs(path string, graph map[string][]string, memo map[string]int) (numPaths int) {
	node := head(path)
	key := node

	// fmt.Println(node, key, path, graph)

	cnt, ok := memo[key]
	if ok {
		return cnt
	}

	numChildren := len(graph[node])
	if numChildren < 2 {
		fmt.Println(node, numChildren)
	}

	numPaths += 2 - numChildren
	for _, n := range graph[node] {
		newPath := path + " " + n
		numPaths += dfs(newPath, graph, memo)
	}

	memo[key] = numPaths
	return numPaths
}

func parse(file io.Reader) (lines []string) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

func mkKey(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func parents(x, y int, graph map[string][]string) (parents []string) {
	for i := y - 1; i >= 0; i-- {

		// parent directly above doesn't cause additional splits
		_, ok := graph[mkKey(x, i)]
		if ok {
			return parents
		}

		// left parent
		_, ok = graph[mkKey(x-1, i)]
		if ok {
			parents = append(parents, mkKey(x-1, i))
		}

		// right parent
		_, ok = graph[mkKey(x+1, i)]
		if ok {
			parents = append(parents, mkKey(x+1, i))
		}
	}
	return parents
}

func head(path string) string {
	pathList := strings.Fields(path)
	return pathList[len(pathList)-1]
}
