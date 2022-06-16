package main

import "fmt"

func main() {
	ch := generator()

	for i := 0; i < 5; i++ {
		fmt.Println("Test 1:", <-ch)
		value := <-ch
		fmt.Println("Test 2:", value)
		fmt.Println("Test 3:", <-ch)
	}
}

func generator() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 1; ; i++ {
			ch <- i
		}
	}()

	return ch
}
