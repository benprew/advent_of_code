package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	grid := `...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9`
	expectedResult := 2
	gridTest(t, grid, expectedResult)
}

func TestSolveFull(t *testing.T) {
	grid := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`
	expectedResult := 36
	gridTest(t, grid, expectedResult)
}

func TestSolve3(t *testing.T) {
	grid := `
..90..9
...1.98
...2..7
6543456
765.987
876....
987....`
	expectedResult := 4
	gridTest(t, grid, expectedResult)
}

func gridTest(t *testing.T, grid string, expected int) {
	testValue := parseGrid(grid)
	result := solve(testValue)
	if result != expected {
		t.Errorf("got %d; want %d", result, expected)
	}
}

func gridTest2(t *testing.T, grid string, expected int) {
	testValue := parseGrid(grid)
	result := solve2(testValue)
	if result != expected {
		t.Errorf("got %d; want %d", result, expected)
	}
}

func TestSolveFull2(t *testing.T) {
	grid := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`
	expectedResult := 81
	gridTest2(t, grid, expectedResult)
}
