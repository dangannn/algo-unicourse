package main

import (
	"fmt"
)

type Heap struct {
	heap []int
}

func (h *Heap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.heap[parentIndex] < h.heap[index] {
			h.heap[parentIndex], h.heap[index] = h.heap[index], h.heap[parentIndex]
			index = parentIndex
		} else {
			break
		}
	}
}

func (h *Heap) heapifyDown(index int) {
	leftChildIndex := 2*index + 1
	rightChildIndex := 2*index + 2
	largestIndex := index

	if leftChildIndex < len(h.heap) && h.heap[leftChildIndex] > h.heap[largestIndex] {
		largestIndex = leftChildIndex
	}
	if rightChildIndex < len(h.heap) && h.heap[rightChildIndex] > h.heap[largestIndex] {
		largestIndex = rightChildIndex
	}
	if largestIndex != index {
		h.heap[index], h.heap[largestIndex] = h.heap[largestIndex], h.heap[index]
		h.heapifyDown(largestIndex)
	}
}

func (h *Heap) Insert(value int) {
	h.heap = append(h.heap, value)
	h.heapifyUp(len(h.heap) - 1)
}

func (h *Heap) Search(value int) bool {
	for _, v := range h.heap {
		if v == value {
			return true
		}
	}
	return false
}

func (h *Heap) Remove(value int) bool {
	for i, v := range h.heap {
		if v == value {
			h.heap[i] = h.heap[len(h.heap)-1]
			h.heap = h.heap[:len(h.heap)-1]
			h.heapifyDown(i)
			return true
		}
	}
	return false
}

func (h *Heap) HeapSort() {
	originalSize := len(h.heap)
	for i := len(h.heap)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
	for i := len(h.heap) - 1; i > 0; i-- {
		h.heap[0], h.heap[i] = h.heap[i], h.heap[0]
		h.heap = h.heap[:i]
		h.heapifyDown(0)
	}
	h.heap = h.heap[:originalSize]
}

func (h *Heap) Print() {
	for _, v := range h.heap {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	var size int
	fmt.Print("Введите размер кучи: ")
	fmt.Scan(&size)

	heap := &Heap{}

	fmt.Printf("Введите %d элементов: ", size)
	for i := 0; i < size; i++ {
		var element int
		fmt.Scan(&element)
		heap.Insert(element)
	}

	fmt.Print("Куча: ")
	heap.Print()

	var insertValue int
	fmt.Print("Введите значение для вставки: ")
	fmt.Scan(&insertValue)
	heap.Insert(insertValue)

	var searchValue int
	fmt.Print("Введите значение для поиска: ")
	fmt.Scan(&searchValue)
	if heap.Search(searchValue) {
		fmt.Println("Значение найдено в куче.")
	} else {
		fmt.Println("Значение не найдено в куче.")
	}

	var removeValue int
	fmt.Print("Введите значение для удаления: ")
	fmt.Scan(&removeValue)
	if heap.Remove(removeValue) {
		fmt.Println("Значение удалено из кучи.")
	} else {
		fmt.Println("Значение не найдено в куче.")
	}

	fmt.Print("Куча после удаления: ")
	heap.Print()

	heap.HeapSort()
	fmt.Print("Сортировка массива: ")
	heap.Print()
}
