package main

import (
	"fmt"
	"time"
)

func BubleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func selectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		// Find the minimum element in the unsorted portion
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		// Swap the found minimum element with the first element in the unsorted portion
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
		fmt.Println(arr)
	}
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	final := make([]int, 0)
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for i < len(a) {
		final = append(final, a[i])
		i++
	}
	for j < len(b) {
		final = append(final, b[j])
		j++
	}
	return final
}

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]

	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}

func main() {
	a := []int{1, 2, 8, 5, 4, 7, 3}
	start := time.Now()
	//sortArray(a)
	elapsed := time.Since(start)
	seconds := elapsed.Milliseconds()
	secondsString := fmt.Sprintf("%d", seconds)
	//BubleSort(a)
	selectionSort(a)
	fmt.Println(a, secondsString)
}
