package rbtree

// This file contains all RB tree modification methods implementations

// Insert inserts new node into Red-Black tree. Creates Root if tree is empty
func (tree *RbTree) Insert(z *Node) {
	if z == nil {
		return
	}

	if tree.Root == nil {
		tree.Root = z
		tree.Root.color = Black
		tree.Root.parent = tree.tnil
		tree.Root.left = tree.tnil
		tree.Root.right = tree.tnil
		tree.Root.Size = 1
		return
	}
	y := tree.tnil
	x := tree.Root
	z.Size = 1
	for x != tree.tnil {
		y = x
		y.Size++
		if (*z.Key).LessThan(*x.Key) {
			x = x.left
		} else {
			x = x.right
		}
	}

	z.parent = y
	if y == tree.tnil {
		tree.Root = z
	} else if (*z.Key).LessThan(*y.Key) {
		y.left = z
	} else {
		y.right = z
	}
	z.left = tree.tnil
	z.right = tree.tnil
	z.color = Red
	rbInsertFixup(tree, z)
}

func rbInsertFixup(tree *RbTree, z *Node) {
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
	tree.Root.color = Black
}

// DeleteNode searches and deletes node with key value specified from Red-black tree
// It returns true if node was successfully deleted otherwise false
func (tree *RbTree) DeleteNode(key *Comparable) bool {
	if key == nil {
		return false
	}
	found, ok := tree.Search(key)
	if ok {
		tree.Delete(found)
	}
	return ok
}

// Delete deletes node specified from Red-black tree
func (tree *RbTree) Delete(z *Node) {
	if z == nil || z.parent == nil {
		return
	}

	y := z

	p := z.parent
	for p != tree.tnil {
		p.Size--
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
		y := z.right.Minimum()
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

func rbDeleteFixup(tree *RbTree, x *Node) {
	for x != tree.Root && x.color == Black {
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
				x = tree.Root
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
				x = tree.Root
			}
		}
	}
	x.color = Black
}

func rbTransplant(tree *RbTree, u *Node, v *Node) {
	if u.parent == tree.tnil {
		tree.Root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

func leftRotate(tree *RbTree, x *Node) {
	y := x.right
	x.right = y.left
	if y.left != tree.tnil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == tree.tnil {
		tree.Root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y

	y.Size = x.Size
	x.Size = x.left.Size + x.right.Size + 1
}

func rightRotate(tree *RbTree, x *Node) {
	y := x.left
	x.left = y.right
	if y.right != tree.tnil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == tree.tnil {
		tree.Root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}

	y.right = x
	x.parent = y

	y.Size = x.Size
	x.Size = x.left.Size + x.right.Size + 1
}
