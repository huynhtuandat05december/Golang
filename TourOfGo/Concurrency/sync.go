package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var mutex = sync.RWMutex{}
var counter int = 0

func say() {
	fmt.Println("Hello %v", counter)
	mutex.RUnlock()
	wg.Done()

}

func increase() {
	counter++
	mutex.Unlock()
	wg.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		mutex.RLock() // read lock
		go say()
		mutex.Lock() //write lock
		go increase()
	}
	wg.Wait()

}

// Dùng dùng biến chung cho các goroutines thì phải dùng lock để kiểm soát các biến chung
