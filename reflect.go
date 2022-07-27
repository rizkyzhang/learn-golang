package main

import (
	"fmt"
	"reflect"
)

func learnReflect() {
	x := 2
	valueOfX := reflect.ValueOf(x)
	typeOfX := reflect.TypeOf(x)

	y := &x
	valueOfY := reflect.ValueOf(y)
	typeOfY := reflect.TypeOf(y)

	fmt.Println("Value of x:", valueOfX)
	fmt.Println("Type of x:", typeOfX)

	fmt.Println("Value of y:", valueOfY)
	fmt.Println("Type of y:", typeOfY)
}