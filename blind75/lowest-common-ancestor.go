package blind75

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func newNode(item int) *Node {
	return &Node{data: item, left: nil, right: nil}
}

func lca(root *Node, n1 int, n2 int) *Node {
	for root != nil {
		if root.data > n1 && root.data > n2 {
			root = root.left
		} else if root.data < n1 && root.data < n2 {
			root = root.right
		} else {
			break
		}
	}
	return root
}

func LCABST() {
	tree := &BinaryTree{}
	tree.root = newNode(20)
	tree.root.left = newNode(8)
	tree.root.right = newNode(22)
	tree.root.left.left = newNode(4)
	tree.root.left.right = newNode(12)
	tree.root.left.right.left = newNode(10)
	tree.root.left.right.right = newNode(14)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 3; i++ {
		n1 := rand.Intn(20)
		n2 := rand.Intn(20)
		t := lca(tree.root, n1, n2)
		fmt.Printf("LCA of %d and %d is %d\n", n1, n2, t.data)
	}
}
