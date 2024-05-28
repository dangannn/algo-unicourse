package main

func removeElements(head *ListNode, val int) *ListNode {
	// проверка входных данных
	for head != nil && head.Val == val {
		head = head.Next
	}

	p := head

	// проходимся по всем элементам
	for p != nil && p.Next != nil {
		// если следующий элемент имеет нужое значение цепляем p.Next = p.Next.Next
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}
