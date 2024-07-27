package collections

// Node represents a node in the linked list.
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// LinkedList represents a generic singly linked list.
type LinkedList[T comparable] struct {
	head *Node[T]
	size int
}

// NewLinkedList creates a new instance of a LinkedList.
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Add adds an element to the end of the linked list.
func (ll *LinkedList[T]) Add(value T) {
	newNode := &Node[T]{Value: value}

	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	ll.size++
}

// Remove removes the first occurrence of the specified value from the linked list.
func (ll *LinkedList[T]) Remove(value T) bool {
	if ll.head == nil {
		return false
	}

	if ll.head.Value == value {
		ll.head = ll.head.Next
		ll.size--
		return true
	}

	current := ll.head
	for current.Next != nil && current.Next.Value != value {
		current = current.Next
	}

	if current.Next == nil {
		return false
	}

	current.Next = current.Next.Next
	ll.size--
	return true
}

// Contains checks if the linked list contains the specified value.
func (ll *LinkedList[T]) Contains(value T) bool {
	current := ll.head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

// Size returns the number of elements in the linked list.
func (ll *LinkedList[T]) Size() int {
	return ll.size
}

// Get retrieves the value at the specified index.
func (ll *LinkedList[T]) Get(index int) T {
	if index < 0 || index >= ll.size {
		panic("index out of bounds")
	}

	current := ll.head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Value
}

// ToSlice converts the linked list to a slice.
func (ll *LinkedList[T]) ToSlice() []T {
	slice := make([]T, 0, ll.size)
	current := ll.head
	for current != nil {
		slice = append(slice, current.Value)
		current = current.Next
	}
	return slice
}
