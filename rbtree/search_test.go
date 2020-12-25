package rbtree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_OrderStatisticSelect_ValueAsExpected(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	var tests = []struct {
		order    int64
		expected int
	}{
		{1, 2},
		{2, 3},
		{11, 20},
		{10, 18},
		{6, 9},
	}
	for _, test := range tests {
		// Act
		found, _ := tree.OrderStatisticSelect(test.order)

		// Assert
		ass.NotNil(found)
		ass.Equal(test.expected, GetInt(found.Key()))
	}
}

func Test_OrderStatisticSelectNegativeTests_NullResult(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	var tests = []struct {
		tree  RbTree
		order int64
	}{
		{createIntegerTestTree(), 200},
		{NewRbTree(), 1},
	}
	for _, test := range tests {
		// Act
		found, ok := test.tree.OrderStatisticSelect(test.order)

		// Assert
		ass.Nil(found)
		ass.False(ok)
	}
}

func Test_SearchIntTree_Success(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	v := Int(13)

	// Act
	found, ok := tree.Search(v)

	// Assert
	ass.True(ok)
	ass.NotNil(found)
	ass.Equal(13, GetInt(found))
}

func Test_SearchAllIntTree_Success(t *testing.T) {
	var tests = []struct {
		name       string
		duplicates int
	}{
		{"1", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			ass := assert.New(t)
			tree := NewRbTree()
			tree.Insert(Int(1))

			for i := 0; i < test.duplicates; i++ {
				tree.Insert(Int(2))
			}

			tree.Insert(Int(3))
			tree.Insert(Int(4))
			v := Int(2)

			// Act
			found := tree.SearchAll(v)

			// Assert
			ass.NotNil(found)
			ass.Equal(test.duplicates, len(found))
		})
	}
}

func Test_SearchStringTree_Success(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()
	n := NewString("intel")

	// Act
	found, ok := tree.Search(n)

	// Assert
	ass.True(ok)
	ass.NotNil(found)
	ass.Equal("intel", found.(*String).String())
}

func Test_SearchStringTree_Fail(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()
	n := NewString("abrakadabra")

	// Act
	found, ok := tree.Search(n)

	// Assert
	ass.False(ok)
	ass.Nil(found)
}

func Test_SearchAllStringTree_Fail(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()
	n := NewString("abrakadabra")

	// Act
	found := tree.SearchAll(n)

	// Assert
	ass.Nil(found)
}

func Test_SearchIntTree_Failure(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	var tests = []struct {
		tree RbTree
		key  Comparable
	}{
		{createIntegerTestTree(), Int(22)},
		{createIntegerTestTree(), nil},
		{NewRbTree(), Int(20)},
	}
	for _, test := range tests {
		// Act
		found, ok := test.tree.Search(test.key)

		// Assert
		ass.False(ok)
		ass.Nil(found)
	}
}

func Test_Successor_ReturnSuccessor(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	var tests = []struct {
		node     int
		expected int
	}{
		{13, 15},
		{6, 7},
		{18, 20},
		{2, 3},
	}
	for _, test := range tests {
		v := Int(test.node)
		r, _ := tree.SearchNode(v)

		// Act
		s := r.Successor()

		// Assert
		ass.Equal(test.expected, GetInt(s.Key()))
	}
}

func Test_SuccessorOfMax_ReturnNil(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	v := Int(20)
	r, _ := tree.SearchNode(v)

	// Act
	s := r.Successor()

	// Assert
	ass.Nil(s)
}

func Test_SuccessorOfNil_ReturnNil(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var r *Node

	// Act
	s := r.Successor()

	// Assert
	ass.Nil(s)
}

func Test_PredecessorInTheMiddle_PredecessorFound(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	var tests = []struct {
		node     int
		expected int
	}{
		{13, 9},
		{6, 4},
		{18, 17},
		{3, 2},
	}
	for _, test := range tests {
		v := Int(test.node)
		r, _ := tree.SearchNode(v)

		// Act
		s := r.Predecessor()

		// Assert
		ass.Equal(test.expected, GetInt(s.Key()))
	}
}

func Test_PredecessorOfMin_ReturnNil(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	v := Int(2)
	r, _ := tree.SearchNode(v)

	// Act
	p := r.Predecessor()

	// Assert
	ass.Nil(p)
}

func Test_PredecessorOfNull_ReturnNil(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var r *Node

	// Act
	p := r.Predecessor()

	// Assert
	ass.Nil(p)
}

func Test_Minimum_ValueAndSizeAsExpected(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	// Act
	r := tree.Minimum()

	// Assert
	ass.Equal(2, GetInt(r.Key()))
	ass.Equal(int64(1), r.Size())
}

func Test_MinimumEmptyTree(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := NewRbTree()

	// Act
	r := tree.Minimum()

	// Assert
	ass.Nil(r)
}

func Test_Maximum_ValueAsExpected(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	// Act
	r := tree.Maximum()

	// Assert
	ass.Equal(20, GetInt(r.Key()))
}

func Test_MaximumEmptyTree(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := NewRbTree()

	// Act
	r := tree.Maximum()

	// Assert
	ass.Nil(r)
}
