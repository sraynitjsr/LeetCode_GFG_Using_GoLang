package blind75

import (
	"fmt"
	"sort"
	"strings"
)

func areAnagrams(str1, str2 string) bool {

	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	str1 = strings.ReplaceAll(str1, " ", "")
	str2 = strings.ReplaceAll(str2, " ", "")

	if len(str1) != len(str2) {
		return false
	}

	str1Runes := []rune(str1)
	str2Runes := []rune(str2)

	sort.Slice(str1Runes, func(i, j int) bool {
		return str1Runes[i] < str1Runes[j]
	})

	sort.Slice(str2Runes, func(i, j int) bool {
		return str2Runes[i] < str2Runes[j]
	})

	return string(str1Runes) == string(str2Runes)
}

func Anagram() {
	word1 := "listen"
	word2 := "silent"

	if areAnagrams(word1, word2) {
		fmt.Printf("%s and %s are anagrams.\n", word1, word2)
	} else {
		fmt.Printf("%s and %s are not anagrams.\n", word1, word2)
	}
}
