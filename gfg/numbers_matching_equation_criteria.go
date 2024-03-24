package gfg

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func findAns(a, b, n int) int {
	lcm := (a * b) / gcd(a, b)
	multiples := (n / lcm) + 1
	answer := max(a, b) * multiples
	lastValue := lcm*(n/lcm) + max(a, b)
	if lastValue > n {
		answer -= (lastValue - n - 1)
	}
	return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func CountNumbersMatchingEquationCriteria() {
	a, b, n := 3, 4, 25
	fmt.Println(findAns(a, b, n))
}
