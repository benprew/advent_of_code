package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	A int = iota
	B
	C
)

var output = []int{}

// The adv instruction (opcode 0) performs division. The numerator is the value in the A register. The denominator is found by raising 2 to the power of the instruction's combo operand. (So, an operand of 2 would divide A by 4 (2^2); an operand of 5 would divide A by 2^B.) The result of the division operation is truncated to an integer and then written to the A register.
func adv(stack []int, pc int, registers *[]int) int {
	arg := comboVal(stack[pc+1], *registers)
	(*registers)[A] = (*registers)[A] / int(math.Pow(2, float64(arg)))
	return pc + 2
}

// The bxl instruction (opcode 1) calculates the bitwise XOR of register B and the
// instruction's literal operand, then stores the result in register B.
func bxl(stack []int, pc int, registers *[]int) int {
	arg := stack[pc+1]
	(*registers)[B] = (*registers)[B] ^ arg
	return pc + 2
}
func bst(stack []int, pc int, registers *[]int) int {
	arg := comboVal(stack[pc+1], *registers)
	(*registers)[B] = arg % 8
	return pc + 2
}
func jnz(stack []int, pc int, registers *[]int) int {
	if (*registers)[A] == 0 {
		return pc + 2
	}
	arg := stack[pc+1]
	return arg
}
func bxc(stack []int, pc int, registers *[]int) int {
	(*registers)[B] = (*registers)[B] ^ (*registers)[C]
	return pc + 2
}
func out(stack []int, pc int, registers *[]int) int {
	arg := comboVal(stack[pc+1], *registers)
	output = append(output, arg%8)
	if output[0] != 2 {
		return 1000
	}
	if len(output) > 1 && output[1] != 4 {
		return 1000

	}
	if len(output) > 2 && output[2] != 1 {
		return 1000
	}
	if len(output) > 3 && output[3] != 3 {
		return 1000
	}
	return pc + 2
}
func bdv(stack []int, pc int, registers *[]int) int {
	arg := comboVal(stack[pc+1], *registers)
	(*registers)[B] = (*registers)[A] / int(math.Pow(2, float64(arg)))
	return pc + 2
}
func cdv(stack []int, pc int, registers *[]int) int {
	arg := comboVal(stack[pc+1], *registers)
	(*registers)[C] = (*registers)[A] / int(math.Pow(2, float64(arg)))
	return pc + 2
}

func comboVal(arg int, registers []int) int {
	if arg > 3 {
		return registers[arg-4]
	}
	return arg
}

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

	stack, registers := parseInput(file)

	fmt.Println("Part 1:", solve(stack, registers, 0))
	fmt.Println("Part 2:", solve2(stack, registers))
}

func solve(stack []int, registers []int, count int) string {
	opcodes := map[int]func([]int, int, *[]int) int{
		0: adv,
		1: bxl,
		2: bst,
		3: jnz,
		4: bxc,
		5: out,
		6: bdv,
		7: cdv,
	}

	if count < 1 {
		count = 10000
	}
	output = []int{}
	pc := 0
	i := 0
	for pc < len(stack) && i < count {
		opcode := stack[pc]
		pc = opcodes[opcode](stack, pc, &registers)
		i++
	}
	return joinOutput(output)
}

func solve2(stack []int, registers []int) string {
	i := 4_669_757
	for i < 500_000_000_000 {
		i++
		registers[A] = i
		if len(output) < 5 {
			continue
		}
		solve(stack, registers, 64)
		if slices.Equal(output, stack) {
			fmt.Println(i)
			return joinOutput(output)
		}
	}

	return ""
}

func joinOutput(arr []int) string {
	strArr := make([]string, len(arr))
	for i, num := range arr {
		strArr[i] = strconv.Itoa(num)
	}
	return strings.Join(strArr, ",")
}

func parseInput(reader io.Reader) ([]int, []int) {
	scanner := bufio.NewScanner(reader)
	registers := make([]int, 3)
	program := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Register") {
			parts := strings.Split(line, ": ")
			value, _ := strconv.Atoi(parts[1])
			switch parts[0] {
			case "Register A":
				registers[0] = value
			case "Register B":
				registers[1] = value
			case "Register C":
				registers[2] = value
			}
		} else if strings.HasPrefix(line, "Program") {
			parts := strings.Split(line, ": ")
			programStr := strings.Split(parts[1], ",")
			for _, numStr := range programStr {
				num, _ := strconv.Atoi(numStr)
				program = append(program, num)
			}
		}
	}

	return program, registers
}
