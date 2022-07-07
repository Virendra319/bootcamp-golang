package ExercisesDay1

import "fmt"

type Node struct {
	val         string
	left, right *Node
}

func (node *Node) preOrder() {
	if node == nil {
		return
	}
	fmt.Print(node.val)
	node.left.preOrder()
	node.right.preOrder()
}

func (node *Node) postOrder() {
	if node == nil {
		return
	}
	node.left.postOrder()
	node.right.postOrder()
	fmt.Print(node.val)
}

func Problem2() {
	var node2, node4, node5 Node
	node2.val, node4.val, node5.val = "a", "b", "c"
	node3 := Node{"-", &node4, &node5}
	node1 := Node{"+", &node2, &node3}

	node1.preOrder()
	fmt.Println()
	node1.postOrder()
}
