package main

import "fmt"

func buildMaxHeap(arr []int) {
	n := len(arr)
	for i := (n - 1) / 2; i >= 0; i-- {
		heapify(arr, n, i)
	}
}

func swap(arr []int, index1 int, index2 int) {
	arr[index1], arr[index2] = arr[index1], arr[index1]
}

func heapify(arr []int, n int, i int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < n && arr[largest] < arr[left] {
		swap(arr, left, largest)
	}
	if right < n && arr[i] < arr[right] {
		swap(arr, right, largest)
	}
	if i != largest {
		swap(arr, i, largest)
		heapify(arr, n, largest)
	}
}

func printArr(arr []int) {
	for _, value := range arr {
		fmt.Print(value)
	}
}

func main() {
	arr := []int{1, 2, 3, 2, 3, 5, 7, 1, 4, 5, 2}
	buildMaxHeap(arr)
	printArr(arr)
}
