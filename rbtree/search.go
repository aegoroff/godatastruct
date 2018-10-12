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
	if tree.Root == nil {
		return nil
	}
	return tree.Root.Minimum()
}

// Minimum gets tree's min element
func (n *Node) Minimum() *Node {
	x := n
	for x != nil && x.left != nil && x.left.Key != nil {
		x = x.left
	}
	return x
}

// Maximum gets tree's max element
func (tree *RbTree) Maximum() *Node {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.Maximum()
}

// Maximum gets tree's max element
func (n *Node) Maximum() *Node {
	x := n
	for x != nil && x.right != nil && x.right.Key != nil {
		x = x.right
	}
	return x
}

// Successor gets node specified successor
func (n *Node) Successor() *Node {
	if n != nil && n.right != nil && n.right.Key != nil {
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
	if n != nil && n.left != nil && n.left.Key != nil {
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

// Ascend calls the iterator for every value in the tree until iterator returns false.
func (tree *RbTree) Ascend(iterator KeyIterator) {
	min := tree.Minimum()
	if min == nil {
		return
	}
	max := tree.Maximum()
	tree.AscendRange(min.Key, max.Key, iterator)
}

// Ascend calls the iterator for every value in the tree within the range
// [from, to], until iterator returns false.
func (tree *RbTree) AscendRange(from, to *Comparable, iterator KeyIterator) {
	if tree.Root == nil || tree.Root.Key == nil || from == nil || to == nil {
		return
	}
	tree.Root.ascend(from, to, iterator)
}

// Descend calls the iterator for every value in the tree until iterator returns false.
func (tree *RbTree) Descend(iterator KeyIterator) {
	min := tree.Minimum()
	if min == nil {
		return
	}
	max := tree.Maximum()
	tree.DescendRange(max.Key, min.Key, iterator)
}

// Descend calls the iterator for every value in the tree within the range
// [from, to], until iterator returns false.
func (tree *RbTree) DescendRange(from, to *Comparable, iterator KeyIterator) {
	if tree.Root == nil || tree.Root.Key == nil || from == nil || to == nil {
		return
	}
	tree.Root.descend(from, to, iterator)
}

func (n *Node) ascend(from, to *Comparable, iterator KeyIterator) {
	curr, ok := n.Search(from)
	for ok && curr != nil && curr.Key != nil && ((*curr.Key).LessThan(*to) || (*curr.Key).EqualTo(*to)) {
		ok = iterator(curr.Key)
		if ok {
			curr = curr.Successor()
		}
	}
}

func (n *Node) descend(from, to *Comparable, iterator KeyIterator) {
	curr, ok := n.Search(from)
	for ok && curr != nil && curr.Key != nil && (!(*curr.Key).LessThan(*to) || (*curr.Key).EqualTo(*to)) {
		ok = iterator(curr.Key)
		if ok {
			curr = curr.Predecessor()
		}
	}
}
