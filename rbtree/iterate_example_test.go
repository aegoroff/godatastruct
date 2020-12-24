package rbtree

import (
	"fmt"
)

func ExampleNewWalkInorder() {
	tree := NewRbTree()

	tree.Insert(Int(6))
	tree.Insert(Int(18))
	tree.Insert(Int(3))

	it := NewWalkInorder(tree)

	it.Foreach(func(n Comparable) {
		fmt.Println(n)
	})
	// Output:
	// 3
	// 6
	// 18
}

func ExampleNewWalkPreorder() {
	tree := NewRbTree()

	tree.Insert(Int(6))
	tree.Insert(Int(18))
	tree.Insert(Int(3))

	it := NewWalkPreorder(tree)

	it.Foreach(func(n Comparable) {
		fmt.Println(n)
	})
	// Output:
	// 6
	// 3
	// 18
}

func ExampleNewWalkPostorder() {
	tree := NewRbTree()

	tree.Insert(Int(6))
	tree.Insert(Int(18))
	tree.Insert(Int(3))

	it := NewWalkPostorder(tree)

	it.Foreach(func(n Comparable) {
		fmt.Println(n)
	})
	// Output:
	// 3
	// 18
	// 6
}

func ExampleNewAscend() {
	tree := NewRbTree()

	tree.Insert(Int(6))
	tree.Insert(Int(18))
	tree.Insert(Int(3))

	it := NewAscend(tree)

	it.Foreach(func(n Comparable) {
		fmt.Println(n)
	})
	// Output:
	// 3
	// 6
	// 18
}

func ExampleNewDescend() {
	tree := NewRbTree()

	tree.Insert(Int(6))
	tree.Insert(Int(18))
	tree.Insert(Int(3))

	it := NewDescend(tree)

	it.Foreach(func(n Comparable) {
		fmt.Println(n)
	})
	// Output:
	// 18
	// 6
	// 3
}

func ExampleNewAscendRange() {
	tree := NewRbTree()

	tree.Insert(Int(6))
	tree.Insert(Int(18))
	tree.Insert(Int(3))

	it := NewAscendRange(tree, Int(3), Int(6))

	it.Foreach(func(n Comparable) {
		fmt.Println(n)
	})
	// Output:
	// 3
	// 6
}

func ExampleNewDescendRange() {
	tree := NewRbTree()

	tree.Insert(Int(6))
	tree.Insert(Int(18))
	tree.Insert(Int(3))

	it := NewDescendRange(tree, Int(6), Int(3))

	it.Foreach(func(n Comparable) {
		fmt.Println(n)
	})
	// Output:
	// 6
	// 3
}
