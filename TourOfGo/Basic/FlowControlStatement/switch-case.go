package main

import (
	"fmt"
	"time"
)

func main() {
	var a = 1
	switch a {
	case 0:
		fmt.Print(("0"))
	case 1:
		fmt.Print(("1"))
	default:
		fmt.Print(("default"))
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
