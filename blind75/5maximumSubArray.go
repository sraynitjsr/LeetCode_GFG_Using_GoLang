package blind75

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxSumSubArray() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := maxSubArray(nums)
	fmt.Println("Maximum Sum Subarray:", result)
}
