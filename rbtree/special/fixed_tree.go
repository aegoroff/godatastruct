package special

import "github.com/aegoroff/godatastruct/rbtree"

// maxTree represents Red-black search binary tree
// that stores only limited size of max possible values
type maxTree struct {
	// tree contains underlying Red-black search binary tree
	tree rbtree.RbTree
	size int64
}

func (t *maxTree) Root() *rbtree.Node {
	return t.tree.Root()
}

func (t *maxTree) Len() int64 {
	return t.tree.Len()
}

// Insert inserts node into tree which size is limited
// Only <size> max nodes will be in the tree
func (t *maxTree) Insert(c rbtree.Comparable) {
	min := t.tree.Minimum()
	if t.tree.Len() < t.size || min.Key().Less(c) {
		if t.Len() == t.size {
			t.DeleteNode(min.Key())
		}

		t.tree.Insert(c)
	}
}

func (t *maxTree) ReplaceOrInsert(c rbtree.Comparable) rbtree.Comparable {
	min := t.tree.Minimum()
	if t.tree.Len() < t.size || min.Key().Less(c) {
		if t.Len() == t.size {
			t.DeleteNode(min.Key())
		}

		return t.tree.ReplaceOrInsert(c)
	}
	return nil
}

func (t *maxTree) DeleteNode(c rbtree.Comparable) bool {
	return t.tree.DeleteNode(c)
}

func (t *maxTree) DeleteAllNodes(c rbtree.Comparable) bool {
	return t.tree.DeleteAllNodes(c)
}

func (t *maxTree) Search(value rbtree.Comparable) (rbtree.Comparable, bool) {
	return t.tree.Search(value)
}

func (t *maxTree) SearchAll(value rbtree.Comparable) []rbtree.Comparable {
	return t.tree.SearchAll(value)
}

func (t *maxTree) SearchNode(value rbtree.Comparable) (*rbtree.Node, bool) {
	return t.tree.SearchNode(value)
}

func (t *maxTree) Minimum() *rbtree.Node {
	return t.tree.Minimum()
}

func (t *maxTree) Maximum() *rbtree.Node {
	return t.tree.Maximum()
}

func (t *maxTree) OrderStatisticSelect(i int64) (*rbtree.Node, bool) {
	return t.tree.OrderStatisticSelect(i)
}

// minTree represents Red-black search binary tree
// that stores only limited size of min possible values
type minTree struct {
	// tree contains underlying Red-black search binary tree
	tree rbtree.RbTree
	size int64
}

func (t *minTree) Root() *rbtree.Node {
	return t.tree.Root()
}

func (t *minTree) Len() int64 {
	return t.tree.Len()
}

// Insert inserts node into tree which size is limited
// Only <size> min nodes will be in the tree
func (t *minTree) Insert(c rbtree.Comparable) {
	max := t.tree.Maximum()
	if t.tree.Len() < t.size || !max.Key().Less(c) {
		if t.tree.Len() == t.size {
			t.tree.DeleteNode(max.Key())
		}

		t.tree.Insert(c)
	}
}

func (t *minTree) ReplaceOrInsert(c rbtree.Comparable) rbtree.Comparable {
	max := t.tree.Maximum()
	if t.tree.Len() < t.size || !max.Key().Less(c) {
		if t.tree.Len() == t.size {
			t.tree.DeleteNode(max.Key())
		}

		return t.tree.ReplaceOrInsert(c)
	}
	return nil
}

func (t *minTree) DeleteNode(c rbtree.Comparable) bool {
	return t.tree.DeleteNode(c)
}

func (t *minTree) DeleteAllNodes(c rbtree.Comparable) bool {
	return t.tree.DeleteAllNodes(c)
}

func (t *minTree) Search(value rbtree.Comparable) (rbtree.Comparable, bool) {
	return t.tree.Search(value)
}

func (t *minTree) SearchAll(value rbtree.Comparable) []rbtree.Comparable {
	return t.tree.SearchAll(value)
}

func (t *minTree) SearchNode(value rbtree.Comparable) (*rbtree.Node, bool) {
	return t.tree.SearchNode(value)
}

func (t *minTree) Minimum() *rbtree.Node {
	return t.tree.Minimum()
}

func (t *minTree) Maximum() *rbtree.Node {
	return t.tree.Maximum()
}

func (t *minTree) OrderStatisticSelect(i int64) (*rbtree.Node, bool) {
	return t.tree.OrderStatisticSelect(i)
}

// NewMaxTree creates new fixed size tree that stores <sz> max values
func NewMaxTree(sz int64) rbtree.RbTree {
	return &maxTree{
		tree: rbtree.New(),
		size: sz,
	}
}

// NewMinTree creates new fixed size tree that stores <sz> min values
func NewMinTree(sz int64) rbtree.RbTree {
	return &minTree{
		tree: rbtree.New(),
		size: sz,
	}
}
