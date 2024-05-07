package binarytree

import (
	"fmt"
	"github/yudai2929/go-tree/pkg/tree"
)

// node represents a node in a binary binaryTree.
type node struct {
	Value       int
	Left, Right *node
}

// isEmpty returns true if the node is nil.
func (n *node) isEmpty() bool {
	return n == nil
}

// binaryTree represents a binary binaryTree.
type binaryTree struct {
	root *node
}

// New returns a new binary binaryTree.
func New() tree.Tree {
	return &binaryTree{}
}

// Insert inserts a value into the binaryTree.
func (t *binaryTree) Insert(value int) {
	if t.root.isEmpty() {
		t.root = &node{Value: value}
		return
	}
	insert(t.root, value)
}

func insert(n *node, value int) {
	if value < n.Value {
		if n.Left.isEmpty() {
			n.Left = &node{Value: value}
			return
		}
		insert(n.Left, value)
		return
	}

	if value > n.Value {
		if n.Right.isEmpty() {
			n.Right = &node{Value: value}
			return
		}
		insert(n.Right, value)
		return
	}

	return
}

// Search searches for a value in the binaryTree.
func (t *binaryTree) Search(value int) bool {
	return search(t.root, value)
}

func search(n *node, value int) bool {
	if n.isEmpty() {
		return false
	}

	if value < n.Value {
		return search(n.Left, value)
	}

	if value > n.Value {
		return search(n.Right, value)
	}

	return true
}

// Delete deletes a value from the binaryTree.
func (t *binaryTree) Delete(value int) {
	t.root = delete(t.root, value)
}

func delete(n *node, value int) *node {
	if n.isEmpty() {
		return n
	}

	if value < n.Value {
		n.Left = delete(n.Left, value)
		return n
	}

	if value > n.Value {
		n.Right = delete(n.Right, value)
		return n
	}

	if n.Left.isEmpty() {
		return n.Right
	}

	if n.Right.isEmpty() {
		return n.Left
	}

	min := findMin(n.Right)
	n.Value = min.Value
	n.Right = delete(n.Right, min.Value)
	return n
}

func findMin(n *node) *node {
	if n.Left.isEmpty() {
		return n
	}

	return findMin(n.Left)
}

// InOrder returns the in-order traversal of the binaryTree.
func (t *binaryTree) InOrder() []int {
	return inOrder(t.root)
}

func inOrder(n *node) []int {
	if n.isEmpty() {
		return []int{}
	}

	return append(append(inOrder(n.Left), n.Value), inOrder(n.Right)...)
}

// Update updates a value in the binaryTree.
func (t *binaryTree) Update(oldValue, newValue int) {
	t.Delete(oldValue)
	t.Insert(newValue)
}

// Len returns the number of nodes in the binaryTree.
func (t *binaryTree) Len() int {
	return len(t.InOrder())
}

func (t *binaryTree) Height() int {
	return height(t.root)
}

func height(n *node) int {
	if n.isEmpty() {
		return 0
	}

	leftHeight := height(n.Left)
	rightHeight := height(n.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}

// LeafCount returns the number of leaf nodes in the binaryTree.
func (t *binaryTree) LeafCount() int {
	return leafCount(t.root)
}

func leafCount(n *node) int {
	if n.isEmpty() {
		return 0
	}

	if n.Left.isEmpty() && n.Right.isEmpty() {
		return 1
	}

	return leafCount(n.Left) + leafCount(n.Right)
}

// Print prints the binaryTree.
func (t *binaryTree) Print() {
	if t.root.isEmpty() {
		fmt.Println("binaryTree is empty")
		return
	}

	fmt.Println("Binary binaryTree:")
	printWithChars(t.root, "", true)
	fmt.Println("In-order traversal:", t.InOrder())
	fmt.Println("Height:", t.Height())
}

// printWithChars recursively prints the binary binaryTree with special characters.
func printWithChars(n *node, prefix string, isLeft bool) {
	if n.isEmpty() {
		return
	}

	if !n.Right.isEmpty() {
		printWithChars(n.Right, prefix+ifThenElse(isLeft, "│   ", "    "), false)
	}

	fmt.Printf("%s%s%d\n", prefix, ifThenElse(isLeft, "└── ", "┌── "), n.Value)

	if !n.Left.isEmpty() {
		printWithChars(n.Left, prefix+ifThenElse(isLeft, "    ", "│   "), true)
	}
}

// ifThenElse is a simple helper function to implement ternary operator.
func ifThenElse(condition bool, trueVal, falseVal string) string {
	if condition {
		return trueVal
	}
	return falseVal
}
