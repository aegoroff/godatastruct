package collections

import (
	"fmt"
	"github.com/aegoroff/godatastruct/countingsort"
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

// IntHashSet

func TestIntHashSet_EmptySet_NoKey(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = NewIntHashSet()

	// Act

	// Assert
	_, ok := set[1]
	ass.False(ok)
}

func TestIntHashSet_Add_ItemAdded(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = NewIntHashSet()
	k := 1

	// Act
	set.Add(k)

	// Assert
	_, ok := set[k]
	ass.True(ok)
}

func TestIntHashSet_AddRange_ItemsAdded(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = NewIntHashSet()
	k1 := 1
	k2 := 2

	// Act
	set.AddRange(k1, k2)

	// Assert
	_, ok := set[k1]
	ass.True(ok)
	_, ok = set[k2]
	ass.True(ok)
}

func TestIntHashSet_Remove_ItemRemoved(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = NewIntHashSet()

	k := 1
	set.Add(k)

	// Act
	set.Remove(k)

	// Assert
	_, ok := set[k]
	ass.False(ok)
}

func TestIntHashSet_RemoveUnexist_NothingHappend(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = NewIntHashSet()

	k := 1
	set.Add(k)

	// Act
	set.Remove(2)

	// Assert
	_, ok := set[k]
	ass.True(ok)
}

func TestIntHashSet_Contains_ResultTrue(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = NewIntHashSet()
	set.Add(1)

	// Act
	ok := set.Contains(1)

	// Assert
	ass.True(ok)
}

func TestIntHashSet_ContainsUnexistKey_ResultFalse(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = NewIntHashSet()
	set.Add(1)

	// Act
	ok := set.Contains(2)

	// Assert
	ass.False(ok)
}

func TestIntHashSet_Count_ResultAsSpecified(t *testing.T) {
	var tests = []struct {
		name  string
		items []int
		count int
	}{
		{"one", []int{1}, 1},
		{"empty", []int{}, 0},
		{"many", []int{1, 2}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			ass := assert.New(t)

			var set = NewIntHashSet()
			for _, i := range tt.items {
				set.Add(i)
			}

			// Act
			count := set.Count()

			// Assert
			ass.Equal(tt.count, count)
		})
	}
}

func TestIntHashSet_Items_ResultAsSpecified(t *testing.T) {
	var tests = []struct {
		name  string
		items []int
	}{
		{"one", []int{1}},
		{"empty", []int{}},
		{"many", []int{3, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			ass := assert.New(t)

			var set = NewIntHashSet()
			for _, i := range tt.items {
				set.Add(i)
			}

			// Act
			items := set.Items()

			// Assert
			ass.ElementsMatch(tt.items, items)
		})
	}
}

// Int64HashSet

func TestInt64HashSet_EmptySet_NoKey(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = NewInt64HashSet()

	// Act

	// Assert
	_, ok := set[1]
	ass.False(ok)
}

func TestInt64HashSet_Add_ItemAdded(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = NewInt64HashSet()

	// Act
	set.Add(1)

	// Assert
	_, ok := set[1]
	ass.True(ok)
}

func TestInt64HashSet_AddRange_ItemsAdded(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = NewInt64HashSet()
	k1 := int64(1)
	k2 := int64(2)

	// Act
	set.AddRange(k1, k2)

	// Assert
	_, ok := set[k1]
	ass.True(ok)
	_, ok = set[k2]
	ass.True(ok)
}

func TestInt64HashSet_Remove_ItemRemoved(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = NewInt64HashSet()

	k := int64(1)
	set.Add(k)

	// Act
	set.Remove(k)

	// Assert
	_, ok := set[k]
	ass.False(ok)
}

func TestInt64HashSet_RemoveUnexist_NothingHappened(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = NewInt64HashSet()

	k := int64(1)
	set.Add(k)

	// Act
	set.Remove(int64(2))

	// Assert
	_, ok := set[k]
	ass.True(ok)
}

func TestInt64HashSet_Contains_ResultTrue(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = NewInt64HashSet()
	set.Add(1)

	// Act
	ok := set.Contains(1)

	// Assert
	ass.True(ok)
}

func TestInt64HashSet_ContainsUnexistKey_ResultFalse(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = NewInt64HashSet()
	set.Add(1)

	// Act
	ok := set.Contains(2)

	// Assert
	ass.False(ok)
}

func TestInt64HashSet_Count_ResultAsSpecified(t *testing.T) {
	var tests = []struct {
		name  string
		items []int64
		count int
	}{
		{"one", []int64{1}, 1},
		{"empty", []int64{}, 0},
		{"many", []int64{1, 2}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			ass := assert.New(t)

			var set = NewInt64HashSet()
			for _, i := range tt.items {
				set.Add(i)
			}

			// Act
			count := set.Count()

			// Assert
			ass.Equal(tt.count, count)
		})
	}
}

