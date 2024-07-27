
# Collections Package

The `collections` package provides generic data structures in Go, including a `List`, a `Dictionary`, and a `HashSet`. These types offer various functionalities for managing collections of items in a type-safe and efficient manner.

## Features

- **Generic List:** Holds a collection of items of any comparable type with methods for adding, removing, and accessing items.
- **Generic Dictionary:** A map-like data structure that associates keys with values, with methods for adding, retrieving, and removing key-value pairs.
- **Generic HashSet:** A set-like data structure that holds unique items and provides methods for adding, removing, and checking membership.
- **Concurrent List:** A thread-safe version of the List for concurrent use.
- **Concurrent Dictionary:** A thread-safe version of the Dictionary for concurrent use.

## Installation

To install the `collections` package, include it in your Go module:

```bash
go get github.com/VikashChauhan51/collections
```

## Usage

### List

#### Creating a New List

```go
package main

import (
    "fmt"
    "github.com/VikashChauhan51/collections"
)

func main() {
    list := collections.NewList[int]()
    fmt.Println(list) // Output: &{[]}

    list.Add(1)
    list.AddRange([]int{2, 3})
    fmt.Println(list) // Output: &{[1 2 3]}

    // Access items using Indexer
    fmt.Println(list.Index(1)) // Output: 2

    // Remove an item
    list.Remove(2)
    fmt.Println(list) // Output: &{[1 3]}
}
```

#### Functions and Methods

- **`NewList[T comparable]() *List[T]`**: Initializes a new empty List.
  ```go
  list := collections.NewList[int]()
  fmt.Println(list) // Output: &{[]}
  ```

- **`NewListT[T comparable](items ...T) *List[T]`**: Initializes a new List with the given items.
  ```go
  list := collections.NewListT(1, 2, 3, 4, 5, 6)
  fmt.Println(list) // Output: &{[1 2 3 4 5 6]}
  ```

- **`Count() int`**: Returns the number of items in the List.
  ```go
  list := collections.NewListT(1, 2, 3)
  fmt.Println(list.Count()) // Output: 3
  ```

- **`Clear()`**: Removes all items from the List.
  ```go
  list := collections.NewListT(1, 2, 3)
  list.Clear()
  fmt.Println(list) // Output: &{[]}
  ```

- **`Items() []T`**: Returns a slice of all items in the List.
  ```go
  list := collections.NewListT(1, 2, 3)
  items := list.Items()
  fmt.Println(items) // Output: [1 2 3]
  ```

- **`Add(item T)`**: Appends an item to the List.
  ```go
  list := collections.NewList[int]()
  list.Add(1)
  fmt.Println(list) // Output: &{[1]}
  ```

- **`AddRange(items []T)`**: Appends multiple items to the List.
  ```go
  list := collections.NewList[int]()
  list.AddRange([]int{1, 2, 3})
  fmt.Println(list) // Output: &{[1 2 3]}
  ```

- **`Get(index int) T`**: Retrieves the item at the specified index.
  ```go
  list := collections.NewListT(1, 2, 3)
  item := list.Get(1)
  fmt.Println(item) // Output: 2
  ```

- **`Set(index int, item T)`**: Updates the item at the specified index.
  ```go
  list := collections.NewListT(1, 2, 3)
  list.Set(1, 10)
  fmt.Println(list) // Output: &{[1 10 3]}
  ```

- **`GetIndex(item T) int`**: Retrieves the index of the specified item in the List. Returns -1 if the item is not found.
  ```go
  list := collections.NewListT(1, 2, 3)
  index := list.GetIndex(2)
  fmt.Println(index) // Output: 1
  ```

- **`Remove(item T) bool`**: Removes the first occurrence of the specified item from the List.
  ```go
  list := collections.NewListT(1, 2, 3)
  list.Remove(2)
  fmt.Println(list) // Output: &{[1 3]}
  ```

- **`RemoveAt(index int)`**: Removes the item at the specified index.
  ```go
  list := collections.NewListT(1, 2, 3)
  list.RemoveAt(1)
  fmt.Println(list) // Output: &{[1 3]}
  ```

- **`Filter(predicate func(T) bool) []T`**: Returns a new slice containing all items that match the predicate.
  ```go
  list := collections.NewListT(1, 2, 3, 4, 5)
  evenNumbers := list.Filter(func(item int) bool { return item%2 == 0 })
  fmt.Println(evenNumbers) // Output: [2 4]
  ```

- **`FirstOrDefault(predicate func(T) bool) T`**: Returns the first item that matches the predicate. If no item matches, it returns the zero value of the type.
  ```go
  list := collections.NewListT(1, 2, 3, 4, 5)
  firstEven := list.FirstOrDefault(func(item int) bool { return item%2 == 0 })
  fmt.Println(firstEven) // Output: 2
  ```

