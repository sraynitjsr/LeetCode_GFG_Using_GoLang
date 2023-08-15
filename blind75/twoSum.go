package blind75

import "fmt"

func TwoSum(mySlice []int, target int) []int {
    helperValueIndexMap := make(map[int]int)
    output := []int{}
    for index,element := range mySlice {
        anotherIndex, presentOrNot := helperValueIndexMap[target-element]
        if !presentOrNot {
            helperValueIndexMap[element] = index
        } else {
            output = append(output, index)
            output = append(output, anotherIndex)
            break
        }
    }
    return output
}
