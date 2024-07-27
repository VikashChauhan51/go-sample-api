package collections

import "sort"

// List is a generic type that holds a collection of items of any comparable type.
type List[T comparable] struct {
	collection []T
}

// NewList initializes a new empty List.
//
// Example:
//  list := NewList[int]()
//  fmt.Println(list) // Output: &{[]}
func NewList[T comparable]() *List[T] {
	return &List[T]{
		collection: []T{},
	}
}

// NewListT initializes a new List with the given items.
//
// Example:
//  list := NewListT(1, 2, 3, 4, 5, 6)
//  fmt.Println(list) // Output: &{[1 2 3 4 5 6]}
func NewListT[T comparable](items ...T) *List[T] {
	l := List[T]{
		collection: make([]T, len(items)),
	}

	// Copy the items into the collection
	copy(l.collection, items)

	return &l
}

// Count returns the number of items in the List.
//
// Example:
//  list := NewListT(1, 2, 3)
//  fmt.Println(list.Count()) // Output: 3
func (l *List[T]) Count() int {
	return len(l.collection)
}

// Clear removes all items from the List.
//
// Example:
//  list := NewListT(1, 2, 3)
//  list.Clear()
//  fmt.Println(list) // Output: &{[]}
func (l *List[T]) Clear() {
	l.collection = []T{}
}

// Items returns a slice of all items in the List.
//
// Example:
//  list := NewListT(1, 2, 3)
//  items := list.Items()
//  fmt.Println(items) // Output: [1 2 3]
func (l *List[T]) Items() []T {
	return l.collection
}

// Add appends an item to the List.
//
// Example:
//  list := NewList[int]()
//  list.Add(1)
//  fmt.Println(list) // Output: &{[1]}
func (l *List[T]) Add(item T) {
	l.collection = append(l.collection, item)
}

// AddRange appends multiple items to the List.
//
// Example:
//  list := NewList[int]()
//  list.AddRange([]int{1, 2, 3})
//  fmt.Println(list) // Output: &{[1 2 3]}
func (l *List[T]) AddRange(items []T) {
	l.collection = append(l.collection, items...)
}

// Get retrieves the item at the specified index.
//
// Example:
//  list := NewListT(1, 2, 3)
//  item := list.Get(1)
//  fmt.Println(item) // Output: 2
func (l *List[T]) Get(index int) T {
	if index < 0 || index >= l.Count() {
		panic("Index out of range.")
	}
	return l.collection[index]
}

// GetIndex retrieves the index of the specified item in the List.
// Returns -1 if the item is not found.
//
// Example:
//  list := NewListT(1, 2, 3)
//  index := list.GetIndex(2)
//  fmt.Println(index) // Output: 1
func (l *List[T]) GetIndex(item T) int {
	for i, v := range l.collection {
		if v == item {
			return i
		}
	}
	return -1
}

// Remove removes the first occurrence of the specified item from the List.
//
// Example:
//  list := NewListT(1, 2, 3)
//  list.Remove(2)
//  fmt.Println(list) // Output: &{[1 3]}
func (l *List[T]) Remove(item T) bool {
	for i, v := range l.collection {
		if v == item {
			l.collection = append(l.collection[:i], l.collection[i+1:]...)
			return true
		}
	}
	return false
}

// RemoveAt removes the item at the specified index.
//
// Example:
//  list := NewListT(1, 2, 3)
//  list.RemoveAt(1)
//  fmt.Println(list) // Output: &{[1 3]}
func (l *List[T]) RemoveAt(index int) {
	if index < 0 || index >= l.Count() {
		panic("Index out of range.")
	}

	l.collection = append(l.collection[:index], l.collection[index+1:]...)
}

// OrderBy sorts the list using the provided less function
func (l *List[T]) OrderBy(less func(i, j T) bool) {
	sort.Slice(l.collection, func(i, j int) bool {
		return less(l.collection[i], l.collection[j])
	})
}

