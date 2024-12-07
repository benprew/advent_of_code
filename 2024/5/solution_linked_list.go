package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

// Part 1
// start 9pm
// finish 9:18pm
// Part 2
// start 10:00pm
// stop 11:42pm
// Part 2 9453 -- too high
// Part 2 9206 -- too high still
// no more invalid lists
// start 5 minutes

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	fmt.Println("Part 1", solve(filename, false))
	fmt.Println("Part 2", solve(filename, true))
}

func solve(filename string, reorder bool) (sum int) {
	rules, updates := parseInput(filename)

	for i, update := range updates {
		fmt.Println("Original List")
		listPrint(update)
		valid := true
		curr := update
		origLen := listLen(update)
		numReorders := 0
		for curr != nil {
			pNum := curr.Val
			found := findIntersect(rules[pNum], curr)
			if found == nil {
				curr = curr.Next
				continue
			}
			if reorder {
				numReorders++
				next := curr.Next
				moveAfter(curr, found)
				curr = next
				continue
			} else {
				valid = false
			}

			curr = curr.Next
		}
		if valid {
			update = listRewind(update)
			listValidate(update, rules, i)
			listPrint(update)
			fmt.Println("Made", numReorders, "reorders")
			lLen := listLen(update)
			if origLen != lLen {
				panic("list size changed")
			}
			num := listForward(update, lLen/2).Val
			fmt.Printf("Adding %d (len %d)\n", num, lLen)
			sum += num
		} else {
			fmt.Println("Intersections for ", update.Val)
		}
	}

	return
}

func listValidate(head *ListNode, rules map[int][]int, listNum int) {
	curr := head
	for curr != nil {
		pNum := curr.Val
		found := findIntersect(rules[pNum], curr)
		if found != nil {
			listPrint(head)
			fmt.Println(rules[pNum])
			fmt.Println("Found", pNum, "in list", listNum)
			panic("List not valid (listNum: " + strconv.Itoa(listNum) + ")")
		}
		curr = curr.Next
	}
}

func listRewind(head *ListNode) *ListNode {
	for head.Prev != nil {
		head = head.Prev
	}
	return head
}

func listPrint(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}

func listLen(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

func listForward(head *ListNode, n int) *ListNode {
	for i := 0; i < n; i++ {
		head = head.Next
	}
	return head
}

func findIntersect(list1 []int, head *ListNode) *ListNode {
	// fmt.Printf("Finding %v in %d\n", list1, head.Val)
	for _, num := range list1 {
		if found := findNum(head, num); found != nil {
			// fmt.Println("Found", num, "in", head.Val)
			return found
		}
	}

	return nil
}

func findNum(head *ListNode, num int) *ListNode {
	curr := head
	for curr != nil {
		if curr.Val == num {
			return curr
		}
		curr = curr.Next
	}
	return nil
}

func moveAfter(a, b *ListNode) {
	// fmt.Println("Move", a.Val, "after", b.Val)
	// Remove node A from its current position
	if a.Prev != nil {
		a.Prev.Next = a.Next
	}
	if a.Next != nil {
		a.Next.Prev = a.Prev
	}

	// Insert node A after node B
	a.Next = b.Next
	a.Prev = b
	if b.Next != nil {
		b.Next.Prev = a
	}
	b.Next = a
}

// Prompt:
// how do I scan a file looking for 2 numbers separated by |? Or N numbers separated by commas
// now take the numbers separated by | and add them to a map, where the 2nd number is the key and the 1st number is appended to the value (int array).
// For the numbers separate by commas, convert them to a list of integers and append that to the updates list.
func parseInput(filename string) (rules map[int][]int, updates []*ListNode) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer file.Close()

	rules = make(map[int][]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			if len(parts) == 2 {
				num1, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
				num2, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err1 == nil && err2 == nil {
					rules[num2] = append(rules[num2], num1)
				}
			}
		} else if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			var head, current *ListNode
			for _, part := range parts {
				num, err := strconv.Atoi(strings.TrimSpace(part))
				if err != nil {
					panic(err)
				}

				node := &ListNode{Val: num}
				if head == nil {
					head = node
					current = node
				} else {
					node.Prev = current
					current.Next = node
					current = node
				}
			}
			if head != nil {
				updates = append(updates, head)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return rules, updates
}
