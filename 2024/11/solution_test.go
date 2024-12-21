package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	stones := parse(strings.NewReader("125 17"))
	expected := 55312
	actual := solve(stones, 25)
	if expected != actual {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

}

func TestPart1_6(t *testing.T) {
	stones := parse(strings.NewReader("125 17"))
	expected := 22
	actual := solve(stones, 6)
	if expected != actual {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

}
