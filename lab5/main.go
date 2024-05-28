package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

type BinaryTree struct {
	Root *TreeNode
}

func (t *BinaryTree) Insert(Val int) *BinaryTree {
	if t.Root == nil {
		t.Root = &TreeNode{Val: Val, Left: nil, Right: nil}
	} else {
		t.Root.Insert(Val)
	}
	return t
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		// case 1: this is a leaf node, directly delete
		if root.Left == nil && root.Right == nil {
			return nil
		}

		// case 2: it has only one child, let the one child to replace it
		if root.Left == nil && root.Right != nil {
			return root.Right
		}
		if root.Left != nil && root.Right == nil {
			return root.Left
		}

		// case 3: it has both left and right child
		if root.Left != nil && root.Right != nil {
			// Found the smallest node on the right to replace it
			minSubTreeNode := getMin(root.Right)
			leftSubTree := root.Left
			rightSubTree := deleteNode(root.Right, minSubTreeNode.Val)
			minSubTreeNode.Left = leftSubTree
			minSubTreeNode.Right = rightSubTree
			return minSubTreeNode
		}
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}

	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}

func getMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left == nil {
		return root
	}

	return getMin(root.Left)
}

func (n *TreeNode) Insert(Val int) {
	if n == nil {
		return
	} else if Val <= n.Val {
		if n.Left == nil {
			n.Left = &TreeNode{Val: Val, Left: nil, Right: nil}
		} else {
			n.Left.Insert(Val)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{Val: Val, Left: nil, Right: nil}
		} else {
			n.Right.Insert(Val)
		}
	}
}

func height(Root *TreeNode) int {
	if Root == nil {
		return 0
	}
	lh := height(Root.Left)
	rh := height(Root.Right)
	return max(lh, rh) + 1
}

func helper(Root *TreeNode, ans [][]string, h, r, c int) {
	if Root == nil {
		return
	}

	ans[r][c] = strconv.Itoa(Root.Val)

	lc := c - int(math.Pow(2, float64(h-r-1)))
	rc := c + int(math.Pow(2, float64(h-r-1)))

	helper(Root.Left, ans, h, r+1, lc)
	helper(Root.Right, ans, h, r+1, rc)
}

func printTree(Root *TreeNode) [][]string {
	h := height(Root) - 1
	rows := h + 1
	cols := int(math.Pow(2, float64(h+1))) - 1

	ans := make([][]string, rows)
	for i := 0; i < rows; i++ {
		ans[i] = make([]string, cols)
	}

	r := 0
	c := (cols - 1) / 2

	helper(Root, ans, h, r, c)
	return ans
}

func main() {
	tree := &BinaryTree{}
	tree.Insert(100).
		Insert(-20).
		Insert(-50).
		Insert(-15).
		Insert(-60).
		Insert(50).
		Insert(60).
		Insert(55).
		Insert(85).
		Insert(15).
		Insert(5).
		Insert(-10)

	output := printTree(tree.Root)
	for i := 0; i < len(output); i++ {
		fmt.Println(output[i])
	}
	deleteNode(tree.Root, 50)

	fmt.Println(strings.Repeat("-", len(output[0])))
	output = printTree(tree.Root)
	for i := 0; i < len(output); i++ {
		fmt.Println(output[i])
	}
}
