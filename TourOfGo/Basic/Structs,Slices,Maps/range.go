package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func range1() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func range2() {
	pow := make([]int, 10)
	fmt.Print(pow)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i // dịch trái 1 bit
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func main() {
	range2()
}
