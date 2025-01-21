package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	towels, patterns := parse(file)

	want := 6
	if got := solve(towels, patterns); got != want {
		t.Errorf("Expected %d, got %d", want, got)
	}
}
