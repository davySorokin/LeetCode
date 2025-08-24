package main

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordSet := make(map[string]bool, len(wordList))
	for _, w := range wordList {
		wordSet[w] = true
	}
	if !wordSet[endWord] {
		return [][]string{}
	}

	parents := make(map[string][]string) // child -> parents
	level := map[string]struct{}{beginWord: {}}
	visited := map[string]struct{}{beginWord: {}}
	found := false

	for len(level) > 0 && !found {
		nextLevel := map[string]struct{}{}
		for w := range level {
			for _, v := range neighbors(w, wordSet) {
				if _, seen := visited[v]; seen {
					continue
				}
				parents[v] = append(parents[v], w)
				if _, in := nextLevel[v]; !in {
					nextLevel[v] = struct{}{}
				}
				if v == endWord {
					found = true
				}
			}
		}
		for v := range nextLevel {
			visited[v] = struct{}{}
		}
		level = nextLevel
	}

	if _, ok := parents[endWord]; !ok {
		return [][]string{}
	}

	var res [][]string
	path := []string{endWord}
	var dfs func(string)
	dfs = func(word string) {
		if word == beginWord {
			rev := make([]string, len(path))
			for i := range path {
				rev[i] = path[len(path)-1-i]
			}
			res = append(res, rev)
			return
		}
		for _, p := range parents[word] {
			path = append(path, p)
			dfs(p)
			path = path[:len(path)-1]
		}
	}
	dfs(endWord)
	return res
}

func neighbors(s string, dict map[string]bool) []string {
	res := []string{}
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		orig := b[i]
		for c := byte('a'); c <= 'z'; c++ {
			if c == orig {
				continue
			}
			b[i] = c
			if dict[string(b)] {
				res = append(res, string(b))
			}
		}
		b[i] = orig
	}
	return res
}
