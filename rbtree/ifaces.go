package rbtree

// RbTree represents red-black tree interface
type RbTree interface {
	// Len returns the number of nodes in the tree.
	Len() int64

	// Insert inserts new Node into Red-Black tree. Creates Root if tree is empty
	Insert(n Comparable)

	// DeleteNode searches and deletes Node with key value specified from Red-black tree
	// It returns true if Node was successfully deleted otherwise false
	DeleteNode(c Comparable) bool

	// DeleteAllNodes searches and deletes all found nodes with key value specified from Red-black tree
	// It returns true if nodes was successfully deleted otherwise false
	DeleteAllNodes(c Comparable) bool

	// Search searches value specified within search tree
	Search(value Comparable) (Comparable, bool)

	// SearchAll searches all values with the same key as specified within search tree
	SearchAll(value Comparable) []Comparable

	// SearchNode searches *Node which key is equals value specified
	SearchNode(value Comparable) (*Node, bool)

	// Minimum gets tree's min element
	Minimum() *Node

	// Maximum gets tree's max element
	Maximum() *Node

	// OrderStatisticSelect gets i element from subtree
	// IMPORTANT: numeration starts from 1 not from 0
	OrderStatisticSelect(i int64) (*Node, bool)

	// Root gets tree root Node
	Root() *Node
}

// Comparable defines comparable type interface
type Comparable interface {
	// LessThan gets whether value specified less then current value
	LessThan(y Comparable) bool

	// EqualTo gets whether value specified equal current value
	EqualTo(y Comparable) bool
}

// Enumerable represents tree enumeration interface
type Enumerable interface {
	// Iterator gets underlying Iterator
	Iterator() Iterator

	// Foreach enumerates tree and calls the callback for
	// every value in the tree until callback returns false.
	Foreach(callback NodeAction)
}

// NodeAction defines function prototype that used by an iteration method to iterate over portions of
// the tree.
type NodeAction func(Comparable)

// Iterator is an node iterator.
type Iterator interface {
	// Node gets current Node
	Current() Comparable

	// Next advances the iterator and returns whether
	// the next call to the item method will return a
	// non-nil item.
	//
	// Next should be called prior to any call to the
	// iterator's item retrieval method after the
	// iterator has been obtained
	//
	// The order of iteration is implementation
	// dependent.
	Next() bool
}
