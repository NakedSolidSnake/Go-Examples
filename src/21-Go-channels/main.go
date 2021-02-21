package main

import "fmt"

func SendValue(c chan int) {
	c <- 8
}

func main() {
	fmt.Println("Go Channels Tutorial")

	values := make(chan int)
	defer close(values)

	go SendValue(values)

	value := <-values
	fmt.Println(value)
}
