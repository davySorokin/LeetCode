package main

import (
	"fmt"
)

func kthCharacter(k int) byte {
	word := []byte{'a'}

	for len(word) < k {
		next := make([]byte, len(word))
		for i, ch := range word {
			// Shift character with wraparound
			if ch == 'z' {
				next[i] = 'a'
			} else {
				next[i] = ch + 1
			}
		}
		word = append(word, next...)
	}

	return word[k-1]
}

/*
func main() {
	fmt.Println(string(kthCharacter(5)))  // Output: "b"
	fmt.Println(string(kthCharacter(10))) // Output: "c"
}
*/
