package main

import (
	"container/heap"
	"fmt"
)

// Define a pair struct and a heap to manage these pairs
type Pair struct {
	sum  int
	idx1 int
	idx2 int
}

type MinHeap []Pair

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// kSmallestPairs finds the k pairs with the smallest sums
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return nil
	}

	h := &MinHeap{}
	heap.Init(h)
	
	// Initialise the heap with the smallest element from nums1 paired with all elements in nums2
	for j := 0; j < len(nums2); j++ {
		heap.Push(h, Pair{sum: nums1[0] + nums2[j], idx1: 0, idx2: j})
	}

	var result [][]int
	for i := 0; i < k && h.Len() > 0; i++ {
		smallest := heap.Pop(h).(Pair)
		result = append(result, []int{nums1[smallest.idx1], nums2[smallest.idx2]})
		if smallest.idx1 + 1 < len(nums1) {
			heap.Push(h, Pair{sum: nums1[smallest.idx1+1] + nums2[smallest.idx2], idx1: smallest.idx1 + 1, idx2: smallest.idx2})
		}
	}

	return result
}
