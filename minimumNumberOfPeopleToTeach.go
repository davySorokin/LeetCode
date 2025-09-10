func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
    m := len(languages)
    
    // Convert user languages to a slice of maps for quick lookup
    userLangs := make([]map[int]bool, m)
    for i := 0; i < m; i++ {
        userLangs[i] = make(map[int]bool)
        for _, lang := range languages[i] {
            userLangs[i][lang] = true
        }
    }

    // Find all friendships where communication is not possible
    toFix := [][]int{}
    for _, pair := range friendships {
        u, v := pair[0]-1, pair[1]-1
        canCommunicate := false
        for lang := range userLangs[u] {
            if userLangs[v][lang] {
                canCommunicate = true
                break
            }
        }
        if !canCommunicate {
            toFix = append(toFix, []int{u, v})
        }
    }

    if len(toFix) == 0 {
        return 0
    }

    minTeach := m // max possible is teaching everyone
    for lang := 1; lang <= n; lang++ {
        teachSet := make(map[int]bool)
        for _, pair := range toFix {
            u, v := pair[0], pair[1]
            if !userLangs[u][lang] {
                teachSet[u] = true
            }
            if !userLangs[v][lang] {
                teachSet[v] = true
            }
        }
        if len(teachSet) < minTeach {
            minTeach = len(teachSet)
        }
    }

    return minTeach
}
