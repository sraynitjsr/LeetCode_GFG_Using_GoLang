package blind75

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)
	for _, str := range strs {
		sortedStr := sortString(str)
		if _, ok := anagramMap[sortedStr]; !ok {
			anagramMap[sortedStr] = make([]string, 0)
		}
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}
	groupedAnagrams := make([][]string, 0)
	for _, v := range anagramMap {
		groupedAnagrams = append(groupedAnagrams, v)
	}
	return groupedAnagrams
}

func sortString(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func GroupAnagram() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groupedAnagrams := groupAnagrams(strs)
	for _, group := range groupedAnagrams {
		fmt.Println(group)
	}
}
