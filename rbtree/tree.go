// This package is a Red-black search binary tree implementation
package rbtree

const (
	// Black RB tree node
	Black = iota

	// Red RB tree node
	Red
)

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

// WalkInorder walks tree inorder (left, node, right)
func WalkInorder(root *Node, action func(*Node)) {
	if root != nil && root.Key != nil {
		WalkInorder(root.Left, action)
		action(root)
		WalkInorder(root.Right, action)
	}
}

// WalkPreorder walks tree preorder (node, left, right)
func WalkPreorder(root *Node, action func(*Node)) {
	if root != nil && root.Key != nil {
		action(root)
		WalkPreorder(root.Left, action)
		WalkPreorder(root.Right, action)
	}
}
