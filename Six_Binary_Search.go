package main

import "fmt"

func getIndexOfKey(data []int, key int) int {
	start := 0
	end := len(data) - 1
	mid := (start + end) / 2
	for start <= end && data[mid] != key {
		if key > data[mid] {
			start = mid + 1
		} else if key < data[mid] {
			end = mid - 1
		}
		mid = (start + end) / 2
	}
	if key == data[mid] {
		return mid
	} else {
		return -1
	}
}

func main() {
	fmt.Println("Binary Search on Sorted Array")
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(getIndexOfKey(mySlice, 2))
	fmt.Println(getIndexOfKey(mySlice, 9))
	fmt.Println(getIndexOfKey(mySlice, -100))
	fmt.Println(getIndexOfKey(mySlice, 200))
}
