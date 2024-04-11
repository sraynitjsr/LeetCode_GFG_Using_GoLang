package gfg

import "fmt"

const prime = 101

func search(pattern, text string) []int {
	var result []int

	patternLength := len(pattern)
	textLength := len(text)

	patternHash := hash(pattern, patternLength)
	textHash := hash(text[:patternLength], patternLength)

	for i := 0; i <= textLength-patternLength; i++ {
		if patternHash == textHash && text[i:i+patternLength] == pattern {
			result = append(result, i)
		}
		if i < textLength-patternLength {
			textHash = recalculateHash(text, i, patternLength, textHash)
		}
	}

	return result
}

func hash(str string, length int) int {
	hash := 0
	for i := 0; i < length; i++ {
		hash += int(str[i]) * pow(prime, i)
	}
	return hash
}

func recalculateHash(str string, oldIndex, patternLength, oldHash int) int {
	newHash := oldHash - int(str[oldIndex])
	newHash = newHash/prime + int(str[oldIndex+patternLength])*pow(prime, patternLength-1)
	return newHash
}

func pow(x, y int) int {
	if y == 0 {
		return 1
	}
	result := x
	for i := 2; i <= y; i++ {
		result *= x
	}
	return result
}

func RabinKarpAlgorithm() {
	text := "AABAACAADAABAABA"
	pattern := "AABA"
	indices := search(pattern, text)
	if len(indices) > 0 {
		fmt.Printf("Pattern found at indices: %v\n", indices)
	} else {
		fmt.Println("Pattern not found")
	}
}
