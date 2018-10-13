package countingsort

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

func TestInts(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	items := []int{2, 5, 3, 0, 2, 3, 0, 3, 4, 1}

	// Act
	Ints(items, 5)

	// Assert
	ass.Equal([]int{0, 0, 1, 2, 2, 3, 3, 3, 4, 5}, items)
}

func TestInts64(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	items := []int64{2, 5, 3, 0, 2, 3, 0, 3, 4, 1}

	// Act
	Ints64(items, 5)

	// Assert
	ass.Equal([]int64{0, 0, 1, 2, 2, 3, 3, 3, 4, 5}, items)
}

func BenchmarkInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		items := perm(1024)
		b.StartTimer()
		Ints(items, 1024)
	}
	b.ReportAllocs()
}

func BenchmarkQuickSortInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		items := perm(1024)
		b.StartTimer()
		sort.Ints(items)
	}
	b.ReportAllocs()
}

// perm returns a random permutation of n Int items in the range [0, n).
func perm(n int) (out []int) {
	for _, v := range rand.Perm(n) {
		out = append(out, v)
	}
	return
}
