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
	// Iterate does tree iteration and calls the callback for
	// every value in the tree until callback returns false.
	Iterate(callback NodeValidator)
}

// NodeValidator defines function prototype that used by an iteration method to iterate over portions of
// the tree.  When this function returns false, iteration will stop and the
// associated iteration method function will immediately return.
type NodeValidator func(Node) bool

type rbTree struct {
	root *node
	tnil *node
}

// Node represent red-black tree node interface
type Node interface {
	// Subtree size including node itself
	Size() int64

	// Key gets node's key
	Key() Comparable

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
	LessThan(y interface{}) bool
	EqualTo(y interface{}) bool
}

// Int is the int type key that can be stored as Node key
type Int int

// String is the string type key that can be stored as Node key
type String string

func (n *node) Size() int64 {
	return n.size
}

func (n *node) String() string {
	return n.key.String()
}

func (n *node) Key() Comparable {
	return n.key
}

func (n *node) isNil() bool {
	return n == nil || n.key == nil
}

// LessThan define Comparable interface member for Int
func (x Int) LessThan(y interface{}) bool {
	return x < y.(Int)
}

// EqualTo define Comparable interface member for Int
func (x Int) EqualTo(y interface{}) bool {
	return x == y
}

func (x Int) String() string {
	return fmt.Sprintf("%d", x)
}

// LessThan define Comparable interface member for String
func (x *String) LessThan(y interface{}) bool {
	return *x < *(y.(*String))
}

// EqualTo define Comparable interface member for String
func (x *String) EqualTo(y interface{}) bool {
	return *x == *(y.(*String))
}

func (x *String) String() string {
	return string(*x)
}

// GetInt gets int key value from comparable
func GetInt(c Comparable) int {
	return int(c.(Int))
}

// NewInt creates new Comparable that contains int key
func NewInt(v int) Comparable {
	r := Int(v)
	return r
}

// NewString creates new string Comparable
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
