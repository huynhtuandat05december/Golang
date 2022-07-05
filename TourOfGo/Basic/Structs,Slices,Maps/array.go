package main

import "fmt"

func arrayFunc() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	b := a
	b[0] = "test"
	fmt.Println((b))

	arrayStruct := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Printf("len=%d cap=%d %v\n", len(arrayStruct), cap(arrayStruct), arrayStruct) //length and capital
}

func appendFunction() {
	var s []int
	s = append(s, 0)
	s = append(s, 2, 3, 4)
	fmt.Print((s))
}

func main() {
	arrayFunc()
}
