package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (g Vertex) Abs() float64 {
	g.Y++ // method for object
	return g.X + g.Y
}

func (v *Vertex) Scale() {
	v.Y++
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(v) // not reference, result {3,4}
	v.Scale()
	fmt.Println(v) // reference by scale, result {3,5}
}
