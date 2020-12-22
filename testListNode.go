package main

import (
    "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func newNode(val int, next *ListNode) *ListNode{
    return &ListNode{
        Val: val,
        Next: next,
    }
}

func createList(k int) *ListNode{
	head := newNode(0, nil)
	cur := head
	for i := 0; i < k; i++ {
		cur.Next = newNode(i, nil)
		cur = cur.Next
	}
	return head.Next
}

func main() {
    cur := createList(5)
    for cur != nil {
		fmt.Printf("val: %d. next is nil? %v\n", cur.Val, cur.Next == nil)
		cur = cur.Next
	}
	
}