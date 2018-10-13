// Package countingsort contains counting sort methods that sorts int or in64 slice using counting sort algorithm
package countingsort

// Ints sorts integers slice using counting sort algorithm
func Ints(items []int, max int) {
	sorted := countingSort(items, max)

	for i, x := range sorted {
		items[i] = x
	}
}

// Ints64 sorts int64 slice using counting sort algorithm
func Ints64(items []int64, max int64) {
	sorted := countingSortInt64(items, max)

	for i, x := range sorted {
		items[i] = x
	}
}

func countingSort(a []int, max int) []int {
	b := make([]int, len(a))
	c := make([]int, max+1)

	for _, item := range a {
		c[item] += 1
	}

	for i := 1; i < len(c); i++ {
		c[i] += c[i-1]
	}

	for i := len(a) - 1; i >= 0; i-- {
		b[c[a[i]]-1] = a[i]
		c[a[i]] -= 1
	}

	return b
}

func countingSortInt64(a []int64, max int64) []int64 {
	b := make([]int64, len(a))
	c := make([]int64, max+1)

	for _, item := range a {
		c[item] += 1
	}

	for i := 1; i < len(c); i++ {
		c[i] += c[i-1]
	}

	for i := len(a) - 1; i >= 0; i-- {
		b[c[a[i]]-1] = a[i]
		c[a[i]] -= 1
	}

	return b
}
