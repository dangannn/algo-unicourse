package main

import "fmt"

type List interface {
	Print()
	Fill(args ...int)
	Swap(a, b int) error
	Insert(int) error
	Remove(int) error
	FindByValue(int) error
}

type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

type SinglyLinkedList struct {
	head *ListNode
}

func (ll *SinglyLinkedList) Fill(args ...int) {
	p := ll.head
	for _, v := range args {
		tmp := &ListNode{Val: v}
		p.Next = tmp
		p = p.Next
	}
}

func (ll *SinglyLinkedList) Print() {
	p := ll.head.Next
	for p != nil {
		fmt.Print(p.Val, " ")
		p = p.Next
	}
	fmt.Println()
}

func (ll *SinglyLinkedList) FindByValue(element int) (*ListNode, error) {
	p := ll.head
	for p != nil && p.Val != element {
		p = p.Next
	}
	if (p != nil && p.Val != element) || p == nil {
		return nil, fmt.Errorf("такого элемента нет")
	}
	return p, nil
}

func (ll *SinglyLinkedList) FindByIndex(index int) (*ListNode, error) {
	if index < 0 {
		return nil, fmt.Errorf("индекс выходит за рамки")
	}
	p := ll.head.Next
	for p != nil && index != 0 {
		p = p.Next
		index--
	}
	if p == nil && index > 0 {
		return nil, fmt.Errorf("индекс выходит за рамки")
	}
	return p, nil
}

func (ll *SinglyLinkedList) Swap(a, b int) error {
	if a == b {
		return nil
	}
	node1, err := ll.FindByValue(a)
	if err != nil {
		return err
	}

	node2, err := ll.FindByValue(b)
	if err != nil {
		return err
	}

	p1 := ll.head
	for p1 != nil && p1.Next.Val != node1.Val {
		p1 = p1.Next
	}

	p2 := ll.head
	for p2 != nil && p2.Next.Val != node2.Val {
		p2 = p2.Next
	}

	if p1 == nil || p2 == nil {
		return fmt.Errorf("нет таких элементов")
	}
	if node1.Next != node2 {
		tmp1 := node1.Next
		tmp2 := node2.Next

		p1.Next = node2
		node2.Next = tmp1

		p2.Next = node1
		node1.Next = tmp2
	} else {
		tmp2 := node2.Next

		p1.Next = node2
		node2.Next = node1

		node1.Next = tmp2
	}
	return nil
}

func (ll *SinglyLinkedList) Delete(a int) error {
	node, err := ll.FindByValue(a)
	if err != nil {
		return err
	}

	p := ll.head
	for p != nil && p.Next.Val != node.Val {
		p = p.Next
	}
	p.Next = node.Next
	return nil
}

func (ll *SinglyLinkedList) Insert(value int) error {
	node := &ListNode{Val: value}

	p := ll.head
	for p.Next != nil {
		p = p.Next
	}
	p.Next = node
	return nil
}

func (ll *SinglyLinkedList) SelectionSort() {
	p1 := ll.head.Next
	for p1 != nil {
		tmp := p1.Next
		minElement := p1
		p2 := p1.Next
		for p2 != nil {
			if p2.Val < minElement.Val {
				minElement = p2
			}
			p2 = p2.Next
		}
		err := ll.Swap(p1.Val, minElement.Val)
		if err != nil {
			fmt.Println("ошибка свапа")
		}
		p1 = tmp
	}

}

type DoubleLinkedList struct {
	head *ListNode
}

func (ll *DoubleLinkedList) Fill(args ...int) {
	p := ll.head
	for _, v := range args {
		tmp := &ListNode{Val: v}
		p.Next = tmp
		tmp.Prev = p
		p = p.Next
	}
}

func (ll *DoubleLinkedList) Print() {
	p := ll.head.Next
	for p != nil {
		fmt.Print(p.Val, " ")
		p = p.Next
	}
	fmt.Println()
}

