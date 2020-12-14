package special

import (
	"github.com/aegoroff/godatastruct/rbtree"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func Test_ConcurrencySafeTree_InsertTest(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var wg sync.WaitGroup

	const nodesCount = 200
	tree := NewConcurrencySafeTree()

	// Act
	for i := 1; i <= nodesCount; i++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			tree.Insert(rbtree.NewInt(ix))
		}(i)
	}
	wg.Wait()

	// Assert
	ass.Equal(int64(nodesCount), tree.Len())
}

func Test_WrapTreeToConcurrencySafeTree_InsertTest(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	result := make([]int, 0)
	var wg sync.WaitGroup

	const nodesCount = 200
	top := int64(4)
	mt := NewMaxTree(top)
	tree := WrapTreeToConcurrencySafeTree(mt)

	// Act
	for i := 1; i <= nodesCount; i++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			tree.Insert(rbtree.NewInt(ix))
		}(i)
	}
	wg.Wait()

	// Assert
	ass.Equal(top, tree.Len())
	rbtree.NewDescend(tree).Foreach(func(n rbtree.Node) {
		result = append(result, rbtree.GetInt(n.Key()))
	})
	ass.Equal([]int{200, 199, 198, 197}, result)
}

func Test_ConcurrencySafeTree_DeleteNodeTest(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var wg sync.WaitGroup

	const nodesCount = 200
	tree := NewConcurrencySafeTree()

	for i := 1; i <= nodesCount; i++ {
		tree.Insert(rbtree.NewInt(i))
	}

	// Act
	for i := 1; i <= nodesCount; i++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			tree.DeleteNode(rbtree.NewInt(ix))
		}(i)
	}
	wg.Wait()

	// Assert
	ass.Equal(int64(0), tree.Len())
}

func Test_ConcurrencySafeTree_DeleteAllNodesTest(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var wg sync.WaitGroup

	const nodesCount = 200
	tree := NewConcurrencySafeTree()

	for i := 1; i <= nodesCount; i++ {
		tree.Insert(rbtree.NewInt(i))
	}

	// Act
	for i := 1; i <= nodesCount; i++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			tree.DeleteAllNodes(rbtree.NewInt(ix))
		}(i)
	}
	wg.Wait()

	// Assert
	ass.Equal(int64(0), tree.Len())
}

func Test_ConcurrencySafeTree_ConcurrentModificationAndSearchTest(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var wg sync.WaitGroup

	const nodesCount = 200
	tree := NewConcurrencySafeTree()
	readResultsChan := make(chan bool, nodesCount/2)

	for i := 1; i <= nodesCount; i++ {
		tree.Insert(rbtree.NewInt(i))
	}

	// Act
	for i := 1; i <= nodesCount/2; i++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			tree.DeleteAllNodes(rbtree.NewInt(ix))
		}(i)

		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			_, ok := tree.Search(rbtree.NewInt(nodesCount/2 + ix))
			readResultsChan <- ok
		}(i)
	}
	wg.Wait()
	close(readResultsChan)

	// Assert
	ass.Equal(int64(nodesCount/2), tree.Len())
	for ok := range readResultsChan {
		ass.True(ok)
	}
}

func Test_ConcurrencySafeTree_ConcurrentModificationAndMinimumTest(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var wg sync.WaitGroup

	const nodesCount = 200
	tree := NewConcurrencySafeTree()
	readResultsChan := make(chan rbtree.Node, nodesCount/2)

	for i := 1; i <= nodesCount; i++ {
		tree.Insert(rbtree.NewInt(i))
	}

	// Act
	for i := 1; i <= nodesCount/2; i++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			tree.DeleteAllNodes(rbtree.NewInt(ix))
		}(i)

		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			n := tree.Minimum()
			readResultsChan <- n
		}(i)
	}
	wg.Wait()
	close(readResultsChan)

	// Assert
	ass.Equal(int64(nodesCount/2), tree.Len())
	for n := range readResultsChan {
		ass.NotNil(n)
	}
}

func Test_ConcurrencySafeTree_ConcurrentModificationAndMaximumTest(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var wg sync.WaitGroup

	const nodesCount = 200
	tree := NewConcurrencySafeTree()
	readResultsChan := make(chan rbtree.Node, nodesCount/2)

	for i := 1; i <= nodesCount; i++ {
		tree.Insert(rbtree.NewInt(i))
	}

	// Act
	for i := 1; i <= nodesCount/2; i++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			tree.DeleteAllNodes(rbtree.NewInt(ix))
		}(i)

		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			n := tree.Maximum()
			readResultsChan <- n
		}(i)
	}
	wg.Wait()
	close(readResultsChan)

	// Assert
	ass.Equal(int64(nodesCount/2), tree.Len())
	for n := range readResultsChan {
		ass.NotNil(n)
		ass.Equal(nodesCount, rbtree.GetInt(n.Key()))
	}
}

func Test_ConcurrencySafeTree_ConcurrentModificationAndOrderStatisticSelectTest(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var wg sync.WaitGroup

	const nodesCount = 200
	tree := NewConcurrencySafeTree()
	readResultsChan := make(chan bool, nodesCount/2)

	for i := 1; i <= nodesCount; i++ {
		tree.Insert(rbtree.NewInt(i))
	}

	// Act
	for i := 1; i <= nodesCount/2; i++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			tree.DeleteAllNodes(rbtree.NewInt(ix))
		}(i)

		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			_, ok := tree.OrderStatisticSelect(int64(nodesCount / 2))
			readResultsChan <- ok
		}(i)
	}
	wg.Wait()
	close(readResultsChan)

	// Assert
	ass.Equal(int64(nodesCount/2), tree.Len())
	for ok := range readResultsChan {
		ass.True(ok)
	}
}
