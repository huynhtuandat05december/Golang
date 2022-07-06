package main

import "fmt"

func generatePipeline(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, value := range numbers {
			out <- value
		}
		close(out)
	}()
	return out
}

func fanIn(channels ...<-chan int) <-chan int {
	in := make(chan int)
	go func() {
		for _, c := range channels {
			for n := range c {
				in <- n
			}
		}
		close(in)
	}()
	return in

}

func fanOut(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for value := range in {
			out <- value * value
		}
		close(out)
	}()
	return out

}

func main() {
	n := 100000000
	numbers := []int{}
	for i := 1; i <= n; i++ {
		numbers = append(numbers, i)
	}

	// sum := 0

	// for value := range numbers {
	// 	sum += value * value
	// }

	inputChannel := generatePipeline(numbers)

	c1 := fanOut((inputChannel))
	c2 := fanOut((inputChannel))
	c3 := fanOut((inputChannel))
	c4 := fanOut((inputChannel))

	channelResult := fanIn(c1, c2, c3, c4)

	sum := 0

	for value := range channelResult {
		sum += value
	}

	fmt.Println(sum)

}
