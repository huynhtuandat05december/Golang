package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}

	channel := make(chan int, 5)
	channel2 := make(chan int, 6)
	wg.Add(2)
	// go func(channel <-chan int) {
	// 	for i := range channel {
	// 		fmt.Println((i))
	// 	}
	// 	wg.Done()
	// }(channel)

	// go func(channel chan<- int) {
	// 	channel <- 21
	// 	channel <- 22
	// 	channel <- 23
	// 	close(channel)
	// 	wg.Done()
	// }(channel)

	go func() {
		for {
			select {
			case <-channel:
				v1 := <-channel
				fmt.Println(v1)

			case <-channel2:
				v2 := <-channel2
				fmt.Println(v2)
			default:
				fmt.Println("    .")
				time.Sleep(50 * time.Millisecond)
			}

		}
		wg.Done()

	}()

	go func() {
		channel <- 21
		channel <- 22
		channel <- 23
		close(channel)
		channel2 <- 24
		channel2 <- 25
		channel2 <- 26
		close(channel2)
		wg.Done()
	}()

	wg.Wait()
}
