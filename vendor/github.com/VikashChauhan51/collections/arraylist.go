package collections

import "sort"

// ArrayList is a generic type that holds a collection of items of any type.
type ArrayList[T any] struct {
	collection []T
}

// NewArrayList initializes a new empty ArrayList.
//
// Example:
//  list := NewArrayList[int]()
//  fmt.Println(list) // Output: &{[]}
func NewArrayList[T any]() *ArrayList[T] {
	return &ArrayList[T]{
		collection: []T{},
	}
}

// NewArrayListT initializes a new ArrayList with the given items.
//
// Example:
//  list := NewArrayListT(1, 2, 3, 4, 5, 6)
//  fmt.Println(list) // Output: &{[1 2 3 4 5 6]}
func NewArrayListT[T any](items ...T) *ArrayList[T] {
	l := ArrayList[T]{
		collection: make([]T, len(items)),
	}

	// Copy the items into the collection
	copy(l.collection, items)

	return &l
}

// Count returns the number of items in the ArrayList.
//
// Example:
//  list := NewArrayListT(1, 2, 3)
//  fmt.Println(list.Count()) // Output: 3
func (l *ArrayList[T]) Count() int {
	return len(l.collection)
}

// Clear removes all items from the ArrayList.
//
// Example:
//  list := NewArrayListT(1, 2, 3)
//  list.Clear()
//  fmt.Println(list) // Output: &{[]}
func (l *ArrayList[T]) Clear() {
	l.collection = []T{}
}

// Items returns a slice of all items in the ArrayList.
//
// Example:
//  list := NewArrayListT(1, 2, 3)
//  items := list.Items()
//  fmt.Println(items) // Output: [1 2 3]
func (l *ArrayList[T]) Items() []T {
	return l.collection
}

// Add appends an item to the ArrayList.
//
// Example:
//  list := NewArrayList[int]()
//  list.Add(1)
//  fmt.Println(list) // Output: &{[1]}
func (l *ArrayList[T]) Add(item T) {
	l.collection = append(l.collection, item)
}

// AddRange appends multiple items to the ArrayList.
//
// Example:
//  list := NewArrayList[int]()
//  list.AddRange([]int{1, 2, 3})
//  fmt.Println(list) // Output: &{[1 2 3]}
func (l *ArrayList[T]) AddRange(items []T) {
	l.collection = append(l.collection, items...)
}

// Get retrieves the item at the specified index.
//
// Example:
//  list := NewArrayListT(1, 2, 3)
//  item := list.Get(1)
//  fmt.Println(item) // Output: 2
func (l *ArrayList[T]) Get(index int) T {
	if index < 0 || index >= l.Count() {
		panic("Index out of range.")
	}
	return l.collection[index]
}

// GetIndex retrieves the index of the specified item in the ArrayList.
// Returns -1 if the item is not found.
//
// Example:
//  list := NewArrayListT(1, 2, 3)
//  index := list.GetIndex(2, func(x int) bool { return x == 2 })
//  fmt.Println(index) // Output: 1
func (l *ArrayList[T]) GetIndex(item T, equal func(T, T) bool) int {
	for i, v := range l.collection {
		if equal(v, item) {
			return i
		}
	}
	return -1
}

// Remove removes the first occurrence of the specified item from the ArrayList.
//
// Example:
//  list := NewArrayListT(1, 2, 3)
//  removed := list.Remove(2, func(x, y int) bool { return x == y })
//  fmt.Println(removed) // Output: true
//  fmt.Println(list)    // Output: &{[1 3]}
func (l *ArrayList[T]) Remove(item T, equal func(T, T) bool) bool {
	for i, v := range l.collection {
		if equal(v, item) {
			l.collection = append(l.collection[:i], l.collection[i+1:]...)
			return true
		}
	}
	return false
}

// RemoveAt removes the item at the specified index.
//
// Example:
//  list := NewArrayListT(1, 2, 3)
//  list.RemoveAt(1)
//  fmt.Println(list) // Output: &{[1 3]}
func (l *ArrayList[T]) RemoveAt(index int) {
	if index < 0 || index >= l.Count() {
		panic("Index out of range.")
	}

	l.collection = append(l.collection[:index], l.collection[index+1:]...)
}

// Set updates the item at the specified index.
//
// Example:
//  list := NewArrayListT(1, 2, 3)
//  list.Set(1, 10)
//  fmt.Println(list) // Output: &{[1 10 3]}
func (l *ArrayList[T]) Set(index int, item T) {
	if index < 0 || index >= l.Count() {
		panic("Index out of range.")
	}
	l.collection[index] = item
}

// OrderBy sorts the ArrayList using the provided less function
func (l *ArrayList[T]) OrderBy(less func(i, j T) bool) {
	sort.Slice(l.collection, func(i, j int) bool {
		return less(l.collection[i], l.collection[j])
	})
}

// Filter returns a new slice containing all items that match the predicate.
//
// Example:
//  list := NewArrayListT(1, 2, 3, 4, 5)
//  evenNumbers := list.Filter(func(item int) bool { return item%2 == 0 })
//  fmt.Println(evenNumbers) // Output: [2 4]
func (l *ArrayList[T]) Filter(predicate func(T) bool) []T {
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
//  list := NewArrayListT(1, 2, 3, 4, 5)
//  firstEven := list.First(func(item int) bool { return item%2 == 0 })
//  fmt.Println(firstEven) // Output: 2
func (l *ArrayList[T]) First(predicate func(T) bool) T {
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
//  list := NewArrayListT(1, 2, 3, 4, 5)
//  firstEven := list.FirstOrDefault(func(item int) bool { return item%2 == 0 })
//  fmt.Println(firstEven) // Output: 2
func (l *ArrayList[T]) FirstOrDefault(predicate func(T) bool) T {
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
//  list := NewArrayListT(1, 2, 3, 4, 5)
//  lastEven := list.LastOrDefault(func(item int) bool { return item%2 == 0 })
//  fmt.Println(lastEven) // Output: 4
func (l *ArrayList[T]) LastOrDefault(predicate func(T) bool) T {
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
//  list := NewArrayListT(1, 2, 3, 4, 5)
//  singleEven := list.SingleOrDefault(func(item int) bool { return item == 2 })
//  fmt.Println(singleEven) // Output: 2
func (l *ArrayList[T]) SingleOrDefault(predicate func(T) bool) T {
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
//  list := NewArrayListT(1, 2, 3, 4, 5)
//  singleEven := list.Single(func(item int) bool { return item == 2 })
//  fmt.Println(singleEven) // Output: 2
func (l *ArrayList[T]) Single(predicate func(T) bool) T {
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

// Iterator represents an iterator for the ArrayList.
type ArrayListIterator[T any] struct {
	list  *ArrayList[T]
	index int
}

// NewIterator creates a new iterator for the List.
func (l *ArrayList[T]) NewIterator() *ArrayListIterator[T] {
	return &ArrayListIterator[T]{list: l, index: 0}
}

// HasNext checks if there are more elements to iterate over.
func (it *ArrayListIterator[T]) HasNext() bool {
	return it.index < it.list.Count()
}

// Next returns the next element in the iteration.
func (it *ArrayListIterator[T]) Next() (T, bool) {
	if it.HasNext() {
		item := it.list.Get(it.index)
		it.index++
		return item, true
	}
	var zeroValue T
	return zeroValue, false
}
