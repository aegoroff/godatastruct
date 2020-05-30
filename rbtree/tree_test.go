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
	Node   *node
	NodeID int64
}

func (n GraphNode) ID() int64 {
	return n.NodeID
}

func (n GraphNode) DOTID() string {
	if key, ok := n.Node.Comparable.(*String); ok {
		return fmt.Sprintf("\"%v\"", key)
	}

	if key, ok := n.Node.Comparable.(Int); ok {
		return fmt.Sprintf("\"%d\"", key)
	}

	return ""
}

func (n GraphNode) Attributes() []encoding.Attribute {
	node := *n.Node

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

func Test_InorderWalkTreeInt_AllElementsAscending(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	var result []int

	// Act
	tree.WalkInorder(func(node Comparable) {
		result = append(result, GetInt(node))
	})

	// Assert
	ass.Equal([]int{2, 3, 4, 6, 7, 9, 13, 15, 17, 18, 20}, result)
}

func Test_PreorderAllTreeWalkInt_AllElementsAsSpecified(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	var result []int

	// Act
	tree.WalkPreorder(func(node Comparable) {
		result = append(result, GetInt(node))
	})

	// Assert
	ass.Equal([]int{6, 3, 2, 4, 15, 9, 7, 13, 18, 17, 20}, result)
}

func Test_PostorderAllTreeWalkInt_AllElementsAsSpecified(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	var result []int

	// Act
	tree.WalkPostorder(func(node Comparable) {
		result = append(result, GetInt(node))
	})

	// Assert
	ass.Equal([]int{2, 4, 3, 7, 13, 9, 17, 20, 18, 15, 6}, result)
}

func Test_Ascend(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	var tests = []struct {
		predicate func(Comparable) bool
		expected  []Int
	}{
		{func(c Comparable) bool { return true }, []Int{2, 3, 4, 6, 7, 9, 13, 15, 17, 18, 20}},
		{func(c Comparable) bool { return false }, []Int{2}},
		{func(c Comparable) bool { return c.LessThan(NewInt(15)) }, []Int{2, 3, 4, 6, 7, 9, 13, 15}},
	}
	for _, test := range tests {
		result := []Int{}

		// Act
		tree.Ascend(func(c Comparable) bool {
			result = append(result, c.(Int))
			return test.predicate(c)
		})

		// Assert
		ass.Equal(test.expected, result)
	}
}

func Test_AscendEmptyTree(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := NewRbTree()
	result := []Int{}

	// Act
	tree.Ascend(func(c Comparable) bool {
		result = append(result, c.(Int))
		return true
	})

	// Assert
	ass.Equal([]Int{}, result)
}

func Test_AscendRange(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	var tests = []struct {
		from     int
		to       int
		expected []Int
	}{
		{6, 15, []Int{6, 7, 9, 13, 15}},
		{6, 6, []Int{6}},
		{15, 15, []Int{15}},
		{20, 20, []Int{20}},
		{2, 2, []Int{2}},
		{15, 6, []Int{}},
		{8, 15, []Int{}},
	}
	for _, test := range tests {
		result := []Int{}
		from := NewInt(test.from)
		to := NewInt(test.to)

		// Act
		tree.AscendRange(from, to, func(c Comparable) bool {
			result = append(result, c.(Int))
			return true
		})

		// Assert
		ass.Equal(test.expected, result)
	}
}

func Test_AscendRangeNilTests(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	var tests = []struct {
		from Comparable
		to   Comparable
	}{
		{nil, NewInt(6)},
		{NewInt(6), nil},
		{nil, nil},
	}
	for _, test := range tests {
		result := []Int{}

		// Act
		tree.AscendRange(test.from, test.to, func(c Comparable) bool {
			result = append(result, c.(Int))
			return true
		})

		// Assert
		ass.Equal([]Int{}, result)
	}
}

func Test_Descend(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	var tests = []struct {
		predicate func(Comparable) bool
		expected  []Int
	}{
		{func(c Comparable) bool { return true }, []Int{20, 18, 17, 15, 13, 9, 7, 6, 4, 3, 2}},
		{func(c Comparable) bool { return false }, []Int{20}},
		{func(c Comparable) bool { return !c.LessThan(NewInt(15)) }, []Int{20, 18, 17, 15, 13}},
	}
	for _, test := range tests {
		result := []Int{}

		// Act
		tree.Descend(func(c Comparable) bool {
			result = append(result, c.(Int))
			return test.predicate(c)
		})

		// Assert
		ass.Equal(test.expected, result)
	}
}

func Test_DescendEmptyTree(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := NewRbTree()
	result := []Int{}

	// Act
	tree.Descend(func(c Comparable) bool {
		result = append(result, c.(Int))
		return true
	})

	// Assert
	ass.Equal([]Int{}, result)
}

func Test_DescendRange(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	var tests = []struct {
		from     int
		to       int
		expected []Int
	}{
		{15, 6, []Int{15, 13, 9, 7, 6}},
		{6, 6, []Int{6}},
		{15, 15, []Int{15}},
		{20, 20, []Int{20}},
		{2, 2, []Int{2}},
		{6, 15, []Int{}},
		{14, 6, []Int{}},
	}
	for _, test := range tests {
		result := []Int{}
		from := NewInt(test.from)
		to := NewInt(test.to)

		// Act
		tree.DescendRange(from, to, func(c Comparable) bool {
			result = append(result, c.(Int))
			return true
		})

		// Assert
		ass.Equal(test.expected, result)
	}
}

func Test_DescendRangeNilTests(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	var tests = []struct {
		from Comparable
		to   Comparable
	}{
		{nil, NewInt(6)},
		{NewInt(6), nil},
		{nil, nil},
	}
	for _, test := range tests {
		result := []Int{}

		// Act
		tree.DescendRange(test.from, test.to, func(c Comparable) bool {
			result = append(result, c.(Int))
			return true
		})

		// Assert
		ass.Equal([]Int{}, result)
	}
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

func Test_InorderWalkString_AllElementsAscending(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()
	var result []string

	// Act
	tree.WalkInorder(func(node Comparable) {
		result = append(result, GetString(node))
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
		found, _ := tree.OrderStatisticSelect(test.order)

		// Assert
		ass.NotNil(found)
		ass.Equal(test.expected, GetInt(found))
	}
}

func Test_OrderStatisticSelectNegativeTests_NullResult(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	var tests = []struct {
		tree  RbTree
		order int64
	}{
		{createIntegerTestTree(), 200},
		{NewRbTree(), 1},
	}
	for _, test := range tests {
		// Act
		found, ok := test.tree.OrderStatisticSelect(test.order)

		// Assert
		ass.Nil(found)
		ass.False(ok)
	}
}

func Test_SearchIntTree_Success(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	v := NewInt(13)

	// Act
	found, ok := tree.Search(v)

	// Assert
	ass.True(ok)
	ass.NotNil(found)
	ass.Equal(13, GetInt(found))
}

func Test_SearchStringTree_Success(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()
	n := NewString("intel")

	// Act
	found, ok := tree.Search(n)

	// Assert
	ass.True(ok)
	ass.NotNil(found)
	ass.Equal("intel", GetString(found))
}

func Test_SearchStringTree_Fail(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()
	n := NewString("abrakadabra")

	// Act
	found, ok := tree.Search(n)

	// Assert
	ass.False(ok)
	ass.Nil(found)
	ass.Equal("", GetString(found))
}

func Test_SearchIntTree_Failure(t *testing.T) {
	// Arrange
	ass := assert.New(t)

	var tests = []struct {
		tree RbTree
		key  Comparable
	}{
		{createIntegerTestTree(), NewInt(22)},
		{createIntegerTestTree(), nil},
		{NewRbTree(), NewInt(20)},
	}
	for _, test := range tests {
		// Act
		found, ok := test.tree.Search(test.key)

		// Assert
		ass.False(ok)
		ass.Nil(found)
	}
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
		v := NewInt(test.node)
		r, _ := tree.root.search(v)

		// Act
		s := r.successor()

		// Assert
		ass.Equal(test.expected, GetInt(s))
	}
}

func Test_SuccessorOfMax_ReturnNil(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	v := NewInt(20)
	r, _ := tree.root.search(v)

	// Act
	s := r.successor()

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
		v := NewInt(test.node)
		r, _ := tree.root.search(v)

		// Act
		s := r.predecessor()

		// Assert
		ass.Equal(test.expected, GetInt(s))
	}
}

func Test_PredecessorOfMin_ReturnNil(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()
	v := NewInt(2)
	r, _ := tree.root.search(v)

	// Act
	p := r.predecessor()

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
	ass.Equal(2, GetInt(r))
}

func Test_MinimumEmptyTree(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := NewRbTree()

	// Act
	r := tree.Minimum()

	// Assert
	ass.Nil(r)
}

func Test_Maximum_ValueAsExpected(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createIntegerTestTree()

	// Act
	r := tree.Maximum()

	// Assert
	ass.Equal(20, GetInt(r))
}

func Test_MaximumEmptyTree(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := NewRbTree()

	// Act
	r := tree.Maximum()

	// Assert
	ass.Nil(r)
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
	ass.Equal("root", GetString(x.parent))
	ass.Equal("a", GetString(x.left))
	ass.Equal("y", GetString(x.right))
	ass.Equal("b", GetString(y.left))
	ass.Equal("g", GetString(y.right))
}

func Test_LeftRotate_StructureAsExpected(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	r := newNode(NewString("root"))

	tree := newRbTree()
	tree.Insert(r)

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
	ass.Equal("root", GetString(y.parent))
	ass.Equal("x", GetString(y.left))
	ass.Equal("g", GetString(y.right))
	ass.Equal("a", GetString(x.left))
	ass.Equal("b", GetString(x.right))
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
	found, _ := tree.root.search(n)

	// Act
	tree.Delete(found)

	// Assert
	n = NewInt(28)
	found, _ = tree.root.search(n)
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
		found, _ := tree.Search(n)
		tree.Delete(found)
	}

	// Assert
	ass.Nil(tree.root)
	ass.Equal(int64(0), tree.Len())
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

	tree.WalkPreorder(func(c Comparable) {
		nod := c.(*node)
		gn := &GraphNode{Node: nod, NodeID: id}
		gr.AddNode(gn)
		id++

		for i := id - 2; i >= 0; i-- {
			n := gr.Node(i)
			if nod.parent != nil && n.(*GraphNode).Node == nod.parent {
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
	found, _ := tree.Search(n)

	// Act
	tree.Delete(found)

	// Assert
	found, ok := tree.Search(n)
	ass.False(ok)
	ass.Nil(found)

	found, ok = tree.Search(NewString("microsoft"))
	ass.True(ok)
	ass.Equal("microsoft", GetString(found))
}

func Test_DeleteEmptyTree_NoError(t *testing.T) {
	// Arrange
	tree := NewRbTree()
	n := NewString("intel")

	// Act
	tree.Delete(n)

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

func Test_InsertNil_NothingIserted(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()
	oldSize := tree.Len()

	// Act
	tree.Insert(nil)

	// Assert
	ass.Equal(oldSize, tree.Len())
}

func Test_DeleteNil_NothingDeleted(t *testing.T) {
	// Arrange
	ass := assert.New(t)
	tree := createTestStringTree()
	oldSize := tree.Len()

	// Act
	tree.Delete(nil)

	// Assert
	ass.Equal(oldSize, tree.Len())
}

func createIntegerTestTree() *rbTree {
	nodes := []int{6, 18, 3, 15, 7, 2, 4, 13, 9, 17, 20}
	return createIntTree(nodes)
}

func createTestStringTree() RbTree {
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
