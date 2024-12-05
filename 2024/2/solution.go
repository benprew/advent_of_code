package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	reports := parse(filename)
	solveReports(reports, false)
	solveReports(reports, true)
}

func parse(filename string) (reports [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		var integers []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error parsing integer:", err)
				continue
			}
			integers = append(integers, num)
		}
		reports = append(reports, integers)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return reports
}

func solveReports(reports [][]int, dampenElement bool) {
	validReports := 0
	for _, report := range reports {
		if isValid(report) {
			validReports += 1
		} else if dampenElement {
			// fmt.Println("Original Report: ", report)
			// iterate over report removing a single element to see if we can make a valid report from it
			for i := range report {
				// Create a new slice with the element removed
				newReport := append([]int{}, report[:i]...)
				newReport = append(newReport, report[i+1:]...)

				// fmt.Println("report: ", newReport, " ", i)
				if isValid(newReport) {
					validReports += 1
					break
				}
			}
		}
	}

	fmt.Printf("Valid Reports: %d (dampened: %t)\n", validReports, dampenElement)
}

func isValid(report []int) bool {
	ascending := report[0]-report[1] < 0
	for i := 0; i < len(report)-1; i++ {
		diff := report[i] - report[i+1]
		if abs(diff) > 3 || ascending == (diff > 0) || diff == 0 {
			return false
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
