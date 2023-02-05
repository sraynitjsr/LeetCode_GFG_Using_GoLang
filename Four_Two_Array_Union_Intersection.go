package main

import (
	"fmt"
)

type MyHashSet struct {
	mySlice []int
}

func (hs *MyHashSet) put(val int) {
	for _, v := range hs.mySlice {
		if v == val {
			return
		}
	}
	hs.mySlice = append(hs.mySlice, val)
}

func (hs *MyHashSet) find(val int) bool {
	for _, v := range hs.mySlice {
		if v == val {
			return true
		}
	}
	return false
}

func printUnion(hs *MyHashSet, slice1, slice2 []int) {
	for _, v := range slice1 {
		hs.put(v)
	}
	for _, v := range slice2 {
		hs.put(v)
	}

	fmt.Print("Union => ")

	for _, v := range hs.mySlice {
		fmt.Print(v, " ")
	}

	fmt.Println("")
}

func printIntersection(hs *MyHashSet, slice1, slice2 []int) {
	for _, v := range slice1 {
		hs.put(v)
	}

	fmt.Print("Intersection => ")
	for _, v := range slice2 {
		if hs.find(v) {
			fmt.Print(v, " ")
		}
	}
	fmt.Println("")
}

func main() {
	slice1 := []int{1, 3, 20, -39, 45, -34, 67, 100}
	slice2 := []int{2020, -39, 56, 45, 78, 2, 22, 2}
	printUnion(&MyHashSet{}, slice1, slice2)
	printIntersection(&MyHashSet{}, slice1, slice2)
}

