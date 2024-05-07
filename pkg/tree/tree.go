package tree

// Tree represents a binary binaryTree.
type Tree interface {
	Insert(value int)
	Search(value int) bool
	Delete(value int)
	InOrder() []int
	Print()
}
