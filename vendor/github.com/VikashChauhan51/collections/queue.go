package collections

// Queue represents a generic, thread-safe queue data structure.
type Queue[T comparable] struct {
	elements []T
}

// NewQueue creates a new instance of a Queue.
func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{}
}

// Enqueue adds an element to the end of the queue.
func (q *Queue[T]) Enqueue(element T) {
	q.elements = append(q.elements, element)
}

// Dequeue removes and returns the element from the front of the queue.
// It panics if the queue is empty.
func (q *Queue[T]) Dequeue() T {
	if len(q.elements) == 0 {
		panic("Dequeue from an empty queue")
	}

	element := q.elements[0]
	q.elements = q.elements[1:]
	return element
}

// Peek returns the element at the front of the queue without removing it.
// It panics if the queue is empty.
func (q *Queue[T]) Peek() T {
	if len(q.elements) == 0 {
		panic("Peek from an empty queue")
	}

	return q.elements[0]
}

// IsEmpty returns true if the queue is empty, false otherwise.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

// Size returns the number of elements in the queue.
func (q *Queue[T]) Size() int {
	return len(q.elements)
}
