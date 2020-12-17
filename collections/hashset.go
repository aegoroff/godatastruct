// Package collections contains various containers. Now only int, int64 and string hashsets implemented
package collections

// StringHashSet defines strings hash set
type StringHashSet map[string]struct{}

// IntHashSet defines integers hash set
type IntHashSet map[int]struct{}

// Int64HashSet defines large integers hash set
type Int64HashSet map[int64]struct{}

// Count gets the number of items in the set
func (s *StringHashSet) Count() int {
	return len(*s)
}

// Items gets all set's items
func (s *StringHashSet) Items() []string {
	keys := make([]string, 0, len(*s))
	for k := range *s {
		keys = append(keys, k)
	}
	return keys
}

// ItemsDecorated gets all set's items applying decorator function to each item
func (s *StringHashSet) ItemsDecorated(decorator func(s string) string) []string {
	keys := make([]string, 0, len(*s))
	for k := range *s {
		keys = append(keys, decorator(k))
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

// Count gets the number of items in the set
func (s *IntHashSet) Count() int {
	return len(*s)
}

// Items gets all set's items
func (s *IntHashSet) Items() []int {
	keys := make([]int, 0, len(*s))
	for k := range *s {
		keys = append(keys, k)
	}
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

// Count gets the number of items in the set
func (s *Int64HashSet) Count() int {
	return len(*s)
}

// Items gets all set's items
func (s *Int64HashSet) Items() []int64 {
	keys := make([]int64, 0, len(*s))
	for k := range *s {
		keys = append(keys, k)
	}
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
