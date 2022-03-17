package goroutine

import (
	"pratice/util"
	"time"
)

// 排队与队列

// 设置最大GOMAXPROCS为1，限制任意时刻一个M只能执行一个goroutine，go线程如何调用的

func line_schedule() {
	// runtime.GOMAXPROCS(1)
	go util.PrintNum(1)
	go util.PrintNum(2)
	go util.PrintNum(3)
	time.Sleep(time.Second * 5)
	// 输出结果
	// print_num =>  3
	// print_num =>  1
	// print_num =>  2
	// 原因
	// runnext(待执行的goroutine)  runq(本地goroutine执行队列)
	// 1先进runnext， 2把1挤入runq，3把2挤入runq
	// runnext的值为3，runq的值分别是1，2
	// 所以执行顺序是3，1，2
}
