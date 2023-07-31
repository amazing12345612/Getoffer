package main

func levelOrder(root *TreeNode) (res []int) {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]
		res = append(res, t.Val)
		if t.Left != nil {
			queue = append(queue, t.Left)
		}
		if t.Right != nil {
			queue = append(queue, t.Right)
		}
	}
	return

}
