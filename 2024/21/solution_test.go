package main

import (
	"fmt"
	"os"
	"testing"
)

func TestSolve1(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	door_code := parse(file)

	want := 126384
	if got := solve(door_code); got != want {
		t.Errorf("Expected %d, got %d", want, got)
	}
}

func TestSolveSingle(t *testing.T) {
	door_code := "029A"

	door_grid := [][]rune{
		{'7', '8', '9'},
		{'4', '5', '6'},
		{'1', '2', '3'},
		{'Z', '0', 'A'},
	}
	door_start := Point{2, 3}
	door_cache := make(map[CacheKey][]QItem)

	want := "<A^A^^>AvvvA"
	got := getMoves([]string{door_code}, door_grid, door_start, door_cache)
	if len(got[0]) != len(want) {
		t.Errorf("Expected %s, got %v", want, got)
	}
}

func TestSolveSingle2(t *testing.T) {
	robot_grid := [][]rune{
		{'Z', '^', 'A'},
		{'<', 'v', '>'},
	}
	robot_start := Point{2, 0}
	robot_cache := make(map[CacheKey][]QItem)
	moves := []string{
		"<A^A^^>AvvvA", "<A^A^>^AvvvA", "<A^A>^^AvvvA",
	}
	want := "v<<A>>^A<A>AvA<^AA>A<vAAA>^A"
	got := getMoves(moves, robot_grid, robot_start, robot_cache)
	if len(got[0]) != len(want) {
		t.Errorf("Expected %s (%d), got %s (%d)", want, len(want), got, len(got[0]))
	}
	fmt.Println(got)
}

func TestBFS(t *testing.T) {
	grid := [][]rune{
		{'7', '8', '9'},
		{'4', '5', '6'},
		{'1', '2', '3'},
		{'Z', '0', 'A'},
	}
	start := Point{2, 3}
	end := '7'
	expectedNumPaths := 9
	door_cache := make(map[CacheKey][]QItem)
	results := bfs(grid, start, end, door_cache)
	if len(results) != expectedNumPaths {
		t.Errorf("Expected path %v, got %v (%d)", expectedNumPaths, results, len(results))
	}
}
