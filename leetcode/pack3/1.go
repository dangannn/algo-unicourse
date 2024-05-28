package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Если оба узла пустые, то деревья идентичны
	if p == nil && q == nil {
		return true
	}

	// Если один из узлов пустой, а другой не пустой, или значения узлов не равны,
	// то деревья не идентичны
	if p == nil && q != nil || p != nil && q == nil || p.Val != q.Val {
		return false
	}

	// Рекурсивно проверяем идентичность левых и правых поддеревьев
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
