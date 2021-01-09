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
	var result []int
	const nodesCount = 200

	for i := 1; i <= nodesCount; i++ {
		nodes = append(nodes, i)
	}
	tree := rbtree.New()
	for _, value := range nodes {
		tree.Insert(rbtree.Int(value))
	}

	top := int64(5)
	ft := NewMaxTree(top)

	// Act
	rbtree.NewWalkInorder(tree).Foreach(func(n rbtree.Comparable) {
		ft.Insert(n)
	})

	iterationCount := int64(0)
	rbtree.NewDescend(ft).Foreach(func(n rbtree.Comparable) {
		iterationCount++
		result = append(result, rbtree.GetInt(n))
	})

	// Assert
	ass.Equal(top, ft.Len())
	ass.Equal(top, iterationCount)
	ass.Equal([]int{200, 199, 198, 197, 196}, result)
}

func Test_MinSizeTree_SizeAsExpectedIterationWithoutSideEffects(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	var nodes []int
	var result []int
	const nodesCount = 200

	for i := 1; i <= nodesCount; i++ {
		nodes = append(nodes, i)
	}
	tree := rbtree.New()
	for _, value := range nodes {
		tree.Insert(rbtree.Int(value))
	}

	top := int64(5)
	ft := NewMinTree(top)

	// Act
	rbtree.NewWalkPreorder(tree).Foreach(func(n rbtree.Comparable) {
		ft.Insert(n)
	})

	iterationCount := int64(0)
	rbtree.NewDescend(ft).Foreach(func(n rbtree.Comparable) {
		iterationCount++
		result = append(result, rbtree.GetInt(n))
	})

	// Assert
	ass.Equal(top, ft.Len())
	ass.Equal(top, iterationCount)
	ass.Equal([]int{5, 4, 3, 2, 1}, result)
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
	tree := rbtree.New()
	for _, n := range nodes {
		c := rbtree.NewString(n)
		tree.Insert(c)
	}

	ftSize := int64(10)
	ft := NewMaxTree(ftSize)

	// Act
	rbtree.NewWalkInorder(tree).Foreach(func(n rbtree.Comparable) {
		ft.Insert(n)
	})

	rbtree.NewDescend(ft).Foreach(func(n rbtree.Comparable) {
		result = append(result, n.(*rbtree.String).String())
	})

	// Assert
	max := ft.Maximum()
	pred1 := max.Predecessor()
	pred2 := pred1.Predecessor()

	ass.Equal(ftSize, ft.Len())
	ass.Equal(max.Key().(*rbtree.String).String(), result[0])
	ass.Equal(pred1.Key().(*rbtree.String).String(), result[1])
	ass.Equal(pred2.Key().(*rbtree.String).String(), result[2])
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
	tree := rbtree.New()
	for _, n := range nodes {
		c := rbtree.NewString(n)
		tree.Insert(c)
	}

	ftSize := int64(10)
	ft := NewMinTree(ftSize)

	// Act
	rbtree.NewWalkInorder(tree).Foreach(func(n rbtree.Comparable) {
		ft.Insert(n)
	})

	rbtree.NewAscend(ft).Foreach(func(n rbtree.Comparable) {
		result = append(result, n.(*rbtree.String).String())
	})

	// Assert
	min := ft.Minimum()
	succ1 := min.Successor()
	succ2 := succ1.Successor()

	ass.Equal(ftSize, ft.Len())
	ass.Equal(min.Key().(*rbtree.String).String(), result[0])
	ass.Equal(succ1.Key().(*rbtree.String).String(), result[1])
	ass.Equal(succ2.Key().(*rbtree.String).String(), result[2])
	ass.Equal(ftSize, int64(len(result)))
	ass.Equal(int64(nodesCount), tree.Len())
}

