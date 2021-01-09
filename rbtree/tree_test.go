package rbtree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"
	"math/rand"
	"strings"
	"testing"
)

type GraphNode struct {
	node   *Node
	NodeID int64
}

func (n GraphNode) ID() int64 {
	return n.NodeID
}

func (n GraphNode) DOTID() string {
	if key, ok := n.node.key.(*String); ok {
		return fmt.Sprintf("\"%s\"", *key)
	}

	if key, ok := n.node.key.(Int); ok {
		return fmt.Sprintf("\"%d\"", key)
	}

	return ""
}

func (n GraphNode) Attributes() []encoding.Attribute {
	node := *n.node

	fc := "black"
	if node.color == red {
		fc = "red"
	}

	fillcolor := encoding.Attribute{Key: "fillcolor", Value: fc}
	fontcolor := encoding.Attribute{Key: "fontcolor", Value: "white"}
	style := encoding.Attribute{Key: "style", Value: "filled"}
	shape := encoding.Attribute{Key: "shape", Value: "box"}
	label := encoding.Attribute{Key: "label", Value: fmt.Sprintf(`"%s [%d]"`, strings.Trim(n.DOTID(), `"`), node.size)}
	return []encoding.Attribute{fontcolor, fillcolor, style, label, shape}
}

func Test_Len(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	var tests = []struct {
		tree     RbTree
		expected int64
	}{
		{createIntegerTestTree(), 11},
		{New(), 0},
	}
	for _, test := range tests {
		// Act
		length := test.tree.Len()

		// Assert
		ass.Equal(test.expected, length)
	}
}

func Test_RightRotate_StructureAsExpected(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	r := newNode(NewString("root"))

	tree := newRbTree()
	tree.insert(r)

	y := newNode(NewString("y"))
	x := newNode(NewString("x"))
	a := newNode(NewString("a"))
	b := newNode(NewString("b"))
	g := newNode(NewString("g"))

	r.right = y
	y.parent = r
	y.left = x
	y.right = g
	x.left = a
	x.right = b
	x.parent = y
	g.parent = y
	a.parent = x
	b.parent = x

	// Act
	rightRotate(tree, y)

	// Assert
	ass.Equal("root", x.parent.key.(*String).String())
	ass.Equal("a", x.left.key.(*String).String())
	ass.Equal("y", x.right.key.(*String).String())
	ass.Equal("b", y.left.key.(*String).String())
	ass.Equal("g", y.right.key.(*String).String())
}

func Test_LeftRotate_StructureAsExpected(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	r := newNode(NewString("root"))

	tree := newRbTree()
	tree.insert(r)

	x := newNode(NewString("x"))
	y := newNode(NewString("y"))
	a := newNode(NewString("a"))
	b := newNode(NewString("b"))
	g := newNode(NewString("g"))

	r.right = x
	x.parent = r
	x.left = a
	x.right = y
	y.left = b
	y.right = g
	y.parent = y
	g.parent = y
	a.parent = x
	b.parent = y

	// Act
	leftRotate(tree, x)

	// Assert
	ass.Equal("root", y.parent.key.(*String).String())
	ass.Equal("x", y.left.key.(*String).String())
	ass.Equal("g", y.right.key.(*String).String())
	ass.Equal("a", x.left.key.(*String).String())
	ass.Equal("b", x.right.key.(*String).String())
}

func Test_GraphvizInt(t *testing.T) {
	// Arrange
	tree := createIntegerTestTree()

	// Act
	graphviz := getTreeAsGraphviz(tree)

	// Assert
	t.Log(graphviz)
}

func Test_DeleteFromLargeTree_SpecifiedNodeColorBlack(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var nodes []int

	for i := 1; i < 40; i++ {
		nodes = append(nodes, i)
	}
	tree := createIntTree(nodes)

	n := Int(24)
	found, _ := tree.Search(n)

	// Act
	tree.DeleteNode(found)

	// Assert
	n = Int(28)
	foundAfterDelete, _ := tree.SearchNode(n)
	ass.Equal(black, foundAfterDelete.color)
}

func Test_DeleteAllNodes_EmptyTree(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	const nodesCount = 128
	r := rand.New(rand.NewSource(1000))

	nodes := make([]int, nodesCount-1)

	for i := 1; i < nodesCount; i++ {
		nodes[i-1] = r.Int()
	}
	tree := createIntTree(nodes)

	// Act
	for i := 1; i < nodesCount; i++ {
		n := Int(nodes[i-1])
		found, _ := tree.Search(n)
		tree.DeleteNode(found)
	}

	// Assert
	ass.Equal(int64(0), tree.Len())
}

