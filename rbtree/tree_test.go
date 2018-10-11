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
	Node   *Node
	NodeID int64
}

func (n GraphNode) ID() int64 {
	return n.NodeID
}

func (n GraphNode) DOTID() string {
	if key, ok := (*n.Node.Key).(String); ok {
		return fmt.Sprintf("\"%v\"", key)
	}

	if key, ok := (*n.Node.Key).(Int); ok {
		return fmt.Sprintf("\"%d\"", key)
	}

	return ""
}

func (n GraphNode) Attributes() []encoding.Attribute {
	node := *n.Node

	fc := "black"
	if node.Color == Red {
		fc = "red"
	}

	fillcolor := encoding.Attribute{Key: "fillcolor", Value: fc}
	fontcolor := encoding.Attribute{Key: "fontcolor", Value: "white"}
	style := encoding.Attribute{Key: "style", Value: "filled"}
	shape := encoding.Attribute{Key: "shape", Value: "box"}
	label := encoding.Attribute{Key: "label", Value: fmt.Sprintf(`"%s [%d]"`, strings.Trim(n.DOTID(), `"`), node.Size)}
	return []encoding.Attribute{fontcolor, fillcolor, style, label, shape}
}

func Test_InorderWalkTreeInt_AllElementsAscending(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	var result []int

	// Act
	tree.WalkInorder(func(node *Node) {
		result = append(result, getIntValueOf(node))
	})

	// Assert
	ass.Equal([]int{2, 3, 4, 6, 7, 9, 13, 15, 17, 18, 20}, result)
}

func Test_PreorderAllTreeWalkInt_AllElementsAscending(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	var result []int

	// Act
	tree.WalkPreorder(func(node *Node) {
		result = append(result, getIntValueOf(node))
	})

	// Assert
	ass.Equal([]int{6, 3, 2, 4, 15, 9, 7, 13, 18, 17, 20}, result)
}

func Test_InorderWalkString_AllElementsAscending(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()
	var result []string

	// Act
	tree.Root.WalkInorder(func(node *Node) {
		result = append(result, getStringValueOf(node))
	})

	// Assert
	ass.Equal([]string{"abc", "amd", "cisco", "do", "fake", "intel", "it", "let", "microsoft", "russia", "usa", "xxx", "yyy", "zen"}, result)
}

func Test_OrderStatisticSelect_ValueAsExpected(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	var tests = []struct {
		order    int64
		expected int
	}{
		{1, 2},
		{2, 3},
		{11, 20},
		{10, 18},
		{6, 9},
	}
	for _, test := range tests {
		// Act
		found := tree.OrderStatisticSelect(test.order)

		// Assert
		ass.NotNil(found)
		ass.Equal(test.expected, getIntValueOf(found))
	}
}

func Test_SearchIntTree_Success(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	v := createIntNode(13)

	// Act
	found, ok := tree.Search(v)

	// Assert
	ass.True(ok)
	ass.NotNil(found)
	ass.Equal(13, getIntValueOf(found))
}

func Test_SearchStringTree_Success(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()
	n := createStringNode("intel")

	// Act
	found, ok := tree.Search(n)

	// Assert
	ass.True(ok)
	ass.NotNil(found)
	ass.Equal("intel", getStringValueOf(found))
}

func Test_SearchIntTree_Failure(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	v := createIntNode(22)

	// Act
	found, ok := tree.Search(v)

	// Assert
	ass.False(ok)
	ass.Nil(found)
}

func Test_Successor_ReturnSuccessor(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	var tests = []struct {
		node     int
		expected int
	}{
		{13, 15},
		{6, 7},
		{18, 20},
		{2, 3},
	}
	for _, test := range tests {
		v := createIntNode(test.node)
		r, _ := tree.Search(v)

		// Act
		s := r.Successor()

		// Assert
		ass.Equal(test.expected, getIntValueOf(s))
	}
}

func Test_SuccessorOfMax_ReturnNil(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	v := createIntNode(20)
	r, _ := tree.Search(v)

	// Act
	s := r.Successor()

	// Assert
	ass.Nil(s)
}

func Test_PredecessorInTheMiddle_PredecessorFound(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	var tests = []struct {
		node     int
		expected int
	}{
		{13, 9},
		{6, 4},
		{18, 17},
		{3, 2},
	}
	for _, test := range tests {
		v := createIntNode(test.node)
		r, _ := tree.Search(v)

		// Act
		s := r.Predecessor()

		// Assert
		ass.Equal(test.expected, getIntValueOf(s))
	}
}

func Test_PredecessorOfMin_ReturnNil(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	v := createIntNode(2)
	r, _ := tree.Search(v)

	// Act
	p := r.Predecessor()

	// Assert
	ass.Nil(p)
}

func Test_Minimum_ValueAsExpected(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	// Act
	r := tree.Minimum()

	// Assert
	ass.Equal(2, getIntValueOf(r))
}

func Test_Maximum_ValueAsExpected(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	// Act
	r := tree.Maximum()

	// Assert
	ass.Equal(20, getIntValueOf(r))
}

