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

// Ascend calls the callback for every value in the tree until callback returns false.
func (tree *rbTree) Ascend(callback NodeValidator) {
	max := tree.Maximum()
	if max == nil {
		return
	}

	min := tree.root.minimum()
	min.ascend(max.Key(), callback)
}

// AscendRange calls the callback for every value in the tree within the range
// [from, to], until callback returns false.
func (tree *rbTree) AscendRange(from, to Comparable, callback NodeValidator) {
	if tree.root.isNil() || to == nil {
		return
	}
	curr, ok := tree.root.search(from)
	if ok {
		curr.ascend(to, callback)
	}
}

func (n *node) ascend(to Comparable, callback NodeValidator) {
	curr := n
	ok := true
	for ok && !curr.isNil() && (curr.key.LessThan(to) || curr.key.EqualTo(to)) {
		ok = callback(curr)
		if ok {
			curr = curr.successor()
		}
	}
}

// Descend calls the callback for every value in the tree until callback returns false.
func (tree *rbTree) Descend(callback NodeValidator) {
	min := tree.Minimum()
	if min == nil {
		return
	}
	max := tree.root.maximum()
	max.descend(min.Key(), callback)
}

// DescendRange calls the callback for every value in the tree within the range
// [from, to], until callback returns false.
func (tree *rbTree) DescendRange(from, to Comparable, callback NodeValidator) {
	if tree.root == nil || to == nil {
		return
	}
	curr, ok := tree.root.search(from)
	if ok {
		curr.descend(to, callback)
	}
}

func (n *node) descend(to Comparable, callback NodeValidator) {
	curr := n
	ok := true
	for ok && !curr.isNil() && !curr.key.LessThan(to) {
		ok = callback(curr)
		if ok {
			curr = curr.predecessor()
		}
	}
}
