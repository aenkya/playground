package datastructures

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

type SinglyLinkedListNode struct {
	value any
	next  *SinglyLinkedListNode
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (sll *SinglyLinkedList) Append(val any) {
	node := &SinglyLinkedListNode{value: val}

	if sll.head == nil {
		sll.head = node
		sll.tail = node
		return
	}

	sll.tail.next = node
	sll.tail = node
}

func (sll *SinglyLinkedList) Prepend(val any) {
	node := &SinglyLinkedListNode{value: val}

	if sll.head == nil {
		sll.head = node
		sll.tail = node
		return
	}

	node.next = sll.head
	sll.head = node
}

func (sll *SinglyLinkedList) InsertAt(index int, val any) {
	if index == 0 {
		sll.Prepend(val)
		return
	}

	node := &SinglyLinkedListNode{value: val}
	prev := sll.head

	for i := 0; i < index-1; i++ {
		prev = prev.next
	}

	node.next = prev.next
	prev.next = node
}

func (sll *SinglyLinkedList) RemoveAt(index int) {
	if index == 0 {
		sll.head = sll.head.next
		return
	}

	prev := sll.head
	for i := 0; i < index-1; i++ {
		prev = prev.next
	}

	prev.next = prev.next.next

	if prev.next == nil {
		sll.tail = prev
	}
}

func (sll *SinglyLinkedList) Reverse() {
	if sll.head == nil || sll.head.next == nil {
		return
	}

	first := sll.head
	second := first.next

	for second != nil {
		temp := second.next
		second.next = first
		first = second
		second = temp
	}

	sll.head.next = nil
	sll.tail = sll.head
	sll.head = first
}

func (sll *SinglyLinkedList) ToArray() []any {
	var arr []any
	current := sll.head

	for current != nil {
		arr = append(arr, current.value)
		current = current.next
	}

	return arr
}
