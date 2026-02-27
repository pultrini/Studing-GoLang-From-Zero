package structures

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Rigth *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

func (t *BinaryTree) Insert(val int) {
	if t.Root == nil {
		t.Root = &TreeNode{Value: val}
		return
	} else {
		t.Root.insertNode(val)
	}
}

func (n *TreeNode) insertNode(val int) {
	if n == nil {
		return
	}
	if val < n.Value {
		if n.Left == nil {
			n.Left = &TreeNode{Value: val}
		} else {
			n.Left.insertNode(val)
		}
	} else if val > n.Value {
		n.Rigth.insertNode(val)
	}
}

func (t *BinaryTree) Search(val int) bool {
	return searchRecursive(t.Root, val)
}

func searchRecursive(n *TreeNode, val int) bool {
	if n == nil {
		return false
	}
	if n.Value == val {
		return true
	}
	if val < n.Value {
		return searchRecursive(n.Left, val)
	}
	return searchRecursive(n.Rigth, val)
}

func (t *BinaryTree) PrintInOrder() {
	printRecursive(t.Root)
	fmt.Println()
}

func printRecursive(n *TreeNode) {
	if n == nil {
		return
	}
	printRecursive(n.Left)
	fmt.Printf("%d", n.Value)
	printRecursive(n.Rigth)
}