func Test_OrderStatisticSelect_ValueAsExpected(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	minTree := NewMinTree(3)
	maxTree := NewMaxTree(3)

	for i := 1; i <= 10; i++ {
		minTree.Insert(rbtree.Int(i))
		maxTree.Insert(rbtree.Int(i))
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
		minTree.Insert(rbtree.Int(i))
		maxTree.Insert(rbtree.Int(i))
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
			v := rbtree.Int(test.expected)

			// Act
			found, ok := test.tree.Search(v)

			// Assert
			ass.True(ok)
			ass.NotNil(found)
			ass.Equal(test.expected, rbtree.GetInt(found))
		})
	}
}

func Test_SearchAllIntTree_Success(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	minTree := NewMinTree(3)
	maxTree := NewMaxTree(3)

	for i := 1; i <= 10; i++ {
		minTree.Insert(rbtree.Int(i))
		maxTree.Insert(rbtree.Int(i))
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
			v := rbtree.Int(test.expected)

			// Act
			found := test.tree.SearchAll(v)

			// Assert
			ass.NotNil(found)
			ass.Equal(1, len(found))
		})
	}
}

func Test_SearchNodeIntTree_Success(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	minTree := NewMinTree(3)
	maxTree := NewMaxTree(3)

	for i := 1; i <= 10; i++ {
		minTree.Insert(rbtree.Int(i))
		maxTree.Insert(rbtree.Int(i))
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
			v := rbtree.Int(test.expected)

			// Act
			found, ok := test.tree.SearchNode(v)

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
		minTree.Insert(rbtree.Int(i))
		maxTree.Insert(rbtree.Int(i))
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
			v := rbtree.Int(test.expected)

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

func Test_ReplaceOrInsertMinTreeDuplicateLessThenPrevNodes_PrevNodesDeletedLenLessThenMinLen(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := NewMinTree(3)

	tree.Insert(rbtree.Int(1))
	tree.Insert(rbtree.Int(3))
	tree.Insert(rbtree.Int(4))

	duplicate := rbtree.Int(2)

	// Act
	for i := 0; i <= 5; i++ {
		tree.ReplaceOrInsert(duplicate)
	}

	// Assert
	found, ok := tree.Search(duplicate)
	ass.True(ok)
	ass.Equal(int64(2), tree.Len())
	ass.NotNil(found)
	_, ok = tree.Search(rbtree.Int(3))
	ass.False(ok)
	_, ok = tree.Search(rbtree.Int(4))
	ass.False(ok)
}

func Test_ReplaceOrInsertMinTreeDuplicateGreaterThenPrevNodes_PrevNodesNotDeleteLenEqualThenMinLen(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := NewMinTree(3)

	tree.Insert(rbtree.Int(1))
	tree.Insert(rbtree.Int(2))

	duplicate := rbtree.Int(3)

	// Act
	for i := 0; i <= 5; i++ {
		tree.ReplaceOrInsert(duplicate)
	}

	tree.ReplaceOrInsert(rbtree.Int(4))

	// Assert
	found, ok := tree.Search(duplicate)
	ass.True(ok)
	ass.Equal(int64(3), tree.Len())
	ass.NotNil(found)
	_, ok = tree.Search(rbtree.Int(3))
	ass.True(ok)
	_, ok = tree.Search(rbtree.Int(4))
	ass.False(ok)
}

func Test_ReplaceOrInsertMaxTreeDuplicateLessThenPrevNodes_PrevNodesNotDeletedLenEqualMaxLen(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := NewMaxTree(3)

	tree.Insert(rbtree.Int(1))
	tree.Insert(rbtree.Int(3))
	tree.Insert(rbtree.Int(4))

	duplicate := rbtree.Int(2)

	// Act
	for i := 0; i <= 5; i++ {
		tree.ReplaceOrInsert(duplicate)
	}

	// Assert
	found, ok := tree.Search(duplicate)
	ass.True(ok)
	ass.Equal(int64(3), tree.Len())
	ass.NotNil(found)
	_, ok = tree.Search(rbtree.Int(3))
	ass.True(ok)
	_, ok = tree.Search(rbtree.Int(4))
	ass.True(ok)
}

func Test_ReplaceOrInsertMaxTreeDuplicateGreaterThenPrevNodes_PrevNodesDeletedLenLessMaxLen(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := NewMaxTree(3)

	tree.Insert(rbtree.Int(1))
	tree.Insert(rbtree.Int(2))
	tree.Insert(rbtree.Int(3))

	duplicate := rbtree.Int(4)

	// Act
	for i := 0; i <= 5; i++ {
		tree.ReplaceOrInsert(duplicate)
	}

	// Assert
	found, ok := tree.Search(duplicate)
	ass.True(ok)
	ass.Equal(int64(2), tree.Len())
	ass.NotNil(found)
	_, ok = tree.Search(rbtree.Int(3))
	ass.True(ok)
	_, ok = tree.Search(rbtree.Int(4))
	ass.True(ok)
}

func Test_ReplaceOrInsertMaxTreeDuplicateGreaterThenPrevNodesAndInsertLessNodesAfterwards_PrevNodesDeletedLenEqualMaxLen(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := NewMaxTree(3)

	tree.Insert(rbtree.Int(1))
	tree.Insert(rbtree.Int(2))
	tree.Insert(rbtree.Int(3))

	duplicate := rbtree.Int(4)

	// Act
	for i := 0; i <= 5; i++ {
		tree.ReplaceOrInsert(duplicate)
	}
	tree.ReplaceOrInsert(rbtree.Int(0))
	tree.ReplaceOrInsert(rbtree.Int(1))

	// Assert
	found, ok := tree.Search(duplicate)
	ass.True(ok)
	ass.Equal(int64(3), tree.Len())
	ass.NotNil(found)
	_, ok = tree.Search(rbtree.Int(3))
	ass.True(ok)
	_, ok = tree.Search(rbtree.Int(4))
	ass.True(ok)
	_, ok = tree.Search(rbtree.Int(0))
	ass.False(ok)
	_, ok = tree.Search(rbtree.Int(1))
	ass.True(ok)
}

func Test_DeleteAllNodesNodes_Success(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	minTree := NewMinTree(3)
	maxTree := NewMaxTree(3)

	for i := 1; i <= 10; i++ {
		minTree.Insert(rbtree.Int(i))
		maxTree.Insert(rbtree.Int(i))
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
			v := rbtree.Int(test.expected)

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

func Test_FexedTree_Foreach(t *testing.T) {
	minTree := NewMinTree(3)
	maxTree := NewMaxTree(3)

	for i := 1; i <= 10; i++ {
		minTree.Insert(rbtree.Int(i))
		maxTree.Insert(rbtree.Int(i))
	}

	var tests = []struct {
		name     string
		it       rbtree.Enumerable
		expected []int
	}{
		{"ascend normal min", rbtree.NewAscend(minTree), []int{1, 2, 3}},
		{"descend normal min", rbtree.NewDescend(minTree), []int{3, 2, 1}},
		{"inorder normal min", rbtree.NewWalkInorder(minTree), []int{1, 2, 3}},
		{"preorder normal min", rbtree.NewWalkPreorder(minTree), []int{2, 1, 3}},
		{"postorder normal min", rbtree.NewWalkPostorder(minTree), []int{1, 3, 2}},

		{"ascend normal max", rbtree.NewAscend(maxTree), []int{8, 9, 10}},
		{"descend normal max", rbtree.NewDescend(maxTree), []int{10, 9, 8}},
		{"inorder normal max", rbtree.NewWalkInorder(maxTree), []int{8, 9, 10}},
		{"preorder normal max", rbtree.NewWalkPreorder(maxTree), []int{9, 8, 10}},
		{"postorder normal max", rbtree.NewWalkPostorder(maxTree), []int{8, 10, 9}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			ass := assert.New(t)
			result := make([]int, 0)

			// Act
			test.it.Foreach(func(n rbtree.Comparable) {
				result = append(result, rbtree.GetInt(n))
			})

			// Assert
			ass.Equal(test.expected, result)
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
