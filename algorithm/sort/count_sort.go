package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 3, 6, 5, 4, 55, 23, 12, 12, 33}
	min, max := getMinAndMaxNum(arr)
	count := make([]int, max-min+1)
	aa := count_sort(arr, min, count)
	fmt.Println(aa)
}

func getMinAndMaxNum(slice []int) (int, int) {
	min, max := slice[0], slice[0]
	for i := 0; i < len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
		}
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min, max
}

func count_sort(slice []int, min int, count []int) (res []int) {
	for i := 0; i < len(slice); i++ {
		count[slice[i]-min]++
	}
	for i := 0; i < len(count); i++ {
		if count[i] == 0 {
			continue
		}
		for j := 0; j < count[i]; j++ {
			res = append(res, i+min)
		}
	}
	return res
}
