package blind75

import "fmt"

func productExceptSelf(nums []int) []int {
    n := len(nums)
    result := make([]int, n)

    for i := range result {
        result[i] = 1
    }

    leftProduct := 1
    for i := 0; i < n; i++ {
        result[i] *= leftProduct
        leftProduct *= nums[i]
    }

    rightProduct := 1
    for i := n - 1; i >= 0; i-- {
        result[i] *= rightProduct
        rightProduct *= nums[i]
    }

    return result
}

func ProductExceptSelf() {
    nums := []int{1, 2, 3, 4}
    result := productExceptSelf(nums)

    fmt.Println("Original array:", nums)
    fmt.Println("Product of array except self:", result)
}
