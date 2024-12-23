package cache

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	next  *Node
	prev  *Node
	value int
	key   int
}

func NewLinkedList() LinkedList {
	return LinkedList{}
}

func (l *LinkedList) AddFirst(node *Node) {
	if l.head == nil {
		// If the list is empty, both head and tail point to the new node
		l.head = node
		l.tail = node
	} else {
		// Otherwise, add the new node before the current head
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
	l.length++
}

func (l *LinkedList) Remove(node *Node) {
	if node == l.head {
		// If the node is the head, update the head pointer
		l.head = node.next
	}
	if node == l.tail {
		// If the node is the tail, update the tail pointer
		l.tail = node.prev
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	node.prev = nil
	node.next = nil
	l.length--
}

func (l *LinkedList) Last() *Node {
	return l.tail
}
