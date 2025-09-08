import "strconv"

func getNoZeroIntegers(n int) []int {
    for a := 1; a < n; a++ {
        b := n - a
        if !containsZero(a) && !containsZero(b) {
            return []int{a, b}
        }
    }
    return nil
}

func containsZero(num int) bool {
    return strconv.Itoa(num) != strings.ReplaceAll(strconv.Itoa(num), "0", "")
}
