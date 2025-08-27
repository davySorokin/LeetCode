func largestGoodInteger(num string) string {
    best := ""
    for i := 0; i+2 < len(num); i++ {
        if num[i] == num[i+1] && num[i] == num[i+2] {
            cand := num[i : i+3]
            if best == "" || cand > best {
                best = cand
            }
        }
    }
    return best
}
