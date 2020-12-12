package rbtree

// This file contains all RB tree iteration methods implementations

type enumerable struct {
	it Iterator
}

type iterator struct {
	enumerable
	tree *rbTree
	curr *node
}

type walk struct {
	iterator
	stack []*node
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
	iterator
	next *node
	to   Comparable
}

type descend struct {
	iterator
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

	it := iterator{tree: tree}

	w := walk{
		iterator: it,
		stack:    make([]*node, 0),
	}

	if !tree.root.isNil() {
		w.stack = append(w.stack, tree.root)
	}

	return w
}

// NewAscend creates Enumerable that walks tree in ascending order
func NewAscend(t RbTree) Enumerable {
	e := newAscend(t)

	if !e.tree.root.isNil() {
		e.next = e.tree.root.minimum()
		e.to = t.Maximum().Key()
	}

	return e
}

// NewAscendRange creates Enumerable that walks tree in ascending order within the range [from, to]
func NewAscendRange(t RbTree, from, to Comparable) Enumerable {
	e := newAscend(t)

	if !e.tree.root.isNil() && to != nil {
		e.next, _ = e.tree.root.search(from)
		e.to = to
	}

	return e
}

func newAscend(t RbTree) *ascend {
	tree := t.(*rbTree)
	it := iterator{tree: tree}
	e := &ascend{iterator: it}
	e.it = e
	return e
}

// NewDescend creates Enumerable that walks tree in descending order
func NewDescend(t RbTree) Enumerable {
	e := newDescend(t)

	if !e.tree.root.isNil() {
		e.next = e.tree.root.maximum()
		e.to = t.Minimum().Key()
	}

	return e
}

// NewDescendRange that walks tree in descending order within the range [from, to]
func NewDescendRange(t RbTree, from, to Comparable) Enumerable {
	e := newDescend(t)

	if !e.tree.root.isNil() && to != nil {
		e.next, _ = e.tree.root.search(from)
		e.to = to
	}

	return e
}

func newDescend(t RbTree) *descend {
	tree := t.(*rbTree)
	it := iterator{tree: tree}
	e := &descend{iterator: it}
	e.it = e
	return e
}

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

func (i *ascend) Next() bool {
	result := !i.next.isNil() && (i.next.key.LessThan(i.to) || i.next.key.EqualTo(i.to))
	if result {
		i.curr = i.next
		i.next = i.curr.successor()
	}
	return result
}

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

func (i *iterator) Current() Node { return i.curr }
