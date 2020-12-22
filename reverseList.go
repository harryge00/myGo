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

func reverseList(k int, node *ListNode) *ListNode{
	if node == nil {
		return node
	}
    head := newNode(0, node)
    cur := node
    lastHead := head 
    i := 1
    for cur.Next != nil {
        if i == k {
            lastHead = cur
            cur = cur.Next
            i = 1
            continue
        }
        
		next := cur.Next
		cur.Next = next.Next		
		next.Next = lastHead.Next
		lastHead.Next = next
		i++
	}
	
    return head.Next
}

func main() {
	node5 := newNode(5, nil)
	fmt.Println(node5.Next)
	
    node4 := newNode(4, node5)
    node3 := newNode(3, node4)
    node2 := newNode(2, node3)
	node1 := newNode(1, node2)
    cur := reverseList(2, node1)
    for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
}