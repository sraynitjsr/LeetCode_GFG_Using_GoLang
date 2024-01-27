package blind75

import "fmt"

type Stock struct {
	Buy  int
	Sell int
}

func findBuySell(prices []int) []Stock {
	n := len(prices)
	var result []Stock

	if n < 2 {
		return result
	}

	i := 0
	for i < n-1 {
		for i < n-1 && prices[i+1] <= prices[i] {
			i++
		}

		if i == n-1 {
			break
		}

		buy := i

		i++
		for i < n && prices[i] >= prices[i-1] {
			i++
		}

		sell := i - 1

		result = append(result, Stock{Buy: buy, Sell: sell})
	}

	return result
}

func StockBuySell() {
	prices := []int{100, 180, 260, 310, 40, 535, 695}

	result := findBuySell(prices)

	for i, stock := range result {
		fmt.Printf("Trade %d: Buy on day %d, Sell on day %d\n", i+1, stock.Buy, stock.Sell)
	}
}
