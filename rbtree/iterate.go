package rbtree

// This file contains all RB tree iteration methods implementations

type enumerable struct {
	it Iterator
}

type walk struct {
	enumerable
	tree  *rbTree
	stack []*node
	curr  *node
}

type walkInorder struct {
	walk
	p *node
}

type walkPreorder struct {
	walk
}

type walkPostorder struct {
	walk
	p *node
}

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
	e := &walkInorder{
		walk: newWalk(t),
	}

	if len(e.stack) > 0 {
		e.p = e.stack[0]
	}

	e.it = e
	return e
}

// NewWalkPreorder creates Enumerable that walks tree preorder (node, left, right)
func NewWalkPreorder(t RbTree) Enumerable {
	e := &walkPreorder{
		walk: newWalk(t),
	}

	e.it = e
	return e
}

// NewWalkPostorder creates Enumerable that walks tree postorder (left, right, node)
func NewWalkPostorder(t RbTree) Enumerable {
	e := &walkPostorder{
		walk: newWalk(t),
	}

	if len(e.stack) > 0 {
		e.p = e.stack[0]
	}

	e.it = e
	return e
}

func newWalk(t RbTree) walk {
	tree := t.(*rbTree)

	w := walk{
		tree:  tree,
		stack: make([]*node, 0),
	}

	if !tree.root.isNil() {
		w.stack = append(w.stack, tree.root)
	}

	return w
}

// NewAscend creates Enumerable that walks tree in ascending order
func NewAscend(t RbTree) Enumerable {
	tree := t.(*rbTree)
	e := &ascend{
		tree: tree,
	}

	if !tree.root.isNil() {
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

	if !tree.root.isNil() && to != nil {
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

	if !tree.root.isNil() {
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

	if !tree.root.isNil() && to != nil {
		e.next, _ = tree.root.search(from)
		e.to = to
	}
	e.it = e
	return e
}

func (i *walkInorder) Current() Node { return i.curr }

func (i *walkInorder) Next() bool {
	for len(i.stack) > 0 {
		if !i.p.isNil() {
			i.p = i.p.left
			if !i.p.isNil() {
				i.stack = append(i.stack, i.p)
			}
		} else {
			top := len(i.stack) - 1
			i.p = i.stack[top]
			i.curr = i.p
			i.stack = i.stack[:top]
			i.p = i.p.right

			if !i.p.isNil() {
				i.stack = append(i.stack, i.p)
			}
			return true
		}
	}

	return false
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

func (i *walkPostorder) Current() Node { return i.curr }

func (i *walkPostorder) Next() bool {
	for len(i.stack) > 0 {
		top := len(i.stack) - 1
		next := i.stack[top]

		if next.right == i.p || next.left == i.p || (next.right.isNil() && next.left.isNil()) {
			i.stack = i.stack[:top]
			i.curr = next
			i.p = next
			return true
		} else {
			if !next.right.isNil() {
				i.stack = append(i.stack, next.right)
			}
			if !next.left.isNil() {
				i.stack = append(i.stack, next.left)
			}
		}
	}

	return false
}

func (i *ascend) Current() Node { return i.curr }

func (i *ascend) Next() bool {
	result := !i.next.isNil() && (i.next.key.LessThan(i.to) || i.next.key.EqualTo(i.to))
	if result {
		i.curr = i.next
		i.next = i.curr.successor()
	}
	return result
}

func (i *descend) Current() Node { return i.curr }

func (i *descend) Next() bool {
	result := !i.next.isNil() && !i.next.key.LessThan(i.to)
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
