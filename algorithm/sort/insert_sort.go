package main

import "fmt"

func main() {
	arr := []int{7, 2, 24, 5, 6, 3, 868, 233}
	insert_sort(arr)
	fmt.Println(arr)
}

func insert_sort(slice []int) {
	for i := 1; i < len(slice); i++ {
		for j := i; j > 0; j-- {
			if slice[j] < slice[j-1] {
				tmp := slice[j]
				slice[j] = slice[j-1]
				slice[j-1] = tmp
			} else {
				break
			}
		}
	}
}
