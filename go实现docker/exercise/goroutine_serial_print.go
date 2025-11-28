package main

import (
	"fmt"
	"sync"
	"time"
)

// 多个goroutine顺序打印
func main() {
	case1()
}

func case1() {
	var counter int
	var m = &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock()
			defer m.Unlock()
			counter++
			fmt.Printf("Goroutine: %d, result: %d\n", i, counter)
		}(i)
	}

	time.Sleep(time.Second * 10)
}

func case2() {
	var counter int
	var ch = make(chan struct{}, 1)

	for i := 0; i < 10; i++ {
		go func(i int) {
			ch <- struct{}{}
			counter++
			fmt.Printf("Goroutine: %d, result: %d\n", i, counter)
			<-ch
		}(i)
	}

	time.Sleep(time.Second * 10)
}
