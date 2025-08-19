package main

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return []string{}
	}

	enc := map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	mask := (1 << 20) - 1
	x := 0

	seen := make(map[int]struct{})
	added := make(map[int]struct{})
	res := []string{}

	for i := 0; i < len(s); i++ {
		x = ((x << 2) & mask) | enc[s[i]]
		if i >= 9 {
			if _, ok := seen[x]; ok {
				if _, dup := added[x]; !dup {
					res = append(res, s[i-9:i+1])
					added[x] = struct{}{}
				}
			} else {
				seen[x] = struct{}{}
			}
		}
	}
	return res
}
