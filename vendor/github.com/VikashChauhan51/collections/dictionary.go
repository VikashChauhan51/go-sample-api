package collections

// Dictionary is a generic type that holds a map of items with keys and values of any comparable types.
type Dictionary[K comparable, V any] struct {
    items map[K]V
}

// NewDictionary initializes a new empty Dictionary.
func NewDictionary[K comparable, V any]() *Dictionary[K, V] {
    return &Dictionary[K, V]{
        items: make(map[K]V),
    }
}

// Set adds or updates a key-value pair in the Dictionary.
func (d *Dictionary[K, V]) Set(key K, value V) {
    d.items[key] = value
}

// Get retrieves the value for a given key from the Dictionary.
// Returns the value and a boolean indicating if the key was found.
func (d *Dictionary[K, V]) Get(key K) (V, bool) {
    value, ok := d.items[key]
    return value, ok
}

// Remove removes a key-value pair from the Dictionary by key.
// Returns true if the item was removed, false otherwise.
func (d *Dictionary[K, V]) Remove(key K) bool {
    _, ok := d.items[key]
    if ok {
        delete(d.items, key)
    }
    return ok
}

// Keys returns a slice of all keys in the Dictionary.
func (d *Dictionary[K, V]) Keys() []K {
    keys := make([]K, 0, len(d.items))
    for k := range d.items {
        keys = append(keys, k)
    }
    return keys
}

// Values returns a slice of all values in the Dictionary.
func (d *Dictionary[K, V]) Values() []V {
    values := make([]V, 0, len(d.items))
    for _, v := range d.items {
        values = append(values, v)
    }
    return values
}

// Count returns the number of key-value pairs in the Dictionary.
func (d *Dictionary[K, V]) Count() int {
    return len(d.items)
}
