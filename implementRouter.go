package main

import (
	"sort"
)

type packet struct {
	s, d, t int
}

type tsList struct {
	arr   []int
	start int
}

func (b *tsList) push(t int) {
	b.arr = append(b.arr, t)
}

func (b *tsList) popFront() {
	b.start++
	// compact occasionally
	if b.start > 64 && b.start*2 >= len(b.arr) {
		b.arr = append([]int(nil), b.arr[b.start:]...)
		b.start = 0
	}
}

func (b *tsList) countInRange(lo, hi int) int {
	if b == nil {
		return 0
	}
	n := len(b.arr) - b.start
	if n <= 0 {
		return 0
	}
	// lower_bound for lo
	L := sort.Search(n, func(i int) bool { return b.arr[b.start+i] >= lo })
	// upper_bound for hi
	R := sort.Search(n, func(i int) bool { return b.arr[b.start+i] > hi })
	return R - L
}

type key struct{ s, d, t int }

type Router struct {
	capacity int
	q        []packet
	head     int // index of front in q
	present  map[key]struct{}
	perDest  map[int]*tsList
	size     int
}

func Constructor(memoryLimit int) Router {
	return Router{
		capacity: memoryLimit,
		q:        make([]packet, 0),
		head:     0,
		present:  make(map[key]struct{}),
		perDest:  make(map[int]*tsList),
		size:     0,
	}
}

func (this *Router) compactQueueIfNeeded() {
	if this.head > 64 && this.head*2 >= len(this.q) {
		this.q = append([]packet(nil), this.q[this.head:]...)
		this.head = 0
	}
}

func (this *Router) ensureDest(d int) *tsList {
	if b, ok := this.perDest[d]; ok {
		return b
	}
	b := &tsList{arr: make([]int, 0), start: 0}
	this.perDest[d] = b
	return b
}

func (this *Router) evictOldest() {
	p := this.q[this.head]
	this.head++
	this.size--
	delete(this.present, key{p.s, p.d, p.t})
	this.ensureDest(p.d).popFront()
	this.compactQueueIfNeeded()
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	k := key{source, destination, timestamp}
	if _, ok := this.present[k]; ok {
		return false
	}
	if this.size == this.capacity {
		this.evictOldest()
	}
	this.q = append(this.q, packet{source, destination, timestamp})
	this.size++
	this.present[k] = struct{}{}
	this.ensureDest(destination).push(timestamp)
	return true
}

func (this *Router) ForwardPacket() []int {
	if this.size == 0 {
		return []int{}
	}
	p := this.q[this.head]
	this.head++
	this.size--
	delete(this.present, key{p.s, p.d, p.t})
	this.ensureDest(p.d).popFront()
	this.compactQueueIfNeeded()
	return []int{p.s, p.d, p.t}
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	b, ok := this.perDest[destination]
	if !ok {
		return 0
	}
	return b.countInRange(startTime, endTime)
}
