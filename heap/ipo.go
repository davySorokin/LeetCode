package main

import (
	"container/heap"
	"fmt"
)

// A project structure to hold capital and profit together
type Project struct {
	capital int
	profit  int
}

// MinHeap to manage projects by increasing capital requirement
type MinCapHeap []Project

func (h MinCapHeap) Len() int            { return len(h) }
func (h MinCapHeap) Less(i, j int) bool  { return h[i].capital < h[j].capital }
func (h MinCapHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinCapHeap) Push(x interface{}) { *h = append(*h, x.(Project)) }
func (h *MinCapHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MaxHeap to manage projects by decreasing profit
type MaxProfitHeap []Project

func (h MaxProfitHeap) Len() int            { return len(h) }
func (h MaxProfitHeap) Less(i, j int) bool  { return h[i].profit > h[j].profit }
func (h MaxProfitHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxProfitHeap) Push(x interface{}) { *h = append(*h, x.(Project)) }
func (h *MaxProfitHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	minCapHeap := &MinCapHeap{}
	maxProfitHeap := &MaxProfitHeap{}

	// Push all projects into the MinCapHeap
	for i := range profits {
		heap.Push(minCapHeap, Project{capital: capital[i], profit: profits[i]})
	}

	currentCapital := w
	for i := 0; i < k; i++ {
		// Transfer all feasible projects from minCapHeap to maxProfitHeap
		for minCapHeap.Len() > 0 && (*minCapHeap)[0].capital <= currentCapital {
			heap.Push(maxProfitHeap, heap.Pop(minCapHeap).(Project))
		}

		if maxProfitHeap.Len() == 0 {
			break
		}

		// Select the project with the highest profit
		currentCapital += heap.Pop(maxProfitHeap).(Project).profit
	}

	return currentCapital
}
