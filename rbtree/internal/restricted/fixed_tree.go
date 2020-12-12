package restricted

import "github.com/aegoroff/godatastruct/rbtree"

type fixedTree struct {
	tree rbtree.RbTree
	size int64
}

func newFixedTree(sz int64) *fixedTree {
	return &fixedTree{
		tree: rbtree.NewRbTree(),
		size: sz,
	}
}

// insert inserts node into tree which size is limited
// Only <size> max nodes will be in the tree
func (t *fixedTree) insert(c rbtree.Comparable) {
	min := t.tree.Minimum()
	if t.tree.Len() < t.size || min.Key().LessThan(c) {
		if t.tree.Len() == t.size {
			t.tree.DeleteNode(min.Key())
		}

		t.tree.Insert(c)
	}
}