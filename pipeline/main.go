package main

import "fmt"

func main() {
	integers := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; ; i++ {
			integers <- i
		}
	}()

	go func() {
		for {
			x := <-integers
			squares <- x * x
		}
	}()

	for {
		fmt.Println(<-squares)
	}
}
