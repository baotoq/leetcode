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

func reverseList(head *ListNode) *ListNode {
	var prev, curr, next *ListNode
	curr = head

	for curr != nil {
		next = curr.Next

		curr.Next = prev

		prev = curr
		curr = next
	}

	return prev
}
