package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "Cristiano"
	var age int = 32
	// var age2 = 33
	// age3 := 33
	var version float32 = 1.1
	fmt.Println("Hello, Sir.", name, "Your age is", age, "and the go version is", version)
	fmt.Println("Type name is", reflect.TypeOf(name))
	fmt.Println("Type age is", reflect.TypeOf(age))
	fmt.Println("Type version is", reflect.TypeOf(version))
}
