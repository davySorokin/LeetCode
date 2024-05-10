package main

import (
	"container/heap"
	"fmt"
)

type Fraction struct {
	value float64 // Value of the fraction to compare in the heap
	num   int     // Numerator of the fraction
	denom int     // Denominator of the fraction
}

type MinHeap []Fraction

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Fraction))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	h := &MinHeap{}
	heap.Init(h)

	// Push the first element from each row onto the heap
	for i := 0; i < len(arr)-1; i++ {
		heap.Push(h, Fraction{value: float64(arr[i]) / float64(arr[len(arr)-1]), num: arr[i], denom: arr[len(arr)-1]})
	}

	// Extract the smallest elements k-1 times from the heap
	var current Fraction
	for i := 0; i < k; i++ {
		current = heap.Pop(h).(Fraction)
		// Check if there is another element in the column to add to the heap
		if current.denom != current.num + 1 { // The next element in the column
			for j := len(arr) - 1; j >= 0 && arr[j] > current.num; j-- {
				if arr[j] == current.denom {
					if j > 0 {
						heap.Push(h, Fraction{value: float64(current.num) / float64(arr[j-1]), num: current.num, denom: arr[j-1]})
						break
					}
				}
			}
		}
	}

	return []int{current.num, current.denom}
}
