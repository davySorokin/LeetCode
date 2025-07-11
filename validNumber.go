func isNumber(s string) bool {
    i, n := 0, len(s)

    // Skip leading spaces
    for i < n && s[i] == ' ' {
        i++
    }

    // Parse optional sign
    if i < n && (s[i] == '+' || s[i] == '-') {
        i++
    }

    isNumeric := false
    // Parse digits before decimal point
    for i < n && isDigit(s[i]) {
        i++
        isNumeric = true
    }

    // Parse decimal point and digits after it
    if i < n && s[i] == '.' {
        i++
        for i < n && isDigit(s[i]) {
            i++
            isNumeric = true
        }
    }

    // Parse exponent if any
    if isNumeric && i < n && (s[i] == 'e' || s[i] == 'E') {
        i++
        isNumeric = false
        if i < n && (s[i] == '+' || s[i] == '-') {
            i++
        }
        for i < n && isDigit(s[i]) {
            i++
            isNumeric = true
        }
    }

    return isNumeric && i == n
}

func isDigit(c byte) bool {
    return c >= '0' && c <= '9'
}
