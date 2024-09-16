func longestPalindrome(s string) string {
    if len(s) == 0 {
        return ""
    }

    // Function to expand around the center and return the longest palindrome
    expandAroundCenter := func(s string, left, right int) string {
        for left >= 0 && right < len(s) && s[left] == s[right] {
            left--
            right++
        }
        return s[left+1 : right]
    }

    longest := ""

    for i := 0; i < len(s); i++ {
        // For odd-length palindromes, use a single center
        oddPalindrome := expandAroundCenter(s, i, i)
        // For even-length palindromes, use two centers
        evenPalindrome := expandAroundCenter(s, i, i+1)

        // Compare and update the longest palindrome found
        if len(oddPalindrome) > len(longest) {
            longest = oddPalindrome
        }
        if len(evenPalindrome) > len(longest) {
            longest = evenPalindrome
        }
    }

    return longest
}
