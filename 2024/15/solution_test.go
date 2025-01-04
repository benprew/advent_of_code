package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	gridStr := `##########
#.O.O.OOO#
#........#
#OO......#
#OO@.....#
#O#.....O#
#O.....OO#
#O.....OO#
#OO....OO#
##########
`
	grid, _ := parse(strings.NewReader(gridStr))
	got := scoreGrid(grid)
	want := 10092
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

}
