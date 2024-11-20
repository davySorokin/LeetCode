import "math"

func divide(dividend int, divisor int) int {
    // Handle edge cases
    if divisor == 0 {
        return math.MaxInt32
    }
    if dividend == math.MinInt32 && divisor == -1 {
        return math.MaxInt32 // Overflow
    }

    negative := (dividend < 0) != (divisor < 0)

    absDividend := int(math.Abs(float64(dividend)))
    absDivisor := int(math.Abs(float64(divisor)))

    quotient := 0
    for absDividend >= absDivisor {
        tempDivisor := absDivisor
        multiple := 1

        for absDividend >= (tempDivisor << 1) {
            tempDivisor <<= 1
            multiple <<= 1
        }

        absDividend -= tempDivisor
        quotient += multiple
    }

    // Apply the sign
    if negative {
        quotient = -quotient
    }

    return int(math.Max(float64(math.MinInt32), math.Min(float64(math.MaxInt32), float64(quotient))))
}
