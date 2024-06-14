import (
    "strings"
)

func spellchecker(wordlist []string, queries []string) []string {
    exactMatch := make(map[string]string)
    caseInsensitive := make(map[string]string)
    vowelError := make(map[string]string)

    for _, word := range wordlist {
        exactMatch[word] = word
        lowerWord := strings.ToLower(word)
        if _, exists := caseInsensitive[lowerWord]; !exists {
            caseInsensitive[lowerWord] = word
        }
        vowelWord := replaceVowels(lowerWord)
        if _, exists := vowelError[vowelWord]; !exists {
            vowelError[vowelWord] = word
        }
    }

    result := make([]string, len(queries))
    for i, query := range queries {
        if word, exists := exactMatch[query]; exists {
            result[i] = word
        } else if word, exists := caseInsensitive[strings.ToLower(query)]; exists {
            result[i] = word
        } else if word, exists := vowelError[replaceVowels(strings.ToLower(query))]; exists {
            result[i] = word
        } else {
            result[i] = ""
        }
    }

    return result
}

func replaceVowels(word string) string {
    vowels := "aeiou"
    result := []rune(word)
    for i, char := range result {
        if strings.ContainsRune(vowels, char) {
            result[i] = '*'
        }
    }
    return string(result)
}
