package special

import (
	"fmt"
	"github.com/aegoroff/godatastruct/rbtree"
	"sync"
)

func ExampleNewConcurrencySafeTree() {
	tree := NewConcurrencySafeTree()

	const nodesCount = 10
	for i := 1; i <= nodesCount; i++ {
		tree.Insert(rbtree.Int(i))
	}

	var wg sync.WaitGroup
	for i := 1; i <= nodesCount/2; i++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			tree.DeleteNode(rbtree.Int(ix))
		}(i)

		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			_, ok := tree.Search(rbtree.Int(nodesCount/2 + ix))
			fmt.Println(ok)
		}(i)
	}
	wg.Wait()
	// Output:
	// true
	// true
	// true
	// true
	// true
}

func ExampleWrapTreeToConcurrencySafeTree() {
	tree := NewMaxTree(4)

	const nodesCount = 10
	for i := 1; i <= nodesCount; i++ {
		tree.Insert(rbtree.Int(i))
	}

	safeTree := WrapTreeToConcurrencySafeTree(tree)

	var wg sync.WaitGroup
	fixedLen := int(tree.Len())
	for i := 1; i <= fixedLen/2; i++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			safeTree.DeleteNode(rbtree.Int(nodesCount - ix))
		}(i)

		wg.Add(1)
		go func() {
			defer wg.Done()
			_, ok := safeTree.Search(rbtree.Int(nodesCount))
			fmt.Println(ok)
		}()
	}
	wg.Wait()
	// Output:
	// true
	// true
}
