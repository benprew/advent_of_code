package utils

import "strconv"

// Collection of various Advent of Code utility functions

// Atoi but panic on error
func Toi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
