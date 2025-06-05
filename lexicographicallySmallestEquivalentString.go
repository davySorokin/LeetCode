func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
    parent := make([]int, 26)
    for i := 0; i < 26; i++ {
        parent[i] = i
    }

    find := func(x int) int {
        for x != parent[x] {
            parent[x] = parent[parent[x]]
            x = parent[x]
        }
        return x
    }

    union := func(x, y int) {
        px, py := find(x), find(y)
        if px == py {
            return
        }
        if px < py {
            parent[py] = px
        } else {
            parent[px] = py
        }
    }

    for i := 0; i < len(s1); i++ {
        union(int(s1[i]-'a'), int(s2[i]-'a'))
    }

    res := make([]byte, len(baseStr))
    for i := 0; i < len(baseStr); i++ {
        res[i] = byte(find(int(baseStr[i]-'a')) + 'a')
    }

    return string(res)
}
