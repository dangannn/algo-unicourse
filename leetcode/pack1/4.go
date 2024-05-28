package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swap(node *ListNode, prev *ListNode) {
	if node == nil {
		return
	}
	tmp := node.Next
	if node.Next != nil {
		node.Next = tmp.Next  // Устанавливаем ссылку текущего узла на узел, следующий за следующим узлом
		tmp.Next = node       // Устанавливаем ссылку временного узла на текущий узел
		prev.Next = tmp       // Устанавливаем ссылку предыдущего узла на временный узел
		swap(node.Next, node) // Рекурсивный вызов функции для следующей пары узлов
	}
}

func swapPairs(head *ListNode) *ListNode {
	// проверка ввода
	if head == nil || head.Next == nil {
		return head
	}
	prev := &ListNode{}
	swap(head, prev)
	return prev.Next
}
