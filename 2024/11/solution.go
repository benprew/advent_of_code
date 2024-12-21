package main

import (
	"fmt"
	"io"
	"math"
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
	stones := parse(file)

	fmt.Println("Part 1:", solve(stones, 25))
	fmt.Println("Part 1:", solve(stones, 75))
}

func solve(stones []int, blinks int) (sum int) {
	for _, stone := range stones {
		sum += count(stone, blinks)
	}
	return sum
}

var cache = map[string]int{}

// If the stone is engraved with the number 0, it is replaced by a stone engraved
// with the number 1.

// If the stone is engraved with a number that has an even number of digits, it is
// replaced by two stones. The left half of the digits are engraved on the new left
// stone, and the right half of the digits are engraved on the new right stone. (The
// new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)

// If none of the other rules apply, the stone is replaced by a new stone; the old
// stone's number multiplied by 2024 is engraved on the new stone.
func count(stone, steps int) int {
	if steps == 0 {
		return 1
	}
	if v, ok := cache[fmt.Sprintf("%d|%d", stone, steps)]; ok {
		return v
	}
	if stone == 0 {
		sum := count(1, steps-1)
		cache[fmt.Sprintf("%d|%d", stone, steps)] = sum
		return sum
	}
	if numDigits(stone)%2 == 0 {
		mid := int(math.Pow(10, float64(numDigits(stone)/2)))
		old := stone / mid
		new := stone % mid
		sum := count(old, steps-1) + count(new, steps-1)
		cache[fmt.Sprintf("%d|%d", stone, steps)] = sum
		return sum
	}
	sum := count(stone*2024, steps-1)
	cache[fmt.Sprintf("%d|%d", stone, steps)] = sum
	return sum
}

func parse(r io.Reader) []int {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		panic(err)
	}
	parts := strings.Fields(buf.String())
	result := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		result[i] = num
	}
	return result
}

func numDigits(n int) int {
	return len(fmt.Sprintf("%d", n))
}