func TestInt64HashSet_Items_ResultAsSpecified(t *testing.T) {
	var tests = []struct {
		name  string
		items []int64
	}{
		{"one", []int64{1}},
		{"empty", []int64{}},
		{"many", []int64{3, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			ass := assert.New(t)

			var set = NewInt64HashSet()
			for _, i := range tt.items {
				set.Add(i)
			}

			// Act
			items := set.Items()

			// Assert
			ass.ElementsMatch(tt.items, items)
		})
	}
}

// StringHashSet

func TestStringHashSet_EmptySet_NoKey(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = NewStringHashSet()

	// Act

	// Assert
	_, ok := set["1"]
	ass.False(ok)
}

func TestStringHashSet_Add_ItemAdded(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = NewStringHashSet()
	k := "1"

	// Act
	set.Add(k)

	// Assert
	_, ok := set[k]
	ass.True(ok)
}

func TestStringHashSet_AddRange_ItemsAdded(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = NewStringHashSet()
	k1 := "1"
	k2 := "2"

	// Act
	set.AddRange(k1, k2)

	// Assert
	_, ok := set[k1]
	ass.True(ok)
	_, ok = set[k2]
	ass.True(ok)
}

func TestStringHashSet_Remove_ItemRemoved(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = NewStringHashSet()

	k := "1"
	set.Add(k)

	// Act
	set.Remove(k)

	// Assert
	_, ok := set[k]
	ass.False(ok)
}

func TestStringHashSet_RemoveUnexist_NothingHappened(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = NewStringHashSet()

	k := "1"
	set.Add(k)

	// Act
	set.Remove("2")

	// Assert
	_, ok := set[k]
	ass.True(ok)
}

func TestStringHashSet_Contains_ResultTrue(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = NewStringHashSet()
	set.Add("1")

	// Act
	ok := set.Contains("1")

	// Assert
	ass.True(ok)
}

func TestStringHashSet_ContainsUnexistKey_ResultFalse(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = NewStringHashSet()
	set.Add("1")

	// Act
	ok := set.Contains("2")

	// Assert
	ass.False(ok)
}

func TestStringHashSet_Count_ResultAsSpecified(t *testing.T) {
	var tests = []struct {
		name  string
		items []string
		count int
	}{
		{"one", []string{"1"}, 1},
		{"empty", []string{}, 0},
		{"many", []string{"1", "2"}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			ass := assert.New(t)

			var set = NewStringHashSet()
			for _, i := range tt.items {
				set.Add(i)
			}

			// Act
			count := set.Count()

			// Assert
			ass.Equal(tt.count, count)
		})
	}
}

func TestStringHashSet_Items_ResultAsSpecified(t *testing.T) {
	var tests = []struct {
		name  string
		items []string
	}{
		{"one", []string{"1"}},
		{"empty", []string{}},
		{"many", []string{"3", "1"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			ass := assert.New(t)

			var set = NewStringHashSet()
			for _, i := range tt.items {
				set.Add(i)
			}

			// Act
			items := set.Items()

			// Assert
			ass.ElementsMatch(tt.items, items)
		})
	}
}

func TestStringHashSet_ItemsDecorated_ResultAsSpecified(t *testing.T) {
	items := []string{"a", "b"}

	var tests = []struct {
		name      string
		decorator func(s string) string
		result    []string
	}{
		{"pass through", func(s string) string { return s }, []string{"a", "b"}},
		{"transform", strings.ToUpper, []string{"A", "B"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			ass := assert.New(t)

			var set = NewStringHashSet()
			for _, i := range items {
				set.Add(i)
			}

			// Act
			items := set.ItemsDecorated(tt.decorator)

			// Assert
			ass.ElementsMatch(tt.result, items)
		})
	}
}

func ExampleStringHashSet_Contains() {
	var set = NewStringHashSet()
	set.Add("1")
	fmt.Println(set.Contains("1"))
	fmt.Println(set.Contains("2"))
	// Output:
	// true
	// false
}

func ExampleStringHashSet_ItemsDecorated() {
	var set = NewStringHashSet()
	set.Add("a")

	itemsOriginal := set.ItemsDecorated(func(s string) string { return s })
	itemsUpperCased := set.ItemsDecorated(strings.ToUpper)

	fmt.Println(itemsOriginal)
	fmt.Println(itemsUpperCased)
	// Output:
	// [a]
	// [A]
}

func ExampleStringHashSet_SortedItems() {
	var set = NewStringHashSet()
	set.Add("c")
	set.Add("a")
	set.Add("b")

	items := set.SortedItems(sort.Strings)

	fmt.Println(items)
	// Output:
	// [a b c]
}

func ExampleIntHashSet_SortedItems() {
	var set = NewIntHashSet()
	set.Add(3)
	set.Add(1)
	set.Add(2)

	items := set.SortedItems(sort.Ints)

	fmt.Println(items)
	// Output:
	// [1 2 3]
}

func ExampleInt64HashSet_SortedItems() {
	var set = NewInt64HashSet()
	set.Add(3)
	set.Add(1)
	set.Add(2)

	items := set.SortedItems(func(a []int64) { sort.Sort(countingsort.Int64Slice(a)) })

	fmt.Println(items)
	// Output:
	// [1 2 3]
}