func (ll *DoubleLinkedList) FindByValue(element int) (*ListNode, error) {
	p := ll.head
	for p != nil && p.Val != element {
		p = p.Next
	}
	if (p != nil && p.Val != element) || p == nil {
		return nil, fmt.Errorf("такого элемента нет")
	}
	return p, nil
}

func (ll *DoubleLinkedList) FindByIndex(index int) (*ListNode, error) {
	if index < 0 {
		return nil, fmt.Errorf("индекс выходит за рамки")
	}
	p := ll.head.Next
	for p != nil && index != 0 {
		p = p.Next
		index--
	}
	if p == nil && index > 0 {
		return nil, fmt.Errorf("индекс выходит за рамки")
	}
	return p, nil
}

func (ll *DoubleLinkedList) Swap(a, b int) error {
	if a == b {
		return nil
	}
	node1, err := ll.FindByValue(a)
	if err != nil {
		return err
	}

	node2, err := ll.FindByValue(b)
	if err != nil {
		return err
	}

	p1 := ll.head
	for p1 != nil && p1.Next.Val != node1.Val {
		p1 = p1.Next
	}

	p2 := ll.head
	for p2 != nil && p2.Next.Val != node2.Val {
		p2 = p2.Next
	}

	if p1 == nil || p2 == nil {
		return fmt.Errorf("нет таких элементов")
	}
	if node1.Next != node2 {
		tmpNext1 := node1.Next
		tmpNext2 := node2.Next

		// p1 - node1     p2 - node2
		p1.Next = node2
		node2.Prev = p1
		node2.Next = tmpNext1
		tmpNext1.Prev = node2

		p2.Next = node1
		node1.Prev = p2
		node1.Next = tmpNext2
		tmpNext2.Prev = node2
	} else {
		tmp2 := node2.Next

		p1.Next = node2
		node2.Prev = p1
		node2.Next = node1
		node1.Prev = node2

		node1.Next = tmp2
		tmp2.Prev = node1
	}
	return nil
}

func (ll *DoubleLinkedList) Delete(a int) error {
	node, err := ll.FindByValue(a)
	if err != nil {
		return err
	}

	p := ll.head
	for p != nil && p.Next.Val != node.Val {
		p = p.Next
	}
	p.Next = node.Next
	return nil
}

func (ll *DoubleLinkedList) Insert(value int) error {
	node := &ListNode{Val: value}

	p := ll.head
	for p.Next != nil {
		p = p.Next
	}
	p.Next = node
	node.Prev = p
	return nil
}

func (ll *DoubleLinkedList) SelectionSort() {
	p1 := ll.head.Next
	for p1 != nil {
		tmp := p1.Next
		minElement := p1
		p2 := p1.Next
		for p2 != nil {
			if p2.Val < minElement.Val {
				minElement = p2
			}
			p2 = p2.Next
		}
		err := ll.Swap(p1.Val, minElement.Val)
		if err != nil {
			fmt.Println("ошибка свапа")
		}
		p1 = tmp
	}

}

type BidirectionalLinkedList struct {
	head *ListNode
	tail *ListNode
}

func (ll *BidirectionalLinkedList) Fill(args ...int) {
	p := ll.head
	for _, v := range args {
		tmp := &ListNode{Val: v}
		p.Next = tmp
		tmp.Prev = p
		p = p.Next
	}
}

func (ll *BidirectionalLinkedList) Print() {
	p := ll.head.Next
	for p != nil {
		fmt.Print(p.Val, " ")
		p = p.Next
	}
	fmt.Println()
}

func (ll *BidirectionalLinkedList) FindByValue(element int) (*ListNode, error) {
	p := ll.head
	for p != nil && p.Val != element {
		p = p.Next
	}
	if (p != nil && p.Val != element) || p == nil {
		return nil, fmt.Errorf("такого элемента нет")
	}
	return p, nil
}

func (ll *BidirectionalLinkedList) FindByIndex(index int) (*ListNode, error) {
	if index < 0 {
		return nil, fmt.Errorf("индекс выходит за рамки")
	}
	p := ll.head.Next
	for p != nil && index != 0 {
		p = p.Next
		index--
	}
	if p == nil && index > 0 {
		return nil, fmt.Errorf("индекс выходит за рамки")
	}
	return p, nil
}

