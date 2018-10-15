package main

import "fmt"

func main() {
	integers := make(chan int)
	squares := make(chan int)

	go counter(integers)
	go squarer(squares, integers)
	printer(squares)
}

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
