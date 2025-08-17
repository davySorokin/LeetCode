package main

import (
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	// Use int64 to safely handle abs(INT_MIN)
	n := int64(numerator)
	d := int64(denominator)

	var b strings.Builder

	// Sign
	if (n < 0) != (d < 0) {
		b.WriteByte('-')
	}
	if n < 0 {
		n = -n
	}
	if d < 0 {
		d = -d
	}

	// Integer part
	intPart := n / d
	b.WriteString(strconv.FormatInt(intPart, 10))
	rem := n % d
	if rem == 0 {
		return b.String()
	}

	b.WriteByte('.')

	// remainder -> index in builder string after current content
	seen := make(map[int64]int)
	for rem != 0 {
		if idx, ok := seen[rem]; ok {
			s := b.String()
			return s[:idx] + "(" + s[idx:] + ")"
		}
		seen[rem] = b.Len()
		rem *= 10
		b.WriteByte(byte('0' + (rem/d)))
		rem %= d
	}

	return b.String()
}
