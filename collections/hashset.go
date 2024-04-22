package collections

// HashSet defines a hash set
type HashSet[T comparable] map[T]struct{}

// NewIntHashSet creates new IntHashSet
func NewHashSet[T comparable]() HashSet[T] {
	return make(HashSet[T])
}

// Count gets the number of items in the set
func (s *HashSet[T]) Count() int {
	return len(*s)
}

// Items gets all set's items
func (s *HashSet[T]) Items() []T {
	return s.SortedItems(func(s []T) {})
}

// SortedItems gets all set's items sorted using function specified
func (s *HashSet[T]) SortedItems(sorter func(s []T)) []T {
	keys := make([]T, len(*s))
	i := 0
	for k := range *s {
		keys[i] = k
		i++
	}
	sorter(keys)
	return keys
}

// ItemsDecorated gets all set's items applying decorator function to each item
func (s *HashSet[T]) ItemsDecorated(decorator func(s T) T) []T {
	keys := make([]T, len(*s))
	i := 0
	for k := range *s {
		keys[i] = decorator(k)
		i++
	}
	return keys
}

// Contains gets whether a key is presented within the set
func (s *HashSet[T]) Contains(key T) bool {
	_, ok := (*s)[key]
	return ok
}

// Add adds new item into the set
func (s *HashSet[T]) Add(key T) {
	(*s)[key] = struct{}{}
}

// AddRange adds several new items into the set
func (s *HashSet[T]) AddRange(keys ...T) {
	for _, key := range keys {
		s.Add(key)
	}
}

// Remove removes item from the set
// If there is no such element, Remove is a no-op.
func (s *HashSet[T]) Remove(key T) {
	delete(*s, key)
}
