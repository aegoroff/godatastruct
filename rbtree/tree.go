// Package rbtree is a Red-black search binary tree implementation
package rbtree

const (
	// Black RB tree node
	Black = iota

	// Red RB tree node
	Red
)

// RbTree represents red-black tree structure
type RbTree struct {
	Root *Node
	tnil *Node
}

// Node represent red-black tree node
type Node struct {
	// Node key (data)
	Key    *Comparable
	Parent *Node
	Left   *Node
	Right  *Node
	// Node color (red or black)
	Color int
	// Subtree size including node itself
	Size int64
}

// Comparable defines comparable type interface
type Comparable interface {
	LessThan(y interface{}) bool
	EqualTo(y interface{}) bool
}

// NewRbTree creates new Red-Black empty tree
func NewRbTree() *RbTree {
	tnil := Node{Color: Black}
	return &RbTree{tnil: &tnil}
}

// NewNode creates new node
func NewNode(si Comparable) *Node {
	return &Node{Key: &si}
}

// WalkInorder walks subtree inorder (left, node, right)
func (root *Node) WalkInorder(action func(*Node)) {
	if root != nil && root.Key != nil {
		root.Left.WalkInorder(action)
		action(root)
		root.Right.WalkInorder(action)
	}
}

// WalkInorder walks tree inorder (left, node, right)
func (tree *RbTree) WalkInorder(action func(*Node)) {
	tree.Root.WalkInorder(action)
}

// WalkPreorder walks subtree preorder (node, left, right)
func (root *Node) WalkPreorder(action func(*Node)) {
	if root != nil && root.Key != nil {
		action(root)
		root.Left.WalkPreorder(action)
		root.Right.WalkPreorder(action)
	}
}

// WalkInorder walks tree inorder (left, node, right)
func (tree *RbTree) WalkPreorder(action func(*Node)) {
	tree.Root.WalkPreorder(action)
}
