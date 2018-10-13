package countingsort

import (
	"github.com/stretchr/testify/assert"
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
		items := []int{2, 5, 3, 0, 2, 3, 0, 3, 4, 1}
		Ints(items, 5)
	}
	b.ReportAllocs()
}

func BenchmarkQuickSortInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		items := []int{2, 5, 3, 0, 2, 3, 0, 3, 4, 1}
		sort.Ints(items)
	}
	b.ReportAllocs()
}
