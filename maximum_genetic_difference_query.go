package main

import (
	"fmt"
)

// TrieNode represents a node in the Trie
type TrieNode struct {
	children [2]*TrieNode
	count    int
}

// Trie represents a Trie (prefix tree) for binary numbers
type Trie struct {
	root *TrieNode
}

// NewTrie creates a new Trie
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{}}
}

// Insert adds a number to the Trie
func (t *Trie) Insert(num int) {
	node := t.root
	for i := 31; i >= 0; i-- {
		bit := (num >> i) & 1
		if node.children[bit] == nil {
			node.children[bit] = &TrieNode{}
		}
		node = node.children[bit]
		node.count++
	}
}

// Remove removes a number from the Trie
func (t *Trie) Remove(num int) {
	node := t.root
	for i := 31; i >= 0; i-- {
		bit := (num >> i) & 1
		node = node.children[bit]
		node.count--
	}
}

// MaxXOR finds the maximum XOR for a given number
func (t *Trie) MaxXOR(num int) int {
	node := t.root
	xor := 0
	for i := 31; i >= 0; i-- {
		bit := (num >> i) & 1
		if node.children[1-bit] != nil && node.children[1-bit].count > 0 {
			xor |= (1 << i)
			node = node.children[1-bit]
		} else {
			node = node.children[bit]
		}
	}
	return xor
}

func maxGeneticDifference(parents []int, queries [][]int) []int {
	n := len(parents)
	children := make([][]int, n)
	root := -1

	// Build the tree from the parents array
	for i, p := range parents {
		if p == -1 {
			root = i
		} else {
			children[p] = append(children[p], i)
		}
	}

	// Map queries to nodes
	queryMap := make(map[int][][]int)
	for i, q := range queries {
		node, val := q[0], q[1]
		queryMap[node] = append(queryMap[node], []int{i, val})
	}

	// Result array
	res := make([]int, len(queries))

	// DFS with Trie to answer queries
	trie := NewTrie()
	var dfs func(node int)
	dfs = func(node int) {
		// Insert the current node's value into the Trie
		trie.Insert(node)
		// Process all queries for the current node
		for _, q := range queryMap[node] {
			idx, val := q[0], q[1]
			res[idx] = trie.MaxXOR(val)
		}
		// Visit all children
		for _, child := range children[node] {
			dfs(child)
		}
		// Remove the current node's value from the Trie before backtracking
		trie.Remove(node)
	}

	// Start DFS from the root
	dfs(root)
	return res
}
