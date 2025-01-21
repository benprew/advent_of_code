package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0
`
	points := parse(strings.NewReader(input))

	want := 22
	printGrid(points[:12], 7)
	if got := solve(points[:12], 7); got != want {
		t.Errorf("Expected %d, got %d", want, got)
	}
}
