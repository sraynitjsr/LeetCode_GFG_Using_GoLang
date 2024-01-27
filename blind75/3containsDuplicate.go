package blind75

import "fmt"

func findDuplicates(arr []int) []int {
	duplicates := make([]int, 0)
	visited := make(map[int]bool)

	for _, num := range arr {
		if visited[num] {
			duplicates = append(duplicates, num)
		} else {
			visited[num] = true
		}
	}
	return duplicates
}

func ContainsDuplicate () {
	array := []int{1, 2, 3, 4, 5, 2, 7, 8, 5, 1}

	duplicates := findDuplicates(array)

	if len(duplicates) > 0 {
		fmt.Println("Duplicate elements:", duplicates)
	} else {
		fmt.Println("No duplicate elements found.")
	}
}
