package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Token struct {
	Type  string
	Value string
	X     int
	Y     int
}

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	// Part 1
	sum := solve(filename, false)
	fmt.Println("Part 1 Sum:", sum)
	// part 2
	sum = solve(filename, true)
	fmt.Println("Part 2 Sum:", sum)
}

func solve(filename string, trackDo bool) (sum int) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return sum
	}

	// tokenize input
	tokens := tokenize(string(fileContent))

	// run state machine over tokens
	doing := true
	for _, token := range tokens {
		if token.Type == "DO" {
			doing = true
		} else if token.Type == "DONT" && trackDo {
			doing = false
		} else if token.Type == "MULT" {
			if doing {
				sum += token.X * token.Y
			}
		}
	}
	return sum
}

func tokenize(input string) []Token {
	var tokens []Token

	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatchIndex(input, -1)

	for _, match := range matches {
		if match[2] == -1 && match[4] == -1 { // do() or don't()
			tokenValue := input[match[0]:match[1]]
			if tokenValue == "do()" {
				tokens = append(tokens, Token{Type: "DO", Value: "do()"})
			} else if tokenValue == "don't()" {
				tokens = append(tokens, Token{Type: "DONT", Value: "don't()"})
			}
		} else { // mult(x,y)
			x, err1 := strconv.Atoi(input[match[2]:match[3]])
			y, err2 := strconv.Atoi(input[match[4]:match[5]])
			if err1 == nil && err2 == nil {
				tokens = append(tokens, Token{Type: "MULT", Value: input[match[0]:match[1]], X: x, Y: y})
			}
		}
	}

	return tokens
}
