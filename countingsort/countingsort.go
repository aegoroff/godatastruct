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
	c := make([]T, max+1)

	for _, item := range items {
		c[item]++
	}

	var b T
	b = 0
	var i T
	var j T
	for i = 0; i < max+1; i++ {
		for j = 0; j < c[i]; j++ {
			items[b] = i
			b++
		}
	}
}

// GetSorted returns sorted slice
func GetSorted[T types.Integer](a []T, max T) []T {
	b := make([]T, len(a))
	c := make([]T, max+1)

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
