package btree

// node represents a node in a binary binaryTree.
type node struct {
	Values   values
	Children children
}

type values []int

func (v values) isEmpty() bool {
	return len(v) == 0
}

type children []*node

func (c children) isEmpty() bool {
	return len(c) == 0
}

// isEmpty returns true if the node is empty.
func (n *node) isEmpty() bool {
	return n == nil
}

// binaryTree represents a binary binaryTree.
type bTree struct {
	root      *node
	minDegree int
}

func New(minDegree int) *bTree {
	return &bTree{minDegree: minDegree}
}

func (t *bTree) Insert(value int) {
	if t.root.isEmpty() {
		t.root = &node{Values: values{value}}
		return
	}
	insert(t.root, value, t.minDegree)
}

func insert(n *node, value int, minDegree int) {
	if n.Values.isEmpty() {
		n.Values = append(n.Values, value)
		return
	}

	if n.Children.isEmpty() {
		insertIntoNode(n, value)
		return
	}

	// find the child to insert into
	var i int
	for i = 0; i < len(n.Values); i++ {
		if value < n.Values[i] {
			break
		}
	}
	insert(n.Children[i], value, minDegree)
}

func insertIntoNode(n *node, value int) {
	for i := 0; i < len(n.Values); i++ {
		if value < n.Values[i] {
			n.Values = append(n.Values[:i], append([]int{value}, n.Values[i:]...)...)
			return
		}
	}
	n.Values = append(n.Values, value)
}

func (t *bTree) Search(value int) bool {
	return search(t.root, value)
}

func search(n *node, value int) bool {
	if n.Values.isEmpty() {
		return false
	}

	for i := 0; i < len(n.Values); i++ {
		if value == n.Values[i] {
			return true
		}
		if value < n.Values[i] {
			if n.Children.isEmpty() {
				return false
			}
			return search(n.Children[i], value)
		}
	}

	if n.Children.isEmpty() {
		return false
	}
	return search(n.Children[len(n.Children)-1], value)
}

func (t *bTree) Delete(value int) {
	t.root = delete(t.root, value)
}

func delete(n *node, value int) *node {
	if n.Values.isEmpty() {
		return n
	}

	for i := 0; i < len(n.Values); i++ {
		if value == n.Values[i] {
			if n.Children.isEmpty() {
				n.Values = append(n.Values[:i], n.Values[i+1:]...)
				return n
			}
			// replace the value with the predecessor
			n.Values[i] = deletePredecessor(n.Children[i])
			return n
		}
		if value < n.Values[i] {
			if n.Children.isEmpty() {
				return n
			}
			n.Children[i] = delete(n.Children[i], value)
			return n
		}
	}

	if n.Children.isEmpty() {
		return n
	}
	n.Children[len(n.Children)-1] = delete(n.Children[len(n.Children)-1], value)
	return n
}

func deletePredecessor(n *node) int {
	if n.Children.isEmpty() {
		predecessor := n.Values[len(n.Values)-1]
		n.Values = n.Values[:len(n.Values)-1]
		return predecessor
	}
	return deletePredecessor(n.Children[len(n.Children)-1])
}

func (t *bTree) InOrder() []int {
	return inOrder(t.root)
}

func inOrder(n *node) []int {
	if n.isEmpty() {
		return nil
	}

	var result []int
	for i := 0; i < len(n.Values); i++ {
		result = append(result, inOrder(n.Children[i])...)
		result = append(result, n.Values[i])
	}
	result = append(result, inOrder(n.Children[len(n.Children)-1])...)
	return result
}
