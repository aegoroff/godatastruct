package rbtree

// This file contains all RB tree search methods implementations

// Search searches value specified within search tree
func (tree *rbTree) Search(value Comparable) (Node, bool) {
	if tree.root.isNil() {
		return nil, false
	}
	n, ok := tree.root.search(value)
	if !ok {
		return nil, ok
	}
	return n, ok
}

func (n *node) search(value Comparable) (*node, bool) {
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
	if tree.root.isNil() {
		return nil
	}
	return tree.root.minimum()
}

func (n *node) minimum() *node {
	x := n
	for !x.isNil() && !x.left.isNil() {
		x = x.left
	}
	return x
}

// Maximum gets tree's max element
func (tree *rbTree) Maximum() Node {
	if tree.root.isNil() {
		return nil
	}
	return tree.root.maximum()
}

func (n *node) maximum() *node {
	x := n
	for !x.isNil() && !x.right.isNil() {
		x = x.right
	}
	return x
}

// Successor gets Node's successor
func (n *node) Successor() Node {
	if n == nil {
		return n
	}
	return n.successor()
}

func (n *node) successor() *node {
	x := n
	if !x.right.isNil() {
		return x.right.minimum()
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

// Predecessor gets Node's predecessor
func (n *node) Predecessor() Node {
	if n == nil {
		return n
	}
	return n.predecessor()
}

func (n *node) predecessor() *node {
	x := n
	if !x.left.isNil() {
		return x.left.maximum()
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
// IMPORTANT: numeration starts from 1 not from 0
func (tree *rbTree) OrderStatisticSelect(i int64) (Node, bool) {
	if tree.root.isNil() {
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
