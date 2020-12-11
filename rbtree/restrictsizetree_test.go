package rbtree

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func Test_RestrictedSizeTree_SizeAsExpectedIterationWithoutSideEffects(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	topTree := NewRbTree()

	var nodes []int
	var result []string
	const nodesCount = 200

	for i := 1; i <= nodesCount; i++ {
		nodes = append(nodes, i)
	}
	tree := createIntTree(nodes)
	top := int64(5)

	// Act
	NewWalkInorder(tree).Foreach(func(n Node) bool {
		insertTo(topTree, top, n)
		return true
	})

	iterationCount := int64(0)
	NewDescend(topTree).Foreach(func(n Node) bool {
		iterationCount++
		result = append(result, n.String())
		return true
	})

	// Assert
	ass.Equal(top, topTree.Len())
	ass.Equal(top, iterationCount)
	ass.Equal([]string{"200", "199", "198", "197", "196"}, result)
}

func TestRestrictedSizeRandomTree_SizeAsExpectedIterationWithoutSideEffects(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	topTree := newRbTree()

	var nodes []string
	var result []string
	const nodesCount = 500

	for i := 1; i <= nodesCount; i++ {
		l := 1 + rand.Intn(50)
		nodes = append(nodes, randomString(l))
	}
	tree := NewRbTree()
	for _, n := range nodes {
		c := NewString(n)
		tree.Insert(c)
	}

	top := int64(10)

	// Act
	NewWalkInorder(tree).Foreach(func(n Node) bool {
		insertTo(topTree, top, n)
		return true
	})

	NewDescend(topTree).Foreach(func(n Node) bool {
		result = append(result, n.String())
		return true
	})

	// Assert
	max := topTree.Maximum()
	pred1 := max.Predecessor()
	pred2 := pred1.Predecessor()

	ass.Equal(top, topTree.Len())
	ass.Equal(max.String(), result[0])
	ass.Equal(pred1.String(), result[1])
	ass.Equal(pred2.String(), result[2])
	ass.Equal(top, int64(len(result)))
	ass.Equal(int64(nodesCount), tree.Len())
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

func insertTo(tree RbTree, size int64, c Comparable) {
	if tree.Len() < size {
		tree.Insert(c)
		return
	}

	min := tree.Minimum()

	k := min

	if k.LessThan(c) {
		tree.DeleteNode(k)
		tree.Insert(c)
	}
}
