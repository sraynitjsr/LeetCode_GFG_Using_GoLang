package main

import (
	"fmt"
	"math"
)

func startEight() {
	fmt.Println("\nInside Module 8")

	myMap := make(map[int]int)

	mySlice := []int{1, 8, 30, 40, 100}

	diff := -60

	for _, val := range mySlice {
		myMap[val] = math.MinInt
	}

	for _, val := range mySlice {
		_, present := myMap[val+diff]
		if present {
			fmt.Println(val, val+diff)
		}
	}
}
