package main

import "fmt"

func checkSubArray(mySmallSlice []int, myLargeSlice []int, myMap map[int]int) {
	for index, val := range mySmallSlice {
		myMap[val] = index
	}
	for i := 0; i < len(myLargeSlice); i++ {
		value, present := myMap[myLargeSlice[i]]
		if present && value != -1 {
			myMap[myLargeSlice[i]] = -1
		}
	}
	for _, v := range myMap {
		if v != -1 {
			fmt.Println("Not A Sub Array")
			return
		}
	}
	fmt.Println("Sub Array")
}

func startNine() {
	fmt.Println("\nInside Module 9")
	myFirstSlice := []int{11, 1, 13, 21, 3, 7}
	mySecondSlice := []int{11, 3, 7, 7, 1}
	tempMap := make(map[int]int)
	if len(myFirstSlice) > len(mySecondSlice) {
		checkSubArray(mySecondSlice, myFirstSlice, tempMap)
	} else {
		checkSubArray(myFirstSlice, mySecondSlice, tempMap)
	}
}
