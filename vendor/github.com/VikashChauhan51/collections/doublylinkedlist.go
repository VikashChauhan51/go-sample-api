package collections

// Node represents a node in the doubly linked list.
type DoublyNode[T comparable] struct {
	Value T
	Prev  *DoublyNode[T]
	Next  *DoublyNode[T]
}

// DoublyLinkedList represents a generic doubly linked list.
type DoublyLinkedList[T comparable] struct {
	head *DoublyNode[T]
	tail *DoublyNode[T]
	size int
}

// NewDoublyLinkedList creates a new instance of a DoublyLinkedList.
func NewDoublyLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

// Add adds an element to the end of the doubly linked list.
func (dll *DoublyLinkedList[T]) Add(value T) {
	newNode := &DoublyNode[T]{Value: value}

	if dll.tail == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.Next = newNode
		newNode.Prev = dll.tail
		dll.tail = newNode
	}
	dll.size++
}

// Remove removes the first occurrence of the specified value from the doubly linked list.
func (dll *DoublyLinkedList[T]) Remove(value T) bool {
	if dll.head == nil {
		return false
	}

	current := dll.head
	for current != nil {
		if current.Value == value {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				dll.head = current.Next
			}

			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				dll.tail = current.Prev
			}

			dll.size--
			return true
		}
		current = current.Next
	}

	return false
}

// Contains checks if the doubly linked list contains the specified value.
func (dll *DoublyLinkedList[T]) Contains(value T) bool {
	current := dll.head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

// Size returns the number of elements in the doubly linked list.
func (dll *DoublyLinkedList[T]) Size() int {
	return dll.size
}

// Get retrieves the value at the specified index.
func (dll *DoublyLinkedList[T]) Get(index int) T {
	if index < 0 || index >= dll.size {
		panic("index out of bounds")
	}

	current := dll.head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Value
}

// ToSlice converts the doubly linked list to a slice.
func (dll *DoublyLinkedList[T]) ToSlice() []T {
	slice := make([]T, 0, dll.size)
	current := dll.head
	for current != nil {
		slice = append(slice, current.Value)
		current = current.Next
	}
	return slice
}

// ReverseToSlice converts the doubly linked list to a slice in reverse order.
func (dll *DoublyLinkedList[T]) ReverseToSlice() []T {
	slice := make([]T, 0, dll.size)
	current := dll.tail
	for current != nil {
		slice = append(slice, current.Value)
		current = current.Prev
	}
	return slice
}
