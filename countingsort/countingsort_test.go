package countingsort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

func TestInts(t *testing.T) {
	var tests = []struct {
		name     string
		algo     func([]int, int)
		input    []int
		expected []int
	}{
		{"normal", Ints, []int{2, 5, 3, 0, 2, 3, 0, 3, 4, 1}, []int{0, 0, 1, 2, 2, 3, 3, 3, 4, 5}},
		{"normal stable", IntsStable, []int{2, 5, 3, 0, 2, 3, 0, 3, 4, 1}, []int{0, 0, 1, 2, 2, 3, 3, 3, 4, 5}},
		{"one", Ints, []int{2}, []int{2}},
		{"one stable", IntsStable, []int{2}, []int{2}},
		{"zero", Ints, []int{}, []int{}},
		{"zero stable", IntsStable, []int{}, []int{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			ass := assert.New(t)

			// Act
			test.algo(test.input, 5)

			// Assert
			ass.Equal(test.expected, test.input)
		})
	}
}

func TestInts64(t *testing.T) {
	var tests = []struct {
		name     string
		algo     func([]int64, int64)
		input    []int64
		expected []int64
	}{
		{"normal", Ints64, []int64{2, 5, 3, 0, 2, 3, 0, 3, 4, 1}, []int64{0, 0, 1, 2, 2, 3, 3, 3, 4, 5}},
		{"normal stable", Ints64Stable, []int64{2, 5, 3, 0, 2, 3, 0, 3, 4, 1}, []int64{0, 0, 1, 2, 2, 3, 3, 3, 4, 5}},
		{"one", Ints64, []int64{2}, []int64{2}},
		{"one stable", Ints64Stable, []int64{2}, []int64{2}},
		{"zero", Ints64, []int64{}, []int64{}},
		{"zero stable", Ints64Stable, []int64{}, []int64{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			ass := assert.New(t)

			// Act
			test.algo(test.input, 5)

			// Assert
			ass.Equal(test.expected, test.input)
		})
	}
}

func TestInts_Random(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	items := perm(1024)

	// Act
	Ints(items, 1024)

	// Assert
	ass.True(sort.IsSorted(sort.IntSlice(items)))
}

func ExampleIntsStable() {
	items := []int{2, 5, 3, 0, 0}

	IntsStable(items, 5)
	fmt.Println(items)
	// Output:
	// [0 0 2 3 5]
}

func ExampleInts() {
	items := []int{2, 5, 3, 0, 0}

	Ints(items, 5)
	fmt.Println(items)
	// Output:
	// [0 0 2 3 5]
}

func TestInts64Stable(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	items := []int64{2, 5, 3, 0, 2, 3, 0, 3, 4, 1}

	// Act
	Ints64Stable(items, 5)

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

func TestInts64_Random(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	items := perm64(1024)

	// Act
	Ints64(items, 1024)

	// Assert
	ass.True(sort.IsSorted(Int64Slice(items)))
}

func ExampleInts64Stable() {
	items := []int64{2, 5, 3, 0, 0}

	Ints64Stable(items, 5)
	fmt.Println(items)
	// Output:
	// [0 0 2 3 5]
}

func ExampleInts64() {
	items := []int64{2, 5, 3, 0, 0}

	Ints64(items, 5)
	fmt.Println(items)
	// Output:
	// [0 0 2 3 5]
}

func BenchmarkIntsStable(b *testing.B) {
	b.SetBytes(2)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		items := perm(1024)
		b.StartTimer()
		IntsStable(items, 1024)
	}
	b.ReportAllocs()
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
