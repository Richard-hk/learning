package main

import "fmt"

func main() {
	slice := []int{5, 2, 66, 34, 1, 6, 45, 56}
	min, max := getMinAndMaxNum(slice)
	bucket_sort(slice, min, max)
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

func bucket_sort(slice []int, min, max int) {
	bucketSize := (max-min)/len(slice) + 1
	bucketSlice := make([][]int, bucketSize)
	for i := 0; i < len(slice); i++ {
		bucketSlice[slice[i]/bucketSize] = append(bucketSlice[slice[i]/bucketSize], slice[i])
	}
	resSlice := []int{}
	for i := 0; i < len(bucketSlice); i++ {
		if len(bucketSlice[i]) == 0 {
			continue
		}
		select_sort(bucketSlice[i])
		resSlice = append(resSlice, bucketSlice[i]...)
	}
	fmt.Println(bucketSlice)
	fmt.Println(resSlice)
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