- **`LastOrDefault(predicate func(T) bool) T`**: Returns the last item that matches the predicate. If no item matches, it returns the zero value of the type.
  ```go
  list := collections.NewListT(1, 2, 3, 4, 5)
  lastEven := list.LastOrDefault(func(item int) bool { return item%2 == 0 })
  fmt.Println(lastEven) // Output: 4
  ```

- **`SingleOrDefault(predicate func(T) bool) T`**: Returns the single item that matches the predicate. If no item or more than one item matches, it panics.
  ```go
  list := collections.NewListT(1, 2, 3, 4, 5)
  singleEven := list.SingleOrDefault(func(item int) bool { return item == 2 })
  fmt.Println(singleEven) // Output: 2
  ```

- **`Index(index int) T`**: Provides indexing-like access to items in the List. This method allows you to use syntax similar to `list[index]`.
  ```go
  list := collections.NewListT(1, 2, 3)
  item := list.Index(1)
  fmt.Println(item) // Output: 2
  ```

### Dictionary

#### Creating a New Dictionary

```go
package main

import (
    "fmt"
    "github.com/VikashChauhan51/collections"
)

func main() {
    dict := collections.NewDictionary[string, int]()
    dict.Add("one", 1)
    dict.Add("two", 2)
    fmt.Println(dict.Get("two")) // Output: 2

    // Remove a key-value pair
    dict.Remove("one")
    fmt.Println(dict.ContainsKey("one")) // Output: false
}
```

#### Functions and Methods

- **`NewDictionary[K comparable, V any]() *Dictionary[K, V]`**: Initializes a new empty Dictionary.
  ```go
  dict := collections.NewDictionary[string, int]()
  ```

- **`Add(key K, value V)`**: Adds a key-value pair to the Dictionary.
  ```go
  dict.Add("key1", "value1")
  ```

- **`Get(key K) (V, bool)`**: Retrieves the value associated with the key. Returns the value and a boolean indicating whether the key was found.
  ```go
  value, found := dict.Get("key1")
  ```

- **`Remove(key K)`**: Removes the key-value pair associated with the key.
  ```go
  dict.Remove("key1")
  ```

- **`ContainsKey(key K) bool`**: Checks if the Dictionary contains the specified key.
  ```go
  exists := dict.ContainsKey("key1")
  ```

- **`Count() int`**: Returns the number of key-value pairs in the Dictionary.
  ```go
  count := dict.Count()
  ```

- **`Clear()`**: Removes all key-value pairs from the Dictionary.
  ```go
  dict.Clear()
  ```

- **`Keys() []K`**: Returns a slice of all keys in the Dictionary.
  ```go
  keys := dict.Keys()
  ```

- **`Values() []V`**: Returns a slice of all values in the Dictionary.
  ```go
  values := dict.Values()
  ```

### HashSet

#### Creating a New HashSet

```go
package main

import (
    "fmt"
   

 "github.com/VikashChauhan51/collections"
)

func main() {
    set := collections.NewHashSet[int]()
    set.Add(1)
    set.Add(2)
    set.Add(3)
    fmt.Println(set.Contains(2)) // Output: true

    // Remove an item
    set.Remove(2)
    fmt.Println(set.Contains(2)) // Output: false

    // Print all items
    fmt.Println(set.Items()) // Output: [1 3]
}
```

#### Functions and Methods

- **`NewHashSet[T comparable]() *HashSet[T]`**: Initializes a new empty HashSet.
  ```go
  set := collections.NewHashSet[int]()
  ```

- **`Add(item T) bool`**: Adds an item to the HashSet. Returns true if the item was added, false if it was already present.
  ```go
  added := set.Add(1)
  ```

- **`Remove(item T) bool`**: Removes an item from the HashSet. Returns true if the item was removed, false if it was not present.
  ```go
  removed := set.Remove(1)
  ```

- **`Contains(item T) bool`**: Checks if the HashSet contains the specified item.
  ```go
  contains := set.Contains(1)
  ```

- **`Count() int`**: Returns the number of items in the HashSet.
  ```go
  count := set.Count()
  ```

- **`Clear()`**: Removes all items from the HashSet.
  ```go
  set.Clear()
  ```

- **`Items() []T`**: Returns a slice of all items in the HashSet.
  ```go
  items := set.Items()
  ```

## Contributing

If you would like to contribute to this package, please fork the repository and submit a pull request. Ensure that your code passes all tests and follows the project's coding style.

## License

This package is licensed under the [MIT License](LICENSE).

---

Feel free to adjust the paths and examples according to your actual module and usage scenarios.
