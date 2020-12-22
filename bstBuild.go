// Build BST
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: nums[0],
	}
	// 用栈保存一个递增序列
	// 如7 5 3
	stack := []*TreeNode{root}
	for i := 1; i < len(nums); i++ {
		child := &TreeNode{
			Val: nums[i],
		}

		topNode := stack[len(stack)-1]
		for len(stack) > 0 && stack[len(stack)-1].Val < child.Val {
			// 删除栈顶
			topNode = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
		}
		if topNode.Val < child.Val {
			topNode.Right = child
		} else {
			topNode.Left = child
		}
		stack = append(stack, child)
	}
	return root
}

func main() {
	root := buildBST([]int{7, 5, 3, 6, 10, 9, 12})
	fmt.Println(root, root.Left, root.Left.Left, root.Left.Right, root.Right)
}
