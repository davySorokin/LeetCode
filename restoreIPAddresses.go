package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	var result []string
	if len(s) < 4 || len(s) > 12 {
		return result
	}

	var backtrack func(start int, dots int, current string)
	backtrack = func(start int, dots int, current string) {
		// If we've placed 3 dots and we're at the end of the string, it's valid
		if dots == 3 {
			// Remaining part must be a valid number
			part := s[start:]
			if isValid(part) {
				result = append(result, current+"."+part)
			}
			return
		}

		// Try adding the dot after every 1, 2, or 3 digit substring
		for i := start + 1; i <= len(s) && i <= start+3; i++ {
			part := s[start:i]
			if isValid(part) {
				// Add the dot and recursively process the next part
				backtrack(i, dots+1, current+part+".")
			}
		}
	}

	// Initialize backtracking
	backtrack(0, 0, "")
	return result
}

func isValid(part string) bool {
	// If the part is longer than 3 digits, it's invalid
	if len(part) > 3 {
		return false
	}
	// If the part has leading zeros, it's invalid
	if len(part) > 1 && part[0] == '0' {
		return false
	}
	// Convert part to an integer and check if it's within the valid range
	num, err := strconv.Atoi(part)
	if err != nil {
		return false
	}
	return num >= 0 && num <= 255
}

/*
func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("101023"))
}
*/
