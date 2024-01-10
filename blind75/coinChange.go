package blind75

import "fmt"

func coinChangeWays(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func CoinChange() {
	coins := []int{1, 2, 5}
	amount := 5
	ways := coinChangeWays(coins, amount)
	fmt.Printf("Number of ways to make %d using coins %v: %d\n", amount, coins, ways)
}
