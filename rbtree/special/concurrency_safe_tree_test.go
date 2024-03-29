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
			tree.Insert(rbtree.Int(ix))
		}(i)
	}
	wg.Wait()

	// Assert
	ass.Equal(int64(nodesCount), tree.Len())
}

func Test_WrapToConcurrencySafe_InsertTest(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	result := make([]int, 0)
	var wg sync.WaitGroup

	const nodesCount = 200
	top := int64(4)
	mt := NewMaxTree(top)
	tree := WrapToConcurrencySafe(mt)

	// Act
	for i := 1; i <= nodesCount; i++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			tree.Insert(rbtree.Int(ix))
		}(i)
	}
	wg.Wait()

	// Assert
	ass.Equal(top, tree.Len())
	rbtree.NewDescend(tree).Foreach(func(n rbtree.Comparable) {
		result = append(result, rbtree.GetInt(n))
	})
	ass.Equal([]int{200, 199, 198, 197}, result)
}

func Test_WrapToConcurrencySafe_ReplaceOrInsertTest(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	result := make([]int, 0)
	var wg sync.WaitGroup

	const nodesCount = 200
	top := int64(4)
	mt := NewMaxTree(top)
	tree := WrapToConcurrencySafe(mt)

	// Act
	for i := 1; i <= nodesCount; i++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			tree.ReplaceOrInsert(rbtree.Int(ix))
		}(i)
	}
	wg.Wait()

	// Assert
	ass.Equal(top, tree.Len())
	rbtree.NewDescend(tree).Foreach(func(n rbtree.Comparable) {
		result = append(result, rbtree.GetInt(n))
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
		tree.Insert(rbtree.Int(i))
	}

	// Act
	for i := 1; i <= nodesCount; i++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			tree.Delete(rbtree.Int(ix))
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
		tree.Insert(rbtree.Int(i))
	}

	// Act
	for i := 1; i <= nodesCount; i++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			tree.DeleteAll(rbtree.Int(ix))
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
		tree.Insert(rbtree.Int(i))
	}

	// Act
	for i := 1; i <= nodesCount/2; i++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			tree.DeleteAll(rbtree.Int(ix))
		}(i)

		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			_, ok := tree.Search(rbtree.Int(nodesCount/2 + ix))
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

func Test_ConcurrencySafeTree_ConcurrentModificationAndFloorTest(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var wg sync.WaitGroup

	const nodesCount = 200
	tree := NewConcurrencySafeTree()
	readResultsChan := make(chan bool, nodesCount/2)

	for i := 1; i <= nodesCount; i++ {
		tree.Insert(rbtree.Int(i))
	}

	// Act
	for i := 1; i <= nodesCount/2; i++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			tree.DeleteAll(rbtree.Int(ix))
		}(i)

		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			_, ok := tree.Floor(rbtree.Int(nodesCount/2 + ix))
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

func Test_ConcurrencySafeTree_ConcurrentModificationAndCeilingTest(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var wg sync.WaitGroup

	const nodesCount = 200
	tree := NewConcurrencySafeTree()
	readResultsChan := make(chan bool, nodesCount/2)

	for i := 1; i <= nodesCount; i++ {
		tree.Insert(rbtree.Int(i))
	}

	// Act
	for i := 1; i <= nodesCount/2; i++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			tree.DeleteAll(rbtree.Int(ix))
		}(i)

		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			_, ok := tree.Ceiling(rbtree.Int(nodesCount/2 + ix))
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

func Test_ConcurrencySafeTree_ConcurrentModificationAndSearchAllTest(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var wg sync.WaitGroup

	const nodesCount = 200
	tree := NewConcurrencySafeTree()
	readResultsChan := make(chan bool, nodesCount/2)

	for i := 1; i <= nodesCount; i++ {
		tree.Insert(rbtree.Int(i))
	}

	// Act
	for i := 1; i <= nodesCount/2; i++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			tree.DeleteAll(rbtree.Int(ix))
		}(i)

		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			r := tree.SearchAll(rbtree.Int(nodesCount/2 + ix))
			readResultsChan <- len(r) > 0
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

func Test_ConcurrencySafeTree_ConcurrentModificationAndSearchNodeTest(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var wg sync.WaitGroup

	const nodesCount = 50
	tree := NewConcurrencySafeTree()
	readResultsChan := make(chan bool, nodesCount/2)

	for i := 1; i <= nodesCount; i++ {
		tree.Insert(rbtree.Int(i))
	}

	// Act
	for i := 1; i <= nodesCount/2; i++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			tree.DeleteAll(rbtree.Int(ix))
		}(i)

		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			_, ok := tree.SearchNode(rbtree.Int(nodesCount/2 + ix))
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

	const nodesCount = 50
	tree := NewConcurrencySafeTree()
	readResultsChan := make(chan *rbtree.Node, nodesCount/2)

	for i := 1; i <= nodesCount; i++ {
		tree.Insert(rbtree.Int(i))
	}

	// Act
	for i := 1; i <= nodesCount/2; i++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			tree.DeleteAll(rbtree.Int(ix))
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

	const nodesCount = 50
	tree := NewConcurrencySafeTree()
	readResultsChan := make(chan *rbtree.Node, nodesCount/2)

	for i := 1; i <= nodesCount; i++ {
		tree.Insert(rbtree.Int(i))
	}

	// Act
	for i := 1; i <= nodesCount/2; i++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			tree.DeleteAll(rbtree.Int(ix))
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

	const nodesCount = 50
	tree := NewConcurrencySafeTree()

	for i := 1; i <= nodesCount; i++ {
		tree.Insert(rbtree.Int(i))
	}
	var mu sync.Mutex
	res := true

	// Act
	wg.Add(nodesCount)
	for i := 1; i <= nodesCount/2; i++ {
		go func(ix int) {
			ix_ := ix
			tree.Delete(rbtree.Int(ix_))
			wg.Done()
		}(i)

		go func() {
			_, ok := tree.OrderStatisticSelect(1)
			mu.Lock()
			res = res && ok
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	// Assert
	ass.Equal(int64(nodesCount/2), tree.Len())
	ass.True(res)
}

func Test_ConcurrencySafeTree_Foreach(t *testing.T) {
	tree := NewConcurrencySafeTree()
	tree.Insert(rbtree.Int(6))
	tree.Insert(rbtree.Int(18))
	tree.Insert(rbtree.Int(3))

	var tests = []struct {
		name     string
		it       rbtree.Enumerable
		expected []int
	}{
		{"ascend normal", rbtree.NewAscend(tree), []int{3, 6, 18}},
		{"descend normal", rbtree.NewDescend(tree), []int{18, 6, 3}},
		{"inorder normal", rbtree.NewWalkInorder(tree), []int{3, 6, 18}},
		{"preorder normal", rbtree.NewWalkPreorder(tree), []int{6, 3, 18}},
		{"postorder normal", rbtree.NewWalkPostorder(tree), []int{3, 18, 6}},
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
