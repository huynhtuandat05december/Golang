package main

import (
	"fmt"
	"strconv"
)

var c, python, java bool

func main() {
	var i int
	var z string = "hello"
	// or
	newString := "hello1"

	const test = "a"
	fmt.Println(i, c, python, java, z, newString)

	number := 42
	numberToString := strconv.Itoa(number)
	fmt.Println((numberToString))
}
