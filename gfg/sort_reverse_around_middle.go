package gfg

import (
    "fmt"
    "sort"
)

func reverseAroundMiddle(arr []int) {
    mid := len(arr) / 2
    for i := 0; i < mid/2; i++ {
        arr[i], arr[mid-i-1] = arr[mid-i-1], arr[i]
    }
    for i := mid + 1; i < (len(arr)+mid)/2; i++ {
        arr[i], arr[len(arr)-i+mid] = arr[len(arr)-i+mid], arr[i]
    }
}

func  SortReverseAroudMiddle() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Println("Original array:", arr)
    sort.Slice(arr, func(i, j int) bool {
        mid := len(arr) / 2
        if i < mid && j >= mid {
            return arr[i] < arr[len(arr)-j-1+mid]
        } else if i >= mid && j < mid {
            return arr[i] > arr[len(arr)-i-1+mid]
        }
        return false
    })
    fmt.Println("Sorted array with reverse around the middle:", arr)
}
