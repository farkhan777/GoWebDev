package main

import "fmt"

func main() {
	c := counter()

	for i := 0; ; i++ {
		fmt.Println("Value 1: ", <-c)
		fmt.Println("Value 2: ", <-c)
		fmt.Println("Value 3: ", <-c)
	}
}

func counter() chan int {
	c := make(chan int)

	go func() {
		for i := 1; ; i++ {
			c <- i
		}
	}()

	return c
}
