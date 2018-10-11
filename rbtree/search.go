// Package rbtree - this file contains all RB tree search methods implementations
package rbtree

// Search searches value specified within search tree
func (tree *RbTree) Search(value *Comparable) (*Node, bool) {
	return tree.Root.Search(value)
}

// Search searches value specified within search tree
func (root *Node) Search(value *Comparable) (*Node, bool) {
	var x *Node
	x = root
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
func (root *Node) Minimum() *Node {
	x := root
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
func (root *Node) Maximum() *Node {
	x := root
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
func (root *Node) OrderStatisticSelect(i int64) *Node {
	r := root.Left.Size + 1
	if i == r {
		return root
	} else if i < r {
		return root.Left.OrderStatisticSelect(i)
	} else {
		return root.Right.OrderStatisticSelect(i - r)
	}
}
