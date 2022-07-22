package main

import "fmt"

func learnStruct() {
	type person struct {
		name string
		age int
	}

	// type alias
	type newPerson = person

	p1 := newPerson{
		name: "John",
		age: 16,
	}

	// anonymous struct
	s1 := struct {
		person // embedded struct
		score float64
	} {
		person: p1,
		score: 67.70,
	}

	fmt.Println(s1)
}