func (ll *BidirectionalLinkedList) Swap(a, b int) error {
	if a == b {
		return nil
	}
	node1, err := ll.FindByValue(a)
	if err != nil {
		return err
	}

	node2, err := ll.FindByValue(b)
	if err != nil {
		return err
	}

	p1 := ll.head
	for p1 != nil && p1.Next.Val != node1.Val {
		p1 = p1.Next
	}

	p2 := ll.head
	for p2 != nil && p2.Next.Val != node2.Val {
		p2 = p2.Next
	}

	if p1 == nil || p2 == nil {
		return fmt.Errorf("нет таких элементов")
	}
	if node1.Next != node2 {
		tmpNext1 := node1.Next
		tmpNext2 := node2.Next

		// p1 - node1     p2 - node2
		p1.Next = node2
		node2.Prev = p1
		node2.Next = tmpNext1
		tmpNext1.Prev = node2

		p2.Next = node1
		node1.Prev = p2
		node1.Next = tmpNext2
		tmpNext2.Prev = node2
	} else {
		tmp2 := node2.Next

		p1.Next = node2
		node2.Prev = p1
		node2.Next = node1
		node1.Prev = node2

		node1.Next = tmp2
		tmp2.Prev = node1
	}
	return nil
}

func (ll *BidirectionalLinkedList) Delete(a int) error {
	node, err := ll.FindByValue(a)
	if err != nil {
		return err
	}

	p := ll.head
	for p != nil && p.Next.Val != node.Val {
		p = p.Next
	}
	p.Next = node.Next
	return nil
}

func (ll *BidirectionalLinkedList) Insert(value int) error {
	node := &ListNode{Val: value}

	p := ll.head
	for p.Next != nil {
		p = p.Next
	}
	p.Next = node
	node.Prev = p
	return nil
}

func (ll *BidirectionalLinkedList) SelectionSort() {
	p1 := ll.head.Next
	for p1 != nil {
		tmp := p1.Next
		minElement := p1
		p2 := p1.Next
		for p2 != nil {
			if p2.Val < minElement.Val {
				minElement = p2
			}
			p2 = p2.Next
		}
		err := ll.Swap(p1.Val, minElement.Val)
		if err != nil {
			fmt.Println("ошибка свапа")
		}
		p1 = tmp
	}

}

func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	curr := head

	// Проходим по всему списку
	for curr != nil {
		// Сохраняем ссылку на следующий узел, так как после вставки текущего узла он будет удален из списка
		next := curr.Next
		// Ищем место для вставки текущего узла в отсортированную часть списка
		prev := dummy
		for prev.Next != nil && prev.Next.Val < curr.Val {
			prev = prev.Next
		}
		// Вставляем текущий узел в отсортированную часть списка
		curr.Next = prev.Next
		prev.Next = curr
		// Переходим к следующему узлу
		curr = next
	}

	// Возвращаем начало отсортированного списка
	return dummy.Next
}

func main() {
	list1 := SinglyLinkedList{head: &ListNode{}}

	fmt.Println("Неотсортированный список: ")
	list1.Fill(76, 2, 4, 3, 1)
	list1.Print()

	fmt.Println("Отсортированный список: ")
	insertionSortList(list1.head)
	list1.Print()

	fmt.Println("Swap элементов 2 и 3 список: ")
	a, b := 2, 3
	list1.Swap(a, b)
	list1.Print()

	element, err := list1.FindByIndex(3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(element.Val)
	}

	list1.Delete(-1)
	list1.Print()
	list1.Insert(-1)
	list1.Print()
	list1.SelectionSort()
	list1.Print()

	fmt.Println("Двусвязный")
	list2 := DoubleLinkedList{head: &ListNode{}}
	list2.Fill(1, 2, 3, 4, 76)
	list2.Print()
	list2.Swap(3, 4)
	list2.Print()
	list2.SelectionSort()
	list2.Print()
}
