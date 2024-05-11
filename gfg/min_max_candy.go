package gfg

import (
	"fmt"
	"sort"
)

func findMinimum(arr []int, n, k int) int {
	res := 0
	for i := 0; i < n; i++ {
		res += arr[i]
		n = n - k
	}
	return res
}

func findMaximum(arr []int, n, k int) int {
	res, index := 0, 0

	for i := n - 1; i >= index; i-- {
		res += arr[i]
		index += k
	}
	return res
}

func MinMaxCandy() {
	arr := []int{3, 2, 1, 4}
	n := len(arr)
	k := 2
	sort.Ints(arr)
	fmt.Println(findMinimum(arr, n, k), findMaximum(arr, n, k))
}
