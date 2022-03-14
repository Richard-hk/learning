package timer

import (
	"fmt"
	"time"
)

func timer() {
	fmt.Println(time.Now().Unix())
	t := time.NewTicker(time.Second * 3)
	select {
	case <-t.C:
		t.Stop()
	}
	fmt.Println(time.Now().Unix())
	t.Reset(time.Second * 5)

	fmt.Println(time.Now().Unix())
	select {}
}
