package special

import (
	"github.com/aegoroff/godatastruct/rbtree"
	"sync"
)

type concurrencySafeTree struct {
	// tree contains underlying Red-black search binary tree
	tree rbtree.RbTree
	mu   sync.RWMutex
}

func (t *concurrencySafeTree) Root() *rbtree.Node {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.Root()
}

func (t *concurrencySafeTree) Len() int64 {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.Len()
}

func (t *concurrencySafeTree) Insert(n rbtree.Comparable) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tree.Insert(n)
}

func (t *concurrencySafeTree) ReplaceOrInsert(n rbtree.Comparable) rbtree.Comparable {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.tree.ReplaceOrInsert(n)
}

func (t *concurrencySafeTree) DeleteNode(c rbtree.Comparable) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.tree.DeleteNode(c)
}

func (t *concurrencySafeTree) DeleteAllNodes(c rbtree.Comparable) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.tree.DeleteAllNodes(c)
}

func (t *concurrencySafeTree) Search(value rbtree.Comparable) (rbtree.Comparable, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.Search(value)
}

func (t *concurrencySafeTree) SearchAll(value rbtree.Comparable) []rbtree.Comparable {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.SearchAll(value)
}

func (t *concurrencySafeTree) SearchNode(value rbtree.Comparable) (*rbtree.Node, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.SearchNode(value)
}

func (t *concurrencySafeTree) Minimum() *rbtree.Node {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.Minimum()
}

func (t *concurrencySafeTree) Maximum() *rbtree.Node {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.Maximum()
}

func (t *concurrencySafeTree) OrderStatisticSelect(i int64) (*rbtree.Node, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	n, ok := t.tree.OrderStatisticSelect(i)
	return n, ok
}

// NewConcurrencySafeTree creates new concurrency safe tree that can be used in concurrency scenarios
func NewConcurrencySafeTree() rbtree.RbTree {
	return WrapTreeToConcurrencySafeTree(rbtree.NewRbTree())
}

// WrapTreeToConcurrencySafeTree creates new concurrency safe tree wrapper over existing rbtree.RbTree
// to use it safely in concurrency scenarios
func WrapTreeToConcurrencySafeTree(t rbtree.RbTree) rbtree.RbTree {
	return &concurrencySafeTree{tree: t}
}
