package main

import "fmt"

func main() {
	slice := []int{5, 2, 66, 34, 1, 6, 45, 56}
	bubble_sort(slice)
	fmt.Println(slice)
}

func bubble_sort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				tmp := slice[j+1]
				slice[j+1] = slice[j]
				slice[j] = tmp
			}
		}
	}
}
