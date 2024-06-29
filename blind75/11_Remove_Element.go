package main

import (
	"fmt"
	"strconv"
	"strings"
)

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func main() {
	var input string
	fmt.Print("Enter nums (comma-separated integers): ")
	fmt.Scanln(&input)
	numsStr := strings.Split(input, ",")
	nums := make([]int, len(numsStr))
	for i := 0; i < len(numsStr); i++ {
		num, _ := strconv.Atoi(strings.TrimSpace(numsStr[i]))
		nums[i] = num
	}

	var val int
	fmt.Print("Enter val to remove: ")
	fmt.Scanln(&val)

	fmt.Println("Original nums:", nums)
	k := removeElement(nums, val)
	fmt.Printf("After removing %d, nums becomes %v\n", val, nums[:k])
	fmt.Println("Returned k:", k)
}
