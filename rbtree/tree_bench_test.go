package rbtree

import (
	"github.com/google/btree"
	"math/rand"
	"testing"
)

const treeSizeInsert = 20000
const treeSizeSearchOrIterate = 100000
const bTreeDegree = 16
const searches = 100

func (i Int) Less(y btree.Item) bool {
	return i < y.(Int)
}

func (s *String) Less(y btree.Item) bool {
	return string(*s) < string(*y.(*String))
}

func Benchmark_RbTree_Insert(b *testing.B) {
	ints := perm(treeSizeInsert)
	tree := NewRbTree()
	for i := 0; i < b.N; i++ {
		for _, n := range ints {
			tree.Insert(Int(n))
		}
	}
	b.ReportAllocs()
}

func Benchmark_RbTree_ReplaceOrInsert(b *testing.B) {
	ints := perm(treeSizeInsert)
	tree := NewRbTree()
	for i := 0; i < b.N; i++ {
		for _, n := range ints {
			tree.ReplaceOrInsert(Int(n))
		}
	}
	b.ReportAllocs()
}

func Benchmark_BTree_ReplaceOrInsert(b *testing.B) {
	ints := perm(treeSizeInsert)
	tree := btree.New(bTreeDegree)
	for i := 0; i < b.N; i++ {
		for _, n := range ints {
			tree.ReplaceOrInsert(Int(n))
		}
	}
	b.ReportAllocs()
}

func Benchmark_RbTree_Search(b *testing.B) {
	// Arrange
	tree := NewRbTree()
	nodes := generateRandomStrings(treeSizeSearchOrIterate, 50)

	for i := 0; i < treeSizeSearchOrIterate; i++ {
		tree.Insert(nodes[i])
	}

	unexist := generateRandomStrings(searches, 50)

	off := rand.Intn(treeSizeSearchOrIterate / 2)

	// Act
	for i := 0; i < b.N; i++ {
		for j := 0; j < searches; j++ {
			tree.Search(nodes[j+off])
			tree.Search(unexist[j])
		}
	}
	b.ReportAllocs()
}

func Benchmark_BTree_Search(b *testing.B) {
	// Arrange
	tree := btree.New(bTreeDegree)
	nodes := generateRandomStrings(treeSizeSearchOrIterate, 50)
	for i := 0; i < treeSizeSearchOrIterate; i++ {
		tree.ReplaceOrInsert(nodes[i])
	}

	unexist := generateRandomStrings(searches, 50)

	off := rand.Intn(treeSizeSearchOrIterate / 2)

	for i := 0; i < b.N; i++ {
		for j := 0; j < searches; j++ {
			tree.Has(nodes[j+off])
			tree.Has(unexist[j])
		}
	}
	b.ReportAllocs()
}

func Benchmark_RbTree_Ascend(b *testing.B) {
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

func Benchmark_BTree_Ascend(b *testing.B) {
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

func Benchmark_RbTree_Descend(b *testing.B) {
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

func Benchmark_BTree_Descend(b *testing.B) {
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

func generateRandomStrings(num int, length int) []*String {
	result := make([]*String, num)
	for i := 0; i < num; i++ {
		l := 1 + rand.Intn(length)
		s := randomString(l)
		n := String(s)
		result[i] = &n
	}
	return result
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		ix := rand.Intn(len(letters))
		s[i] = letters[ix]
	}
	return string(s)
}
