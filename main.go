package main

import "fmt"

func main() {
	// This is a single line comment

	/*
		This
		is
		a
		multiline
		comment
	*/

	fmt.Println("Hello World")
	fmt.Println("Golang", "is", "so", "cool!!!")

	var name string = "John"
	age := 25
	height := 178.95
	x, y, z := 1, 2, 3

	_ = 0
	_ = ""

	numberPointer := new(int)

	fmt.Printf("My name is %s \n", name)
	fmt.Printf("I am %d years old \n", age)
	fmt.Printf("My height is %.1fcm \n", height)
	fmt.Println(x, y, z)
	fmt.Println("Number pointer:", numberPointer)
	fmt.Println("Number pointer value:", *numberPointer)
}