// Filter returns a new slice containing all items that match the predicate.
//
// Example:
//  list := NewListT(1, 2, 3, 4, 5)
//  evenNumbers := list.Filter(func(item int) bool { return item%2 == 0 })
//  fmt.Println(evenNumbers) // Output: [2 4]
func (l *List[T]) Filter(predicate func(T) bool) []T {
	results := make([]T, 0)
	for _, val := range l.collection {
		if predicate(val) {
			results = append(results, val)
		}
	}

	return results
}

// First returns the first item that matches the predicate.
// If no item matches, it will panic.
//
// Example:
//  list := NewListT(1, 2, 3, 4, 5)
//  firstEven := list.First(func(item int) bool { return item%2 == 0 })
//  fmt.Println(firstEven) // Output: 2
func (l *List[T]) First(predicate func(T) bool) T {
	for _, val := range l.collection {
		if predicate(val) {
			return val
		}
	}
	panic("Sequence contains no element.")
}

// FirstOrDefault returns the first item that matches the predicate.
// If no item matches, it returns the zero value of the type.
//
// Example:
//  list := NewListT(1, 2, 3, 4, 5)
//  firstEven := list.FirstOrDefault(func(item int) bool { return item%2 == 0 })
//  fmt.Println(firstEven) // Output: 2
func (l *List[T]) FirstOrDefault(predicate func(T) bool) T {
	for _, val := range l.collection {
		if predicate(val) {
			return val
		}
	}

	var zeroValue T
	return zeroValue
}

// LastOrDefault returns the last item that matches the predicate.
// If no item matches, it returns the zero value of the type.
//
// Example:
//  list := NewListT(1, 2, 3, 4, 5)
//  lastEven := list.LastOrDefault(func(item int) bool { return item%2 == 0 })
//  fmt.Println(lastEven) // Output: 4
func (l *List[T]) LastOrDefault(predicate func(T) bool) T {
	var zeroValue T
	for _, val := range l.collection {
		if predicate(val) {
			zeroValue = val
		}
	}

	return zeroValue
}

// SingleOrDefault returns the single item that matches the predicate.
// If no item or more than one item matches, it panics.
//
// Example:
//  list := NewListT(1, 2, 3, 4, 5)
//  singleEven := list.SingleOrDefault(func(item int) bool { return item == 2 })
//  fmt.Println(singleEven) // Output: 2
func (l *List[T]) SingleOrDefault(predicate func(T) bool) T {
	var result T
	count := 0
	for _, val := range l.collection {
		if predicate(val) {
			if count == 0 {
				result = val
				count++
			} else {
				panic("Sequence contains more than one matching element.")
			}
		}
	}

	return result
}

// Single returns the single item that matches the predicate.
// If no item or more than one item matches, it panics.
//
// Example:
//  list := NewListT(1, 2, 3, 4, 5)
//  singleEven := list.Single(func(item int) bool { return item == 2 })
//  fmt.Println(singleEven) // Output: 2
func (l *List[T]) Single(predicate func(T) bool) T {
	var result T
	count := 0
	for _, val := range l.collection {
		if predicate(val) {
			if count == 0 {
				result = val
				count++
			} else {
				panic("Sequence contains more than one matching element.")
			}
		}
	}
	if count == 0 {
		panic("Sequence contains no element.")
	}
	return result
}

// Set updates the item at the specified index.
//
// Example:
//  list := NewListT(1, 2, 3)
//  list.Set(1, 10)
//  fmt.Println(list) // Output: &{[1 10 3]}
func (l *List[T]) Set(index int, item T) {
	if index < 0 || index >= l.Count() {
		panic("Index out of range.")
	}
	l.collection[index] = item
}

// Iterator represents an iterator for the List.
type ListIterator[T comparable] struct {
	list  *List[T]
	index int
}

// NewIterator creates a new iterator for the List.
func (l *List[T]) NewIterator() *ListIterator[T] {
	return &ListIterator[T]{list: l, index: 0}
}

// HasNext checks if there are more elements to iterate over.
func (it *ListIterator[T]) HasNext() bool {
	return it.index < it.list.Count()
}

// Next returns the next element in the iteration.
func (it *ListIterator[T]) Next() (T, bool) {
	if it.HasNext() {
		item := it.list.Get(it.index)
		it.index++
		return item, true
	}
	var zeroValue T
	return zeroValue, false
}
