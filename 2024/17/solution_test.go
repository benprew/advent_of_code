package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	inputStr := `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
`
	stack, registers := parseInput(strings.NewReader(inputStr))
	fmt.Println(stack, registers)
	got := solve(stack, registers, 0)
	want := "4,6,3,5,6,3,5,2,1,0"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

}

func TestPart2(t *testing.T) {
	inputStr := `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0
`
	stack, registers := parseInput(strings.NewReader(inputStr))
	fmt.Println(stack, registers)
	got := solve2(stack, registers)
	want := "0,3,5,4,3,0"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

}
