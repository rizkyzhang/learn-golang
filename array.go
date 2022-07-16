package main

import "fmt"

func learnArray() {
	nums := [5]int{1, 2, 3, 4, 5} // Fixed capacity
	evenNums := [...]int{2, 4, 6, 8, 10, 12} // Dynamic capacity

	fmt.Println(nums)
	fmt.Println(evenNums)

	oddNums := make([]int, 2)
	oddNums[0] = 1
	oddNums[1] = 3

	fmt.Println("old oddNums:", oddNums)

	for _, num := range nums {
		fmt.Println(num)
	}

	oddNumsCopy := oddNums
	oddNumsCopy[0] = 1
	oddNumsCopy[1] = 5

	fmt.Println("new oddNums:", oddNums)
	fmt.Println("copy oddNums:", oddNumsCopy)
}