package main

import (
	"testing"
)

func TestSolvePart1(t *testing.T) {
	expected := 143
	result, _ := solve("test.txt")

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	expected := 123
	_, result := solve("test.txt")

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSolvePart2Full(t *testing.T) {
	expected := 4077 // something below 9206
	_, result := solve("input.txt")

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSolvePart2Perm(t *testing.T) {
	expected := 238
	_, result := solve("test_perm_input.txt")

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