func Test_DeleteAllNodesWhenTreeContainsSameElements_TreeLenAsExpected(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	const nodesCount = 5

	var tests = []struct {
		input       []string
		expectedlen int64
	}{
		{[]string{"tst"}, 0},
		{[]string{"tst", "www"}, nodesCount},
	}
	for _, test := range tests {
		tree := New()

		for _, in := range test.input {
			k := NewString(in)
			for i := 0; i < nodesCount; i++ {
				tree.Insert(k)
			}
		}

		// Act
		res := tree.DeleteAllNodes(NewString(test.input[0]))

		// Assert
		ass.True(res)
		ass.Equal(test.expectedlen, tree.Len())
	}
}

func Test_SameKeyInsertDeleteLen_TreeLenAsExpected(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := New()

	k := NewString("tst")

	// Act
	tree.Insert(k)
	ass.Equal(int64(1), tree.Len())
	tree.Insert(k)
	ass.Equal(int64(2), tree.Len())
	tree.Insert(k)
	ass.Equal(int64(3), tree.Len())

	tree.DeleteNode(k)
	ass.Equal(int64(2), tree.Len())
	tree.DeleteNode(k)
	ass.Equal(int64(1), tree.Len())
	tree.DeleteNode(k)
	ass.Equal(int64(0), tree.Len())

	// Assert
}

func Test_DifferentKeyInsertDeleteSameOrderLen_TreeLenAsExpected(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := New()

	k1 := NewString("tst1")
	k2 := NewString("tst2")
	k3 := NewString("tst3")

	// Act
	tree.Insert(k1)
	ass.Equal(int64(1), tree.Len())
	tree.Insert(k2)
	ass.Equal(int64(2), tree.Len())
	tree.Insert(k3)
	ass.Equal(int64(3), tree.Len())

	tree.DeleteNode(k1)
	ass.Equal(int64(2), tree.Len())
	tree.DeleteNode(k2)
	ass.Equal(int64(1), tree.Len())
	tree.DeleteNode(k3)
	ass.Equal(int64(0), tree.Len())

	// Assert
}

func Test_DifferentKeyInsertDeleteReverseOrderLen_TreeLenAsExpected(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := New()

	k1 := NewString("tst1")
	k2 := NewString("tst2")
	k3 := NewString("tst3")

	// Act
	tree.Insert(k1)
	ass.Equal(int64(1), tree.Len())
	tree.Insert(k2)
	ass.Equal(int64(2), tree.Len())
	tree.Insert(k3)
	ass.Equal(int64(3), tree.Len())

	tree.DeleteNode(k3)
	ass.Equal(int64(2), tree.Len())
	tree.DeleteNode(k2)
	ass.Equal(int64(1), tree.Len())
	tree.DeleteNode(k1)
	ass.Equal(int64(0), tree.Len())

	// Assert
}

func Test_GraphvizString(t *testing.T) {
	// Arrange
	tree := createTestStringTree()

	// Act
	graphviz := getTreeAsGraphviz(tree)

	// Assert
	t.Log(graphviz)
}

func getTreeAsGraphviz(tree RbTree) string {
	b := strings.Builder{}
	gr := simple.NewUndirectedGraph()

	var id int64

	it := NewWalkPreorder(tree).Iterator()

	for it.Next() {
		nod := it.(*walkPreorder).current()
		gn := &GraphNode{node: nod, NodeID: id}
		gr.AddNode(gn)
		id++

		for i := id - 2; i >= 0; i-- {
			n := gr.Node(i)
			if nod.parent != nil && n.(*GraphNode).node == nod.parent {
				edge := gr.NewEdge(n, gn)
				gr.SetEdge(edge)
				break
			}
		}
	}

	data, _ := dot.Marshal(gr, "", " ", " ")

	b.Write(data)

	return b.String()
}

func Test_Delete_NodeDeleted(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()
	n := NewString("intel")
	found, _ := tree.root.search(n)

	// Act
	tree.delete(found)

	// Assert
	found, ok := tree.root.search(n)
	ass.False(ok)
	ass.Nil(found)

	found, ok = tree.root.search(NewString("microsoft"))
	ass.True(ok)
	ass.Equal("microsoft", found.key.(*String).String())
}

func Test_DeleteNil_NothingDeleted(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()

	// Act
	tree.delete(nil)

	// Assert
	found, ok := tree.root.search(NewString("microsoft"))
	ass.True(ok)
	ass.Equal("microsoft", found.key.(*String).String())
}

func Test_DeleteEmptyTree_NoError(t *testing.T) {
	// Arrange
	tree := New()
	n := NewString("intel")

	// Act
	tree.DeleteNode(n)

	// Assert
}

