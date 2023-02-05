package main

import "fmt"

func missingNumber(nums []int) int {
	actualSum := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		actualSum = actualSum + nums[i]
	}
	expectedSum := (length * (length + 1)) / 2
	return expectedSum - actualSum
}

func main() {
	data := []int{0, 3, 1}
	fmt.Println(missingNumber(data))
}
