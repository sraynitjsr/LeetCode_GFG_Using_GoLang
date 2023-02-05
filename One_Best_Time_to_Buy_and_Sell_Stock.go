package main

import "fmt"

func maxProfit(prices []int) int {
	buy := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if buy > prices[i] {
			buy = prices[i]
		}
		if prices[i]-buy > maxProfit {
			maxProfit = prices[i] - buy
		}
	}
	return maxProfit
}

func main() {
	mySlice := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(mySlice))
}
