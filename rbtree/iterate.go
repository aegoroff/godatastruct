package rbtree

// This file contains all RB tree iteration methods implementations

type walkInorder struct{ tree *rbTree }
type walkPreorder struct{ tree *rbTree }
type walkPostorder struct{ tree *rbTree }
type ascend struct{ tree *rbTree }
type ascendRange struct {
	tree *rbTree
	from Comparable
	to   Comparable
}
type descend struct{ tree *rbTree }
type descendRange struct {
	tree *rbTree
	from Comparable
	to   Comparable
}

// NewWalkInorder walks tree inorder (left, node, right)
func NewWalkInorder(t RbTree) Iterator { return &walkInorder{tree: t.(*rbTree)} }

// NewWalkPreorder walks tree preorder (node, left, right)
func NewWalkPreorder(t RbTree) Iterator { return &walkPreorder{tree: t.(*rbTree)} }

// NewWalkPostorder walks tree postorder (left, right, node)
func NewWalkPostorder(t RbTree) Iterator { return &walkPostorder{tree: t.(*rbTree)} }

// NewAscend calls the callback for every value in the tree until callback returns false.
func NewAscend(t RbTree) Iterator { return &ascend{tree: t.(*rbTree)} }

// NewAscendRange calls the callback for every value in the tree within the range
// [from, to], until callback returns false.
func NewAscendRange(t RbTree, from, to Comparable) Iterator {
	return &ascendRange{
		tree: t.(*rbTree),
		from: from,
		to:   to,
	}
}

// NewDescend calls the callback for every value in the tree until callback returns false.
func NewDescend(t RbTree) Iterator { return &descend{tree: t.(*rbTree)} }

// NewDescendRange calls the callback for every value in the tree within the range
// [from, to], until callback returns false.
func NewDescendRange(t RbTree, from, to Comparable) Iterator {
	return &descendRange{
		tree: t.(*rbTree),
		from: from,
		to:   to,
	}
}

// Iterate does tree iteration and calls the callback for
// every value in the tree until callback returns false.
func (i *walkInorder) Iterate(callback NodeValidator) {
	n := i.tree.root
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
			if !callback(p) {
				return
			}
			stack = stack[:top]
			p = p.right
		}

		if !p.isNil() {
			stack = append(stack, p)
		}
	}
}

// Iterate does tree iteration and calls the callback for
// every value in the tree until callback returns false.
func (i *walkPreorder) Iterate(callback NodeValidator) {
	n := i.tree.root
	if n.isNil() {
		return
	}

	var stack []*node
	p := n
	stack = append(stack, p)
	for len(stack) > 0 {
		top := len(stack) - 1
		p = stack[top]
		if !callback(p) {
			return
		}
		stack = stack[:top]

		if !p.right.isNil() {
			stack = append(stack, p.right)
		}

		if !p.left.isNil() {
			stack = append(stack, p.left)
		}
	}
}

// Iterate does tree iteration and calls the callback for
// every value in the tree until callback returns false.
func (i *walkPostorder) Iterate(callback NodeValidator) {
	i.tree.walkPostorder(i.tree.root, callback)
}

func (tree *rbTree) walkPostorder(n *node, callback NodeValidator) {
	if !n.isNil() {
		tree.walkPostorder(n.left, callback)
		tree.walkPostorder(n.right, callback)

		if !callback(n) {
			return
		}
	}
}

// Iterate does tree iteration and calls the callback for
// every value in the tree until callback returns false.
func (i *ascend) Iterate(callback NodeValidator) {
	max := i.tree.Maximum()
	if max == nil {
		return
	}

	min := i.tree.root.minimum()
	min.ascend(max.Key(), callback)
}

// Iterate does tree iteration and calls the callback for
// every value in the tree until callback returns false.
func (i *ascendRange) Iterate(callback NodeValidator) {
	if i.tree.root.isNil() || i.to == nil {
		return
	}
	curr, ok := i.tree.root.search(i.from)
	if ok {
		curr.ascend(i.to, callback)
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

// Iterate does tree iteration and calls the callback for
// every value in the tree until callback returns false.
func (i *descend) Iterate(callback NodeValidator) {
	min := i.tree.Minimum()
	if min == nil {
		return
	}
	max := i.tree.root.maximum()
	max.descend(min.Key(), callback)
}

// Iterate does tree iteration and calls the callback for
// every value in the tree until callback returns false.
func (i *descendRange) Iterate(callback NodeValidator) {
	if i.tree.root == nil || i.to == nil {
		return
	}
	curr, ok := i.tree.root.search(i.from)
	if ok {
		curr.descend(i.to, callback)
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
