package main

import (
	"fmt"
	"math"
)

func printMaximumSumSubArray(mySlice []int) {
	maxSoFar := math.MinInt
	maxTillNow := 0
	start := 0
	end := 0
	currentStart := 0
	for i := 0; i < len(mySlice); i++ {
		maxTillNow = maxTillNow + mySlice[i]
		if maxSoFar < maxTillNow {
			maxSoFar = maxTillNow
			start = currentStart
			end = i
		}
		if maxTillNow < 0 {
			maxTillNow = 0
			currentStart = i + 1
		}
	}
	fmt.Println("Maximum Sum =>", maxSoFar)
	fmt.Print("Maximum Sum Sub Array is => ")
	for i := start; i <= end; i++ {
		fmt.Print(mySlice[i], " ")
	}
	fmt.Println("")
}

func startTen() {
	mySlice := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	printMaximumSumSubArray(mySlice)
}
