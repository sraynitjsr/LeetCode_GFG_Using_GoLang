package main

import (
	"fmt"
	"math"
)

func startTen() {
	mySlice := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	maxSoFar := math.MinInt
	maxTillNow := 0
	for i := 0; i < len(mySlice); i++ {
		maxTillNow = maxTillNow + mySlice[i]
		if maxTillNow > maxSoFar {
			maxSoFar = maxTillNow
		}
		if maxTillNow < 0 {
			maxTillNow = 0
		}
	}
	fmt.Println("Maximum Sum => ", maxSoFar)
}
