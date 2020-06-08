package rbtree

// This file contains all RB tree iteration methods implementations

// WalkInorder walks tree inorder (left, node, right)
func (tree *rbTree) WalkInorder(action func(Node)) {
	n := tree.root
	if n.isNil() {
		return
	}

	var stack []*node
	p := n
	stack = append(stack, p)
	for len(stack) > 0 {
		if !p.isNil() {
			p = p.left
		} else {
			top := len(stack) - 1
			p = stack[top]
			action(p)
			stack = stack[:top]
			p = p.right
		}

		if !p.isNil() {
			stack = append(stack, p)
		}
	}
}

// WalkPreorder walks tree preorder (node, left, right)
func (tree *rbTree) WalkPreorder(action func(Node)) {
	n := tree.root
	if n.isNil() {
		return
	}

	var stack []*node
	p := n
	stack = append(stack, p)
	for len(stack) > 0 {
		top := len(stack) - 1
		p = stack[top]
		action(p)
		stack = stack[:top]

		if !p.right.isNil() {
			stack = append(stack, p.right)
		}

		if !p.left.isNil() {
			stack = append(stack, p.left)
		}
	}
}

// WalkPostorder walks tree postorder (left, right, node)
func (tree *rbTree) WalkPostorder(action func(Node)) {
	tree.walkPostorder(tree.root, func(n *node) { action(n) })
}

func (tree *rbTree) walkPostorder(n *node, action func(*node)) {
	if !n.isNil() {
		tree.walkPostorder(n.left, action)
		tree.walkPostorder(n.right, action)
		action(n)
	}
}

// Ascend calls the iterator for every value in the tree until iterator returns false.
func (tree *rbTree) Ascend(iterator NodeIterator) {
	max := tree.Maximum()
	if max == nil {
		return
	}

	min := tree.minimum(tree.root)
	tree.ascend(min, max.Key(), iterator)
}

// AscendRange calls the iterator for every value in the tree within the range
// [from, to], until iterator returns false.
func (tree *rbTree) AscendRange(from, to Comparable, iterator NodeIterator) {
	if tree.root.isNil() || to == nil {
		return
	}
	curr, ok := tree.search(tree.root, from)
	if ok {
		tree.ascend(curr, to, iterator)
	}
}

func (tree *rbTree) ascend(n *node, to Comparable, iterator NodeIterator) {
	curr := n
	ok := true
	for ok && !curr.isNil() && (curr.key.LessThan(to) || curr.key.EqualTo(to)) {
		ok = iterator(curr)
		if ok {
			curr = tree.successor(curr)
		}
	}
}

// Descend calls the iterator for every value in the tree until iterator returns false.
func (tree *rbTree) Descend(iterator NodeIterator) {
	min := tree.Minimum()
	if min == nil {
		return
	}
	max := tree.maximum(tree.root)
	tree.descend(max, min.Key(), iterator)
}

// DescendRange calls the iterator for every value in the tree within the range
// [from, to], until iterator returns false.
func (tree *rbTree) DescendRange(from, to Comparable, iterator NodeIterator) {
	if tree.root == nil || to == nil {
		return
	}
	curr, ok := tree.search(tree.root, from)
	if ok {
		tree.descend(curr, to, iterator)
	}
}

func (tree *rbTree) descend(n *node, to Comparable, iterator NodeIterator) {
	curr := n
	ok := true
	for ok && !curr.isNil() && !curr.key.LessThan(to) {
		ok = iterator(curr)
		if ok {
			curr = tree.predecessor(curr)
		}
	}
}
