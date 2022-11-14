package main

import "fmt"

func main() {
	arr := []int{2, 24, 5, 6, 3, 868, 233}
	select_sort(arr)
	fmt.Println(arr)
}

func select_sort(slice []int) {
	for i := 0; i < len(slice); i++ {
		minIndex := i
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[minIndex] {
				minIndex = j
			}
		}
		min := slice[minIndex]
		slice[minIndex] = slice[i]
		slice[i] = min
	}
}