func Test_DeleteNode_ResultAsExpected(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	var tests = []struct {
		key    Comparable
		result bool
		tree   RbTree
	}{
		{NewString("intel"), true, createTestStringTree()},
		{NewString("vff"), false, createTestStringTree()},
		{nil, false, createTestStringTree()},
		{NewString("intel"), false, New()},
	}

	for _, test := range tests {
		// Act
		ok := test.tree.DeleteNode(test.key)

		// Assert
		ass.Equal(test.result, ok)
	}
}

func Test_DeleteNodeDeleteSeveralNodesWithTheSameKey_ResultAsExpected(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()
	k := NewString("vff")
	tree.Insert(k)
	tree.Insert(k)

	// Act
	ok1 := tree.DeleteNode(k)
	ok2 := tree.DeleteNode(k)
	ok3 := tree.DeleteNode(k)

	// Assert
	ass.True(ok1)
	ass.True(ok2)
	ass.False(ok3)
}

func Test_InsertAndCheckLen(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	tree := New()

	// Act
	tree.Insert(NewString("pp"))
	ass.Equal(int64(1), tree.Len())

	tree.Insert(NewString("nnn"))
	ass.Equal(int64(2), tree.Len())

	tree.Insert(NewString("gg"))
	ass.Equal(int64(3), tree.Len())

	tree.Insert(NewString("s"))
	ass.Equal(int64(4), tree.Len())

	tree.Insert(NewString("22"))
	ass.Equal(int64(5), tree.Len())

	// Assert
}

func Test_InsertNil_NothingInserted(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()
	oldSize := tree.Len()

	// Act
	tree.Insert(nil)

	// Assert
	ass.Equal(oldSize, tree.Len())
	ass.Nil(tree.tnil.parent)
}

func Test_ReplaceOrInsertInsertNil_NothingInserted(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()
	oldSize := tree.Len()

	// Act
	tree.ReplaceOrInsert(nil)

	// Assert
	ass.Equal(oldSize, tree.Len())
	ass.Nil(tree.tnil.parent)
}

func Test_InsertIntoEmpty_Inserted(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := New()

	// Act
	tree.Insert(NewString("1"))

	// Assert
	ass.Equal(int64(1), tree.Len())
}

func Test_ReplaceOrInsertIntoEmpty_Inserted(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := New()

	// Act
	r := tree.ReplaceOrInsert(NewString("1"))

	// Assert
	ass.Equal(int64(1), tree.Len())
	ass.Nil(r)
}

func Test_ReplaceOrInsertThatAlreadyInserted_InsertedOldDeletedAndReturned(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := New()
	k := NewString("1")
	tree.Insert(k)

	// Act
	r := tree.ReplaceOrInsert(NewString("1"))

	// Assert
	ass.Equal(int64(1), tree.Len())
	ass.NotNil(r)
	ass.Equal(k, r)
}

func Test_InsertIntoNotEmpty_Inserted(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := New()
	tree.Insert(NewString("1"))

	// Act
	tree.Insert(NewString("2"))

	// Assert
	ass.Equal(int64(2), tree.Len())
}

func Test_DeleteNodeNil_NothingDeleted(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()
	oldSize := tree.Len()

	// Act
	tree.DeleteNode(nil)

	// Assert
	ass.Equal(oldSize, tree.Len())
	ass.Nil(tree.tnil.parent)
}

func TestGetInt(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	c := Int(3)

	// Act
	i := GetInt(c)

	// Assert
	ass.Equal(3, i)
}

func TestGetInt64(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	c := Int64(3)

	// Act
	i := GetInt64(c)

	// Assert
	ass.Equal(int64(3), i)
}

func Test_Int64Tree(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := New()
	tree.Insert(Int64(4))
	tree.Insert(Int64(45))
	tree.Insert(Int64(3))

	// Act
	found, ok := tree.Search(Int64(4))

	// Assert
	ass.True(ok)
	ass.Equal(int64(4), GetInt64(found))
}

// []int{6, 18, 3, 15, 7, 2, 4, 13, 9, 17, 20}
func createIntegerTestTree() RbTree {
	nodes := []int{6, 18, 3, 15, 7, 2, 4, 13, 9, 17, 20}
	return createIntTree(nodes)
}

func createTestStringTree() *rbTree {
	nodes := []string{"abc", "amd", "cisco", "do", "fake", "intel", "it", "let", "microsoft", "russia", "usa", "xxx", "yyy", "zen"}
	return createStringTree(nodes)
}

func createIntTree(nodes []int) RbTree {
	tree := New()
	for _, n := range nodes {
		tree.Insert(Int(n))
	}
	return tree
}

func createStringTree(nodes []string) *rbTree {
	tree := newRbTree()
	for _, n := range nodes {
		tree.Insert(NewString(n))
	}
	return tree
}
