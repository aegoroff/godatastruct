// Package rbtree is a Red-black search binary tree implementation with support ordered statistic on the tree
package rbtree

import "fmt"

const (
	// Black RB tree node
	Black = iota

	// Red RB tree node
	Red
)

// RbTree represents red-black tree interface
type RbTree interface {
	// Len returns the number of nodes in the tree.
	Len() int64

	// Insert inserts new node into Red-Black tree. Creates Root if tree is empty
	Insert(n Comparable)

	// DeleteNode searches and deletes node with key value specified from Red-black tree
	// It returns true if node was successfully deleted otherwise false
	DeleteNode(c Comparable) bool

	// DeleteAllNodes searches and deletes all found nodes with key value specified from Red-black tree
	// It returns true if nodes was successfully deleted otherwise false
	DeleteAllNodes(c Comparable) bool

	// Search searches value specified within search tree
	Search(value Comparable) (Node, bool)

	// Minimum gets tree's min element
	Minimum() Node

	// Maximum gets tree's max element
	Maximum() Node

	// OrderStatisticSelect gets i element from subtree
	// IMPORTANT: numeration starts from 1 not from 0
	OrderStatisticSelect(i int64) (Node, bool)
}

// Iterator represents tree iteration interface
type Iterator interface {
	// Foreach does tree iteration and calls the callback for
	// every value in the tree until callback returns false.
	Foreach(callback NodeEvaluator)
}

// NodeEvaluator defines function prototype that used by an iteration method to iterate over portions of
// the tree.  When this function returns false, iteration will stop and the
// associated iteration method function will immediately return.
type NodeEvaluator func(Node) bool

type rbTree struct {
	root *node
	tnil *node
}

// Node represent red-black tree node interface
type Node interface {
	Comparable

	// Subtree size including node itself
	Size() int64

	// Successor gets node's successor
	Successor() Node

	// Predecessor gets node's predecessor
	Predecessor() Node
}

// node represent red-black tree node implementation
type node struct {
	key Comparable

	// Subtree size including node itself
	size int64

	color  int
	parent *node
	left   *node
	right  *node
}

// Comparable defines comparable type interface
type Comparable interface {
	fmt.Stringer
	LessThan(y Comparable) bool
	EqualTo(y Comparable) bool
}

// Int is the int type key that can be stored as Node
type Int int

// Int64 is the int64 type key that can be stored as Node
type Int64 int64

// String is the string type key that can be stored as Node
type String string

func (n *node) Size() int64 {
	return n.size
}

func (n *node) String() string {
	return n.key.String()
}

// LessThan define Comparable interface member for *node
func (n *node) LessThan(y Comparable) bool {
	yn, ok := y.(*node)
	if ok {
		return n.key.LessThan(yn.key)
	}
	return n.key.LessThan(y)
}

// EqualTo define Comparable interface member for *node
func (n *node) EqualTo(y Comparable) bool {
	yn, ok := y.(*node)
	if ok {
		return n.key.EqualTo(yn.key)
	}
	return n.key.EqualTo(y)
}

func (n *node) isNil() bool {
	return n == nil || n.key == nil
}

// LessThan define Comparable interface member for Int
func (x Int) LessThan(y Comparable) bool {
	return x < y.(Int)
}

// EqualTo define Comparable interface member for Int
func (x Int) EqualTo(y Comparable) bool {
	return x == y
}

func (x Int) String() string {
	return fmt.Sprintf("%d", x)
}

// LessThan define Comparable interface member for Int64
func (x Int64) LessThan(y Comparable) bool {
	return x < y.(Int64)
}

// EqualTo define Comparable interface member for Int64
func (x Int64) EqualTo(y Comparable) bool {
	return x == y
}

func (x Int64) String() string {
	return fmt.Sprintf("%d", x)
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
	n, ok := c.(*node)
	if ok {
		return int(n.key.(Int))
	}
	return int(c.(Int))
}

// GetInt64 gets int value from Comparable
func GetInt64(c Comparable) int64 {
	n, ok := c.(*node)
	if ok {
		return int64(n.key.(Int64))
	}
	return int64(c.(Int64))
}

// NewInt creates new Node that contains int key
func NewInt(v int) Comparable {
	return Int(v)
}

// NewInt64 creates new Node that contains int64 key
func NewInt64(v int64) Comparable {
	return Int64(v)
}

// NewString creates new string Node
func NewString(v string) Comparable {
	s := String(v)
	return &s
}

// NewRbTree creates new Red-Black empty tree
func NewRbTree() RbTree {
	return newRbTree()
}

func newRbTree() *rbTree {
	tnil := node{color: Black}
	return &rbTree{tnil: &tnil}
}

// Len returns the number of nodes in the tree.
func (tree *rbTree) Len() int64 {
	if tree.root == nil || tree.root == tree.tnil {
		return 0
	}

	return tree.root.size
}
