package main

import (
	"fmt"
	"math/rand"
	"sort"
)

const SIZE = 10

type ArrayOrdered struct {
	array [SIZE]int
}

func (a *ArrayOrdered) fillArray() {
	tmpArray := make([]int, SIZE)
	for i := 0; i < SIZE; i++ {
		tmpArray[i] = rand.Intn(100)
	}
	sort.Ints(tmpArray)
	copy(a.array[:], tmpArray)
}

func (a *ArrayOrdered) searchElement(element int) (int, error) {
	l := 0
	r := len(a.array) - 1
	mid := 0
	for l <= r {
		mid = (l + r) / 2
		if a.array[mid] >= element {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l != -1 {
		return l, nil
	}
	return -1, fmt.Errorf("в массиве нет такого элемента: %d", element)
}

func (a *ArrayOrdered) insertElement(element int) error {
	idx := -1
	for i := range a.array {
		if a.array[i] < element {
			i++
		} else {
			idx = i
			break
		}
	}
	if idx == -1 {
		return fmt.Errorf("вставка данного элемента невозможна: %d", element)
	}

	for i := len(a.array) - 1; i > idx; i-- {
		a.array[i] = a.array[i-1]
	}
	a.array[idx] = element
	return nil
}

func (a *ArrayOrdered) deleteElement(element int) error {
	idx, err := a.searchElement(element)
	if err != nil {
		return err
	}
	for i := idx; i < len(a.array)-1; i++ {
		a.array[i] = a.array[i+1]
	}
	a.array[len(a.array)-1] = 0
	return nil
}
func main() {
	arrays := ArrayOrdered{}
	arrays.fillArray()
	fmt.Printf("Массив:%v\n", arrays.array)
	for {
		fmt.Println("Введите команду")
		var command string
		fmt.Scan(&command)
		switch command {
		case "найти":
			fmt.Println("Введите элемент для поиска")
			var element int
			fmt.Scan(&element)
			idx, err := arrays.searchElement(element)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("Позиция элемента: ", idx)
		case "удалить":
			fmt.Println("Введите элемент для удаления")
			var element int
			fmt.Scan(&element)
			err := arrays.deleteElement(element)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(arrays.array)
		case "вставить":
			fmt.Println("Введите элемент для вставки")
			var element int
			fmt.Scan(&element)
			err := arrays.insertElement(element)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(arrays.array)
		case "закончить":
			fmt.Println("Конец программы")
			return
		default:
			fmt.Println("Несуществующая команда")
		}
	}
}
