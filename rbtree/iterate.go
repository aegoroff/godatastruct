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

// NewWalkInorder creates Enumerable that walks tree inorder (left, node, right)
func NewWalkInorder(t RbTree) Enumerable { return &walkInorder{tree: t.(*rbTree)} }

// NewWalkPreorder creates Enumerable that walks tree preorder (node, left, right)
func NewWalkPreorder(t RbTree) Enumerable { return &walkPreorder{tree: t.(*rbTree)} }

// NewWalkPostorder creates Enumerable that walks tree postorder (left, right, node)
func NewWalkPostorder(t RbTree) Enumerable { return &walkPostorder{tree: t.(*rbTree)} }

// NewAscend creates Enumerable that walks tree in ascending order
func NewAscend(t RbTree) Enumerable { return &ascend{tree: t.(*rbTree)} }

// NewAscendRange creates Enumerable that walks tree in ascending order within the range [from, to]
func NewAscendRange(t RbTree, from, to Comparable) Enumerable {
	return &ascendRange{
		tree: t.(*rbTree),
		from: from,
		to:   to,
	}
}

// NewDescend creates Enumerable that walks tree in descending order
func NewDescend(t RbTree) Enumerable { return &descend{tree: t.(*rbTree)} }

// NewDescendRange that walks tree in descending order within the range [from, to]
func NewDescendRange(t RbTree, from, to Comparable) Enumerable {
	return &descendRange{
		tree: t.(*rbTree),
		from: from,
		to:   to,
	}
}

// Foreach does tree iteration and calls the callback for
// every value in the tree until callback returns false.
func (i *walkInorder) Foreach(callback NodeEvaluator) {
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

// Foreach does tree iteration and calls the callback for
// every value in the tree until callback returns false.
func (i *walkPreorder) Foreach(callback NodeEvaluator) {
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

// Foreach does tree iteration and calls the callback for
// every value in the tree until callback returns false.
func (i *walkPostorder) Foreach(callback NodeEvaluator) {
	n := i.tree.root
	if n.isNil() {
		return
	}

	var stack []*node
	p := n
	stack = append(stack, p)

	for len(stack) > 0 {
		top := len(stack) - 1
		next := stack[top]

		if next.right == p || next.left == p || (next.right.isNil() && next.left.isNil()) {
			stack = stack[:top]
			if !callback(next) {
				return
			}
			p = next
		} else {
			if !next.right.isNil() {
				stack = append(stack, next.right)
			}
			if !next.left.isNil() {
				stack = append(stack, next.left)
			}
		}
	}
}

// Foreach does tree iteration and calls the callback for
// every value in the tree until callback returns false.
func (i *ascend) Foreach(callback NodeEvaluator) {
	max := i.tree.Maximum()
	if max == nil {
		return
	}

	min := i.tree.root.minimum()
	min.ascend(max.Key(), callback)
}

// Foreach does tree iteration and calls the callback for
// every value in the tree until callback returns false.
func (i *ascendRange) Foreach(callback NodeEvaluator) {
	if i.tree.root.isNil() || i.to == nil {
		return
	}
	curr, ok := i.tree.root.search(i.from)
	if ok {
		curr.ascend(i.to, callback)
	}
}

func (n *node) ascend(to Comparable, callback NodeEvaluator) {
	curr := n
	ok := true
	for ok && !curr.isNil() && (curr.key.LessThan(to) || curr.key.EqualTo(to)) {
		ok = callback(curr)
		if ok {
			curr = curr.successor()
		}
	}
}

// Foreach does tree iteration and calls the callback for
// every value in the tree until callback returns false.
func (i *descend) Foreach(callback NodeEvaluator) {
	min := i.tree.Minimum()
	if min == nil {
		return
	}
	max := i.tree.root.maximum()
	max.descend(min.Key(), callback)
}

// Foreach does tree iteration and calls the callback for
// every value in the tree until callback returns false.
func (i *descendRange) Foreach(callback NodeEvaluator) {
	if i.tree.root == nil || i.to == nil {
		return
	}
	curr, ok := i.tree.root.search(i.from)
	if ok {
		curr.descend(i.to, callback)
	}
}

func (n *node) descend(to Comparable, callback NodeEvaluator) {
	curr := n
	ok := true
	for ok && !curr.isNil() && !curr.key.LessThan(to) {
		ok = callback(curr)
		if ok {
			curr = curr.predecessor()
		}
	}
}
