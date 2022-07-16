package main

import "fmt"

func learnVariable() {
	var name string = "James"
	var weight = 65
	height := 178.65

	name = "Andi"
	height = 180.29

	x, y, z := 1, 2, 3

	const age = 25
	const (
		const1 = "a"
		const2 string = "b"
	) 

	fmt.Printf("My name is %s \n", name)
	fmt.Printf("My weight is %dkg \n", weight)
	fmt.Printf("My height is %.1fcm \n", height)
	fmt.Printf("I am %d years old \n", age)
	fmt.Println(x, y, z)
	fmt.Println(const1, const2)
}