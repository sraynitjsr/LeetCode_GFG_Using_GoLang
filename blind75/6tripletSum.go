package blind75

import "fmt"

func find3Numbers(A []int, arrSize int, sum int) bool {
	for i := 0; i < arrSize-2; i++ {
		s := make(map[int]bool)
		currSum := sum - A[i]
		for j := i + 1; j < arrSize; j++ {
			requiredValue := currSum - A[j]
			if s[requiredValue] {
				fmt.Printf("Triplet is %d, %d, %d\n", A[i], A[j], requiredValue)
				return true
			}
			s[A[j]] = true
		}
	}
	return false
}

func TrpletSum() {
	A := []int{1, 4, 45, 6, 10, 8}
	sum := 22
	arrSize := len(A)
	if !find3Numbers(A, arrSize, sum) {
		fmt.Println("No triplet found with the given sum.")
	}
}
