package rbtree

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

	// Root gets tree root node
	Root() Node
}

// Node represent red-black tree node interface
type Node interface {
	// Key gets node's key
	Key() Comparable

	// Subtree size including node itself
	Size() int64

	// Successor gets Node's successor
	Successor() Node

	// Predecessor gets Node's predecessor
	Predecessor() Node
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
type NodeAction func(Node)

// Iterator is an node iterator.
type Iterator interface {
	// Node gets current Node
	Current() Node

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
