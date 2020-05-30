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

func ExampleIntHashSet_Contains() {
	var set = make(StringHashSet)
	set.Add("1")
	fmt.Println(set.Contains("1"))
	fmt.Println(set.Contains("2"))
	// Output
	// true
	// false
}
