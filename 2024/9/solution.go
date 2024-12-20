package main

import (
	"fmt"
	"os"
	"strconv"
)

// part 1
// start 10pm
// finish 11:30pm

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	blocksStr, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// if you only have 2 values you can solve it with a binary string
	// in a for loop. IE. to generate all combinations of a length k list
	// you need to iterate 2^k times.
	// You can generate the binary string using Sprintf "%0*b"
	fmt.Println("Part 1:", solve(string(blocksStr)))
	fmt.Println("Part 2:", solve2(string(blocksStr)))
}

func solve(blocksStr string) int {
	values := parse(blocksStr)

	start := 0
	end := len(values) - 1

	for start < end {
		if values[start] == -1 {
			for values[end] == -1 {
				end--
			}
			if end < start {
				break
			}
			values[start], values[end] = values[end], values[start]
		}
		start++
	}

	if !validate(values) {
		fmt.Println("start", start, "end", end)
		panic("invalid")
	}

	checkSum := 0
	for i := range values {
		if values[i] == -1 {
			break
		}
		checkSum += values[i] * i
	}
	return checkSum
}

func validate(values []int) (valid bool) {
	valid = true
	end := false
	for i := range values {
		if end && values[i] != -1 {
			fmt.Println("invalid", i, values[i])
			valid = false
		}
		if values[i] == -1 {
			end = true
		}
	}
	return valid
}

func parse(blocksStr string) (blocks []int) {
	for i, c := range string(blocksStr) {
		val, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}

		fileId := -1
		if i%2 == 0 {
			fileId = i / 2
		}

		for range val {
			blocks = append(blocks, fileId)
		}
	}

	return blocks
}

type Block struct {
	pos  int
	size int
}

// 2 pointers, one at the start, one at the end. When start pointer finds empty space, it consumes from end
// when start and end meet, you're done
// also this means no file/space can be larnger than 9
func solve2(str string) int {
	files, freeList := parse2(str)

	// starting from the largest file id, fill in the free blocks
	// if there are no free blocks, move to the next file
	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]
		// fmt.Println("file", file, "freeList", freeList)
		// fmt.Println("files", files)
		for i := range freeList {
			if freeList[i].pos > file.pos {
				break
			}

			if freeList[i].size >= file.size {
				file.pos = freeList[i].pos
				// fill in the free block
				freeList[i].size -= file.size
				freeList[i].pos += file.size
				break
			}
		}
		files[i] = file
		// fmt.Println("file", file, "freeList", freeList)
		// break
	}

	// calculate checksum
	checksum := 0
	for k, v := range files {
		fmt.Println(k, v)
		for i := 0; i < v.size; i++ {
			checksum += k * (v.pos + i)
		}
	}

	return checksum
}

// instead of returning a list of ints, we return a map of file_ids -> (size, pos)
// and a list free blocks (size, pos)
func parse2(blocksStr string) (files map[int]Block, freeList []Block) {
	files = make(map[int]Block)
	pos := 0
	for i, c := range string(blocksStr) {
		val, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}

		if i%2 == 0 {
			// add file
			fileId := i / 2
			files[fileId] = Block{size: val, pos: pos}

		} else {
			// add free block
			freeList = append(freeList, Block{size: val, pos: pos})
		}
		pos += val
	}
	fmt.Println(files)
	fmt.Println(freeList)
	return
}
