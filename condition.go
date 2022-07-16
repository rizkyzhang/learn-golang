package main

import "fmt"

func learnCondition() {
	

	if n := 5; n > 5 {
		fmt.Println("> 5")
	} else if n < 5 {
		fmt.Println("< 5")
	} else {
		fmt.Println("== 5")
	}

	x := 0

	switch x {
		case 1, 2, 3: 
			fmt.Println("1 or 2 or 3")
		default: 
			fmt.Println("0")
	}

	switch {
		case x < 0: {
			fmt.Println("Negative")
		}
		case x > 0: {
			fmt.Println("Positive")
		}
		default: {
			fmt.Println("Zero")
		}
	}
}