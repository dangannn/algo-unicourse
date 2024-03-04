package main

import (
	"fmt"
	"math/rand"
)

const SIZE = 10

type ArrayRandom struct {
	array [SIZE]int
}

func (a *ArrayRandom) fillArray() {
	for i := 0; i < SIZE; i++ {
		a.array[i] = rand.Intn(100)
	}
}

func (a *ArrayRandom) searchElement(element int) (int, error) {
	for i, v := range a.array {
		if v == element {
			return i, nil
		}
	}
	return -1, fmt.Errorf("в массиве нет такого элемента: %d", element)
}

func (a *ArrayRandom) insertElement(idx int, element int) error {
	if idx < 0 || idx > len(a.array) {
		return fmt.Errorf("некорректный индекс: %d", idx)
	}

	for i := len(a.array) - 1; i > idx; i-- {
		a.array[i] = a.array[i-1]
	}
	a.array[idx] = element
	return nil
}

func (a *ArrayRandom) deleteElement(element int) error {
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
	arrays := ArrayRandom{}
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
			fmt.Println("Введите позицию для вставки")
			var idx int
			fmt.Scan(&idx)
			err := arrays.insertElement(idx, element)
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
