package main

import "fmt"

func slice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	slice := primes[1:4]
	var end []int = primes[2:]

	primes[3] = 1234 // slice reference
	fmt.Println(slice, primes, end)

	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func makeFunction() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5) // type,len,cap
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main() {
	slice()

}
