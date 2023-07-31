package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {

	res := []int{}
	p := head
	var dfs func(root *ListNode)
	dfs = func(root *ListNode) {
		if root == nil {
			return
		}
		dfs(root.Next)
		res = append(res, root.Val)
		return
	}
	dfs(p)
	return res
}
