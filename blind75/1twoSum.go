package blind75

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, ok := numMap[complement]; ok {
			return []int{index, i}
		}
		numMap[num] = i
	}

	return nil
}

func TwoSum() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	if result != nil {
		fmt.Printf("Two numbers at indices %v and %v add up to %v\n", result[0], result[1], target)
	} else {
		fmt.Println("No such pair found.")
	}
}
