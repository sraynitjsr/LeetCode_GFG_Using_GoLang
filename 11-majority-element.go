package main

import (
	"fmt"
	"sort"
)

func startEleven() {
	fmt.Println("Inside Module 11")
	mySlice := []int{1, 1, 2, 1, 3, 5, 1}
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i] < mySlice[j]
	})
	count := 0
	element := mySlice[0]
	halfOfSlice := len(mySlice) / 2
	for i := 0; i < len(mySlice); i++ {
		if count > halfOfSlice {
			fmt.Println("Majority Element is =>", element)
			return
		}
		if element == mySlice[i] {
			count++
		} else {
			element = mySlice[i]
			count = 0
		}
	}
	fmt.Println("No Majority Element Found")
}