func Test_RightRotate_StructureAsExpected(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	r := Node{Key: createStringNode("root")}

	tree := NewRbTree()
	tree.Insert(&r)

	y := Node{Key: createStringNode("y")}
	x := Node{Key: createStringNode("x")}
	a := Node{Key: createStringNode("a")}
	b := Node{Key: createStringNode("b")}
	g := Node{Key: createStringNode("g")}

	r.Right = &y
	y.Parent = &r
	y.Left = &x
	y.Right = &g
	x.Left = &a
	x.Right = &b
	x.Parent = &y
	g.Parent = &y
	a.Parent = &x
	b.Parent = &x

	// Act
	rightRotate(tree, &y)

	// Assert
	ass.Equal("root", getStringValueOf(x.Parent))
	ass.Equal("a", getStringValueOf(x.Left))
	ass.Equal("y", getStringValueOf(x.Right))
	ass.Equal("b", getStringValueOf(y.Left))
	ass.Equal("g", getStringValueOf(y.Right))
}

func Test_LeftRotate_StructureAsExpected(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	r := Node{Key: createStringNode("root")}

	tree := NewRbTree()
	tree.Insert(&r)

	x := Node{Key: createStringNode("x")}
	y := Node{Key: createStringNode("y")}
	a := Node{Key: createStringNode("a")}
	b := Node{Key: createStringNode("b")}
	g := Node{Key: createStringNode("g")}

	r.Right = &x
	x.Parent = &r
	x.Left = &a
	x.Right = &y
	y.Left = &b
	y.Right = &g
	y.Parent = &y
	g.Parent = &y
	a.Parent = &x
	b.Parent = &y

	// Act
	leftRotate(tree, &x)

	// Assert
	ass.Equal("root", getStringValueOf(y.Parent))
	ass.Equal("x", getStringValueOf(y.Left))
	ass.Equal("g", getStringValueOf(y.Right))
	ass.Equal("a", getStringValueOf(x.Left))
	ass.Equal("b", getStringValueOf(x.Right))
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

	n := createIntNode(24)
	found, _ := tree.Search(n)

	// Act
	tree.Delete(found)

	// Assert
	n = createIntNode(28)
	found, _ = tree.Root.Search(n)
	ass.Equal(Black, found.Color)
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
		n := createIntNode(nodes[i-1])
		found, _ := tree.Search(n)
		tree.Delete(found)
	}

	// Assert
	ass.Nil(tree.Root.Key)
	ass.Equal(int64(0), tree.Root.Size)
}

func Test_GraphvizString(t *testing.T) {
	// Arrange
	tree := createTestStringTree()

	// Act
	graphviz := getTreeAsGraphviz(tree)

	// Assert
	t.Log(graphviz)
}

func getTreeAsGraphviz(tree *RbTree) string {
	b := strings.Builder{}
	gr := simple.NewUndirectedGraph()

	var id int64

	tree.Root.WalkPreorder(func(node *Node) {

		gn := &GraphNode{Node: node, NodeID: id}
		gr.AddNode(gn)
		id++

		for i := id - 2; i >= 0; i-- {
			n := gr.Node(i)
			if node.Parent.Key != nil && n.(*GraphNode).Node.Key == node.Parent.Key {
				edge := gr.NewEdge(n, gn)
				gr.SetEdge(edge)
				break
			}
		}
	})

	data, _ := dot.Marshal(gr, "", " ", " ", false)

	b.Write(data)

	return b.String()
}

func Test_Delete_NodeDeleted(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()
	n := createStringNode("intel")
	found, _ := tree.Search(n)

	// Act
	tree.Delete(found)

	// Assert
	found, ok := tree.Search(n)
	ass.False(ok)
	ass.Nil(found)

	found, ok = tree.Search(createStringNode("microsoft"))
	ass.True(ok)
	ass.Equal("microsoft", getStringValueOf(found))
}

func Test_InsertNil_NothingIserted(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()
	oldSize := tree.Root.Size

	// Act
	tree.Insert(nil)

	// Assert
	ass.Equal(oldSize, tree.Root.Size)
}

func Test_DeleteNil_NothingDeleted(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()
	oldSize := tree.Root.Size

	// Act
	tree.Delete(nil)

	// Assert
	ass.Equal(oldSize, tree.Root.Size)
}

func createIntegerTestTree() *RbTree {
	nodes := []int{6, 18, 3, 15, 7, 2, 4, 13, 9, 17, 20}
	return createIntTree(nodes)
}

func createTestStringTree() *RbTree {
	nodes := []string{"abc", "amd", "cisco", "do", "fake", "intel", "it", "let", "microsoft", "russia", "usa", "xxx", "yyy", "zen"}
	return createStringTree(nodes)
}

func createIntTree(nodes []int) *RbTree {
	tree := NewRbTree()
	for _, n := range nodes {
		tree.Insert(NewNode(*createIntNode(n)))
	}
	return tree
}

func createStringTree(nodes []string) *RbTree {
	tree := NewRbTree()
	for _, n := range nodes {
		tree.Insert(NewNode(*createStringNode(n)))
	}
	return tree
}

func getIntValueOf(node *Node) int {
	return int((*node.Key).(Int))
}

func getStringValueOf(node *Node) string {
	return string((*node.Key).(String))
}

func createIntNode(v int) *Comparable {
	var r Comparable
	r = Int(v)
	return &r
}

func createStringNode(v string) *Comparable {
	var r Comparable
	r = String(v)
	return &r
}
