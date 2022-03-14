package goroutine

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	// "os"
	"pratice/util"
	"sync"
	"time"
)

func printNum() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i)
	}
}

func goroutine() {
	spec := "*/1 * * * * *"
	cmd := func() {
		printNum()
	}
	util.SetCronTaskWithSecond(spec, cmd)

}

func test() {
	// 合起来写
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}

var wg sync.WaitGroup

func sync_one() {
	defer wg.Done()
	fmt.Println("one")
	spec := "*/10 * * * * *"
	cmd := func() {
		fmt.Println("one ", time.Now().Format("20060102 15:04:05"))
	}
	util.SetCronTaskWithSecond(spec, cmd)
}

func sync_two() {
	wg.Done()
	fmt.Println("two")
	spec := "*/14 * * * * *"
	cmd := func() {
		fmt.Println("two ", time.Now().Format("20060102 15:04:05"))
	}
	util.SetCronTaskWithSecond(spec, cmd)
}

func syncGroup() {
	wg.Add(2)
	go sync_one()
	go sync_two()
	wg.Wait()
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-sig

	fmt.Println("main goroutine done!", s)
	os.Exit(0)
}
