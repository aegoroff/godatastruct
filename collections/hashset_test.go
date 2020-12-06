package collections

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// IntHashSet

func TestIntHashSet_EmptySet_NoKey(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = make(IntHashSet)

	// Act

	// Assert
	_, ok := set[1]
	ass.False(ok)
}

func TestIntHashSet_AddItem_ItemAdded(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = make(IntHashSet)

	// Act
	set.Add(1)

	// Assert
	_, ok := set[1]
	ass.True(ok)
}

func TestIntHashSet_Contains_ResultTrue(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = make(IntHashSet)
	set.Add(1)

	// Act
	ok := set.Contains(1)

	// Assert
	ass.True(ok)
}

func TestIntHashSet_ContainsUnexistKey_ResultFalse(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = make(IntHashSet)
	set.Add(1)

	// Act
	ok := set.Contains(2)

	// Assert
	ass.False(ok)
}

func TestIntHashSet_Count_ResultAsSpecified(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	var tests = []struct {
		items []int
		count int
	}{
		{[]int{1}, 1},
		{[]int{}, 0},
		{[]int{1, 2}, 2},
	}

	for _, tt := range tests {
		var set = make(IntHashSet)
		for _, i := range tt.items {
			set.Add(i)
		}

		// Act
		count := set.Count()

		// Assert
		ass.Equal(tt.count, count)
	}
}

func TestIntHashSet_Items_ResultAsSpecified(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	var tests = []struct {
		items []int
	}{
		{[]int{1}},
		{[]int{}},
		{[]int{3, 1}},
	}

	for _, tt := range tests {
		var set = make(IntHashSet)
		for _, i := range tt.items {
			set.Add(i)
		}

		// Act
		items := set.Items()

		// Assert
		ass.ElementsMatch(tt.items, items)
	}
}

// Int64HashSet

func TestInt64HashSet_EmptySet_NoKey(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = make(Int64HashSet)

	// Act

	// Assert
	_, ok := set[1]
	ass.False(ok)
}

func TestInt64HashSet_AddItem_ItemAdded(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = make(Int64HashSet)

	// Act
	set.Add(1)

	// Assert
	_, ok := set[1]
	ass.True(ok)
}

func TestInt64HashSet_Contains_ResultTrue(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = make(Int64HashSet)
	set.Add(1)

	// Act
	ok := set.Contains(1)

	// Assert
	ass.True(ok)
}

func TestInt64HashSet_ContainsUnexistKey_ResultFalse(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = make(Int64HashSet)
	set.Add(1)

	// Act
	ok := set.Contains(2)

	// Assert
	ass.False(ok)
}

func TestInt64HashSet_Count_ResultAsSpecified(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	var tests = []struct {
		items []int64
		count int
	}{
		{[]int64{1}, 1},
		{[]int64{}, 0},
		{[]int64{1, 2}, 2},
	}

	for _, tt := range tests {
		var set = make(Int64HashSet)
		for _, i := range tt.items {
			set.Add(i)
		}

		// Act
		count := set.Count()

		// Assert
		ass.Equal(tt.count, count)
	}
}

func TestInt64HashSet_Items_ResultAsSpecified(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	var tests = []struct {
		items []int64
	}{
		{[]int64{1}},
		{[]int64{}},
		{[]int64{3, 1}},
	}

	for _, tt := range tests {
		var set = make(Int64HashSet)
		for _, i := range tt.items {
			set.Add(i)
		}

		// Act
		items := set.Items()

		// Assert
		ass.ElementsMatch(tt.items, items)
	}
}

// StringHashSet

func TestStringHashSet_EmptySet_NoKey(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = make(StringHashSet)

	// Act

	// Assert
	_, ok := set["1"]
	ass.False(ok)
}

func TestStringHashSet_AddItem_ItemAdded(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = make(StringHashSet)

	// Act
	set.Add("1")

	// Assert
	_, ok := set["1"]
	ass.True(ok)
}

func TestStringHashSet_Contains_ResultTrue(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = make(StringHashSet)
	set.Add("1")

	// Act
	ok := set.Contains("1")

	// Assert
	ass.True(ok)
}

func TestStringHashSet_ContainsUnexistKey_ResultFalse(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var set = make(StringHashSet)
	set.Add("1")

	// Act
	ok := set.Contains("2")

	// Assert
	ass.False(ok)
}

func TestStringHashSet_Count_ResultAsSpecified(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	var tests = []struct {
		items []string
		count int
	}{
		{[]string{"1"}, 1},
		{[]string{}, 0},
		{[]string{"1", "2"}, 2},
	}

	for _, tt := range tests {
		var set = make(StringHashSet)
		for _, i := range tt.items {
			set.Add(i)
		}

		// Act
		count := set.Count()

		// Assert
		ass.Equal(tt.count, count)
	}
}

func TestStringHashSet_Items_ResultAsSpecified(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	var tests = []struct {
		items []string
	}{
		{[]string{"1"}},
		{[]string{}},
		{[]string{"3", "1"}},
	}

	for _, tt := range tests {
		var set = make(StringHashSet)
		for _, i := range tt.items {
			set.Add(i)
		}

		// Act
		items := set.Items()

		// Assert
		ass.ElementsMatch(tt.items, items)
	}
}

func TestStringHashSet_ItemsDecorated_ResultAsSpecified(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	var tests = []struct {
		items     []string
		decorator func(s string) string
		result    []string
	}{
		{[]string{"a"}, func(s string) string { return s }, []string{"a"}},
		{[]string{"a", "b"}, func(s string) string { return strings.ToUpper(s) }, []string{"A", "B"}},
	}

	for _, tt := range tests {
		var set = make(StringHashSet)
		for _, i := range tt.items {
			set.Add(i)
		}

		// Act
		items := set.ItemsDecorated(tt.decorator)

		// Assert
		ass.ElementsMatch(tt.result, items)
	}
}

func ExampleStringHashSet_Contains() {
	var set = make(StringHashSet)
	set.Add("1")
	fmt.Println(set.Contains("1"))
	fmt.Println(set.Contains("2"))
	// Output
	// true
	// false
}

func ExampleStringHashSet_ItemsDecorated() {
	var set = make(StringHashSet)
	set.Add("a")
	set.Add("b")

	itemsOriginal := set.ItemsDecorated(func(s string) string { return s })
	itemsUpperCased := set.ItemsDecorated(func(s string) string { return strings.ToUpper(s) })

	fmt.Println(itemsOriginal)
	fmt.Println(itemsUpperCased)

	// Output
	// [a b]
	// [A B]
}
