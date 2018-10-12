package rbtree

// This file contains all RB tree iteration methods implementations

// WalkInorder walks subtree inorder (left, node, right)
func (n *Node) WalkInorder(action func(*Node)) {
	if n != nil && n.Key != nil {
		n.left.WalkInorder(action)
		action(n)
		n.right.WalkInorder(action)
	}
}

// WalkInorder walks tree inorder (left, node, right)
func (tree *RbTree) WalkInorder(action func(*Node)) {
	tree.Root.WalkInorder(action)
}

// WalkPreorder walks subtree preorder (node, left, right)
func (n *Node) WalkPreorder(action func(*Node)) {
	if n != nil && n.Key != nil {
		action(n)
		n.left.WalkPreorder(action)
		n.right.WalkPreorder(action)
	}
}

// WalkPostorder walks tree postorder (left, right, node)
func (tree *RbTree) WalkPostorder(action func(*Node)) {
	tree.Root.WalkPostorder(action)
}

// WalkPostorder walks subtree postorder (left, right, node)
func (n *Node) WalkPostorder(action func(*Node)) {
	if n != nil && n.Key != nil {
		n.left.WalkPostorder(action)
		n.right.WalkPostorder(action)
		action(n)
	}
}

// WalkPreorder walks tree preorder (node, left, right)
func (tree *RbTree) WalkPreorder(action func(*Node)) {
	tree.Root.WalkPreorder(action)
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

// AscendRange calls the iterator for every value in the tree within the range
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

// DescendRange calls the iterator for every value in the tree within the range
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
