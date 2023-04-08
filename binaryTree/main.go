package main

import (
	"fmt"
	"math"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func insert(root *Node, data int) *Node {
	if root == nil {
		return &Node{data: data}
	}
	if data < root.data {
		root.left = insert(root.left, data)
	} else {
		root.right = insert(root.right, data)
	}
	return root
}

func preOrder(root *Node) {
	if root != nil {
		fmt.Printf("%d ", root.data)
		preOrder(root.left)
		preOrder(root.right)
	}
}

func InOrder(root *Node) {
	if root != nil {
		InOrder(root.left)
		fmt.Printf("%d ", root.data)
		InOrder(root.right)
	}
}

func PostOrder(root *Node) {
	if root != nil {
		PostOrder(root.left)
		PostOrder(root.right)
		fmt.Printf("%d ", root.data)
	}
}

func findMax(root *Node) int {
	var max, left, right = math.MinInt, 0, 0
	if root != nil {
		left = findMax(root.left)
		right = findMax(root.right)
		if left > right {
			max = left
		} else {
			max = right
		}
		if root.data > max {
			max = root.data
		}
	}
	return max
}

func findMin(root *Node) int {
	var min, right, left = math.MaxInt, 0, 0
	if root != nil {
		left = findMin(root.left)
		right = findMax(root.right)
		if left < right {
			min = left
		} else {
			min = right
		}
		if root.data < min {
			min = root.data
		}
	}
	return min
}

func findElInTree(root *Node, data int) bool {
	temp := false
	if root == nil {
		return false
	}
	if data == root.data {
		return true
	}
	temp = findElInTree(root.left, data)
	if temp {
		return temp
	}
	return findElInTree(root.right, data)
}

func sizeOfTree(root *Node) int {
	if root == nil {
		return 0
	}
	return sizeOfTree(root.left) + 1 + sizeOfTree(root.right)
}

func heightOfTree(root *Node) int {
	var lh, rh int
	if root == nil {
		return 0
	}
	lh = heightOfTree(root.left)
	rh = heightOfTree(root.right)
	if lh > rh {
		return lh + 1
	}
	return rh + 1
}

func main() {
	root := insert(nil, 3)
	insert(root, 6)
	insert(root, 25)
	insert(root, 4)
	insert(root, 5)

	fmt.Print("preorder traversal: ")
	preOrder(root)
	fmt.Printf("\n")
	fmt.Print("inorder traversal: ")
	InOrder(root)
	fmt.Printf("\n")
	fmt.Print("postorder traversal: ")
	PostOrder(root)
	fmt.Printf("\n")
	fmt.Printf("max value of the tree is: %d\n", findMax(root))
	fmt.Printf("min value of the tree is: %d\n", findMin(root))
	fmt.Printf("1 in tree: %t\n", findElInTree(root, 1))
	fmt.Printf("4 in tree: %t\n", findElInTree(root, 4))
	fmt.Printf("size of tree: %d\n", sizeOfTree(root))
	fmt.Printf("height of the tree: %d\n", heightOfTree(root))
}
