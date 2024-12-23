package linked_list_cycle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	// Create a linked list with a cycle
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 // Creates a cycle

	act := hasCycle(node1)

	assert.True(t, act)
}

func Test2(t *testing.T) {
	// Create a linked list with a cycle
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 // Creates a cycle

	act := hasCycle2(node1)

	assert.True(t, act)
}
