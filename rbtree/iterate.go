package rbtree

// This file contains all RB tree iteration methods implementations

type walkInorder struct {
	tree *rbTree
	curr *node
	next *node
}

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
func NewWalkInorder(t RbTree) Enumerable {
	tree := t.(*rbTree)
	next := tree.root

	e := &walkInorder{
		tree: tree,
		next: next,
	}

	if !next.isNil() {
		e.nextAsDeepestLeft()
	}

	return e
}

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
func (i *walkInorder) Foreach(callback NodeAction) {
	for i.Next() {
		callback(i.Current())
	}
}

func (i *walkInorder) Current() Node {
	return i.curr
}

func (i *walkInorder) Next() bool {
	p := i.next

	if !p.isNil() {
		if !i.next.right.isNil() {
			i.next = i.next.right
			i.nextAsDeepestLeft()
			i.curr = p
			return true
		}

		for true {
			if i.next.parent.isNil() {
				i.next = nil
				i.curr = p
				return true
			}

			if i.next.parent.left == i.next {
				i.next = i.next.parent
				i.curr = p
				return true
			}
			i.next = i.next.parent
		}
	}

	return false
}

func (i *walkInorder) nextAsDeepestLeft() {
	for !i.next.left.isNil() {
		i.next = i.next.left
	}
}

// Foreach does tree iteration and calls the callback for
// every value in the tree until callback returns false.
func (i *walkPreorder) Foreach(callback NodeAction) {
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
		callback(p)
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
func (i *walkPostorder) Foreach(callback NodeAction) {
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
			callback(next)
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
func (i *ascend) Foreach(callback NodeAction) {
	max := i.tree.Maximum()
	if max == nil {
		return
	}

	min := i.tree.root.minimum()
	min.ascend(max.Key(), callback)
}

// Foreach does tree iteration and calls the callback for
// every value in the tree until callback returns false.
func (i *ascendRange) Foreach(callback NodeAction) {
	if i.tree.root.isNil() || i.to == nil {
		return
	}
	curr, ok := i.tree.root.search(i.from)
	if ok {
		curr.ascend(i.to, callback)
	}
}

func (n *node) ascend(to Comparable, callback NodeAction) {
	curr := n
	for !curr.isNil() && (curr.key.LessThan(to) || curr.key.EqualTo(to)) {
		callback(curr)
		curr = curr.successor()
	}
}

// Foreach does tree iteration and calls the callback for
// every value in the tree until callback returns false.
func (i *descend) Foreach(callback NodeAction) {
	min := i.tree.Minimum()
	if min == nil {
		return
	}
	max := i.tree.root.maximum()
	max.descend(min.Key(), callback)
}

// Foreach does tree iteration and calls the callback for
// every value in the tree until callback returns false.
func (i *descendRange) Foreach(callback NodeAction) {
	if i.tree.root == nil || i.to == nil {
		return
	}
	curr, ok := i.tree.root.search(i.from)
	if ok {
		curr.descend(i.to, callback)
	}
}

func (n *node) descend(to Comparable, callback NodeAction) {
	curr := n
	for !curr.isNil() && !curr.key.LessThan(to) {
		callback(curr)
		curr = curr.predecessor()
	}
}
