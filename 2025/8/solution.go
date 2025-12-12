package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
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

	fmt.Println("Part 1:", solve(lines, 10))
	fmt.Println("Part 2:", solve2(lines))
}

// 1k lines
// binomial join is N^2 - seems reasonable
// for i in lst:
//   for j in lst[i:]:
//     pair = lst[i], lst[j]

// find distance of all pairs to each other
// sort
// make connections as graph adjacency list (map node->neighbors)
// dfs traversal of adjacency list to find size
// - track previous nodes visited to prevent counting cycles

type edge [2]TPoint

type Connection struct {
	Edge     edge
	Distance float64
}

var seen map[TPoint]bool
var graph map[TPoint][]TPoint

func solve(lines []string, turns int) int {
	var points []TPoint
	for _, n := range lines {
		points = append(points, mkPoint(n))
	}

	// build all connections ordered because we need the top N shortest
	var conns []Connection
	for i, row1 := range points {
		for _, row2 := range points[i+1:] {
			e := edge{row1, row2}
			d := dist(e)
			conns = append(conns, Connection{Edge: e, Distance: d})
		}
	}
	sort.Slice(conns, func(i, j int) bool { return conns[i].Distance < conns[j].Distance })
	fmt.Println("conns len", len(conns))

	// construct graph
	graph = make(map[TPoint][]TPoint)
	for _, c := range conns[:turns] {
		graph[c.Edge[0]] = append(graph[c.Edge[0]], c.Edge[1])
		graph[c.Edge[1]] = append(graph[c.Edge[1]], c.Edge[0])
	}
	for k, v := range graph {
		fmt.Println(k, v)
	}

	// walk graph to get 3 largest circuits
	seen = make(map[TPoint]bool)
	graphs := []int{1, 1, 1} // list of graph sizes min 3 circuits
	for k, v := range graph {
		if seen[k] {
			continue
		}
		size := traverse(k, v)
		fmt.Println(k, size)
		graphs = append(graphs, size)
	}

	// multiply together the sizes of the three largest circuits
	// graphs is initalized as 3
	sort.Slice(graphs, func(i, j int) bool { return graphs[i] > graphs[j] })
	fmt.Println(graphs)
	return graphs[0] * graphs[1] * graphs[2]
}

func solve2(lines []string) int {
	var points []TPoint
	for _, n := range lines {
		points = append(points, mkPoint(n))
	}

	// build all connections ordered because we need the top N shortest
	var conns []Connection
	for i, row1 := range points {
		for _, row2 := range points[i+1:] {
			e := edge{row1, row2}
			d := dist(e)
			conns = append(conns, Connection{Edge: e, Distance: d})
		}
	}
	sort.Slice(conns, func(i, j int) bool { return conns[i].Distance < conns[j].Distance })
	fmt.Println("conns len", len(conns))

	graphs := []int{0}
	target := len(lines)
	left, right := 0, len(conns)
	mid := left + (right-left)/2
	for left <= right {
		mid = left + (right-left)/2
		fmt.Println(left, right, mid)

		// construct graph
		graph = make(map[TPoint][]TPoint)
		for _, c := range conns[:mid] {
			graph[c.Edge[0]] = append(graph[c.Edge[0]], c.Edge[1])
			graph[c.Edge[1]] = append(graph[c.Edge[1]], c.Edge[0])
		}

		graphs = []int{1, 1, 1} // list of graph sizes min 3 circuits
		// walk graph to get 3 largest circuits
		seen = make(map[TPoint]bool)
		for k, v := range graph {
			if seen[k] {
				continue
			}
			size := traverse(k, v)
			graphs = append(graphs, size)
		}

		// binary search to find shortest number of connections to connect all junctions
		sort.Slice(graphs, func(i, j int) bool { return graphs[i] > graphs[j] })
		longest := graphs[0]
		if longest == target {
			right = mid - 1
		} else if longest < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return conns[mid].Edge[0].X * conns[mid].Edge[1].X
}

func traverse(node TPoint, neighbors []TPoint) int {
	if seen[node] {
		return 0
	}
	seen[node] = true
	size := 1
	for _, n := range neighbors {
		size += traverse(n, graph[n])
	}

	return size
}

type TPoint struct {
	X, Y, Z int
}

func dist(e edge) float64 {
	x := e[0].X - e[1].X
	y := e[0].Y - e[1].Y
	z := e[0].Z - e[1].Z
	return math.Sqrt(float64(x*x + y*y + z*z))
}

func mkPoint(s string) TPoint {
	vals := strings.Split(s, ",")
	return TPoint{toi(vals[0]), toi(vals[1]), toi(vals[2])}
}

func parse(file io.Reader) (lines []string) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

func toi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
