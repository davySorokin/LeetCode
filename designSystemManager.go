package main

import "container/heap"

type entry struct {
	priority int
	taskId   int
}

type maxHeap []entry

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Less(i, j int) bool {
	// Max-heap by priority; tie-break by larger taskId
	if h[i].priority != h[j].priority {
		return h[i].priority > h[j].priority
	}
	return h[i].taskId > h[j].taskId
}
func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(entry)) }
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type taskInfo struct {
	userId   int
	priority int
}

type TaskManager struct {
	info map[int]taskInfo // taskId -> (userId, priority)
	h    maxHeap
}

func Constructor(tasks [][]int) TaskManager {
	tm := TaskManager{
		info: make(map[int]taskInfo),
		h:    maxHeap{},
	}
	for _, t := range tasks {
		userId, taskId, priority := t[0], t[1], t[2]
		tm.info[taskId] = taskInfo{userId: userId, priority: priority}
		heap.Push(&tm.h, entry{priority: priority, taskId: taskId})
	}
	heap.Init(&tm.h)
	return tm
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	this.info[taskId] = taskInfo{userId: userId, priority: priority}
	heap.Push(&this.h, entry{priority: priority, taskId: taskId})
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	ti := this.info[taskId]
	ti.priority = newPriority
	this.info[taskId] = ti
	heap.Push(&this.h, entry{priority: newPriority, taskId: taskId})
}

func (this *TaskManager) Rmv(taskId int) {
	delete(this.info, taskId) // lazy delete
}

func (this *TaskManager) ExecTop() int {
	for this.h.Len() > 0 {
		e := heap.Pop(&this.h).(entry)
		ti, ok := this.info[e.taskId]
		if !ok {
			continue // removed
		}
		if ti.priority == e.priority {
			delete(this.info, e.taskId)
			return ti.userId
		}
		// else stale due to edit; skip
	}
	return -1
}
