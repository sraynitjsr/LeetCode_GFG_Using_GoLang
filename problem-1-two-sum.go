package main

import "fmt"

func twoSum() {
	fmt.Println("Two Sum Array")
	mySlice := []int{2, 7, 11, 15}
	target := 17
	myMap := make(map[int]int)
	for i := 0; i < len(mySlice); i++ {
		myMap[mySlice[i]] = i
	}
	for i := 0; i < len(mySlice); i++ {
		index, present := myMap[target-mySlice[i]]
		if present {
			fmt.Println("Pair Found For", target)
			fmt.Println("First  Index Data =>", mySlice[i], "at Index", i)
			fmt.Println("Second Index Data =>", mySlice[index], "at Index", index)
			return
		}
	}
	fmt.Println("No Pair Found")
}
