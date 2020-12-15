package rbtree

// This file contains all RB tree iteration methods implementations

type enumerable struct{ it Iterator }

type iterator struct {
	enumerable
	tree RbTree
	curr Node
}

type walk struct {
	iterator
	stack []*node
}

type walkPreorder struct{ walk }

type walkInorder struct {
	walk
	p *node
}

type walkPostorder struct {
	walk
	p *node
}

type ascend struct{ ordered }

type descend struct{ ordered }

type ordered struct {
	iterator
	next Node
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

// NewAscend creates Enumerable that walks tree in ascending order
func NewAscend(t RbTree) Enumerable {
	e := newAscend(t)

	min := e.tree.Minimum()
	if min != nil {
		e.next = min.(*node)
		e.to = t.Maximum().Key()
	}

	return e
}

// NewAscendRange creates Enumerable that walks tree in ascending order within the range [from, to]
func NewAscendRange(t RbTree, from, to Comparable) Enumerable {
	e := newAscend(t)

	n, ok := e.tree.Search(from)
	if ok && to != nil {
		e.next = n.(*node)
		e.to = to
	}

	return e
}

// NewDescend creates Enumerable that walks tree in descending order
func NewDescend(t RbTree) Enumerable {
	e := newDescend(t)

	max := e.tree.Maximum()
	if max != nil {
		e.next = max.(*node)
		e.to = t.Minimum().Key()
	}

	return e
}

// NewDescendRange that walks tree in descending order within the range [from, to]
func NewDescendRange(t RbTree, from, to Comparable) Enumerable {
	e := newDescend(t)

	n, ok := e.tree.Search(from)
	if ok && to != nil {
		e.next = n.(*node)
		e.to = to
	}

	return e
}

func (i *walkInorder) Next() bool {
	for len(i.stack) > 0 {
		if i.p.isNotNil() {
			i.p = i.p.left
			if i.p.isNotNil() {
				i.stack = append(i.stack, i.p)
			}
		} else {
			top := len(i.stack) - 1
			i.p = i.stack[top]
			i.curr = i.p
			i.stack = i.stack[:top]
			i.p = i.p.right

			if i.p.isNotNil() {
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

		if p.right.isNotNil() {
			i.stack = append(i.stack, p.right)
		}

		if p.left.isNotNil() {
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
		}

		if next.right.isNotNil() {
			i.stack = append(i.stack, next.right)
		}
		if next.left.isNotNil() {
			i.stack = append(i.stack, next.left)
		}
	}

	return false
}

func (i *ascend) Next() bool {
	result := i.next != nil && (i.next.Key().LessThan(i.to) || i.next.Key().EqualTo(i.to))
	if result {
		i.curr = i.next
		i.next = i.curr.Successor()
	}
	return result
}

func (i *descend) Next() bool {
	result := i.next != nil && !i.next.Key().LessThan(i.to)
	if result {
		i.curr = i.next
		i.next = i.curr.Predecessor()
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

func (e *enumerable) Iterator() Iterator { return e.it }

func (i *iterator) Current() Node { return i.curr }

func newWalk(t RbTree) walk {
	it := iterator{tree: t}

	w := walk{
		iterator: it,
		stack:    make([]*node, 0),
	}

	if t.Len() > 0 {
		w.stack = append(w.stack, t.Root().(*node))
	}

	return w
}

func newAscend(t RbTree) *ascend {
	ordered := newOrdered(t)
	e := &ascend{ordered: ordered}
	e.it = e
	return e
}

func newDescend(t RbTree) *descend {
	ordered := newOrdered(t)
	e := &descend{ordered: ordered}
	e.it = e
	return e
}

func newOrdered(t RbTree) ordered {
	it := iterator{tree: t}
	return ordered{iterator: it}
}
