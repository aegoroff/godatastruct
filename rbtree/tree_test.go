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
	node   *node
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
	if node.color == Red {
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
		{NewRbTree(), 0},
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
	ass.Equal("root", x.parent.String())
	ass.Equal("a", x.left.String())
	ass.Equal("y", x.right.String())
	ass.Equal("b", y.left.String())
	ass.Equal("g", y.right.String())
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
	ass.Equal("root", y.parent.String())
	ass.Equal("x", y.left.String())
	ass.Equal("g", y.right.String())
	ass.Equal("a", x.left.String())
	ass.Equal("b", x.right.String())
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

	n := NewInt(24)
	found, _ := tree.search(tree.root, n)

	// Act
	tree.delete(found)

	// Assert
	n = NewInt(28)
	found, _ = tree.search(tree.root, n)
	ass.Equal(Black, found.color)
}

func Test_DeleteAllNodes_EmptyTree(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	var nodes []int
	const nodesCount = 128
	r := rand.New(rand.NewSource(1000))

	for i := 1; i < nodesCount; i++ {
		nodes = append(nodes, r.Int())
	}
	tree := createIntTree(nodes)

	// Act

	for i := 1; i < nodesCount; i++ {
		n := NewInt(nodes[i-1])
		found, _ := tree.search(tree.root, n)
		tree.delete(found)
	}

	// Assert
	ass.Nil(tree.root.key)
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
		tree := NewRbTree()

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

func Test_GraphvizString(t *testing.T) {
	// Arrange
	tree := createTestStringTree()

	// Act
	graphviz := getTreeAsGraphviz(tree)

	// Assert
	t.Log(graphviz)
}

func getTreeAsGraphviz(tree *rbTree) string {
	b := strings.Builder{}
	gr := simple.NewUndirectedGraph()

	var id int64

	tree.walkPreorder(tree.root, func(nod *node) {
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
	})

	data, _ := dot.Marshal(gr, "", " ", " ")

	b.Write(data)

	return b.String()
}

func Test_Delete_NodeDeleted(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()
	n := NewString("intel")
	found, _ := tree.search(tree.root, n)

	// Act
	tree.delete(found)

	// Assert
	found, ok := tree.search(tree.root, n)
	ass.False(ok)
	ass.Nil(found)

	found, ok = tree.search(tree.root, NewString("microsoft"))
	ass.True(ok)
	ass.Equal("microsoft", found.String())
}

func Test_DeleteNil_NothingDeleted(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()

	// Act
	tree.delete(nil)

	// Assert
	found, ok := tree.search(tree.root, NewString("microsoft"))
	ass.True(ok)
	ass.Equal("microsoft", found.String())
}

func Test_DeleteEmptyTree_NoError(t *testing.T) {
	// Arrange
	tree := NewRbTree()
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
		{NewString("intel"), false, NewRbTree()},
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

func Test_InsertIntoEmpty_Inserted(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := NewRbTree()

	// Act
	tree.Insert(NewString("1"))

	// Assert
	ass.Equal(int64(1), tree.Len())
}

func Test_InsertIntoNotEmpty_Inserted(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := NewRbTree()
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

func createIntegerTestTree() *rbTree {
	nodes := []int{6, 18, 3, 15, 7, 2, 4, 13, 9, 17, 20}
	return createIntTree(nodes)
}

func createTestStringTree() *rbTree {
	nodes := []string{"abc", "amd", "cisco", "do", "fake", "intel", "it", "let", "microsoft", "russia", "usa", "xxx", "yyy", "zen"}
	return createStringTree(nodes)
}

func createIntTree(nodes []int) *rbTree {
	tree := newRbTree()
	for _, n := range nodes {
		tree.Insert(NewInt(n))
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
