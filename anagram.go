package main

import (
	"fmt"
	"strings"
)

func checkAnagram() {

	str1 := "Race"

	str2 := "Care"

	str1 = strings.ToUpper(str1)

	str2 = strings.ToUpper(str2)

	if len(str1) != len(str2) {
		fmt.Println("Not Anagram, Length Not Same")
		return
	}

	myFirstMap := make(map[string]int)

	for _, val := range str1 {
		frequency, present := myFirstMap[string(val)]
		if !present {
			myFirstMap[string(val)] = 1
		} else {
			myFirstMap[string(val)] = frequency + 1
		}
	}

	mySecondMap := make(map[string]int)

	for _, val := range str2 {
		frequency, present := mySecondMap[string(val)]
		if !present {
			mySecondMap[string(val)] = 1
		} else {
			mySecondMap[string(val)] = frequency + 1
		}
	}

	for key, val := range myFirstMap {
		valueFromSecondMap, present := mySecondMap[key]
		if !present || val != valueFromSecondMap {
			fmt.Println("Not Anagram, Different Strings")
			return
		}
	}

	fmt.Println("Anagram")
}
