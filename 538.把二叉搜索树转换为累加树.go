package main

func convertBST(root *TreeNode) *TreeNode {

	var dfs func(root *TreeNode, temp int)
	dfs = func(root *TreeNode, temp int) {
		if root == nil {
			return
		}
		rcount := Count(root.Right)
		dfs(root.Left, rcount+temp)
		dfs(root.Right, 0)

		root.Val = rcount + root.Val + temp

		return

	}
	dfs(root, 0)
	return root
}

func Count(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := Count(root.Left)
	r := Count(root.Right)
	return root.Val + l + r
}
