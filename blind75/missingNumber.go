package blind75

import "fmt"

func MissingNumber(nums []int) int {
    fmt.Println(nums)
    actualSum := 0
    length := len(nums)
    fmt.Println(length)
    for i := 0; i < length; i++ {
        actualSum = actualSum + nums[i]
        fmt.Println(actualSum)
    }
    length = length + 1
    expectedSum := (length * (length + 1))/2
    return expectedSum - actualSum
}
