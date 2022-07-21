package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func learnFunc() {
	rand.Seed(time.Now().Unix()) // make sure the random number is truly random

	randomInt := getRandomInt(1, 100)
	fmt.Println("randomInt:", randomInt)

	area, circumference := calcCircle(15.76)
	fmt.Println("area: ", area, "circumference: ", circumference)

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(avg(nums...))

	// Anonymous function
	anonFunc := func() {
		fmt.Println("Anon func")
	}

	anonFunc()

	// IIFE
	func(name string) {
		fmt.Printf("Hello %s \n", name)
	}("James")
}

func getRandomInt(min, max int) int {
	return rand.Int() % (max - min + 1) + min
}

func calcCircle(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d / 2, 2)
	circumference = math.Pi * d

	return
}

// Variadic parameter
func avg(nums ...int) int {
	sum := 0
	totalNums := len(nums)

	for _, num := range nums {
		sum += num
	}

	result := sum / totalNums

	return result
}

