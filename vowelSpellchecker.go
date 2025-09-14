package main

import (
	"fmt"
	"strings"
)

func spellchecker(wordlist []string, queries []string) []string {
	vowels := "aeiou"

	devowel := func(word string) string {
		word = strings.ToLower(word)
		var sb strings.Builder
		for _, ch := range word {
			if strings.ContainsRune(vowels, ch) {
				sb.WriteByte('*')
			} else {
				sb.WriteRune(ch)
			}
		}
		return sb.String()
	}

	exact := make(map[string]bool)
	caseInsensitive := make(map[string]string)
	vowelMap := make(map[string]string)

	for _, word := range wordlist {
		exact[word] = true
		lower := strings.ToLower(word)
		vowel := devowel(word)
		if _, ok := caseInsensitive[lower]; !ok {
			caseInsensitive[lower] = word
		}
		if _, ok := vowelMap[vowel]; !ok {
			vowelMap[vowel] = word
		}
	}

	ans := make([]string, len(queries))
	for i, q := range queries {
		if exact[q] {
			ans[i] = q
		} else if word, ok := caseInsensitive[strings.ToLower(q)]; ok {
			ans[i] = word
		} else if word, ok := vowelMap[devowel(q)]; ok {
			ans[i] = word
		} else {
			ans[i] = ""
		}
	}
	return ans
}


/* testing
func main() {
	wordlist := []string{"KiTe", "kite", "hare", "Hare"}
	queries := []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}
	fmt.Println(spellchecker(wordlist, queries))
}
*/
