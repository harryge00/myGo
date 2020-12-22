class Solution:
    def bstFromPreorder(self, preorder: List[int]) -> TreeNode:
        n = len(preorder)
        if not n:
            return None
        
        root = TreeNode(preorder[0])         
        stack = [root, ] 
        
        for i in range(1, n):
            # take the last element of the stack as a parent
            # and create a child from the next preorder element
            node, child = stack[-1], TreeNode(preorder[i])
            # adjust the parent 
            while stack and stack[-1].val < child.val: 
                node = stack.pop()
             
            # follow BST logic to create a parent-child link
            if node.val < child.val:
                node.right = child 
            else:
                node.left = child 
            # add the child into stack
            stack.append(child)
  
        return root