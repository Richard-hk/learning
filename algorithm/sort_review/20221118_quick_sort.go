package main

import "fmt"

func main() {
	slice := []int{5, 2, 66, 34, 1, 6, 45, 56}
	quick_sort(slice, 0, len(slice)-1)
	fmt.Println(slice)
}

func quick_sort(slice []int, left, right int) {
	if left < right {
		mid := partition(slice, left, right)
		quick_sort(slice, left, mid-1)
		quick_sort(slice, mid+1, right)
	}

}

func partition(slice []int, left, right int) int {
	pivot := slice[left]
	for left < right {
		for left < right && slice[right] > pivot {
			right--
		}
		slice[left] = slice[right]
		for left < right && slice[left] < pivot {
			left++
		}
		slice[right] = slice[left]

	}
	slice[left] = pivot
	return left
}
