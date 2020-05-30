// Package rbtree is a Red-black search binary tree implementation with support ordered statistic on the tree
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
	Key Comparable

	// Subtree size including node itself
	Size int64

	color  int
	parent *Node
	left   *Node
	right  *Node
}

// KeyIterator allows callers of Ascend* to iterate in-order over portions of
// the tree.  When this function returns false, iteration will stop and the
// associated Ascend* function will immediately return.
type KeyIterator func(c Comparable) bool

// Comparable defines comparable type interface
type Comparable interface {
	LessThan(y interface{}) bool
	EqualTo(y interface{}) bool
}

// Int is the int type key that can be stored as Node key
type Int int

// String is the string type key that can be stored as Node key
type String string

// LessThan define Comparable interface member for Int
func (x Int) LessThan(y interface{}) bool {
	return x < y.(Int)
}

// EqualTo define Comparable interface member for Int
func (x Int) EqualTo(y interface{}) bool {
	return x == y
}

// LessThan define Comparable interface member for String
func (x *String) LessThan(y interface{}) bool {
	return *x < *(y.(*String))
}

// EqualTo define Comparable interface member for String
func (x *String) EqualTo(y interface{}) bool {
	return *x == *(y.(*String))
}

// GetIntKey gets int key value from tree node
func (n *Node) GetIntKey() int {
	return int(n.Key.(Int))
}

// GetStringKey gets string key value from tree node
func (n *Node) GetStringKey() string {
	if n == nil || n.Key == nil {
		return ""
	}
	return string(*n.Key.(*String))
}

// NewIntKey creates new int key to be stores as tree node key
func NewIntKey(v int) Comparable {
	var r Comparable
	r = Int(v)
	return r
}

// NewIntNode creates new node that contains int key
func NewIntNode(v int) *Node {
	return NewNode(NewIntKey(v))
}

// NewStringKey creates new string key to be stores as tree node key
func NewStringKey(v string) Comparable {
	var r Comparable
	s := String(v)
	r = &s
	return r
}

// NewStringNode creates new node that contains string key
func NewStringNode(v string) *Node {
	return NewNode(NewStringKey(v))
}

// NewRbTree creates new Red-Black empty tree
func NewRbTree() *RbTree {
	tnil := Node{color: Black}
	return &RbTree{tnil: &tnil}
}

// NewNode creates new node
func NewNode(k Comparable) *Node {
	return &Node{Key: k}
}

// Len returns the number of nodes in the tree.
func (tree *RbTree) Len() int64 {
	if tree.Root == nil {
		return 0
	}

	return tree.Root.Size
}
