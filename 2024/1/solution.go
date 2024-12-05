package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	a, b := parse("input.txt")
	if len(a) != len(b) {
		fmt.Println("ERROR: Arrays are not of the same length")
		return
	}

	solveDistance(a, b)
	solveSimilarity(a, b)
}

func parse(filename string) (a []int, b []int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var x, y int
		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &x, &y)
		if err != nil {
			fmt.Println("Error parsing line:", err)
			continue
		}
		a = append(a, x)
		b = append(b, y)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return a, b
}

func solveDistance(a []int, b []int) {
	sort.Ints(a)
	sort.Ints(b)

	distance := 0
	for i := 0; i < len(a); i++ {
		distance += abs(a[i] - b[i])
	}

	fmt.Println("Distance:")
	fmt.Println(distance)
}

func solveSimilarity(a []int, b []int) {
	sim := 0
	bMap := make(map[int]int)

	for _, v := range b {
		bMap[v] += 1
	}

	for _, v := range a {
		sim += v * bMap[v]
	}

	fmt.Println("Similarity:")
	fmt.Println(sim)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
