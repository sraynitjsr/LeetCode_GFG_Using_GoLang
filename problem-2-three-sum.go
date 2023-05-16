package main

import "fmt"

func hashingSolution(slice []int, target int) {
	fmt.Print("Using Hashing, ")

	myMap := make(map[int]int)
	for i := 0; i < len(slice); i++ {
		myMap[slice[i]] = i
	}

	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			index, present := myMap[target-slice[i]-slice[j]]
			if present {
				fmt.Println("Found Triplet =>", slice[index], slice[i], slice[j])
				return
			}
		}
	}
	fmt.Println("No Triplets Found")
}

func sortingSolution(slice []int, target int) {
	fmt.Println("Using Sorting")
}

func threeSum() {
	fmt.Println("Three Sum Array")
	mySlice := []int{1, 4, 5, 2, 3}
	target := 8
	hashingSolution(mySlice, target)
	sortingSolution(mySlice, target)
}
