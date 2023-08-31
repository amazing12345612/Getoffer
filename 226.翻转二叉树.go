package main

func invertTree(root *TreeNode) *TreeNode {

	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		l := dfs(root.Left)
		r := dfs(root.Right)
		root.Left = r
		root.Right = l
		return root

	}
	return dfs(root)
}
