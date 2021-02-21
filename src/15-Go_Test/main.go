package main

import "fmt"

func Calculate(value int) (result int) {
	result = value + 2
	return result
}

func main() {
	fmt.Println("Go Testing Tutorial")
	result := Calculate(2)
	fmt.Println(result)
}
