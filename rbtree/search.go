package rbtree

// This file contains all RB tree search methods implementations

// Search searches value specified within search tree
func (tree *rbTree) Search(value Comparable) (Node, bool) {
	if tree.root == nil {
		return nil, false
	}
	n, ok := tree.search(tree.root, value)
	if !ok {
		return nil, ok
	}
	return n, ok
}

func (*rbTree) search(n *node, value Comparable) (*node, bool) {
	if value == nil {
		return nil, false
	}
	var x *node
	x = n
	for !x.isNil() && !value.EqualTo(x.key) {
		if value.LessThan(x.key) {
			x = x.left
		} else {
			x = x.right
		}
	}
	ok := !x.isNil()

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
	return tree.minimum(tree.root)
}

// Minimum gets tree's min element
func (*rbTree) minimum(n *node) *node {
	x := n
	for !x.isNil() && !x.left.isNil() {
		x = x.left
	}
	return x
}

// Maximum gets tree's max element
func (tree *rbTree) Maximum() Node {
	if tree.root == nil {
		return nil
	}
	return tree.maximum(tree.root)
}

// Maximum gets tree's max element
func (*rbTree) maximum(n *node) *node {
	x := n
	for !x.isNil() && !x.right.isNil() {
		x = x.right
	}
	return x
}

// Successor gets node specified successor
func (tree *rbTree) Successor(n Node) Node {
	// TODO: think over invalid casting here in case of custom Node implementation
	return tree.successor(n.(*node))
}

func (tree *rbTree) successor(n *node) *node {
	x := n
	if !x.right.isNil() {
		return tree.minimum(x.right)
	}

	y := x.parent
	for !y.isNil() && x == y.right {
		x = y
		y = y.parent
	}

	if y.isNil() {
		return nil
	}

	return y
}

// Predecessor gets node specified predecessor
func (tree *rbTree) Predecessor(n Node) Node {
	// TODO: think over invalid casting here in case of custom Node implementation
	return tree.predecessor(n.(*node))
}

func (tree *rbTree) predecessor(n *node) *node {
	x := n
	if !x.left.isNil() {
		return tree.maximum(x.left)
	}

	y := x.parent
	for !y.isNil() && x == y.left {
		x = y
		y = y.parent
	}

	if y.isNil() {
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
