package rbtree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Foreach_Normal(t *testing.T) {
	tree := createIntegerTestTree()
	var tests = []struct {
		name     string
		it       Enumerable
		expected []int
	}{
		{"ascend normal", NewAscend(tree), []int{2, 3, 4, 6, 7, 9, 13, 15, 17, 18, 20}},

		{"ascend range 6 to 15", NewAscendRange(tree, NewInt(6), NewInt(15)), []int{6, 7, 9, 13, 15}},
		{"ascend range 6 to 6", NewAscendRange(tree, NewInt(6), NewInt(6)), []int{6}},
		{"ascend range 15 to 15", NewAscendRange(tree, NewInt(15), NewInt(15)), []int{15}},
		{"ascend range 20 to 20", NewAscendRange(tree, NewInt(20), NewInt(20)), []int{20}},
		{"ascend range 2 to 2", NewAscendRange(tree, NewInt(2), NewInt(2)), []int{2}},
		{"ascend range 15 to 6", NewAscendRange(tree, NewInt(15), NewInt(6)), []int{}},
		{"ascend range 8 to 15", NewAscendRange(tree, NewInt(8), NewInt(15)), []int{}},

		{"ascend range nil to val", NewAscendRange(tree, nil, NewInt(6)), []int{}},
		{"ascend range val to nil", NewAscendRange(tree, NewInt(6), nil), []int{}},
		{"ascend range nil to nil", NewAscendRange(tree, nil, nil), []int{}},

		{"descend normal", NewDescend(tree), []int{20, 18, 17, 15, 13, 9, 7, 6, 4, 3, 2}},

		{"descend range 15 to 6", NewDescendRange(tree, NewInt(15), NewInt(6)), []int{15, 13, 9, 7, 6}},
		{"descend range 6 to 6", NewDescendRange(tree, NewInt(6), NewInt(6)), []int{6}},
		{"descend range 15 to 15", NewDescendRange(tree, NewInt(15), NewInt(15)), []int{15}},
		{"descend range 20 to 20", NewDescendRange(tree, NewInt(20), NewInt(20)), []int{20}},
		{"descend range 2 to 2", NewDescendRange(tree, NewInt(2), NewInt(2)), []int{2}},
		{"descend range 6 to 15", NewDescendRange(tree, NewInt(6), NewInt(15)), []int{}},
		{"descend range 14 to 6", NewDescendRange(tree, NewInt(14), NewInt(6)), []int{}},

		{"descend range nil to val", NewDescendRange(tree, nil, NewInt(6)), []int{}},
		{"descend range val to nil", NewDescendRange(tree, NewInt(6), nil), []int{}},
		{"descend range nil to nil", NewDescendRange(tree, nil, nil), []int{}},

		{"inorder normal", NewWalkInorder(tree), []int{2, 3, 4, 6, 7, 9, 13, 15, 17, 18, 20}},
		{"preorder normal", NewWalkPreorder(tree), []int{6, 3, 2, 4, 15, 9, 7, 13, 18, 17, 20}},
		{"postorder normal", NewWalkPostorder(tree), []int{2, 4, 3, 7, 13, 9, 17, 20, 18, 15, 6}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			ass := assert.New(t)
			result := make([]int, 0)

			// Act
			test.it.Foreach(func(n Node) {
				result = append(result, GetInt(n.Key()))
			})

			// Assert
			ass.Equal(test.expected, result)
		})
	}
}

func Test_Foreach_SpecialCases(t *testing.T) {
	var tests = []struct {
		name     string
		itFunc   func(t RbTree) Enumerable
		input    []int
		expected []int
	}{
		{"ascend all eq three", NewAscend, []int{2, 2, 2}, []int{2, 2, 2}},
		{"ascend all eq two", NewAscend, []int{2, 2}, []int{2, 2}},
		{"ascend all not eq two", NewAscend, []int{1, 2}, []int{1, 2}},
		{"ascend all eq one", NewAscend, []int{2}, []int{2}},
		{"ascend all eq zero", NewAscend, []int{}, []int{}},

		{"descend all eq three", NewDescend, []int{2, 2, 2}, []int{2, 2, 2}},
		{"descend all eq two", NewDescend, []int{2, 2}, []int{2, 2}},
		{"descend all not eq two", NewDescend, []int{1, 2}, []int{2, 1}},
		{"descend all eq one", NewDescend, []int{2}, []int{2}},
		{"descend all eq zero", NewDescend, []int{}, []int{}},

		{"inorder all eq three", NewWalkInorder, []int{2, 2, 2}, []int{2, 2, 2}},
		{"inorder all eq two", NewWalkInorder, []int{2, 2}, []int{2, 2}},
		{"inorder all not eq two", NewWalkInorder, []int{1, 2}, []int{1, 2}},
		{"inorder all eq one", NewWalkInorder, []int{2}, []int{2}},
		{"inorder all eq zero", NewWalkInorder, []int{}, []int{}},

		{"preorder all eq three", NewWalkPreorder, []int{2, 2, 2}, []int{2, 2, 2}},
		{"preorder all eq two", NewWalkPreorder, []int{2, 2}, []int{2, 2}},
		{"preorder all not eq two", NewWalkPreorder, []int{1, 2}, []int{1, 2}},
		{"preorder all eq one", NewWalkPreorder, []int{2}, []int{2}},
		{"preorder all eq zero", NewWalkPreorder, []int{}, []int{}},

		{"postorder all eq three", NewWalkPostorder, []int{2, 2, 2}, []int{2, 2, 2}},
		{"postorder all eq two", NewWalkPostorder, []int{2, 2}, []int{2, 2}},
		{"postorder all not eq two", NewWalkPostorder, []int{1, 2}, []int{2, 1}},
		{"postorder all eq one", NewWalkPostorder, []int{2}, []int{2}},
		{"postorder all eq zero", NewWalkPostorder, []int{}, []int{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			ass := assert.New(t)

			tree := NewRbTree()
			for _, n := range test.input {
				tree.Insert(NewInt(n))
			}

			it := test.itFunc(tree)

			result := make([]int, 0)

			// Act
			it.Foreach(func(n Node) {
				result = append(result, GetInt(n.Key()))
			})

			// Assert
			ass.Equal(test.expected, result)
		})
	}
}

func Test_IteratorsWithInterruption_Normal(t *testing.T) {
	tree := createIntegerTestTree()
	var tests = []struct {
		name     string
		enum     Enumerable
		expected []int
	}{
		{"ascend", NewAscend(tree), []int{2, 3, 4, 6}},
		{"descend", NewDescend(tree), []int{}},
		{"inorder", NewWalkInorder(tree), []int{2, 3, 4, 6}},
		{"preorder", NewWalkPreorder(tree), []int{6, 3, 2, 4}},
		{"postorder", NewWalkPostorder(tree), []int{2, 4, 3}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			ass := assert.New(t)
			result := make([]int, 0)
			it := test.enum.Iterator()

			// Act
			for it.Next() {
				curr := GetInt(it.Current().Key())
				if curr > 6 {
					break
				}
				result = append(result, curr)
			}

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
	it.Foreach(func(n Node) {
		result = append(result, n.String())
	})

	// Assert
	ass.Equal([]string{"abc", "amd", "cisco", "do", "fake", "intel", "it", "let", "microsoft", "russia", "usa", "xxx", "yyy", "zen"}, result)
}

func Test_Foreach_EmptyTree(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := NewRbTree()
	var result []string

	var tests = []struct {
		name string
		it   Enumerable
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
			test.it.Foreach(func(n Node) {
				result = append(result, n.String())
			})

			// Assert
			ass.Equal(0, len(result))
		})
	}
}
