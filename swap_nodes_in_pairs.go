package main

import (
	"fmt"
)

// Definition for singly-linked list is already provided by the precompiled package.

func swapPairs(head *ListNode) *ListNode {
	// Dummy node to handle edge cases such as swapping the head
	dummy := &ListNode{Next: head}
	current := dummy

	for current.Next != nil && current.Next.Next != nil {
		first := current.Next
		second := current.Next.Next

		// Swapping the nodes
		first.Next = second.Next
		second.Next = first
		current.Next = second

		// Move to the next pair
		current = first
	}

	return dummy.Next
}

// Helper function to print the linked list
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

// Helper function to create a linked list from a slice
func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, value := range values[1:] {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}
	return head
}
