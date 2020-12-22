//评测题目: 无
// 1) 编写一个函数对输入的 单向链表 进行排序，要求算法越快越好，内存使用越少越好。
// 例如输入： 5 -> 3 -> 9 -> 10 -> 2 -> ···，
// 要求输出： 2 -> 3 -> 5 -> 9 -> 10 -> ··
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//  5 -> 3 -> 9 -> 10 -> 2 -> 1
// 5 3
// 3 10
// 9 1
//
// 归并排序
func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, quick := head, head.Next
	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
	}
	// 现在quick是尾部, slow是中间节点
	// 把链表分成两部分以便于处理
	mid := slow.Next
	slow.Next = nil

	left := mergeSort(head)
	right := mergeSort(mid)

	// head 是新的头结点
	if left.Val < right.Val {
		head = left
		left = left.Next
	} else {
		head = right
		right = right.Next
	}
	// cur用来遍历左右两部分链表
	cur := head

	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			cur = cur.Next
			left = left.Next
		} else {
			cur.Next = right
			cur = cur.Next
			right = right.Next
		}
	}
	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return head
}

func main() {
	// 5 -> 3 -> 9 -> 10 -> 2 -> 1
	n5 := &ListNode{
		Val: 100,
	}
	n4 := &ListNode{
		Val:  23,
		Next: n5,
	}
	n3 := &ListNode{
		Val:  44,
		Next: n4,
	}
	n2 := &ListNode{
		Val:  9,
		Next: n3,
	}
	n1 := &ListNode{
		Val:  23,
		Next: n2,
	}
	res := mergeSort(n1)
	fmt.Println("After sorting")
	for cur := res; cur != nil; cur = cur.Next {
		fmt.Printf("%d ", cur.Val)
	}
	fmt.Println("")
}
