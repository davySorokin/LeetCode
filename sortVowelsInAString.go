package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortVowels(s string) string {
	vowelsSet := "aeiouAEIOU"
	isVowel := func(c byte) bool {
		return strings.ContainsRune(vowelsSet, rune(c))
	}

	// Extract vowels
	var vowels []byte
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			vowels = append(vowels, s[i])
		}
	}

	// Sort vowels by ASCII
	sort.Slice(vowels, func(i, j int) bool {
		return vowels[i] < vowels[j]
	})

	// Build result
	result := []byte(s)
	vowelIdx := 0
	for i := 0; i < len(result); i++ {
		if isVowel(result[i]) {
			result[i] = vowels[vowelIdx]
			vowelIdx++
		}
	}

	return string(result)
}

/*
func main() {
	fmt.Println(sortVowels("lEetcOde")) // Output: "lEOtcede"
	fmt.Println(sortVowels("lYmpH"))    // Output: "lYmpH"
}
*/
