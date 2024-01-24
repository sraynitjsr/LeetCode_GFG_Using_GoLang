package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	prefixProduct := 1
	for i := 0; i < n; i++ {
		result[i] = prefixProduct
		prefixProduct *= nums[i]
	}
	suffixProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffixProduct
		suffixProduct *= nums[i]
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf(nums)
	fmt.Println("Input:", nums)
	fmt.Println("Output:", result)
}
