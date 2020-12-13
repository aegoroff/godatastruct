package special

import (
	"github.com/aegoroff/godatastruct/rbtree"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func Test_MaxSizeTree_SizeAsExpectedIterationWithoutSideEffects(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	var nodes []int
	var result []string
	const nodesCount = 200

	for i := 1; i <= nodesCount; i++ {
		nodes = append(nodes, i)
	}
	tree := rbtree.NewRbTree()
	for _, value := range nodes {
		tree.Insert(rbtree.NewInt(value))
	}

	top := int64(5)
	ft := NewMaxTree(top)

	// Act
	rbtree.NewWalkInorder(tree).Foreach(func(n rbtree.Node) {
		ft.Insert(n.Key())
	})

	iterationCount := int64(0)
	rbtree.NewDescend(ft).Foreach(func(n rbtree.Node) {
		iterationCount++
		result = append(result, n.String())
	})

	// Assert
	ass.Equal(top, ft.Len())
	ass.Equal(top, iterationCount)
	ass.Equal([]string{"200", "199", "198", "197", "196"}, result)
}

func Test_MinSizeTree_SizeAsExpectedIterationWithoutSideEffects(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	var nodes []int
	var result []string
	const nodesCount = 200

	for i := 1; i <= nodesCount; i++ {
		nodes = append(nodes, i)
	}
	tree := rbtree.NewRbTree()
	for _, value := range nodes {
		tree.Insert(rbtree.NewInt(value))
	}

	top := int64(5)
	ft := NewMinTree(top)

	// Act
	rbtree.NewWalkPreorder(tree).Foreach(func(n rbtree.Node) {
		ft.Insert(n.Key())
	})

	iterationCount := int64(0)
	rbtree.NewDescend(ft).Foreach(func(n rbtree.Node) {
		iterationCount++
		result = append(result, n.String())
	})

	// Assert
	ass.Equal(top, ft.Len())
	ass.Equal(top, iterationCount)
	ass.Equal([]string{"5", "4", "3", "2", "1"}, result)
}

func TestMaxSizeRandomTree_SizeAsExpectedIterationWithoutSideEffects(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	var nodes []string
	var result []string
	const nodesCount = 500

	for i := 1; i <= nodesCount; i++ {
		l := 1 + rand.Intn(50)
		nodes = append(nodes, randomString(l))
	}
	tree := rbtree.NewRbTree()
	for _, n := range nodes {
		c := rbtree.NewString(n)
		tree.Insert(c)
	}

	ftSize := int64(10)
	ft := NewMaxTree(ftSize)

	// Act
	rbtree.NewWalkInorder(tree).Foreach(func(n rbtree.Node) {
		ft.Insert(n.Key())
	})

	rbtree.NewDescend(ft).Foreach(func(n rbtree.Node) {
		result = append(result, n.String())
	})

	// Assert
	max := ft.Maximum()
	pred1 := max.Predecessor()
	pred2 := pred1.Predecessor()

	ass.Equal(ftSize, ft.Len())
	ass.Equal(max.String(), result[0])
	ass.Equal(pred1.String(), result[1])
	ass.Equal(pred2.String(), result[2])
	ass.Equal(ftSize, int64(len(result)))
	ass.Equal(int64(nodesCount), tree.Len())
}

func TestMinSizeRandomTree_SizeAsExpectedIterationWithoutSideEffects(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	var nodes []string
	var result []string
	const nodesCount = 500

	for i := 1; i <= nodesCount; i++ {
		l := 1 + rand.Intn(50)
		nodes = append(nodes, randomString(l))
	}
	tree := rbtree.NewRbTree()
	for _, n := range nodes {
		c := rbtree.NewString(n)
		tree.Insert(c)
	}

	ftSize := int64(10)
	ft := NewMinTree(ftSize)

	// Act
	rbtree.NewWalkInorder(tree).Foreach(func(n rbtree.Node) {
		ft.Insert(n.Key())
	})

	rbtree.NewAscend(ft).Foreach(func(n rbtree.Node) {
		result = append(result, n.String())
	})

	// Assert
	min := ft.Minimum()
	succ1 := min.Successor()
	succ2 := succ1.Successor()

	ass.Equal(ftSize, ft.Len())
	ass.Equal(min.String(), result[0])
	ass.Equal(succ1.String(), result[1])
	ass.Equal(succ2.String(), result[2])
	ass.Equal(ftSize, int64(len(result)))
	ass.Equal(int64(nodesCount), tree.Len())
}

func Test_OrderStatisticSelect_ValueAsExpected(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	minTree := NewMinTree(3)
	maxTree := NewMaxTree(3)

	for i := 1; i <= 10; i++ {
		minTree.Insert(rbtree.NewInt(i))
		maxTree.Insert(rbtree.NewInt(i))
	}

	var tests = []struct {
		name     string
		tree     rbtree.RbTree
		order    int64
		expected int
	}{
		{"Min tree 1", minTree, 1, 1},
		{"Min tree 3", minTree, 3, 3},
		{"Max tree 1", maxTree, 1, 8},
		{"Max tree 3", maxTree, 3, 10},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Act
			found, _ := test.tree.OrderStatisticSelect(test.order)

			// Assert
			ass.NotNil(found)
			ass.Equal(test.expected, rbtree.GetInt(found.Key()))
		})
	}
}

func Test_SearchIntTree_Success(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	minTree := NewMinTree(3)
	maxTree := NewMaxTree(3)

	for i := 1; i <= 10; i++ {
		minTree.Insert(rbtree.NewInt(i))
		maxTree.Insert(rbtree.NewInt(i))
	}

	var tests = []struct {
		name     string
		tree     rbtree.RbTree
		expected int
	}{
		{"Min tree", minTree, 1},
		{"Max tree", maxTree, 8},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := rbtree.NewInt(test.expected)

			// Act
			found, ok := test.tree.Search(v)

			// Assert
			ass.True(ok)
			ass.NotNil(found)
			ass.Equal(test.expected, rbtree.GetInt(found.Key()))
		})
	}
}

func Test_DeleteNode_Success(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	minTree := NewMinTree(3)
	maxTree := NewMaxTree(3)

	for i := 1; i <= 10; i++ {
		minTree.Insert(rbtree.NewInt(i))
		maxTree.Insert(rbtree.NewInt(i))
	}

	var tests = []struct {
		name     string
		tree     rbtree.RbTree
		expected int
	}{
		{"Min tree", minTree, 1},
		{"Max tree", maxTree, 8},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := rbtree.NewInt(test.expected)

			// Act
			ok := test.tree.DeleteNode(v)

			// Assert
			ass.True(ok)
			found, ok := test.tree.Search(v)
			ass.False(ok)
			ass.Nil(found)
		})
	}
}

func Test_DeleteAllNodesNodes_Success(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	minTree := NewMinTree(3)
	maxTree := NewMaxTree(3)

	for i := 1; i <= 10; i++ {
		minTree.Insert(rbtree.NewInt(i))
		maxTree.Insert(rbtree.NewInt(i))
	}

	var tests = []struct {
		name     string
		tree     rbtree.RbTree
		expected int
	}{
		{"Min tree", minTree, 1},
		{"Max tree", maxTree, 8},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := rbtree.NewInt(test.expected)

			// Act
			ok := test.tree.DeleteAllNodes(v)

			// Assert
			ass.True(ok)
			found, ok := test.tree.Search(v)
			ass.False(ok)
			ass.Nil(found)
		})
	}
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	ix := rand.Intn(len(letters))
	for i := range s {
		s[i] = letters[ix]
	}
	return string(s)
}
