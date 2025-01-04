package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	gridStr := `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############
`
	maze := parse(strings.NewReader(gridStr))
	got := solve(maze)
	want := 7036
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

}
