package main

import (
	"fmt"
	"github.com/sraynitjsr/blind75"
)

func main() {
	fmt.Println("LeetCode and DSA Using GoLang")
	
	intSlice := []int{2, 5, 3, 6, 4}
	
	fmt.Println("Checking for missing number is =>", blind75.MissingNumber(intSlice))

	fmt.Println("Two Sum Pair =>", blind75.TwoSum(intSlice, 11))

	blind75.MaxSumSubArray()
}
