package rbtree

// This file contains all RB tree search methods implementations

// Search searches value specified within search tree
func (tree *rbTree) Search(value Comparable) (Node, bool) {
	if tree.root == nil {
		return nil, false
	}
	n, ok := tree.root.search(value)
	if !ok {
		return nil, ok
	}
	return n, ok
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
	ok := x != nil && x.key != nil

	if !ok {
		return nil, ok
	}

	return x, ok
}

// Minimum gets tree's min element
func (tree *rbTree) Minimum() Node {
	if tree.root == nil {
		return nil
	}
	return tree.root.minimum()
}

// Minimum gets tree's min element
func (n *node) minimum() *node {
	x := n
	for x != nil && x.left != nil && x.left.key != nil {
		x = x.left
	}
	return x
}

// Maximum gets tree's max element
func (tree *rbTree) Maximum() Node {
	if tree.root == nil {
		return nil
	}
	return tree.root.maximum()
}

// Maximum gets tree's max element
func (n *node) maximum() *node {
	x := n
	for x != nil && x.right != nil && x.right.key != nil {
		x = x.right
	}
	return x
}

// Successor gets node specified successor
func (n *node) successor() *node {
	if n != nil && n.right != nil && n.right.key != nil {
		return n.right.minimum()
	}

	y := n.parent
	for y != nil && y.key != nil && n == y.right {
		n = y
		y = y.parent
	}

	if y == nil || y.key == nil {
		return nil
	}

	return y
}

// Predecessor gets node specified predecessor
func (n *node) predecessor() *node {
	if n != nil && n.left != nil && n.left.key != nil {
		return n.left.maximum()
	}

	y := n.parent
	for y != nil && y.key != nil && n == y.left {
		n = y
		y = y.parent
	}

	if y == nil || y.key == nil {
		return nil
	}

	return y
}

// OrderStatisticSelect gets i element from subtree
func (tree *rbTree) OrderStatisticSelect(i int64) (Node, bool) {
	if tree.root == nil {
		return nil, false
	}

	return tree.root.orderStatisticSelect(i)
}

func (n *node) orderStatisticSelect(i int64) (*node, bool) {
	if n.left == nil {
		return nil, false
	}
	r := n.left.size + 1
	if i == r {
		return n, true
	} else if i < r {
		return n.left.orderStatisticSelect(i)
	} else {
		return n.right.orderStatisticSelect(i - r)
	}
}
