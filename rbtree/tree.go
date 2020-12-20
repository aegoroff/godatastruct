// Package rbtree is a Red-black search binary tree implementation with support ordered statistic on the tree
package rbtree

const (
	// Black RB tree node
	Black = iota

	// Red RB tree node
	Red
)

type rbTree struct {
	root *Node
	tnil *Node
}

// Node represent red-black tree node
type Node struct {
	key Comparable

	// Subtree size including node itself
	size int64

	color  int
	parent *Node
	left   *Node
	right  *Node
}

// Int is the int type key that can be stored as Node's key
type Int int

// Int64 is the int64 type key that can be stored as Node's key
type Int64 int64

// String is the string type key that can be stored as Node's key
type String string

// Key gets Node's key
func (n *Node) Key() Comparable {
	return n.key
}

// Size gets subtree size including node itself
func (n *Node) Size() int64 {
	return n.size
}

func (n *Node) isNil() bool {
	return n == nil || n.key == nil
}

func (n *Node) isNotNil() bool {
	return n != nil && n.key != nil
}

// LessThan define Comparable interface member for Int
func (x Int) LessThan(y Comparable) bool {
	return x < y.(Int)
}

// EqualTo define Comparable interface member for Int
func (x Int) EqualTo(y Comparable) bool {
	return x == y
}

// LessThan define Comparable interface member for Int64
func (x Int64) LessThan(y Comparable) bool {
	return x < y.(Int64)
}

// EqualTo define Comparable interface member for Int64
func (x Int64) EqualTo(y Comparable) bool {
	return x == y
}

// LessThan define Comparable interface member for String
func (x *String) LessThan(y Comparable) bool {
	return *x < *(y.(*String))
}

// EqualTo define Comparable interface member for String
func (x *String) EqualTo(y Comparable) bool {
	return *x == *(y.(*String))
}

func (x *String) String() string {
	return string(*x)
}

// GetInt gets int value from Comparable
func GetInt(c Comparable) int {
	return int(c.(Int))
}

// GetInt64 gets int64 value from Comparable
func GetInt64(c Comparable) int64 {
	return int64(c.(Int64))
}

// NewInt creates new int Comparable implementation
func NewInt(v int) Comparable {
	return Int(v)
}

// NewInt64 creates new int64 Comparable implementation
func NewInt64(v int64) Comparable {
	return Int64(v)
}

// NewString creates new string Comparable implementation
func NewString(v string) Comparable {
	s := String(v)
	return &s
}

// NewRbTree creates new Red-Black empty tree
func NewRbTree() RbTree {
	return newRbTree()
}

func newRbTree() *rbTree {
	tnil := Node{color: Black}
	return &rbTree{tnil: &tnil}
}

// Len returns the number of nodes in the tree.
func (tree *rbTree) Len() int64 {
	if tree.root.isNil() {
		return 0
	}

	return tree.root.size
}

func (tree *rbTree) Root() *Node {
	return tree.root
}
