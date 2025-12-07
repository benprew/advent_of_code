package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var cache = make(map[string]int)

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	fmt.Println(filename)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	turns := parse(file)
	fmt.Println("Part 1:", solve(turns))
	fmt.Println("Part 2:", solve2(turns))
}

func solve(turns []string) (numZeros int) {
	// count the number of times a turn ends on 0
	curr := 50
	for _, turn := range turns {
		amt, err := strconv.Atoi(turn[1:])
		if err != nil {
			panic(err)
		}
		amt = amt % 100
		if turn[0] == 'L' {
			amt = amt * -1
		}
		curr = (curr + amt) % 100
		if curr < 0 {
			curr = curr + 100
		}
		fmt.Println(turn, " ", curr)
		if curr == 0 {
			numZeros++
		}
	}
	return
}

func solve2(turns []string) (numZeros int) {
	// count the number of times a turn ends on 0
	curr := 50
	pCurr := curr
	for _, turn := range turns {
		amt, err := strconv.Atoi(turn[1:])
		if err != nil {
			panic(err)
		}
		numZeros += amt / 100
		amt = amt % 100
		if turn[0] == 'L' {
			amt = amt * -1
		}
		curr = (curr + amt)
		if curr > 100 {
			numZeros++
		}
		curr = curr % 100
		if curr < 0 {
			if pCurr != 0 {
				numZeros++
			}
			curr = curr + 100
		}
		if curr == 0 {
			numZeros++
		}
		fmt.Println(turn, " ", curr, " ", numZeros)
		pCurr = curr
	}
	return
}

func parse(file io.Reader) (lines []string) {

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}
