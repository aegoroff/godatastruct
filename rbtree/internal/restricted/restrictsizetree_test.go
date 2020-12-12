package restricted

import (
	"github.com/aegoroff/godatastruct/rbtree"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func Test_RestrictedSizeTree_SizeAsExpectedIterationWithoutSideEffects(t *testing.T) {
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
	ft := newFixedTree(top)

	// Act
	rbtree.NewWalkInorder(tree).Foreach(func(n rbtree.Node) {
		ft.insert(n.Key())
	})

	iterationCount := int64(0)
	rbtree.NewDescend(ft.tree).Foreach(func(n rbtree.Node) {
		iterationCount++
		result = append(result, n.String())
	})

	// Assert
	ass.Equal(top, ft.tree.Len())
	ass.Equal(top, iterationCount)
	ass.Equal([]string{"200", "199", "198", "197", "196"}, result)
}

func TestRestrictedSizeRandomTree_SizeAsExpectedIterationWithoutSideEffects(t *testing.T) {
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
	ft := newFixedTree(ftSize)

	// Act
	rbtree.NewWalkInorder(tree).Foreach(func(n rbtree.Node) {
		ft.insert(n.Key())
	})

	rbtree.NewDescend(ft.tree).Foreach(func(n rbtree.Node) {
		result = append(result, n.String())
	})

	// Assert
	max := ft.tree.Maximum()
	pred1 := max.Predecessor()
	pred2 := pred1.Predecessor()

	ass.Equal(ftSize, ft.tree.Len())
	ass.Equal(max.String(), result[0])
	ass.Equal(pred1.String(), result[1])
	ass.Equal(pred2.String(), result[2])
	ass.Equal(ftSize, int64(len(result)))
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
