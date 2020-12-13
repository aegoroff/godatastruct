// Package special contains specialized Red-black search binary tree implementations
package special

import "github.com/aegoroff/godatastruct/rbtree"

// MaxTree represents Red-black search binary tree
// that stores only limited size of max possible values
type MaxTree struct {
	// Tree contains underlying Red-black search binary tree
	Tree rbtree.RbTree
	size int64
}

// MinTree represents Red-black search binary tree
// that stores only limited size of min possible values
type MinTree struct {
	// Tree contains underlying Red-black search binary tree
	Tree rbtree.RbTree
	size int64
}

// NewMaxTree creates new fixed size tree that stores <sz> max values
func NewMaxTree(sz int64) *MaxTree {
	return &MaxTree{
		Tree: rbtree.NewRbTree(),
		size: sz,
	}
}

// NewMinTree creates new fixed size tree that stores <sz> min values
func NewMinTree(sz int64) *MinTree {
	return &MinTree{
		Tree: rbtree.NewRbTree(),
		size: sz,
	}
}

// Insert inserts node into tree which size is limited
// Only <size> max nodes will be in the tree
func (t *MaxTree) Insert(c rbtree.Comparable) {
	min := t.Tree.Minimum()
	if t.Tree.Len() < t.size || min.Key().LessThan(c) {
		if t.Tree.Len() == t.size {
			t.Tree.DeleteNode(min.Key())
		}

		t.Tree.Insert(c)
	}
}

// Insert inserts node into tree which size is limited
// Only <size> min nodes will be in the tree
func (t *MinTree) Insert(c rbtree.Comparable) {
	max := t.Tree.Maximum()
	if t.Tree.Len() < t.size || !max.Key().LessThan(c) {
		if t.Tree.Len() == t.size {
			t.Tree.DeleteNode(max.Key())
		}

		t.Tree.Insert(c)
	}
}
