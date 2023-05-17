package main

import (
	"fmt"
	"math"
)

func maximumSumSubArray() {
	myData := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	var max_so_far = math.MinInt
	var max_ending_here = 0
	var start = 0
	var end = 0
	var s = 0
	for i := 0; i < len(myData); i++ {
		max_ending_here += myData[i]

		if max_so_far < max_ending_here {
			max_so_far = max_ending_here
			start = s
			end = i
		}

		if max_ending_here < 0 {
			max_ending_here = 0
			s = i + 1
		}
	}
	fmt.Println("Maximum Sum Sub Array is =>", max_so_far)
	fmt.Println("Starting Index is =>", start)
	fmt.Println("Ending Index =>", end)
}
