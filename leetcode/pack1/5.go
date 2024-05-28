package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tmp := result
	for l1 != nil || l2 != nil {
		// Если первый список еще не достиг конца
		if l1 != nil {
			tmp.Val += l1.Val // Добавляем значение текущего узла первого списка к значению временного узла
			l1 = l1.Next      // Переходим к следующему узлу первого списка
		}
		// Если второй список еще не достиг конца
		if l2 != nil {
			tmp.Val += l2.Val // Добавляем значение текущего узла второго списка к значению временного узла
			l2 = l2.Next      // Переходим к следующему узлу второго списка
		}
		// Если сумма значений текущих узлов больше 9
		if tmp.Val > 9 {
			tmp.Val -= 10         // Вычитаем 10 из значения временного узла
			tmp.Next = &ListNode{ // Создаем новый узел с значением 1
				Val: 1,
			}
		} else if l1 != nil || l2 != nil { // Если есть еще узлы для обработки
			tmp.Next = &ListNode{} // Создаем новый пустой узел
		}
		tmp = tmp.Next // Переходим к следующему узлу временного списка
	}
	return result
}
