package rbtree

import (
	"fmt"
)

func ExampleNew() {
	tree := New()
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
	tree := New()

	tree.Insert(Int(6))
	tree.Insert(Int(18))
	tree.Insert(Int(3))

	found, ok := tree.OrderStatisticSelect(2)
	fmt.Println(found.Key())
	fmt.Println(ok)
	// Output:
	// 6
	// true
}

func ExampleNode_Size() {
	tree := New()

	tree.Insert(Int(6))
	tree.Insert(Int(18))
	tree.Insert(Int(3))

	root := tree.Root()

	size := root.Size()
	fmt.Println(size)
	// Output:
	// 3
}

func ExampleNode_Successor() {
	tree := New()

	tree.Insert(Int(6))
	tree.Insert(Int(18))
	tree.Insert(Int(3))

	root := tree.Root()

	n := root.Successor()
	fmt.Println(n.Key())
	// Output:
	// 18
}

func ExampleNode_Predecessor() {
	tree := New()

	tree.Insert(Int(6))
	tree.Insert(Int(18))
	tree.Insert(Int(3))

	root := tree.Root()

	n := root.Predecessor()
	fmt.Println(n.Key())
	// Output:
	// 3
}
