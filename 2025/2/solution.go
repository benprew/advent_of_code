package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

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

	ids := parse(file)

	fmt.Println("Part 1:", solve(ids))
	fmt.Println("Part 2:", solve2(ids))
}

func solve(ids []string) (total int) {
	for _, id := range ids {
		idRange := strings.Split(id, "-")
		start := toi(idRange[0])
		end := toi(idRange[1])
		for i := start; i <= end; i++ {
			n := strconv.Itoa(i)
			first := n[0 : len(n)/2]
			last := n[len(n)/2:]
			// fmt.Println(n, first, last, len(n)/2)
			if first == last {
				fmt.Println(i)
				total += i
			}
		}
	}
	return total
}

func solve2(ids []string) (total int) {
	for _, id := range ids {
		idRange := strings.Split(id, "-")
		start := toi(idRange[0])
		end := toi(idRange[1])
		fmt.Printf("%d-%d\n", start, end)
		for i := start; i <= end; i++ {
			strI := strconv.Itoa(i)
			if len(strI) > 1 && strI == strings.Repeat(string(strI[0]), len(strI)) {
				fmt.Println(i)
				total += i
				continue
			}
			// check maximal divisors to see if they repeat. If a maximal divisor doesn't match
			// the other divisors of it won't match either (ex. if repeat 4 doesn't match, repeat 2 won't match)
			for _, n := range factors(len(strI)) {
				rep := strings.Repeat(strI[0:n], len(strI)/n)

				if strI == rep {
					fmt.Println(i)
					total += i
					break
				}
			}
		}
	}
	return total
}

func factors(n int) (f []int) {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			f = append(f, n/i)
			f = append(f, i)
		}
	}

	return f
}

func toi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func parse(file io.Reader) (ids []string) {
	scanner := bufio.NewScanner(file)
	// should only be a single line
	for scanner.Scan() {
		line := scanner.Text()
		ids = strings.Split(line, ",")
	}
	return
}
