package cache

type LinkedList struct {
	head   *Node
	length int
}

type Node struct {
	next *Node
	data int
}

func NewLinkedList() LinkedList {
	return LinkedList{}
}

func (l *LinkedList) AddFirst(data int) {
	node := &Node{data: data}

	if l.head != nil {
		node.next = l.head
	}

	l.head = node
	l.length++
}

func (l *LinkedList) Remove(node *Node) {
	if l.head == node {
		l.head = node.next
		l.length--
		return
	}

	head := l.head
	for head.next != nil {
		if head.next == node {
			head.next = node.next
			l.length--
			return
		}
		head = head.next
	}
}

func (l *LinkedList) RemoveByData(data int) {
	if l.head == nil {
		return
	}

	if l.head.data == data {
		l.head = l.head.next
		l.length--
		return
	}

	head := l.head
	for head.next != nil {
		if head.next.data == data {
			head.next = head.next.next
			l.length--
			return
		}
		head = head.next
	}
}

func (l *LinkedList) RemoveFirst() {
	l.Remove(l.head)
}

func (l *LinkedList) RemoveLast() {
	head := l.head
	for head.next != nil {
		head = head.next
	}
	l.Remove(head)
}

func (l *LinkedList) Length() int {
	return l.length
}

func (l *LinkedList) Head() *Node {
	return l.head
}

func (l *LinkedList) Last() *Node {
	head := l.head
	for head.next != nil {
		head = head.next
	}
	return head
}
