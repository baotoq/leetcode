package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	// Initialize two pointers
	slow := head
	fast := head

	// Traverse the list
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // Move slow pointer by 1 step
		fast = fast.Next.Next // Move fast pointer by 2 steps

		// If slow and fast pointers meet, there is a cycle
		if slow == fast {
			return true
		}
	}

	// If we exit the loop, there is no cycle
	return false
}

func hasCycle2(head *ListNode) bool {
	visited := make(map[*ListNode]struct{})

	current := head
	for current != nil {
		// If the node is already in the map, we found a cycle
		if _, ok := visited[current]; ok {
			return true
		}
		// Mark the node as visited
		visited[current] = struct{}{}
		current = current.Next
	}

	// If we reach here, there is no cycle
	return false
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
