package main

import "fmt"

func main() {
	slice := []int{5, 2, 66, 34, 1, 6, 45, 56}
	min, max := getMinAndMaxNum(slice)
	resSlice := count_sort(slice, min, max)
	fmt.Println(resSlice)
}

func getMinAndMaxNum(slice []int) (min, max int) {
	min, max = slice[0], slice[0]
	for i := 0; i < len(slice); i++ {
		if slice[i] < min {
			min = slice[i]
		}
		if slice[i] > max {
			max = slice[i]
		}
	}
	return min, max
}

func count_sort(slice []int, min, max int) []int {
	countSlice := make([]int, max-min+1)
	for i := 0; i < len(slice); i++ {
		countSlice[slice[i]-min] += 1
	}

	resSlice := []int{}
	for i := 0; i < len(countSlice); i++ {
		if countSlice[i] == 0 {
			continue
		}
		for j := 0; j < countSlice[i]; j++ {
			resSlice = append(resSlice, i+min)
		}
	}
	return resSlice
}
