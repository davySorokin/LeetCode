package main

import (
	"fmt"
)

type LockingTree struct {
	parent    []int
	locked    []int
	children  [][]int
}

func Constructor(parent []int) LockingTree {
	n := len(parent)
	locked := make([]int, n)
	for i := range locked {
		locked[i] = -1
	}
	children := make([][]int, n)
	for i, p := range parent {
		if p != -1 {
			children[p] = append(children[p], i)
		}
	}
	return LockingTree{parent, locked, children}
}

func (this *LockingTree) Lock(num int, user int) bool {
	if this.locked[num] == -1 {
		this.locked[num] = user
		return true
	}
	return false
}

func (this *LockingTree) Unlock(num int, user int) bool {
	if this.locked[num] == user {
		this.locked[num] = -1
		return true
	}
	return false
}

func (this *LockingTree) Upgrade(num int, user int) bool {
	if this.locked[num] != -1 {
		return false
	}

	// Check if any ancestor is locked
	for ancestor := this.parent[num]; ancestor != -1; ancestor = this.parent[ancestor] {
		if this.locked[ancestor] != -1 {
			return false
		}
	}

	// Check if there is at least one locked descendant and unlock them
	lockedDescendant := false
	var unlockDescendants func(node int)
	unlockDescendants = func(node int) {
		if this.locked[node] != -1 {
			this.locked[node] = -1
			lockedDescendant = true
		}
		for _, child := range this.children[node] {
			unlockDescendants(child)
		}
	}
	unlockDescendants(num)

	if lockedDescendant {
		this.locked[num] = user
		return true
	}
	return false
}
