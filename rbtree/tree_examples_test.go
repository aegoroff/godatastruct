package rbtree

import (
	"fmt"
)

func ExampleNewRbTree() {
	tree := NewRbTree()
	node := NewString("a")
	tree.Insert(node)

	size := tree.Len()
	fmt.Println(size)

	n, ok := tree.Search(node)
	fmt.Println(n)
	fmt.Println(ok)

	n, ok = tree.Search(NewString("b"))
	fmt.Println(n)
	fmt.Println(ok)
	// Output:
	// 1
	// a
	// true
	// <nil>
	// false
}

func ExampleRbTree_OrderStatisticSelect() {
	tree := NewRbTree()

	tree.Insert(NewInt(6))
	tree.Insert(NewInt(18))
	tree.Insert(NewInt(3))

	found, ok := tree.OrderStatisticSelect(2)
	fmt.Println(found.Key())
	fmt.Println(ok)
	// Output:
	// 6
	// true
}

func ExampleNode_Size() {
	tree := NewRbTree()

	tree.Insert(NewInt(6))
	tree.Insert(NewInt(18))
	tree.Insert(NewInt(3))

	root := tree.Root()

	size := root.Size()
	fmt.Println(size)
	// Output:
	// 3
}

func ExampleNode_Successor() {
	tree := NewRbTree()

	tree.Insert(NewInt(6))
	tree.Insert(NewInt(18))
	tree.Insert(NewInt(3))

	root := tree.Root()

	n := root.Successor()
	fmt.Println(n.Key())
	// Output:
	// 18
}

func ExampleNode_Predecessor() {
	tree := NewRbTree()

	tree.Insert(NewInt(6))
	tree.Insert(NewInt(18))
	tree.Insert(NewInt(3))

	root := tree.Root()

	n := root.Predecessor()
	fmt.Println(n.Key())
	// Output:
	// 3
}
