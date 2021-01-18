package rbtree

// This file contains all RB tree search methods implementations

// Search searches value specified within search tree
func (tree *rbTree) Search(value Comparable) (Comparable, bool) {
	n, ok := tree.SearchNode(value)
	if !ok {
		return nil, ok
	}
	return n.key, ok
}

func (tree *rbTree) Floor(value Comparable) (Comparable, bool) {
	if tree.root.isNil() {
		return nil, false
	}
	n, ok := tree.root.floor(value)
	if !ok {
		return nil, ok
	}
	return n.key, ok
}

func (tree *rbTree) Ceiling(value Comparable) (Comparable, bool) {
	if tree.root.isNil() {
		return nil, false
	}
	n, ok := tree.root.ceiling(value)
	if !ok {
		return nil, ok
	}
	return n.key, ok
}

func (tree *rbTree) SearchAll(value Comparable) []Comparable {
	var result []Comparable
	n, ok := tree.SearchNode(value)
	if ok {
		result = append(result, n.key)
		s := n.Successor()
		for s.isNotNil() && s.key.Equal(value) {
			result = append(result, s.key)
			s = s.Successor()
		}
	}
	return result
}

// SearchNode searches *Node which key is equals value specified
func (tree *rbTree) SearchNode(value Comparable) (*Node, bool) {
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
	for x.isNotNil() && !value.Equal(x.key) {
		if value.Less(x.key) {
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

func (n *Node) floor(value Comparable) (*Node, bool) {
	if value == nil {
		return nil, false
	}
	var min *Node
	var x *Node
	x = n
	for x.isNotNil() && !value.Equal(x.key) {
		if value.Less(x.key) {
			if min.isNil() && x.left.isNil() {
				min = x
			}
			x = x.left
		} else {
			min = x
			x = x.right
		}
	}

	if x.isNotNil() {
		return x, true
	}

	return min, true
}

func (n *Node) ceiling(value Comparable) (*Node, bool) {
	if value == nil {
		return nil, false
	}
	var max *Node
	var x *Node
	x = n
	for x.isNotNil() && !value.Equal(x.key) {
		if value.Less(x.key) {
			max = x
			x = x.left
		} else {
			if max.isNil() && x.right.isNil() {
				max = x
			}
			x = x.right
		}
	}

	if x.isNotNil() {
		return x, true
	}

	return max, true
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

	x := tree.root
	r := x.left.size + 1

	for i != r {
		if i < r {
			x = x.left
		} else {
			i = i - r
			x = x.right
		}
		if x.left == nil {
			return nil, false
		}
		r = x.left.size + 1
	}
	return x, true
}
