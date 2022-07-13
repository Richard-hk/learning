package main

import "fmt"

func main() {
	arr_double := []int{2, 24, 5, 6, 3, 868, 233}
	quick_sort_double(arr_double, 0, len(arr_double)-1)
	fmt.Println("double_quick_sort", arr_double)
	arr_single := []int{2, 24, 5, 6, 3, 555, 221}
	quick_sort_single(arr_single, 0, len(arr_single)-1)
	fmt.Println("single_quick_sort", arr_single)
}

func quick_sort_double(arr []int, left, right int) {
	if left < right {
		mid := partition_double(arr, left, right)
		quick_sort_double(arr, left, mid-1)
		quick_sort_double(arr, mid+1, right)
	}
}

// 双边循环
func partition_double(arr []int, left, right int) int {
	pivot := arr[left]
	for left < right {
		for arr[right] >= pivot && left < right {
			right--
		}
		arr[left] = arr[right]
		for arr[left] < pivot && left < right {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = pivot
	return left
}

func quick_sort_single(arr []int, startIndex, endIndex int) {
	if startIndex < endIndex {
		mid := partition_single(arr, startIndex, endIndex)
		quick_sort_single(arr, startIndex, mid-1)
		quick_sort_single(arr, mid+1, endIndex)
	}
}

//单边循环
func partition_single(arr []int, startIndex, endIndex int) int {
	pivot := arr[startIndex]
	mark := startIndex
	for i := startIndex + 1; i <= endIndex; i++ {
		if arr[i] < pivot {
			mark++
			tmp := arr[mark]
			arr[mark] = arr[i]
			arr[i] = tmp
		}
	}
	arr[startIndex] = arr[mark]
	arr[mark] = pivot
	return mark
}
