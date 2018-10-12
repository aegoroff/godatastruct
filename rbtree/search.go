package rbtree

// This file contains all RB tree search methods implementations

// Search searches value specified within search tree
func (tree *RbTree) Search(value *Comparable) (*Node, bool) {
	return tree.Root.Search(value)
}

// Search searches value specified within search tree
func (n *Node) Search(value *Comparable) (*Node, bool) {
	var x *Node
	x = n
	for x != nil && x.Key != nil && !(*value).EqualTo(*x.Key) {
		if (*value).LessThan(*x.Key) {
			x = x.left
		} else {
			x = x.right
		}
	}
	ok := x != nil && x.Key != nil

	if !ok {
		return nil, ok
	}

	return x, ok
}

// Minimum gets tree's min element
func (tree *RbTree) Minimum() *Node {
	return tree.Root.Minimum()
}

// Minimum gets tree's min element
func (n *Node) Minimum() *Node {
	x := n
	for x.left != nil && x.left.Key != nil {
		x = x.left
	}
	return x
}

// Maximum gets tree's max element
func (tree *RbTree) Maximum() *Node {
	return tree.Root.Maximum()
}

// Maximum gets tree's max element
func (n *Node) Maximum() *Node {
	x := n
	for x.right != nil && x.right.Key != nil {
		x = x.right
	}
	return x
}

// Successor gets node specified successor
func (n *Node) Successor() *Node {
	if n.right != nil && n.right.Key != nil {
		return n.right.Minimum()
	}

	y := n.parent
	for y != nil && y.Key != nil && n == y.right {
		n = y
		y = y.parent
	}

	if y.Key == nil {
		return nil
	}

	return y
}

// Predecessor gets node specified predecessor
func (n *Node) Predecessor() *Node {
	if n.left != nil && n.left.Key != nil {
		return n.left.Maximum()
	}

	y := n.parent
	for y != nil && y.Key != nil && n == y.left {
		n = y
		y = y.parent
	}

	if y.Key == nil {
		return nil
	}

	return y
}

// OrderStatisticSelect gets i element from subtree
func (tree *RbTree) OrderStatisticSelect(i int64) *Node {
	return tree.Root.OrderStatisticSelect(i)
}

// OrderStatisticSelect gets i element from subtree
func (n *Node) OrderStatisticSelect(i int64) *Node {
	r := n.left.Size + 1
	if i == r {
		return n
	} else if i < r {
		return n.left.OrderStatisticSelect(i)
	} else {
		return n.right.OrderStatisticSelect(i - r)
	}
}
