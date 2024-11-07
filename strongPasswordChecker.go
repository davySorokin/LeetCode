import "unicode"

func strongPasswordChecker(password string) int {
    length := len(password)
    
    // tracking - missing character types
    hasLower, hasUpper, hasDigit := false, false, false
    for _, ch := range password {
        if unicode.IsLower(ch) {
            hasLower = true
        } else if unicode.IsUpper(ch) {
            hasUpper = true
        } else if unicode.IsDigit(ch) {
            hasDigit = true
        }
    }
    missingTypes := 0
    if !hasLower {
        missingTypes++
    }
    if !hasUpper {
        missingTypes++
    }
    if !hasDigit {
        missingTypes++
    }

    replace := 0
    oneSeq, twoSeq := 0, 0
    i := 2
    for i < length {
        if password[i] == password[i-1] && password[i] == password[i-2] {
            lengthOfSeq := 2
            for i < length && password[i] == password[i-1] {
                lengthOfSeq++
                i++
            }
            replace += lengthOfSeq / 3
            if lengthOfSeq%3 == 0 {
                oneSeq++
            } else if lengthOfSeq%3 == 1 {
                twoSeq++
            }
        } else {
            i++
        }
    }

    if length < 6 {
        return max(6-length, missingTypes)
    } else if length <= 20 {
        return max(replace, missingTypes)
    } else {
        deleteCount := length - 20

        replace -= min(deleteCount, oneSeq*1) / 1
        replace -= min(max(deleteCount-oneSeq, 0), twoSeq*2) / 2
        replace -= max(deleteCount-oneSeq-twoSeq*2, 0) / 3

        return deleteCount + max(replace, missingTypes)
    }
}

// helper - function to find maximum of two numbers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// helper - function to find minimum of two numbers
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
