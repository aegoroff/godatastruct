package rbtree

// This file contains all RB tree search methods implementations

// Search searches value specified within search tree
func (tree *rbTree) Search(value Comparable) (*Node, bool) {
	if tree.root.isNil() {
		return nil, false
	}
	n, ok := tree.root.search(value)
	if !ok {
		return nil, ok
	}
	return n, ok
}

func (n *Node) search(value Comparable) (*Node, bool) {
	if value == nil {
		return nil, false
	}
	var x *Node
	x = n
	for x.isNotNil() && !value.EqualTo(x.key) {
		if value.LessThan(x.key) {
			x = x.left
		} else {
			x = x.right
		}
	}

	if x.isNil() {
		return nil, false
	}

	return x, true
}

// Minimum gets tree's min element
func (tree *rbTree) Minimum() *Node {
	if tree.root.isNil() {
		return nil
	}
	return tree.root.minimum()
}

func (n *Node) minimum() *Node {
	x := n
	for x.isNotNil() && x.left.isNotNil() {
		x = x.left
	}
	return x
}

// Maximum gets tree's max element
func (tree *rbTree) Maximum() *Node {
	if tree.root.isNil() {
		return nil
	}
	return tree.root.maximum()
}

func (n *Node) maximum() *Node {
	x := n
	for x.isNotNil() && x.right.isNotNil() {
		x = x.right
	}
	return x
}

// Successor gets Node's successor
func (n *Node) Successor() *Node {
	if n.isNil() {
		return nil
	}

	x := n
	if x.right.isNotNil() {
		return x.right.minimum()
	}

	y := x.parent
	for y.isNotNil() && x == y.right {
		x = y
		y = y.parent
	}

	if y.isNil() {
		return nil
	}

	return y
}

// Predecessor gets Node's predecessor
func (n *Node) Predecessor() *Node {
	if n.isNil() {
		return nil
	}

	x := n
	if x.left.isNotNil() {
		return x.left.maximum()
	}

	y := x.parent
	for y.isNotNil() && x == y.left {
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
func (tree *rbTree) OrderStatisticSelect(i int64) (*Node, bool) {
	if tree.root.isNil() {
		return nil, false
	}

	return tree.root.orderStatisticSelect(i)
}

func (n *Node) orderStatisticSelect(i int64) (*Node, bool) {
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
