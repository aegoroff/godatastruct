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
	tree.WalkInorder(func(n Node) {
		insertTo(topTree, top, n.Key())
	})

	iterationCount := int64(0)
	topTree.Descend(func(n Node) bool {
		iterationCount++
		result = append(result, n.Key().String())
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
	const nodesCount = 44

	for i := 1; i <= nodesCount; i++ {
		l := 1 + rand.Intn(5)
		nodes = append(nodes, randomString(l))
	}
	tree := createStringTree(nodes)
	top := int64(3)

	// Act
	tree.WalkInorder(func(n Node) {
		insertTo(topTree, top, n.Key())
	})

	iterationCount := int64(0)
	topTree.Descend(func(n Node) bool {
		iterationCount++
		result = append(result, n.Key().String())
		return true
	})

	// Assert
	max := topTree.Maximum()
	pred1 := topTree.Predecessor(max)
	pred2 := topTree.Predecessor(pred1)

	ass.Equal(top, topTree.Len())
	ass.Equal(max.Key().String(), result[0])
	ass.Equal(pred1.Key().String(), result[1])
	ass.Equal(pred2.Key().String(), result[2])
	ass.Equal(top, iterationCount)
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

	k := min.Key()

	if k.LessThan(c) {
		tree.DeleteAllNodes(k)
		tree.Insert(c)
	}
}
