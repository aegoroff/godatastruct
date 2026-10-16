package countingsort

import "github.com/aegoroff/godatastruct/types"

// Stable sorts slice using stable counting sort algorithm,
// but it allocates more memory and works slower
func Stable[T types.Integer](items []T, max T) {
	sorted := GetSorted(items, max)
	copy(items, sorted)
}

// Sort sorts slice using counting sort algorithm that is
// less stable but in the most cases much faster due to less memory allocations
func Sort[T types.Integer](items []T, max T) {
	c := counters(max)

	for _, item := range items {
		c[item]++
	}

	b := 0
	for i, n := range c {
		for range n {
			items[b] = T(i)
			b++
		}
	}
}

// GetSorted returns sorted slice
func GetSorted[T types.Integer](a []T, max T) []T {
	b := make([]T, len(a))
	c := counters(max)

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

// counters creates counters slice for values in [0, max].
// Counters are int because they are bounded by items count, not by T range,
// so T would overflow on them (e.g. 300 items of uint8)
func counters[T types.Integer](max T) []int {
	return make([]int, uint64(max)+1)
}
