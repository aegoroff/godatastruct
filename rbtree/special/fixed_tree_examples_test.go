package special

import (
	"fmt"
	"github.com/aegoroff/godatastruct/rbtree"
)

func ExampleNewMaxTree() {
	tree := NewMaxTree(3)
	for i := 1; i <= 10; i++ {
		tree.Insert(rbtree.NewInt(i))
	}

	size := tree.Len()
	fmt.Println(size)

	found, ok := tree.Search(rbtree.NewInt(8))

	fmt.Println(found.Key())
	fmt.Println(ok)

	found, ok = tree.Search(rbtree.NewInt(1))

	fmt.Println(found)
	fmt.Println(ok)
	// Output:
	// 3
	// 8
	// true
	// <nil>
	// false
}

func ExampleNewMinTree() {
	tree := NewMinTree(3)
	for i := 1; i <= 10; i++ {
		tree.Insert(rbtree.NewInt(i))
	}

	size := tree.Len()
	fmt.Println(size)

	found, ok := tree.Search(rbtree.NewInt(8))

	fmt.Println(found)
	fmt.Println(ok)

	found, ok = tree.Search(rbtree.NewInt(1))

	fmt.Println(found.Key())
	fmt.Println(ok)
	// Output:
	// 3
	// <nil>
	// false
	// 1
	// true
}
