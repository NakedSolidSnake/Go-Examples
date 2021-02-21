package main

import (
	"fmt"

	_ "30-MultipleFiles/domain"
)

func main() {

	var person Person
	person.Age = 34
	person.Name = "Cristiano"

	fmt.Println(person)

}
