package main

import (
	"container/heap"
)

type Class struct {
	p, t int
}

func gain(c Class) float64 {
	return float64(c.p+1)/float64(c.t+1) - float64(c.p)/float64(c.t)
}

type MaxHeap []Class

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	// Max-heap by gain
	return gain(h[i]) > gain(h[j])
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Class))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := &MaxHeap{}
	for _, c := range classes {
		*h = append(*h, Class{p: c[0], t: c[1]})
	}
	heap.Init(h)

	for i := 0; i < extraStudents; i++ {
		c := heap.Pop(h).(Class)
		c.p++
		c.t++
		heap.Push(h, c)
	}

	total := 0.0
	for _, c := range *h {
		total += float64(c.p) / float64(c.t)
	}
	return total / float64(len(classes))
}
