package gfg

import "fmt"

func minFibonacciTermsSum(K int) int {
	fib := []int{1, 1}
	for fib[len(fib)-1] < K {
		next := fib[len(fib)-1] + fib[len(fib)-2]
		fib = append(fib, next)
	}
	count := 0
	for i := len(fib) - 1; i >= 0 && K > 0; i-- {
		if fib[i] <= K {
			K -= fib[i]
			count++
		}
	}
	return count
}

func MinimumFibonacciNumbersEqualToK() {
	K := 10
	fmt.Printf("Minimum number of Fibonacci terms with sum equal to %d: %d\n", K, minFibonacciTermsSum(K))
}
