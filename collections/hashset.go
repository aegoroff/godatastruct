package collections

// StringHashSet defines strings hash set
type StringHashSet map[string]struct{}

// IntHashSet defines integers hash set
type IntHashSet map[int]struct{}

// Int64HashSet defines large integers hash set
type Int64HashSet map[int64]struct{}

// NewIntHashSet creates new IntHashSet
func NewIntHashSet() IntHashSet {
	return make(IntHashSet)
}

// NewInt64HashSet creates new Int64HashSet
func NewInt64HashSet() Int64HashSet {
	return make(Int64HashSet)
}

// NewStringHashSet creates new StringHashSet
func NewStringHashSet() StringHashSet {
	return make(StringHashSet)
}

// Count gets the number of items in the set
func (s *StringHashSet) Count() int {
	return len(*s)
}

// Items gets all set's items
func (s *StringHashSet) Items() []string {
	return s.SortedItems(func(s []string) {})
}

// SortedItems gets all set's items sorted using function specified
func (s *StringHashSet) SortedItems(sorter func(s []string)) []string {
	keys := make([]string, len(*s))
	i := 0
	for k := range *s {
		keys[i] = k
		i++
	}
	sorter(keys)
	return keys
}

// ItemsDecorated gets all set's items applying decorator function to each item
func (s *StringHashSet) ItemsDecorated(decorator func(s string) string) []string {
	keys := make([]string, len(*s))
	i := 0
	for k := range *s {
		keys[i] = decorator(k)
		i++
	}
	return keys
}

// Contains gets whether a key is presented within the set
func (s *StringHashSet) Contains(key string) bool {
	_, ok := (*s)[key]
	return ok
}

// Add adds new item into the set
func (s *StringHashSet) Add(key string) {
	(*s)[key] = struct{}{}
}

// AddRange adds several new items into the set
func (s *StringHashSet) AddRange(keys ...string) {
	for _, key := range keys {
		s.Add(key)
	}
}

// Remove removes item from the set
// If there is no such element, Remove is a no-op.
func (s *StringHashSet) Remove(key string) {
	delete(*s, key)
}

// Count gets the number of items in the set
func (s *IntHashSet) Count() int {
	return len(*s)
}

// Items gets all set's items
func (s *IntHashSet) Items() []int {
	return s.SortedItems(func(s []int) {})
}

// SortedItems gets all set's items sorted using function specified
func (s *IntHashSet) SortedItems(sorter func(s []int)) []int {
	keys := make([]int, len(*s))
	i := 0
	for k := range *s {
		keys[i] = k
		i++
	}
	sorter(keys)
	return keys
}

// Contains gets whether a key is presented within the set
func (s *IntHashSet) Contains(key int) bool {
	_, ok := (*s)[key]
	return ok
}

// Add adds new item into the set
func (s *IntHashSet) Add(key int) {
	(*s)[key] = struct{}{}
}

// AddRange adds several new items into the set
func (s *IntHashSet) AddRange(keys ...int) {
	for _, key := range keys {
		s.Add(key)
	}
}

// Remove removes item from the set
// If there is no such element, Remove is a no-op.
func (s *IntHashSet) Remove(key int) {
	delete(*s, key)
}

// Count gets the number of items in the set
func (s *Int64HashSet) Count() int {
	return len(*s)
}

// Items gets all set's items
func (s *Int64HashSet) Items() []int64 {
	return s.SortedItems(func(s []int64) {})
}

// SortedItems gets all set's items sorted using function specified
func (s *Int64HashSet) SortedItems(sorter func(s []int64)) []int64 {
	keys := make([]int64, len(*s))
	i := 0
	for k := range *s {
		keys[i] = k
		i++
	}
	sorter(keys)
	return keys
}

// Contains gets whether a key is presented within the set
func (s *Int64HashSet) Contains(key int64) bool {
	_, ok := (*s)[key]
	return ok
}

// Add adds new item into the set
func (s *Int64HashSet) Add(key int64) {
	(*s)[key] = struct{}{}
}

// AddRange adds several new items into the set
func (s *Int64HashSet) AddRange(keys ...int64) {
	for _, key := range keys {
		s.Add(key)
	}
}

// Remove removes item from the set
// If there is no such element, Remove is a no-op.
func (s *Int64HashSet) Remove(key int64) {
	delete(*s, key)
}
