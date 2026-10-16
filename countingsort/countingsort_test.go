package countingsort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	var tests = []struct {
		name     string
		input    []int
		expected []int
	}{
		{"normal", []int{2, 5, 3, 0, 2, 3, 0, 3, 4, 1}, []int{0, 0, 1, 2, 2, 3, 3, 3, 4, 5}},
		{"one", []int{2}, []int{2}},
		{"zero", []int{}, []int{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			ass := assert.New(t)

			// Act
			Sort(test.input, 5)

			// Assert
			ass.Equal(test.expected, test.input)
		})
	}
}

func TestSortInt64(t *testing.T) {
	var tests = []struct {
		name     string
		input    []int64
		expected []int64
	}{
		{"normal", []int64{2, 5, 3, 0, 2, 3, 0, 3, 4, 1}, []int64{0, 0, 1, 2, 2, 3, 3, 3, 4, 5}},
		{"one", []int64{2}, []int64{2}},
		{"zero", []int64{}, []int64{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			ass := assert.New(t)

			// Act
			Sort(test.input, 5)

			// Assert
			ass.Equal(test.expected, test.input)
		})
	}
}

func TestSortStable(t *testing.T) {
	var tests = []struct {
		name     string
		input    []int
		expected []int
	}{
		{"normal", []int{2, 5, 3, 0, 2, 3, 0, 3, 4, 1}, []int{0, 0, 1, 2, 2, 3, 3, 3, 4, 5}},
		{"one", []int{2}, []int{2}},
		{"zero", []int{}, []int{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			ass := assert.New(t)

			// Act
			Stable(test.input, 5)

			// Assert
			ass.Equal(test.expected, test.input)
		})
	}
}

func TestSortSmallTypes_NoOverflow(t *testing.T) {
	manyUint8 := func() []uint8 {
		items := make([]uint8, 300)
		for i := range items {
			items[i] = uint8(i % 3)
		}
		return items
	}
	var tests = []struct {
		name string
		sort func(t *testing.T)
	}{
		{"Sort: more uint8 items than uint8 max", func(t *testing.T) {
			items := manyUint8()
			Sort(items, 2)
			assert.True(t, sort.SliceIsSorted(items, func(i, j int) bool { return items[i] < items[j] }))
		}},
		{"Stable: more uint8 items than uint8 max", func(t *testing.T) {
			items := manyUint8()
			Stable(items, 2)
			assert.True(t, sort.SliceIsSorted(items, func(i, j int) bool { return items[i] < items[j] }))
		}},
		{"Sort: max is uint8 max", func(t *testing.T) {
			items := []uint8{255, 0, 128}
			Sort(items, 255)
			assert.Equal(t, []uint8{0, 128, 255}, items)
		}},
		{"Stable: max is int8 max", func(t *testing.T) {
			items := []int8{127, 0, 64}
			Stable(items, 127)
			assert.Equal(t, []int8{0, 64, 127}, items)
		}},
	}
	for _, test := range tests {
		t.Run(test.name, test.sort)
	}
}

func TestInts_Random(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	items := perm(1024)

	// Act
	Sort(items, 1024)

	// Assert
	ass.True(sort.IsSorted(sort.IntSlice(items)))
}

func ExampleStable() {
	items := []int{2, 5, 3, 0, 0}

	Stable(items, 5)
	fmt.Println(items)
	// Output:
	// [0 0 2 3 5]
}

func ExampleSort() {
	items := []int{2, 5, 3, 0, 0}

	Sort(items, 5)
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
		Stable(items, 1024)
	}
	b.ReportAllocs()
}

func BenchmarkInts(b *testing.B) {
	b.SetBytes(2)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		items := perm(1024)
		b.StartTimer()
		Sort(items, 1024)
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
