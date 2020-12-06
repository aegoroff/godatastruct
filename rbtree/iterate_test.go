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
	it := NewWalkInorder(tree)

	// Act
	it.Iterate(func(n Node) bool {
		result = append(result, n.Key().String())
		return true
	})

	// Assert
	ass.Equal([]string{"abc", "amd", "cisco", "do", "fake", "intel", "it", "let", "microsoft", "russia", "usa", "xxx", "yyy", "zen"}, result)
}

func Test_InorderEmptyTree_NothingHappened(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := NewRbTree()
	var result []string
	it := NewWalkInorder(tree)

	// Act
	it.Iterate(func(n Node) bool {
		result = append(result, n.Key().String())
		return true
	})

	// Assert
	ass.Equal(0, len(result))
}

func Test_InorderWalkTreeInt_AllElementsAscending(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	var result []int
	it := NewWalkInorder(tree)

	// Act
	it.Iterate(func(n Node) bool {
		result = append(result, GetInt(n.Key()))
		return true
	})

	// Assert
	ass.Equal([]int{2, 3, 4, 6, 7, 9, 13, 15, 17, 18, 20}, result)
}

func Test_InorderWalkTreeIntIterateBreaking_AllElementsAscending(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	var result []int
	it := NewWalkInorder(tree)

	// Act
	it.Iterate(func(n Node) bool {
		i := GetInt(n.Key())
		result = append(result, i)
		return i <= 9
	})

	// Assert
	ass.Equal([]int{2, 3, 4, 6, 7, 9, 13}, result)
}

func Test_PreorderAllTreeWalkInt_AllElementsAsSpecified(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	var result []int
	it := NewWalkPreorder(tree)

	// Act
	it.Iterate(func(n Node) bool {
		result = append(result, GetInt(n.Key()))
		return true
	})

	// Assert
	ass.Equal([]int{6, 3, 2, 4, 15, 9, 7, 13, 18, 17, 20}, result)
}

func Test_PreorderAllTreeWithBreaking_AllElementsAsSpecified(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	var result []int
	it := NewWalkPreorder(tree)

	// Act
	it.Iterate(func(n Node) bool {
		i := GetInt(n.Key())
		result = append(result, i)
		return i <= 9
	})

	// Assert
	ass.Equal([]int{6, 3, 2, 4, 15}, result)
}

func Test_PreorderEmptyTree_NothingHappened(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := NewRbTree()
	var result []string
	it := NewWalkPreorder(tree)

	// Act
	it.Iterate(func(n Node) bool {
		result = append(result, n.Key().String())
		return true
	})

	// Assert
	ass.Equal(0, len(result))
}

func Test_PostorderAllTreeWalkInt_AllElementsAsSpecified(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	var result []int
	it := NewWalkPostorder(tree)

	// Act
	it.Iterate(func(n Node) bool {
		result = append(result, GetInt(n.Key()))
		return true
	})

	// Assert
	ass.Equal([]int{2, 4, 3, 7, 13, 9, 17, 20, 18, 15, 6}, result)
}

func Test_PostorderAllTreeWithBreaking_AllElementsAsSpecified(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	var result []int
	it := NewWalkPostorder(tree)

	// Act
	it.Iterate(func(n Node) bool {
		i := GetInt(n.Key())
		result = append(result, i)
		return i <= 6
	})

	// Assert
	ass.Equal([]int{2, 4, 3, 7, 13, 9, 17, 20, 18, 15, 6}, result)
}

func Test_PostorderEmptyTree_NothingHappened(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := NewRbTree()
	var result []string
	it := NewWalkPostorder(tree)

	// Act
	it.Iterate(func(n Node) bool {
		result = append(result, n.Key().String())
		return true
	})

	// Assert
	ass.Equal(0, len(result))
}

func Test_Ascend(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	it := NewAscend(tree)

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
		it.Iterate(func(n Node) bool {
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
	it := NewAscend(tree)

	// Act
	it.Iterate(func(n Node) bool {
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
		it := NewAscendRange(tree, from, to)

		// Act
		it.Iterate(func(n Node) bool {
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
		it := NewAscendRange(tree, test.from, test.to)

		// Act
		it.Iterate(func(n Node) bool {
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
	it := NewDescend(tree)

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
		it.Iterate(func(n Node) bool {
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
	it := NewDescend(tree)

	// Act
	it.Iterate(func(n Node) bool {
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
		it := NewDescendRange(tree, from, to)

		// Act
		it.Iterate(func(n Node) bool {
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
		it := NewDescendRange(tree, test.from, test.to)

		// Act
		it.Iterate(func(n Node) bool {
			result = append(result, n.Key().(Int))
			return true
		})

		// Assert
		ass.Equal([]Int{}, result)
	}
}
