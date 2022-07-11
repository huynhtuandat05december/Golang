package main

import (
	"fmt"
	"singleton_pattern/singleton"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%p \n", singleton.GetInstance())
	}
}
