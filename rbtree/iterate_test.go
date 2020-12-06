package rbtree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Iterate_Normal(t *testing.T) {
	tree := createIntegerTestTree()
	allTrue := func(c Comparable) bool { return true }
	allFalse := func(c Comparable) bool { return false }
	var tests = []struct {
		name      string
		it        Iterator
		predicate func(Comparable) bool
		expected  []Int
	}{
		{"ascend normal", NewAscend(tree), allTrue, []Int{2, 3, 4, 6, 7, 9, 13, 15, 17, 18, 20}},
		{"ascend with breaking immediately", NewAscend(tree), allFalse, []Int{2}},
		{"ascend with break complex condition", NewAscend(tree), func(c Comparable) bool { return c.LessThan(NewInt(15)) }, []Int{2, 3, 4, 6, 7, 9, 13, 15}},

		{"ascend range 6 to 15", NewAscendRange(tree, NewInt(6), NewInt(15)), allTrue, []Int{6, 7, 9, 13, 15}},
		{"ascend range 6 to 6", NewAscendRange(tree, NewInt(6), NewInt(6)), allTrue, []Int{6}},
		{"ascend range 15 to 15", NewAscendRange(tree, NewInt(15), NewInt(15)), allTrue, []Int{15}},
		{"ascend range 20 to 20", NewAscendRange(tree, NewInt(20), NewInt(20)), allTrue, []Int{20}},
		{"ascend range 2 to 2", NewAscendRange(tree, NewInt(2), NewInt(2)), allTrue, []Int{2}},
		{"ascend range 15 to 6", NewAscendRange(tree, NewInt(15), NewInt(6)), allTrue, []Int{}},
		{"ascend range 8 to 15", NewAscendRange(tree, NewInt(8), NewInt(15)), allTrue, []Int{}},

		{"ascend range nil to val", NewAscendRange(tree, nil, NewInt(6)), allTrue, []Int{}},
		{"ascend range val to nil", NewAscendRange(tree, NewInt(6), nil), allTrue, []Int{}},
		{"ascend range nil to nil", NewAscendRange(tree, nil, nil), allTrue, []Int{}},

		{"descend normal", NewDescend(tree), allTrue, []Int{20, 18, 17, 15, 13, 9, 7, 6, 4, 3, 2}},
		{"descend with breaking immediately", NewDescend(tree), allFalse, []Int{20}},
		{"descend complex condition", NewDescend(tree), func(c Comparable) bool { return !c.LessThan(NewInt(15)) }, []Int{20, 18, 17, 15, 13}},

		{"descend range 15 to 6", NewDescendRange(tree, NewInt(15), NewInt(6)), allTrue, []Int{15, 13, 9, 7, 6}},
		{"descend range 6 to 6", NewDescendRange(tree, NewInt(6), NewInt(6)), allTrue, []Int{6}},
		{"descend range 15 to 15", NewDescendRange(tree, NewInt(15), NewInt(15)), allTrue, []Int{15}},
		{"descend range 20 to 20", NewDescendRange(tree, NewInt(20), NewInt(20)), allTrue, []Int{20}},
		{"descend range 2 to 2", NewDescendRange(tree, NewInt(2), NewInt(2)), allTrue, []Int{2}},
		{"descend range 6 to 15", NewDescendRange(tree, NewInt(6), NewInt(15)), allTrue, []Int{}},
		{"descend range 14 to 6", NewDescendRange(tree, NewInt(14), NewInt(6)), allTrue, []Int{}},

		{"descend range nil to val", NewDescendRange(tree, nil, NewInt(6)), allTrue, []Int{}},
		{"descend range val to nil", NewDescendRange(tree, NewInt(6), nil), allTrue, []Int{}},
		{"descend range nil to nil", NewDescendRange(tree, nil, nil), allTrue, []Int{}},

		{"inorder normal", NewWalkInorder(tree), allTrue, []Int{2, 3, 4, 6, 7, 9, 13, 15, 17, 18, 20}},
		{"inorder with breaking", NewWalkInorder(tree), func(c Comparable) bool { return GetInt(c) <= 9 }, []Int{2, 3, 4, 6, 7, 9, 13}},
		{"preorder normal", NewWalkPreorder(tree), allTrue, []Int{6, 3, 2, 4, 15, 9, 7, 13, 18, 17, 20}},
		{"preorder with breaking", NewWalkPreorder(tree), func(c Comparable) bool { return GetInt(c) <= 9 }, []Int{6, 3, 2, 4, 15}},
		{"postorder normal", NewWalkPostorder(tree), allTrue, []Int{2, 4, 3, 7, 13, 9, 17, 20, 18, 15, 6}},
		{"postorder with breaking", NewWalkPostorder(tree), func(c Comparable) bool { return GetInt(c) <= 6 }, []Int{2, 4, 3, 7, 13, 9, 17, 20, 18, 15, 6}},
		{"postorder with breaking immediately", NewWalkPostorder(tree), allFalse, []Int{2, 4, 3, 7, 13, 9, 17, 20, 18, 15, 6}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			ass := assert.New(t)
			result := []Int{}

			// Act
			test.it.Iterate(func(n Node) bool {
				c := n.(*node).key
				result = append(result, c.(Int))
				return test.predicate(c)
			})

			// Assert
			ass.Equal(test.expected, result)
		})
	}
}

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

func Test_Iterate_EmptyTree(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := NewRbTree()
	var result []string

	var tests = []struct {
		name string
		it   Iterator
	}{
		{"inorder", NewWalkInorder(tree)},
		{"preorder", NewWalkPreorder(tree)},
		{"postorder", NewWalkPostorder(tree)},
		{"ascend", NewAscend(tree)},
		{"descend", NewDescend(tree)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Act
			test.it.Iterate(func(n Node) bool {
				result = append(result, n.Key().String())
				return true
			})

			// Assert
			ass.Equal(0, len(result))
		})
	}
}
