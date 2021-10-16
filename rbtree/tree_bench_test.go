package rbtree

import (
	"bytes"
	"fmt"
	"github.com/aegoroff/godatastruct/collections"
	"github.com/google/btree"
	"math/rand"
	"testing"
)

const treeSizeInsert = 20000
const treeSizeSearchOrIterate = 100000
const bTreeDegree = 256
const searches = 100

type bint int
type bstring string

func (x bint) Less(y btree.Item) bool {
	return x < y.(bint)
}

func (x *bstring) Less(y btree.Item) bool {
	return string(*x) < string(*y.(*bstring))
}

func Benchmark_RbTree_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		tree := New()
		ints := perm(treeSizeInsert)
		b.StartTimer()

		for _, n := range ints {
			tree.Insert(Int(n))
		}
	}
	b.ReportAllocs()
}

func Benchmark_RbTree_ReplaceOrInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ints := perm(treeSizeInsert)
		tree := New()
		b.StartTimer()

		for _, n := range ints {
			tree.ReplaceOrInsert(Int(n))
		}
	}
	b.ReportAllocs()
}

func Benchmark_BTree_ReplaceOrInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ints := perm(treeSizeInsert)
		tree := btree.New(bTreeDegree)
		b.StartTimer()

		for _, n := range ints {
			tree.ReplaceOrInsert(bint(n))
		}
	}
	b.ReportAllocs()
}

func Benchmark_RbTree_Search(b *testing.B) {
	// Arrange
	tree := New()
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
	nodes := generateRandomBStrings(treeSizeSearchOrIterate, 50)
	for i := 0; i < treeSizeSearchOrIterate; i++ {
		tree.ReplaceOrInsert(nodes[i])
	}

	unexist := generateRandomBStrings(searches, 50)

	off := rand.Intn(treeSizeSearchOrIterate / 2)

	for i := 0; i < b.N; i++ {
		for j := 0; j < searches; j++ {
			tree.Has(nodes[j+off])
			tree.Has(unexist[j])
		}
	}
	b.ReportAllocs()
}

func Benchmark_StringHashSet_Search(b *testing.B) {
	// Arrange
	hs := collections.StringHashSet{}
	nodes := generateRandomStrings(treeSizeSearchOrIterate, 50)

	for i := 0; i < treeSizeSearchOrIterate; i++ {
		hs.Add(string(*nodes[i]))
	}

	unexist := generateRandomStrings(searches, 50)

	off := rand.Intn(treeSizeSearchOrIterate / 2)

	// Act
	for i := 0; i < b.N; i++ {
		for j := 0; j < searches; j++ {
			hs.Contains(string(*nodes[j+off]))
			hs.Contains(string(*unexist[j]))
		}
	}
	b.ReportAllocs()
}

func Benchmark_RbTree_Ascend(b *testing.B) {
	ints := perm(treeSizeSearchOrIterate)
	tree := New()
	for _, n := range ints {
		tree.Insert(Int(n))
	}
	it := NewAscend(tree)
	x := 0
	for i := 0; i < b.N; i++ {
		it.Foreach(func(c Comparable) {
			x = int(c.(Int))
		})
	}
	b.ReportAllocs()
	_, _ = fmt.Fprintf(bytes.NewBuffer(nil), "%v", x)
}

func Benchmark_BTree_Ascend(b *testing.B) {
	ints := perm(treeSizeSearchOrIterate)
	tree := btree.New(bTreeDegree)
	for _, n := range ints {
		tree.ReplaceOrInsert(bint(n))
	}

	x := 0
	for i := 0; i < b.N; i++ {
		tree.Ascend(func(i btree.Item) bool {
			x = int(i.(bint))
			return true
		})
	}
	b.ReportAllocs()
	_, _ = fmt.Fprintf(bytes.NewBuffer(nil), "%v", x)
}

func Benchmark_RbTree_Descend(b *testing.B) {
	ints := perm(treeSizeSearchOrIterate)
	tree := New()
	for _, n := range ints {
		tree.Insert(Int(n))
	}
	it := NewDescend(tree)
	x := 0
	for i := 0; i < b.N; i++ {
		it.Foreach(func(c Comparable) {
			x = int(c.(Int))
		})
	}
	b.ReportAllocs()
	_, _ = fmt.Fprintf(bytes.NewBuffer(nil), "%v", x)
}

func Benchmark_BTree_Descend(b *testing.B) {
	ints := perm(treeSizeSearchOrIterate)
	tree := btree.New(bTreeDegree)
	for _, n := range ints {
		tree.ReplaceOrInsert(bint(n))
	}
	x := 0
	for i := 0; i < b.N; i++ {
		tree.Descend(func(i btree.Item) bool {
			x = int(i.(bint))
			return true
		})
	}
	b.ReportAllocs()
	_, _ = fmt.Fprintf(bytes.NewBuffer(nil), "%v", x)
}

// perm returns a random permutation of n Int items in the range [0, n).
func perm(n int) (out []int) {
	out = append(out, rand.Perm(n)...)
	return
}

func generateRandomStrings(num int, length int) []*String {
	result := make([]*String, num)
	rnd := generateRandomStringSlice(num, length)
	for i, s := range rnd {
		bs := String(s)
		result[i] = &bs
	}
	return result
}

func generateRandomBStrings(num int, length int) []*bstring {
	result := make([]*bstring, num)
	rnd := generateRandomStringSlice(num, length)
	for i, s := range rnd {
		bs := bstring(s)
		result[i] = &bs
	}
	return result
}

func generateRandomStringSlice(num int, length int) []string {
	result := make([]string, num)
	for i := 0; i < num; i++ {
		l := 1 + rand.Intn(length)
		s := randomString(l)
		result[i] = s
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
