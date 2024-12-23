package merge_two_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	curr := &ListNode{}
	r := curr

	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			curr = p1
			p1 = p1.Next
		} else {
			curr = p2
			p2 = p2.Next
		}

		r = curr
		curr = curr.Next
	}

	return r
}
