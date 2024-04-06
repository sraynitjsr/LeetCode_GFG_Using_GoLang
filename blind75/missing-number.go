package blind75

import "fmt"

func findMissingNumber(nums []int) int {
	n := len(nums)
	missing := n
	for i := 0; i < n; i++ {
		missing ^= i ^ nums[i]
	}
	return missing
}

func MissingNumber() {
	nums := []int{3, 0, 1, 4, 6, 2}
	missingNumber := findMissingNumber(nums)
	fmt.Println("The missing number is:", missingNumber)
}
