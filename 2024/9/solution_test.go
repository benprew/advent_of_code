package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	testValue := "2333133121414131402"
	expectedResult := 1928

	result := solve(testValue)
	if result != expectedResult {
		t.Errorf("solve2(%s) = %d; want %d", testValue, result, expectedResult)
	}
}

func TestSolve2(t *testing.T) {
	testValue := "2333133121414131402"
	expectedResult := 2858

	result := solve2(testValue)
	if result != expectedResult {
		t.Errorf("solve2(%s) = %d; want %d", testValue, result, expectedResult)
	}
}
