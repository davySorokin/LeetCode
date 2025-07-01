func possibleStringCount(word string) int {

    count := 1 // The word itself is one valid possibility

    // Group characters
    i := 0
    for i < len(word) {
        j := i
        for j < len(word) && word[j] == word[i] {
            j++
        }
        groupLen := j - i
        if groupLen >= 2 {
            count += groupLen - 1
        }
        i = j
    }

    return count
}
