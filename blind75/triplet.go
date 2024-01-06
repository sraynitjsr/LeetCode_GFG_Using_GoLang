package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

func Triplet() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	targetSum := 0
	result := threeSum(nums, targetSum)

	fmt.Printf("Triplets that sum to %d:\n", targetSum)
	for _, triplet := range result {
		fmt.Println(triplet)
	}
}
