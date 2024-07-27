package collections

// HashSet is a generic type that holds a set of unique items.
type HashSet[T comparable] struct {
    items map[T]struct{}
}

// NewHashSet initializes a new empty HashSet.
func NewHashSet[T comparable]() *HashSet[T] {
    return &HashSet[T]{
        items: make(map[T]struct{}),
    }
}

// Add adds an item to the HashSet.
// Returns true if the item was added, false if it was already present.
func (s *HashSet[T]) Add(item T) bool {
    if _, exists := s.items[item]; exists {
        return false
    }
    s.items[item] = struct{}{}
    return true
}

// Remove removes an item from the HashSet.
// Returns true if the item was removed, false if it was not present.
func (s *HashSet[T]) Remove(item T) bool {
    if _, exists := s.items[item]; !exists {
        return false
    }
    delete(s.items, item)
    return true
}

// Contains checks if an item is present in the HashSet.
// Returns true if the item is present, false otherwise.
func (s *HashSet[T]) Contains(item T) bool {
    _, exists := s.items[item]
    return exists
}

// Count returns the number of items in the HashSet.
func (s *HashSet[T]) Count() int {
    return len(s.items)
}

// Clear removes all items from the HashSet.
func (s *HashSet[T]) Clear() {
    s.items = make(map[T]struct{})
}

// Items returns a slice of all items in the HashSet.
func (s *HashSet[T]) Items() []T {
    keys := make([]T, 0, len(s.items))
    for k := range s.items {
        keys = append(keys, k)
    }
    return keys
}
