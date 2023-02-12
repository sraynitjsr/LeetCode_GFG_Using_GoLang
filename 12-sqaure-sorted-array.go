package main

import "fmt"

func sortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1
	index := 0
	output := []int{}
	for index < len(nums) {
		var leftAbs int = nums[left]
		var rightAbs int = nums[right]
		if nums[left] < 0 {
			leftAbs = nums[left] * -1
		}
		if nums[right] < 0 {
			rightAbs = nums[right] * -1
		}
		if leftAbs > rightAbs {
			output = append(output, nums[left]*nums[left])
			left++
		} else {
			output = append(output, nums[right]*nums[right])
			right--
		}
		index++
	}
	j := 0
	for i := len(nums) - 1; i >= 0; i-- {
		nums[j] = output[i]
		j++
	}
	return nums
}

func startTwelve() {
	fmt.Println("Inside Module 12")
	mySlice := []int{-6, -3, -1, 2, 4, 5}
	fmt.Println(sortedSquares(mySlice))
}
