package rbtree

import (
	"github.com/google/btree"
	"math/rand"
	"testing"
)

func (i Int) Less(y btree.Item) bool {
	return i < y.(Int)
}

func BenchmarkRbTree_Insert(b *testing.B) {
	ints := perm(100000)
	tree := NewRbTree()
	for i := 0; i < b.N; i++ {
		for _, n := range ints {
			tree.Insert(Int(n))
		}
	}
	b.ReportAllocs()
}

func BenchmarkBTree_Insert(b *testing.B) {
	ints := perm(100000)
	tree := btree.New(16)
	for i := 0; i < b.N; i++ {
		for _, n := range ints {
			tree.ReplaceOrInsert(Int(n))
		}
	}
	b.ReportAllocs()
}

func BenchmarkRbTree_Search(b *testing.B) {
	ints := perm(100000)
	tree := NewRbTree()
	for _, n := range ints {
		tree.Insert(Int(n))
	}
	k := Int(ints[len(ints)/2])
	k1 := Int(ints[len(ints)/4])
	for i := 0; i < b.N; i++ {
		tree.Search(k)
		tree.Search(k1)
	}
	b.ReportAllocs()
}

func BenchmarkBTree_Search(b *testing.B) {
	ints := perm(100000)
	tree := btree.New(16)
	for _, n := range ints {
		tree.ReplaceOrInsert(Int(n))
	}

	k := Int(ints[len(ints)/2])
	k1 := Int(ints[len(ints)/4])
	for i := 0; i < b.N; i++ {
		tree.Has(k)
		tree.Has(k1)
	}
	b.ReportAllocs()
}

// perm returns a random permutation of n Int items in the range [0, n).
func perm(n int) (out []int) {
	out = append(out, rand.Perm(n)...)
	return
}
