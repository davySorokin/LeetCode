func nthMagicalNumber(n int, a int, b int) int {
    const MOD = int(1e9 + 7)
    lcm := a * b / gcd(a, b)
    low, high := 2, int(1e18)
    
    for low < high {
        mid := (low + high) / 2
        if (mid/a) + (mid/b) - (mid/lcm) < n {
            low = mid + 1
        } else {
            high = mid
        }
    }
    
    return low % MOD
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y
    }
    return x
}
