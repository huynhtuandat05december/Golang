package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}
func multiResult(x, y string) (string, string) {
	return x, y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(5, 6))
	fmt.Println((multiResult("a", "b")))
	fmt.Println((split(7)))
}
