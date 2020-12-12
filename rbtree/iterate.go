package rbtree

// This file contains all RB tree iteration methods implementations

type enumerable struct {
	it Iterator
}

type walkInorder struct {
	enumerable
	tree *rbTree
	curr *node
	next *node
}

type walkPreorder struct {
	enumerable
	tree  *rbTree
	stack []*node
	curr  *node
}

type walkPostorder struct{ tree *rbTree }

type ascend struct {
	enumerable
	tree *rbTree
	curr *node
	next *node
	to   Comparable
}

type descend struct {
	enumerable
	tree *rbTree
	curr *node
	next *node
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

	e.it = e
	return e
}

// NewWalkPreorder creates Enumerable that walks tree preorder (node, left, right)
func NewWalkPreorder(t RbTree) Enumerable {
	tree := t.(*rbTree)

	e := &walkPreorder{
		tree:  tree,
		stack: make([]*node, 0),
	}

	if tree.root != nil {
		e.stack = append(e.stack, tree.root)
	}
	e.it = e
	return e
}

// NewWalkPostorder creates Enumerable that walks tree postorder (left, right, node)
func NewWalkPostorder(t RbTree) Enumerable { return &walkPostorder{tree: t.(*rbTree)} }

// NewAscend creates Enumerable that walks tree in ascending order
func NewAscend(t RbTree) Enumerable {
	tree := t.(*rbTree)
	e := &ascend{
		tree: tree,
	}

	if tree.root != nil {
		e.next = tree.root.minimum()
		e.to = t.Maximum().Key()
	}
	e.it = e
	return e
}

// NewAscendRange creates Enumerable that walks tree in ascending order within the range [from, to]
func NewAscendRange(t RbTree, from, to Comparable) Enumerable {
	tree := t.(*rbTree)
	e := &ascend{
		tree: tree,
	}

	if tree.root != nil {
		e.next, _ = tree.root.search(from)
		e.to = to
	}
	e.it = e
	return e
}

// NewDescend creates Enumerable that walks tree in descending order
func NewDescend(t RbTree) Enumerable {
	tree := t.(*rbTree)
	e := &descend{
		tree: tree,
	}

	if tree.root != nil {
		e.next = tree.root.maximum()
		e.to = t.Minimum().Key()
	}
	e.it = e
	return e
}

// NewDescendRange that walks tree in descending order within the range [from, to]
func NewDescendRange(t RbTree, from, to Comparable) Enumerable {
	tree := t.(*rbTree)
	e := &descend{
		tree: tree,
	}

	if tree.root != nil {
		e.next, _ = tree.root.search(from)
		e.to = to
	}
	e.it = e
	return e
}

func (i *walkInorder) Current() Node { return i.curr }

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

func (i *walkPreorder) Current() Node { return i.curr }

func (i *walkPreorder) Next() bool {
	if len(i.stack) > 0 {
		top := len(i.stack) - 1
		p := i.stack[top]
		i.curr = p
		i.stack = i.stack[:top]

		if !p.right.isNil() {
			i.stack = append(i.stack, p.right)
		}

		if !p.left.isNil() {
			i.stack = append(i.stack, p.left)
		}

		return true
	}

	return false
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

func (i *ascend) Current() Node { return i.curr }

func (i *ascend) Next() bool {
	result := !i.next.isNil() && i.to != nil && (i.next.key.LessThan(i.to) || i.next.key.EqualTo(i.to))
	if result {
		i.curr = i.next
		i.next = i.curr.successor()
	}
	return result
}

func (i *descend) Current() Node { return i.curr }

func (i *descend) Next() bool {
	result := !i.next.isNil() && i.to != nil && !i.next.key.LessThan(i.to)
	if result {
		i.curr = i.next
		i.next = i.curr.predecessor()
	}
	return result
}

// Foreach does tree iteration and calls the callback for
// every value in the tree until callback returns false.
func (e *enumerable) Foreach(callback NodeAction) {
	for e.it.Next() {
		callback(e.it.Current())
	}
}
