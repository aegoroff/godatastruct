package countingsort

import (
	"fmt"
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

func TestIntsUnstable(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	items := []int{2, 5, 3, 0, 2, 3, 0, 3, 4, 1}

	// Act
	IntsUnstable(items, 5)

	// Assert
	ass.Equal([]int{0, 0, 1, 2, 2, 3, 3, 3, 4, 5}, items)
}

func TestIntsUnstable_Random(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	items := perm(1024)

	// Act
	IntsUnstable(items, 1024)

	// Assert
	ass.True(sort.IsSorted(sort.IntSlice(items)))
}

func ExampleInts() {
	items := []int{2, 5, 3, 0, 0}

	Ints(items, 5)
	fmt.Println(items)
	// Output:
	// [0 0 2 3 5]
}

func ExampleIntsUnstable() {
	items := []int{2, 5, 3, 0, 0}

	IntsUnstable(items, 5)
	fmt.Println(items)
	// Output:
	// [0 0 2 3 5]
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

func TestInts64Unstable(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	items := []int64{2, 5, 3, 0, 2, 3, 0, 3, 4, 1}

	// Act
	Ints64Unstable(items, 5)

	// Assert
	ass.Equal([]int64{0, 0, 1, 2, 2, 3, 3, 3, 4, 5}, items)
}

func TestInts64Slice_Sort(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	items := []int64{2, 5, 3, 0, 2, 3, 0, 3, 4, 1}

	// Act
	sort.Sort(Int64Slice(items))

	// Assert
	ass.Equal([]int64{0, 0, 1, 2, 2, 3, 3, 3, 4, 5}, items)
}

func TestInts64Unstable_Random(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	items := perm64(1024)

	// Act
	Ints64Unstable(items, 1024)

	// Assert
	ass.True(sort.IsSorted(Int64Slice(items)))
}

func ExampleInts64() {
	items := []int64{2, 5, 3, 0, 0}

	Ints64(items, 5)
	fmt.Println(items)
	// Output:
	// [0 0 2 3 5]
}

func ExampleInts64Unstable() {
	items := []int64{2, 5, 3, 0, 0}

	Ints64Unstable(items, 5)
	fmt.Println(items)
	// Output:
	// [0 0 2 3 5]
}

func BenchmarkInts(b *testing.B) {
	b.SetBytes(2)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		items := perm(1024)
		b.StartTimer()
		Ints(items, 1024)
	}
	b.ReportAllocs()
}

func BenchmarkIntsUnstable(b *testing.B) {
	b.SetBytes(2)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		items := perm(1024)
		b.StartTimer()
		IntsUnstable(items, 1024)
	}
	b.ReportAllocs()
}

func BenchmarkQuickSortInt(b *testing.B) {
	b.SetBytes(2)
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
	out = append(out, rand.Perm(n)...)
	return
}

// perm returns a random permutation of n Int items in the range [0, n).
func perm64(n int) (out []int64) {
	for _, i := range rand.Perm(n) {
		out = append(out, int64(i))
	}
	return
}
