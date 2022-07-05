package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, s)
		time.Sleep(time.Second)
	}
}

func main() {
	go say("hello")
	say("world")
}

// goroutines tạo ra 1 tiến trình độc lập => say("hello") và say("world") sẽ chạy song song
// main cũng là 1 goroutines, khi hàm kết thúc thì goroutines đó sẽ dừng, không quan tâm đến các goroutines khác
