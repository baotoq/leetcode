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
