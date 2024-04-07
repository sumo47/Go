package main

import (
	"fmt"
	myPackage "mylearning/package"
)

func main() {
	fmt.Println("Hello world!")
	myPackage.PrintFunction("hello")

	var variable = "sumit"
	myPackage.PrintFunction(variable)

	var bool bool = false
	fmt.Println(bool)

	Public := "kumar"
	fmt.Println(Public)

	const age = "kumar"
	fmt.Println(age)

	//* print statement
	fmt.Println("Hello world!") // added new line after code completion
	fmt.Printf("Hello world!")  // exactly works like C language printf function
}
