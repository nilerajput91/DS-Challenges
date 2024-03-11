//self balance tree

package main

import "fmt"

func main() {

	tree := &AVLTree{}
	keys := []int{10, 20, 30, 40, 50, 25}

	for _, keys := range keys {
		tree.
			insert(keys)
	}

	fmt.Println("Inorder Traversal:")
	tree.InorderTrversal()
}

type Node struct {
	Key    int
	Height int
	Left   *Node
	Right  *Node
}

type AVLTree struct {
	Root *Node
}

func NewNode(key int) *Node {
	return &Node{
		Key:    key,
		Height: 1,
	}
}

func Height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b

}

func balanceFactor(node *Node) int {
	return Height(node.Left) - Height(node.Right)
}

func updateHeight(node *Node) {
	node.Height = 1 + (max(Height(node.Left), Height(node.Right)))
}

func rigthRoatate(y *Node) *Node {
	x := y.Left
	T2 := x.Right

	//Perform roatations
	x.Right = y
	y.Left = T2

	//Update Height

	updateHeight(y)
	updateHeight(x)

	return x

}

func LeftRotate(x *Node) *Node {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	updateHeight(x)
	updateHeight(y)

	return y

}

func insert(root *Node, key int) *Node {
	if root == nil {
		return NewNode(key)
	}

	if key < root.Key {
		root.Left = insert(root.Left, key)
	} else if key > root.Key {
		root.Right = insert(root.Right, key)
	} else {
		return root
	}

	updateHeight(root)

	balance := balanceFactor(root)

	//left left case
	if balance > 1 && key < root.Left.Key {
		return rigthRoatate(root)
	}

	//right right case
	if balance < -1 && key > root.Right.Key {
		return LeftRotate(root)
	}

	//left right case

	if balance > 1 && key > root.Left.Key {
		root.Left = LeftRotate(root.Left)

		return rigthRoatate(root)
	}

	// right left case

	if balance < -1 && key < root.Right.Key {
		root.Right = rigthRoatate(root.Right)
		return LeftRotate(root)

	}

	return root
}

func (t *AVLTree) insert(key int) {
	t.Root = insert(t.Root, key)
}

func inorderTrversal(root *Node) {
	if root != nil {
		inorderTrversal(root.Left)
		fmt.Printf("%d ", root.Key)

		inorderTrversal(root.Right)
	}
}

func (t *AVLTree) InorderTrversal() {
	inorderTrversal(t.Root)
	fmt.Println()

}
