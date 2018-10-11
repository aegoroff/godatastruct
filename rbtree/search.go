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
			x = x.Left
		} else {
			x = x.Right
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
	for x.Left != nil && x.Left.Key != nil {
		x = x.Left
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
	for x.Right != nil && x.Right.Key != nil {
		x = x.Right
	}
	return x
}

// Successor gets node specified successor
func (n *Node) Successor() *Node {
	if n.Right != nil && n.Right.Key != nil {
		return n.Right.Minimum()
	}

	y := n.Parent
	for y != nil && y.Key != nil && n == y.Right {
		n = y
		y = y.Parent
	}

	if y.Key == nil {
		return nil
	}

	return y
}

// Predecessor gets node specified predecessor
func (n *Node) Predecessor() *Node {
	if n.Left != nil && n.Left.Key != nil {
		return n.Left.Maximum()
	}

	y := n.Parent
	for y != nil && y.Key != nil && n == y.Left {
		n = y
		y = y.Parent
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
	r := n.Left.Size + 1
	if i == r {
		return n
	} else if i < r {
		return n.Left.OrderStatisticSelect(i)
	} else {
		return n.Right.OrderStatisticSelect(i - r)
	}
}
