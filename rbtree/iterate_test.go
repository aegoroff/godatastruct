package rbtree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_InorderWalkString_AllElementsAscending(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()
	var result []string

	// Act
	tree.WalkInorder(func(n Node) {
		result = append(result, n.Key().String())
	})

	// Assert
	ass.Equal([]string{"abc", "amd", "cisco", "do", "fake", "intel", "it", "let", "microsoft", "russia", "usa", "xxx", "yyy", "zen"}, result)
}

func Test_InorderWalkTreeInt_AllElementsAscending(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	var result []int

	// Act
	tree.WalkInorder(func(n Node) {
		result = append(result, GetInt(n.Key()))
	})

	// Assert
	ass.Equal([]int{2, 3, 4, 6, 7, 9, 13, 15, 17, 18, 20}, result)
}

func Test_PreorderAllTreeWalkInt_AllElementsAsSpecified(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	var result []int

	// Act
	tree.WalkPreorder(func(n Node) {
		result = append(result, GetInt(n.Key()))
	})

	// Assert
	ass.Equal([]int{6, 3, 2, 4, 15, 9, 7, 13, 18, 17, 20}, result)
}

func Test_PostorderAllTreeWalkInt_AllElementsAsSpecified(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	var result []int

	// Act
	tree.WalkPostorder(func(n Node) {
		result = append(result, GetInt(n.Key()))
	})

	// Assert
	ass.Equal([]int{2, 4, 3, 7, 13, 9, 17, 20, 18, 15, 6}, result)
}

func Test_Ascend(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	var tests = []struct {
		predicate func(Comparable) bool
		expected  []Int
	}{
		{func(c Comparable) bool { return true }, []Int{2, 3, 4, 6, 7, 9, 13, 15, 17, 18, 20}},
		{func(c Comparable) bool { return false }, []Int{2}},
		{func(c Comparable) bool { return c.LessThan(NewInt(15)) }, []Int{2, 3, 4, 6, 7, 9, 13, 15}},
	}
	for _, test := range tests {
		result := []Int{}

		// Act
		tree.Ascend(func(n Node) bool {
			c := n.(*node).key
			result = append(result, c.(Int))
			return test.predicate(c)
		})

		// Assert
		ass.Equal(test.expected, result)
	}
}

func Test_AscendEmptyTree(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := NewRbTree()
	result := []Int{}

	// Act
	tree.Ascend(func(n Node) bool {
		result = append(result, n.Key().(Int))
		return true
	})

	// Assert
	ass.Equal([]Int{}, result)
}

func Test_AscendRange(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	var tests = []struct {
		from     int
		to       int
		expected []Int
	}{
		{6, 15, []Int{6, 7, 9, 13, 15}},
		{6, 6, []Int{6}},
		{15, 15, []Int{15}},
		{20, 20, []Int{20}},
		{2, 2, []Int{2}},
		{15, 6, []Int{}},
		{8, 15, []Int{}},
	}
	for _, test := range tests {
		result := []Int{}
		from := NewInt(test.from)
		to := NewInt(test.to)

		// Act
		tree.AscendRange(from, to, func(n Node) bool {
			result = append(result, n.Key().(Int))
			return true
		})

		// Assert
		ass.Equal(test.expected, result)
	}
}

func Test_AscendRangeNilTests(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	var tests = []struct {
		from Comparable
		to   Comparable
	}{
		{nil, NewInt(6)},
		{NewInt(6), nil},
		{nil, nil},
	}
	for _, test := range tests {
		result := []Int{}

		// Act
		tree.AscendRange(test.from, test.to, func(n Node) bool {
			result = append(result, n.Key().(Int))
			return true
		})

		// Assert
		ass.Equal([]Int{}, result)
	}
}

func Test_Descend(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	var tests = []struct {
		predicate func(Comparable) bool
		expected  []Int
	}{
		{func(c Comparable) bool { return true }, []Int{20, 18, 17, 15, 13, 9, 7, 6, 4, 3, 2}},
		{func(c Comparable) bool { return false }, []Int{20}},
		{func(c Comparable) bool { return !c.LessThan(NewInt(15)) }, []Int{20, 18, 17, 15, 13}},
	}
	for _, test := range tests {
		result := []Int{}

		// Act
		tree.Descend(func(n Node) bool {
			result = append(result, n.Key().(Int))
			return test.predicate(n.Key())
		})

		// Assert
		ass.Equal(test.expected, result)
	}
}

func Test_DescendEmptyTree(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := NewRbTree()
	result := []Int{}

	// Act
	tree.Descend(func(n Node) bool {
		result = append(result, n.Key().(Int))
		return true
	})

	// Assert
	ass.Equal([]Int{}, result)
}

func Test_DescendRange(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	var tests = []struct {
		from     int
		to       int
		expected []Int
	}{
		{15, 6, []Int{15, 13, 9, 7, 6}},
		{6, 6, []Int{6}},
		{15, 15, []Int{15}},
		{20, 20, []Int{20}},
		{2, 2, []Int{2}},
		{6, 15, []Int{}},
		{14, 6, []Int{}},
	}
	for _, test := range tests {
		result := []Int{}
		from := NewInt(test.from)
		to := NewInt(test.to)

		// Act
		tree.DescendRange(from, to, func(n Node) bool {
			result = append(result, n.Key().(Int))
			return true
		})

		// Assert
		ass.Equal(test.expected, result)
	}
}

func Test_DescendRangeNilTests(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	var tests = []struct {
		from Comparable
		to   Comparable
	}{
		{nil, NewInt(6)},
		{NewInt(6), nil},
		{nil, nil},
	}
	for _, test := range tests {
		result := []Int{}

		// Act
		tree.DescendRange(test.from, test.to, func(n Node) bool {
			result = append(result, n.Key().(Int))
			return true
		})

		// Assert
		ass.Equal([]Int{}, result)
	}
}
