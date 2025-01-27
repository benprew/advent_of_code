package main

import (
	"os"
	"testing"
)

func TestSolve1(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid, start, end := parse(file)

	want := 1
	if got := solve(grid, start, end, 50); got != want {
		t.Errorf("Expected %d, got %d", want, got)
	}
}

func TestSolve2(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid, start, end := parse(file)

	want := 2
	if got := solve2(grid, start, end, 50); got != want {
		t.Errorf("Expected %d, got %d", want, got)
	}
}
