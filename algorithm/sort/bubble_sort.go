package main

import "fmt"

func main() {
	arr := []int{2, 24, 5, 6, 3, 868, 233}
	bubble_sort(arr)
	fmt.Println(arr)
}
func bubble_sort(arr []int) {
	len := len(arr)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[i] {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}
}
