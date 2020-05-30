package rbtree

// This file contains all RB tree search methods implementations

// Search searches value specified within search tree
func (tree *rbTree) Search(value Comparable) (Comparable, bool) {
	if tree.root == nil {
		return nil, false
	}
	n, ok := tree.root.search(value)
	if !ok {
		return nil, ok
	}
	return n.key, ok
}

// Search searches value specified within search tree
func (n *node) search(value Comparable) (*node, bool) {
	if value == nil {
		return nil, false
	}
	var x *node
	x = n
	for x != nil && x.key != nil && !value.EqualTo(x.key) {
		if value.LessThan(x.key) {
			x = x.left
		} else {
			x = x.right
		}
	}
	ok := x != nil

	if !ok {
		return nil, ok
	}

	return x, ok
}

// Minimum gets tree's min element
func (tree *rbTree) Minimum() Comparable {
	if tree.root == nil {
		return nil
	}
	return tree.root.minimum().key
}

// Minimum gets tree's min element
func (n *node) minimum() *node {
	x := n
	for x != nil && x.left != nil {
		x = x.left
	}
	return x
}

// Maximum gets tree's max element
func (tree *rbTree) Maximum() Comparable {
	if tree.root == nil {
		return nil
	}
	return tree.root.maximum().key
}

// Maximum gets tree's max element
func (n *node) maximum() *node {
	x := n
	for x != nil && x.right != nil {
		x = x.right
	}
	return x
}

// Successor gets node specified successor
func (n *node) successor() *node {
	if n != nil && n.right != nil {
		return n.right.minimum()
	}

	y := n.parent
	for y != nil && n == y.right {
		n = y
		y = y.parent
	}

	if y == nil {
		return nil
	}

	return y
}

// Predecessor gets node specified predecessor
func (n *node) predecessor() *node {
	if n != nil && n.left != nil {
		return n.left.maximum()
	}

	y := n.parent
	for y != nil && n == y.left {
		n = y
		y = y.parent
	}

	if y == nil {
		return nil
	}

	return y
}

// OrderStatisticSelect gets i element from subtree
func (tree *rbTree) OrderStatisticSelect(i int64) (Comparable, bool) {
	if tree.root == nil {
		return nil, false
	}

	return tree.root.orderStatisticSelect(i)
}

func (n *node) orderStatisticSelect(i int64) (Comparable, bool) {
	if n.left == nil {
		return nil, false
	}
	r := n.left.size + 1
	if i == r {
		return n.key, true
	} else if i < r {
		return n.left.orderStatisticSelect(i)
	} else {
		return n.right.orderStatisticSelect(i - r)
	}
}
