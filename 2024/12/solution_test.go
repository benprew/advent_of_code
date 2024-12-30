package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	garden := `AAAA
BBCD
BBCC
EEEC
`
	gardenCostIs(t, garden, 140)
}

func TestPart1_2(t *testing.T) {
	garden := `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO
`
	gardenCostIs(t, garden, 772)
}

func TestPart1_3(t *testing.T) {
	garden := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
`
	gardenCostIs(t, garden, 1930)
}

func TestPart2(t *testing.T) {
	garden := `AAAA
BBCD
BBCC
EEEC
`
	gardenCostBulkIs(t, garden, 80)
}

func gardenCostBulkIs(t *testing.T, garden string, expected int) {
	grid := parse(strings.NewReader(garden))
	actual := solve2(grid)

	if expected != actual {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func gardenCostIs(t *testing.T, garden string, expected int) {
	grid := parse(strings.NewReader(garden))
	actual := solve(grid)
	if expected != actual {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
