package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	root := &ListNode{}
	node := root
	// сравниваем 1 элементы списков и сливаем меньший с переходом на следующий
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			node.Next = list1
			node = node.Next
			list1 = list1.Next
		} else {
			node.Next = list2
			node = node.Next
			list2 = list2.Next
		}
	}
	// остатки 1 списка
	for list1 != nil {
		node.Next = list1
		node = node.Next
		list1 = list1.Next
	}
	// остатки 2 списка
	for list2 != nil {
		node.Next = list2
		node = node.Next
		list2 = list2.Next
	}
	return root.Next
}
