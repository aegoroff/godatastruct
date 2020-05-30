package rbtree

// This file contains all RB tree iteration methods implementations

// WalkInorder walks tree inorder (left, node, right)
func (tree *rbTree) WalkInorder(action func(Comparable)) {
	tree.root.walkInorder(action)
}

func (n *node) walkInorder(action func(Comparable)) {
	if n != nil && n.key != nil {
		n.left.walkInorder(action)
		action(n.key)
		n.right.walkInorder(action)
	}
}

// WalkPostorder walks tree postorder (left, right, node)
func (tree *rbTree) WalkPostorder(action func(Comparable)) {
	tree.root.walkPostorder(action)
}

func (n *node) walkPostorder(action func(Comparable)) {
	if n != nil && n.key != nil {
		n.left.walkPostorder(action)
		n.right.walkPostorder(action)
		action(n.key)
	}
}

// WalkPreorder walks tree preorder (node, left, right)
func (tree *rbTree) WalkPreorder(action func(Comparable)) {
	tree.root.walkPreorder(action)
}

func (n *node) walkPreorder(action func(Comparable)) {
	if n != nil && n.key != nil {
		action(n.key)
		n.left.walkPreorder(action)
		n.right.walkPreorder(action)
	}
}

// Ascend calls the iterator for every value in the tree until iterator returns false.
func (tree *rbTree) Ascend(iterator KeyIterator) {
	min := tree.Minimum()
	if min == nil {
		return
	}
	max := tree.Maximum()
	tree.AscendRange(min, max, iterator)
}

// AscendRange calls the iterator for every value in the tree within the range
// [from, to], until iterator returns false.
func (tree *rbTree) AscendRange(from, to Comparable, iterator KeyIterator) {
	if tree.root == nil || tree.root.key == nil || to == nil {
		return
	}
	tree.root.ascend(from, to, iterator)
}

// Descend calls the iterator for every value in the tree until iterator returns false.
func (tree *rbTree) Descend(iterator KeyIterator) {
	min := tree.Minimum()
	if min == nil {
		return
	}
	max := tree.Maximum()
	tree.DescendRange(max, min, iterator)
}

// DescendRange calls the iterator for every value in the tree within the range
// [from, to], until iterator returns false.
func (tree *rbTree) DescendRange(from, to Comparable, iterator KeyIterator) {
	if tree.root == nil || to == nil {
		return
	}
	tree.root.descend(from, to, iterator)
}

func (n *node) ascend(from, to Comparable, iterator KeyIterator) {
	curr, ok := n.search(from)
	for ok && curr != nil && curr.key != nil && (curr.key).LessThan(to) || curr != nil && curr.key != nil && (curr.key).EqualTo(to) {
		ok = iterator(curr.key)
		if ok {
			curr = curr.successor()
		}
	}
}

func (n *node) descend(from, to Comparable, iterator KeyIterator) {
	curr, ok := n.search(from)
	for ok && curr != nil && curr.key != nil && (!curr.key.LessThan(to) || curr.key != nil && curr.key.EqualTo(to)) {
		ok = iterator(curr.key)
		if ok {
			curr = curr.predecessor()
		}
	}
}
