package main

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	var dmerge func(l, r int) *ListNode
	dmerge = func(l, r int) *ListNode {
		if l == r {
			return lists[l]
		}
		mid := (l + r) / 2

		return Merge(dmerge(l, mid), dmerge(mid+1, r))

	}
	return dmerge(0, len(lists)-1)
}

func Merge(l1, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2
	head := new(ListNode)
	ans := head
	for p1 != nil || p2 != nil {
		if p1 == nil {
			head.Next = p2
			return ans.Next
		}
		if p2 == nil {
			head.Next = p1
			return ans.Next
		}
		if p1.Val < p2.Val {
			head.Next = p1
			p1 = p1.Next
		} else {
			head.Next = p2
			p2 = p2.Next
		}
		head = head.Next
	}
	return ans.Next
}
