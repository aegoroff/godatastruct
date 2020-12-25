package rbtree

import (
	"github.com/google/btree"
	"math/rand"
	"testing"
)

const treeSizeInsert = 20000
const treeSizeSearchOrIterate = 100000
const bTreeDegree = 16

func (i Int) Less(y btree.Item) bool {
	return i < y.(Int)
}

func BenchmarkRbTree_Insert(b *testing.B) {
	ints := perm(treeSizeInsert)
	tree := NewRbTree()
	for i := 0; i < b.N; i++ {
		for _, n := range ints {
			tree.Insert(Int(n))
		}
	}
	b.ReportAllocs()
}

func BenchmarkRbTree_ReplaceOrInsert(b *testing.B) {
	ints := perm(treeSizeInsert)
	tree := NewRbTree()
	for i := 0; i < b.N; i++ {
		for _, n := range ints {
			tree.ReplaceOrInsert(Int(n))
		}
	}
	b.ReportAllocs()
}

func BenchmarkBTree_ReplaceOrInsert(b *testing.B) {
	ints := perm(treeSizeInsert)
	tree := btree.New(bTreeDegree)
	for i := 0; i < b.N; i++ {
		for _, n := range ints {
			tree.ReplaceOrInsert(Int(n))
		}
	}
	b.ReportAllocs()
}

func BenchmarkRbTree_Search(b *testing.B) {
	ints := perm(treeSizeSearchOrIterate)
	tree := NewRbTree()
	for _, n := range ints {
		tree.Insert(Int(n))
	}

	for i := 0; i < b.N; i++ {
		tree.Search(Int(i))
	}
	b.ReportAllocs()
}

func BenchmarkBTree_Search(b *testing.B) {
	ints := perm(treeSizeSearchOrIterate)
	tree := btree.New(bTreeDegree)
	for _, n := range ints {
		tree.ReplaceOrInsert(Int(n))
	}

	for i := 0; i < b.N; i++ {
		tree.Has(Int(i))
	}
	b.ReportAllocs()
}

func BenchmarkRbTree_Ascend(b *testing.B) {
	ints := perm(treeSizeSearchOrIterate)
	tree := NewRbTree()
	for _, n := range ints {
		tree.Insert(Int(n))
	}
	it := NewAscend(tree)
	for i := 0; i < b.N; i++ {
		it.Foreach(func(c Comparable) {
			x := int(c.(Int))
			x++
		})
	}
	b.ReportAllocs()
}

func BenchmarkBTree_Ascend(b *testing.B) {
	ints := perm(treeSizeSearchOrIterate)
	tree := btree.New(bTreeDegree)
	for _, n := range ints {
		tree.ReplaceOrInsert(Int(n))
	}

	for i := 0; i < b.N; i++ {
		tree.Ascend(func(i btree.Item) bool {
			x := int(i.(Int))
			x++
			return true
		})
	}
	b.ReportAllocs()
}

func BenchmarkRbTree_Descend(b *testing.B) {
	ints := perm(treeSizeSearchOrIterate)
	tree := NewRbTree()
	for _, n := range ints {
		tree.Insert(Int(n))
	}
	it := NewDescend(tree)
	for i := 0; i < b.N; i++ {
		it.Foreach(func(c Comparable) {
			x := int(c.(Int))
			x++
		})
	}
	b.ReportAllocs()
}

func BenchmarkBTree_Descend(b *testing.B) {
	ints := perm(treeSizeSearchOrIterate)
	tree := btree.New(bTreeDegree)
	for _, n := range ints {
		tree.ReplaceOrInsert(Int(n))
	}

	for i := 0; i < b.N; i++ {
		tree.Descend(func(i btree.Item) bool {
			x := int(i.(Int))
			x++
			return true
		})
	}
	b.ReportAllocs()
}

// perm returns a random permutation of n Int items in the range [0, n).
func perm(n int) (out []int) {
	out = append(out, rand.Perm(n)...)
	return
}
