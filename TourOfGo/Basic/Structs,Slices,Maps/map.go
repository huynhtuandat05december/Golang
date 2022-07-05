package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer1"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	n := m
	fmt.Println(n, m)
	n["Answer1"] = 5
	fmt.Println(n, m)

}

// map cũng reference nếu assign map1 cho map2 (map1:=map2)
