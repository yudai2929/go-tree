package main

import (
	"fmt"
	"github/yudai2929/go-tree/pkg/binarytree"
	"github/yudai2929/go-tree/pkg/btree"
)

func main() {
	binaryTree := binarytree.New()
	binaryTree.Insert(5)

	binaryTree.Insert(3)
	binaryTree.Insert(7)
	binaryTree.Insert(1)
	binaryTree.Insert(4)
	binaryTree.Insert(6)
	binaryTree.Insert(8)

	binaryTree.Print()

	bTree := btree.New(3)

	bTree.Insert(5)
	bTree.Insert(3)
	bTree.Insert(7)
	bTree.Insert(1)
	bTree.Insert(4)
	bTree.Insert(6)
	bTree.Insert(8)

	found := bTree.Search(2)

	fmt.Println(found)

}
