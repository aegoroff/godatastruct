package special

import (
	"fmt"
	"github.com/aegoroff/godatastruct/rbtree"
)

func ExampleNewMaxTree() {
	maxTree := NewMaxTree(3)
	for i := 1; i <= 10; i++ {
		maxTree.Insert(rbtree.NewInt(i))
	}

	size := maxTree.Len()
	fmt.Println(size)

	found, ok := maxTree.Search(rbtree.NewInt(8))

	fmt.Println(found)
	fmt.Println(ok)

	found, ok = maxTree.Search(rbtree.NewInt(1))

	fmt.Println(found)
	fmt.Println(ok)

	// Output
	// 3
	// 8
	// true
	// <nil>
	// false
}

func ExampleNewMinTree() {
	maxTree := NewMinTree(3)
	for i := 1; i <= 10; i++ {
		maxTree.Insert(rbtree.NewInt(i))
	}

	size := maxTree.Len()
	fmt.Println(size)

	found, ok := maxTree.Search(rbtree.NewInt(8))

	fmt.Println(found)
	fmt.Println(ok)

	found, ok = maxTree.Search(rbtree.NewInt(1))

	fmt.Println(found)
	fmt.Println(ok)

	// Output
	// 3
	// <nil>
	// false
	// 1
	// true
}
