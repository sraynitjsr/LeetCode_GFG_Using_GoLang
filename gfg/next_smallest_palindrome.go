package gfg

import (
    "fmt"
    "strconv"
)

func isPalindrome(n int) bool {
    str := strconv.Itoa(n)
    length := len(str)
    for i := 0; i < length/2; i++ {
        if str[i] != str[length-1-i] {
            return false
        }
    }
    return true
}

func nextPalindrome(num int) int {
    for {
        num++
        if isPalindrome(num) {
            return num
        }
    }
}

func NextSmallestPalindrome() {
    num := 123
    nextPal := nextPalindrome(num)
    fmt.Printf("Next smallest palindrome after %d is %d\n", num, nextPal)
}
