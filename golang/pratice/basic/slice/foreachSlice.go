package slice

import "fmt"

func foreachSlice() { // 遍历数组赋值会出错
	slice := []int{0, 1, 2, 4}
	testMap := make(map[int]*int)
	for index, sl := range slice {
		testMap[index] = &sl // 每次赋值都是sl变量的地址，sl最后一次的值是4
	}
	for k, v := range testMap {
		fmt.Println(k, "===>", *v)
	}
}

func foreachSliceWithGo() { // 遍历数组赋值会出错
	slice := []int{0, 1, 2, 4}
	for _, num := range slice {
		go func(num int) {
			fmt.Println(num)
		}(num)
	}
	for {

	}
}

func appendTest() {
	// list := new([]int)
	// list.append(list,1) //new([]int) 之后的 list 是一个 *[]int 类型的指针

	list := make([]int, 10)
	for i, v := range list {
		fmt.Println(i, v)
	}

}
