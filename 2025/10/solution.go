package main

import (
	"advent_of_code/utils"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// You can push each button as many times as you like. However, to save on time,
// you will need to determine the fewest total presses required to correctly
// configure all indicator lights for all machines in your list.

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

	lines := parse(file)

	fmt.Println("Part 1:", solve(lines))
	fmt.Println("Part 2:", solve2(lines))
}

type Machine struct {
	Lights  []bool
	Buttons []Button
	Joltage []int
}

type Button struct {
	Wiring []int // a list of light indexes connected to button
}

func solve(machines []Machine) int {
	var total int
	for i, m := range machines {
		fmt.Println("======>", i, m)
		total += bfs(&m)
	}
	return total
}

// you will need to determine the fewest total presses required to correctly
// configure all indicator lights for all machines in your list.
//
// bfs + filtering:
// Because multiple button presses can result in the same state, skip queuing a
// press if that state has been seen before.
type QueueItem struct {
	Presses    int
	ButtonIdx  int
	LightState int
}

func bfs(m *Machine) (ret int) {
	seen := make(map[int]bool)
	q := []QueueItem{}
	for i, b := range m.Buttons {
		seen[0] = true
		state := update(0, b)
		q = append(q, QueueItem{1, i, state})
		seen[state] = true
	}
	fmt.Println(q)
	fmt.Println(seen)

	curr := q[0]
	q = q[1:]

	target := toInt(m.Lights)

	for target != curr.LightState {
		for i, b := range m.Buttons {
			state := update(curr.LightState, b)
			if seen[state] {
				continue
			}
			q = append(q, QueueItem{curr.Presses + 1, i, state})
			seen[state] = true
		}
		curr = q[0]
		q = q[1:]
		if curr.Presses > 20 {
			return -1
		}
	}
	fmt.Printf("solution: %+v\n", curr)

	return curr.Presses
}

func solve2(machines []Machine) int {
	return -1
}

// bits are little endian
func toInt(bits []bool) (ret int) {
	for i, bit := range bits {
		if bit {
			ret |= (1 << i)
		}
	}
	return
}

// Update light state after button press
func update(state int, button Button) int {
	for _, i := range button.Wiring {
		state = toggleBit(state, i)
	}
	return state
}

// Toggle bit at position i
func toggleBit(n int, i int) int {
	return n ^ (1 << i)
}

func parse(file io.Reader) (machines []Machine) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		pieces := strings.Split(line, " ")
		j := pieces[len(pieces)-1]
		m := Machine{
			Lights:  mkLights(pieces[0]),
			Buttons: mkButtons(pieces[1 : len(pieces)-1]),
			Joltage: toi_list(j[1 : len(j)-1]),
		}
		machines = append(machines, m)
	}
	return
}

func mkLights(l string) (lights []bool) {
	for _, c := range l {
		if c == '.' {
			lights = append(lights, false)
		}
		if c == '#' {
			lights = append(lights, true)
		}

	}
	return lights
}

func mkButtons(pieces []string) (buttons []Button) {
	for _, p := range pieces {
		p = p[1 : len(p)-1] // strip ( and )
		buttons = append(buttons, Button{Wiring: toi_list(p)})
	}
	return
}

// convert a comma-separated string into a list of integers
func toi_list(s string) (ints []int) {
	for _, n := range strings.Split(s, ",") {
		ints = append(ints, utils.Toi(n))
	}
	return
}
