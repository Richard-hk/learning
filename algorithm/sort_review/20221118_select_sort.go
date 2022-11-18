package main

import "fmt"

func main() {
	slice := []int{5, 2, 66, 34, 1, 6, 45, 56}
	select_sort(slice)
	fmt.Println(slice)
}

func select_sort(slice []int) {
	for i := 0; i < len(slice); i++ {
		index := i
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[index] {
				index = j
			}
		}
		tmp := slice[index]
		slice[index] = slice[i]
		slice[i] = tmp
	}
}
