package gfg

import (
	"fmt"
	"sort"
)

func SmallestSubsetWithSumGreater() {
	arr := []int{2, 1, 2}

	halfSum := 0
	for _, num := range arr {
		halfSum += num
	}
	halfSum = halfSum / 2

	sort.Ints(arr)

	res, currSum := 0, 0
	for i := len(arr) - 1; i >= 0; i-- {
		currSum += arr[i]
		res++

		if currSum > halfSum {
			fmt.Println(res)
			return
		}
	}
	fmt.Println(res)
}
