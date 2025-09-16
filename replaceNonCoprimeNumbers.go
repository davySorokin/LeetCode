func replaceNonCoprimes(nums []int) []int {
    gcd := func(a, b int64) int64 {
        for b != 0 {
            a, b = b, a%b
        }
        return a
    }

    stack := make([]int64, 0, len(nums))
    for _, v := range nums {
        cur := int64(v)
        for len(stack) > 0 {
            g := gcd(stack[len(stack)-1], cur)
            if g == 1 {
                break
            }
            cur = (stack[len(stack)-1] / g) * cur // lcm(top, cur)
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, cur)
    }

    res := make([]int, len(stack))
    for i, v := range stack {
        res[i] = int(v)
    }
    return res
}
