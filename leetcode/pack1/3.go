package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	// создаем переменную для хранения начала нового списка
	cur := head
	for cur != nil {
		// сохраняем ссылку на следующий элемент
		tmp := cur.Next
		// цепляем начало нового списка к текущему узлу
		cur.Next = prev
		// передвигаем указатель на новый список на текующий узел
		prev = cur
		// передвигаемся вперед в старом списке
		cur = tmp
	}
	return prev
}
