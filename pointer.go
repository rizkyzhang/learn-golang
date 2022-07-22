package main

import "fmt"

func learnPointer() {
	numberA := 1
	numberB := &numberA

	fmt.Println("numberA")
	fmt.Println("value:", numberA)
	fmt.Println("addres:", &numberA)
	fmt.Println("")

	fmt.Println("numberB")
	fmt.Println("value:", *numberB)
	fmt.Println("addres:", numberB)
	fmt.Println("")

	fmt.Println("Pointer test")
	num := 5

	fmt.Println("old num:", num)
	changeVarValue(&num, 10)
	fmt.Println("new num:", num)
}

func changeVarValue(varPointer *int, newValue int) {
	*varPointer = newValue
}