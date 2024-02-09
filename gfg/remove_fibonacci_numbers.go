package gfg

import "fmt"

func isFibonacci(n int) bool {
	a, b := 0, 1
	for a <= n {
		if a == n {
			return true
		}
		a, b = b, a+b
	}
	return false
}

func removeFibonacci(arr []int) []int {
	result := []int{}
	for _, num := range arr {
		if !isFibonacci(num) {
			result = append(result, num)
		}
	}
	return result
}

func RemoveFibonacciNumbersFromArray() {
	arr := []int{4, 6, 5, 3, 8, 7, 10, 11, 14, 15}
	fmt.Println("Original array:", arr)
	arrWithoutFibonacci := removeFibonacci(arr)
	fmt.Println("Array after removing Fibonacci numbers:", arrWithoutFibonacci)
}
