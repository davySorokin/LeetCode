package main

import (
    "container/heap"
    "fmt"
)

// IntHeap defines a min-heap structure.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

// MaxHeap defines a max-heap structure by embedding IntHeap and overriding Less.
type MaxHeap struct {
    IntHeap // Embedding IntHeap for implementation reuse
}

func (h MaxHeap) Less(i, j int) bool { return h.IntHeap[i] > h.IntHeap[j] }

// MedianFinder is a data structure that finds the median from a data stream.
type MedianFinder struct {
    low  MaxHeap  // Max-heap to store the smaller half of the numbers.
    high IntHeap  // Min-heap to store the larger half of the numbers.
}

// Constructor initializes a new MedianFinder object.
func Constructor() MedianFinder {
    mf := MedianFinder{
        low:  MaxHeap{},
        high: IntHeap{},
    }
    heap.Init(&mf.low)
    heap.Init(&mf.high)
    return mf
}

// AddNum adds a number to the data structure.
func (this *MedianFinder) AddNum(num int) {
    if this.low.Len() == 0 || num <= this.low.IntHeap[0] {
        heap.Push(&this.low, num)
    } else {
        heap.Push(&this.high, num)
    }

    // Balancing the heaps to ensure that the max-heap has the same number or one more element than the min-heap
    if this.low.Len() > this.high.Len()+1 {
        heap.Push(&this.high, heap.Pop(&this.low).(int))
    } else if this.high.Len() > this.low.Len() {
        heap.Push(&this.low, heap.Pop(&this.high).(int))
    }
}

// FindMedian returns the median of all elements so far.
func (this *MedianFinder) FindMedian() float64 {
    if this.low.Len() > this.high.Len() {
        return float64(this.low.IntHeap[0])
    }
    return (float64(this.low.IntHeap[0]) + float64(this.high[0])) / 2.0  // Correctly accessing the elements
}
