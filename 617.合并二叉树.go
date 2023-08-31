package main

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	var dfs func(root1 *TreeNode, root2 *TreeNode) *TreeNode
	dfs = func(root1 *TreeNode, root2 *TreeNode) *TreeNode {

		if root1 == nil && root2 == nil {
			return nil
		} else if root1 != nil && root2 == nil {
			return root1
		} else if root2 != nil && root1 == nil {
			return root2
		}
		root1.Val = root1.Val + root2.Val
		root1.Left = dfs(root1.Left, root2.Left)
		root1.Right = dfs(root1.Right, root2.Right)
		return root1
	}
	return dfs(root1, root2)
}
