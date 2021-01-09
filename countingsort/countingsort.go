package countingsort

// IntsStable sorts integers slice using stable counting sort algorithm
// but it allocates more memory and works slower
func IntsStable(items []int, max int) {
	sorted := GetSortedInts(items, max)
	copy(items, sorted)
}

// Ints sorts integers slice using counting sort algorithm that is
// less stable but in the most cases much faster due to less memory allocations
func Ints(items []int, max int) {
	c := make([]int, max+1)

	for _, item := range items {
		c[item]++
	}

	b := 0
	for i := 0; i < max+1; i++ {
		for j := 0; j < c[i]; j++ {
			items[b] = i
			b++
		}
	}
}

// Ints64Stable sorts int64 slice using stable counting sort algorithm
// but it allocates more memory and works slower
func Ints64Stable(items []int64, max int64) {
	sorted := GetSortedInts64(items, max)
	copy(items, sorted)
}

// Ints64 sorts integers slice using counting sort algorithm that is
// less stable but in the most cases much faster due to less memory allocations
func Ints64(items []int64, max int64) {
	c := make([]int64, max+1)

	for _, item := range items {
		c[item]++
	}

	b := 0
	for i := int64(0); i < max+1; i++ {
		for j := int64(0); j < c[i]; j++ {
			items[b] = i
			b++
		}
	}
}

// GetSortedInts returns sorted integers slice
func GetSortedInts(a []int, max int) []int {
	b := make([]int, len(a))
	c := make([]int, max+1)

	for _, item := range a {
		c[item]++
	}

	for i := 1; i < len(c); i++ {
		c[i] += c[i-1]
	}

	for i := len(a) - 1; i >= 0; i-- {
		c[a[i]]--
		b[c[a[i]]] = a[i]
	}

	return b
}

// GetSortedInts64 returns sorted int64 slice
func GetSortedInts64(a []int64, max int64) []int64 {
	b := make([]int64, len(a))
	c := make([]int64, max+1)

	for _, item := range a {
		c[item]++
	}

	for i := 1; i < len(c); i++ {
		c[i] += c[i-1]
	}

	for i := len(a) - 1; i >= 0; i-- {
		c[a[i]]--
		b[c[a[i]]] = a[i]
	}

	return b
}

// Int64Slice attaches the methods of Interface to []int64, sorting in increasing order.
type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
