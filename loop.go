package main

import "fmt"

func learnLooping() {
	x := 1

	for i := 0; i <= 10; i++ {
		if (i % 2 == 0) {
			fmt.Println(i)
		}
	}

	for x < 6 {
		fmt.Println(x)
		x++
	}
}