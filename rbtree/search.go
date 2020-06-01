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

func (tree *rbTree) search(n *node, value Comparable) (*node, bool) {
	if value == nil {
		return nil, false
	}
	var x *node
	x = n
	for x != nil && x.key != nil && x != tree.tnil && !value.EqualTo(x.key) {
		if value.LessThan(x.key) {
			x = x.left
		} else {
			x = x.right
		}
	}
	ok := x != nil && x.key != nil && x != tree.tnil

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
func (tree *rbTree) minimum(n *node) *node {
	x := n
	for x != nil && x.left != nil && x.left.key != nil && x.left != tree.tnil {
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
func (tree *rbTree) maximum(n *node) *node {
	x := n
	for x != nil && x.right != nil && x.right.key != nil && x.right != tree.tnil {
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
	if x.right != nil && x.right.key != nil && x.right != tree.tnil {
		return tree.minimum(x.right)
	}

	y := x.parent
	for y != nil && y.key != nil && x == y.right && y != tree.tnil {
		x = y
		y = y.parent
	}

	if y == nil || y.key == nil || y == tree.tnil {
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
	if x.left != nil && x.left.key != nil && x.left != tree.tnil {
		return tree.maximum(x.left)
	}

	y := x.parent
	for y != nil && y.key != nil && x == y.left && y != tree.tnil {
		x = y
		y = y.parent
	}

	if y == nil || y.key == nil || y == tree.tnil {
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
