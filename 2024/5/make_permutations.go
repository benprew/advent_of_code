package main

// func main() {
// 	list := []int{0, 1, 2, 3, 4}
// 	permutations := generatePermutations(list)
// 	for _, perm := range permutations {
// 		fmt.Printf("%d,%d,%d,%d,%d\n", perm[0], perm[1], perm[2], perm[3], perm[4])
// 	}
// }

func generatePermutations(list []int) [][]int {
	var result [][]int
	permute(list, 0, &result)
	return result
}

func permute(list []int, start int, result *[][]int) {
	if start == len(list)-1 {
		perm := make([]int, len(list))
		copy(perm, list)
		*result = append(*result, perm)
		return
	}

	for i := start; i < len(list); i++ {
		list[start], list[i] = list[i], list[start]
		permute(list, start+1, result)
		list[start], list[i] = list[i], list[start] // backtrack
	}
}
