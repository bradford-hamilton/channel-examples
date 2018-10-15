package main

import "fmt"

func main() {
	integers := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			integers <- i
		}
		close(integers)
	}()

	go func() {
		for x := range integers {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}
