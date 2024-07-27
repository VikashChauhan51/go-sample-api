package collections

// Stack represents a generic stack data structure.
type Stack[T comparable] struct {
	elements []T
}

// NewStack creates a new instance of a Stack.
func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

// Pop removes and returns the top element of the stack.
// It panics if the stack is empty.
func (s *Stack[T]) Pop() T {
	if len(s.elements) == 0 {
		panic("Pop from an empty stack")
	}

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element
}

// Peek returns the top element of the stack without removing it.
// It panics if the stack is empty.
func (s *Stack[T]) Peek() T {
	if len(s.elements) == 0 {
		panic("Peek from an empty stack")
	}

	return s.elements[len(s.elements)-1]
}

// IsEmpty returns true if the stack is empty, false otherwise.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return len(s.elements)
}
