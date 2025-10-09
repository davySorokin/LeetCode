func minTime(skill []int, mana []int) int64 {
    n := len(skill)
    f := make([]int64, n)

    for _, x := range mana {
        now := f[0]
        for i := 1; i < n; i++ {
            candidate := now + int64(skill[i-1])*int64(x)
            if candidate > f[i] {
                now = candidate
            } else {
                now = f[i]
            }
        }
        f[n-1] = now + int64(skill[n-1])*int64(x)
        for i := n - 2; i >= 0; i-- {
            f[i] = f[i+1] - int64(skill[i+1])*int64(x)
        }
    }

    return f[n-1]
}
