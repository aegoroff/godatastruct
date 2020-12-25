package rbtree

// This file contains all RB tree modification methods implementations

// Insert inserts new node into Red-Black tree. Creates Root if tree is empty
func (tree *rbTree) Insert(z Comparable) {
	if z == nil {
		return
	}
	n := newNode(z)
	tree.insert(n)
}

// Insert inserts new node into Red-Black tree. Creates Root if tree is empty
func (tree *rbTree) ReplaceOrInsert(z Comparable) Comparable {
	if z == nil {
		return nil
	}

	var r Comparable
	n, ok := tree.SearchNode(z)
	if ok {
		tree.delete(n)
		r = n.key
	}

	tree.insert(newNode(z))
	return r
}

func newNode(z Comparable) *Node {
	return &Node{key: z}
}

func (tree *rbTree) insert(z *Node) {
	if tree.root.isNil() {
		tree.root = z
		tree.root.color = Black
		tree.root.parent = tree.tnil
		tree.root.left = tree.tnil
		tree.root.right = tree.tnil
		tree.root.size = 1
		return
	}
	y := tree.tnil
	x := tree.root
	z.size = 1
	for x != tree.tnil {
		y = x
		y.size++
		if z.key.LessThan(x.key) {
			x = x.left
		} else {
			x = x.right
		}
	}

	z.parent = y
	if z.key.LessThan(y.key) {
		y.left = z
	} else {
		y.right = z
	}
	z.left = tree.tnil
	z.right = tree.tnil
	z.color = Red
	rbInsertFixup(tree, z)
}

func rbInsertFixup(tree *rbTree, z *Node) {
	for z.parent.color == Red {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			if y.color == Red {
				z.parent.color = Black
				y.color = Black
				z.parent.parent.color = Red
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					leftRotate(tree, z)
				}

				z.parent.color = Black
				z.parent.parent.color = Red
				rightRotate(tree, z.parent.parent)
			}
		} else {
			y := z.parent.parent.left
			if y.color == Red {
				z.parent.color = Black
				y.color = Black
				z.parent.parent.color = Red
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					rightRotate(tree, z)
				}

				z.parent.color = Black
				z.parent.parent.color = Red
				leftRotate(tree, z.parent.parent)
			}
		}
	}
	tree.root.color = Black
}

// DeleteNode searches and deletes first found node with key value specified from Red-black tree
// It returns true if node was successfully deleted otherwise false
func (tree *rbTree) DeleteNode(c Comparable) bool {
	found, ok := tree.root.search(c)
	if ok {
		tree.delete(found)
	}
	return ok
}

// DeleteAllNodes searches and deletes all found nodes with key value specified from Red-black tree
// It returns true if nodes was successfully deleted otherwise false
func (tree *rbTree) DeleteAllNodes(c Comparable) bool {
	ok := tree.DeleteNode(c)
	res := ok
	for ok {
		ok = tree.DeleteNode(c)
	}
	return res
}

func (tree *rbTree) delete(z *Node) {
	if z == nil || z.parent == nil {
		return
	}

	y := z

	p := z.parent
	for p != tree.tnil {
		p.size--
		p = p.parent
	}

	var x *Node
	yOriginalColor := y.color
	if z.left == tree.tnil {
		x = z.right
		rbTransplant(tree, z, z.right)
	} else if z.right == tree.tnil {
		x = z.left
		rbTransplant(tree, z, z.left)
	} else {
		y := z.right.minimum()
		yOriginalColor = y.color
		x = y.right
		if y.parent == z {
			x.parent = y
		} else {
			rbTransplant(tree, y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		rbTransplant(tree, z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}
	if yOriginalColor == Black {
		rbDeleteFixup(tree, x)
	}
}

func rbDeleteFixup(tree *rbTree, x *Node) {
	for x != tree.root && x.color == Black {
		if x == x.parent.left {
			w := x.parent.right
			if w.color == Red {
				w.color = Black
				x.parent.color = Red
				leftRotate(tree, x.parent)
				w = x.parent.right
			}

			if w.left.color == Black && w.right.color == Black {
				w.color = Red
				x = x.parent
			} else {
				if w.right.color == Black {
					w.left.color = Black
					w.color = Red
					rightRotate(tree, w)
					w = x.parent.right
				}

				w.color = x.parent.color
				x.parent.color = Black
				w.right.color = Black
				leftRotate(tree, x.parent)
				x = tree.root
			}
		} else {
			w := x.parent.left
			if w.color == Red {
				w.color = Black
				x.parent.color = Red
				rightRotate(tree, x.parent)
				w = x.parent.left
			}

			if w.right.color == Black && w.left.color == Black {
				w.color = Red
				x = x.parent
			} else {
				if w.left.color == Black {
					w.right.color = Black
					w.color = Red
					leftRotate(tree, w)
					w = x.parent.left
				}

				w.color = x.parent.color
				x.parent.color = Black
				w.left.color = Black
				rightRotate(tree, x.parent)
				x = tree.root
			}
		}
	}
	x.color = Black
}

func rbTransplant(tree *rbTree, u *Node, v *Node) {
	if u.parent == tree.tnil {
		tree.root = v
		tree.root.size = u.size - 1
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

func leftRotate(tree *rbTree, x *Node) {
	y := x.right
	x.right = y.left
	if y.left != tree.tnil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == tree.tnil {
		tree.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y

	y.size = x.size
	x.size = x.left.size + x.right.size + 1
}

func rightRotate(tree *rbTree, x *Node) {
	y := x.left
	x.left = y.right
	if y.right != tree.tnil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == tree.tnil {
		tree.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}

	y.right = x
	x.parent = y

	y.size = x.size
	x.size = x.left.size + x.right.size + 1
}
