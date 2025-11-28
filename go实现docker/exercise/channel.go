package main

import "fmt"

func main() {
	ch := make(chan int, 100)
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)

	for {
		i, ok := <-ch
		fmt.Println(i, ok)
		if !ok {
			break
		}
	}
}
