package blind75

import "fmt"

func findMissingNumberXOR(nums []int) int {
	n := len(nums)
	missing := n

	for i := 0; i < n; i++ {
		missing ^= i ^ nums[i]
	}

	return missing
}

func findMissingNumberSum(nums []int) int {
	n := len(nums)
	expectedSum := (n * (n + 1)) / 2
	actualSum := 0

	for _, num := range nums {
		actualSum += num
	}

	return expectedSum - actualSum
}

func MissingNumber() {
	nums := []int{3, 0, 1, 4, 6, 2}

	missingNumberXOR := findMissingNumberXOR(nums)
	fmt.Println("Missing number using XOR approach:", missingNumberXOR)

	missingNumberSum := findMissingNumberSum(nums)
	fmt.Println("Missing number using sum formula approach:", missingNumberSum)
}
