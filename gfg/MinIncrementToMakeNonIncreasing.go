package gfg

import "fmt"

func MinIncrementToMakeNonIncreasing() {
	arr := []int{3, 2, 1, 5, 4}
	count := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] {
			diff := arr[i-1] - arr[i] + 1
			count += diff
			arr[i] += diff
		}
	}
	fmt.Println(count)
}
