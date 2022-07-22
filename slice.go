package main

import "fmt"

func learnSlice() {
	nums1 := []int{1, 2, 3, 4, 5}
	fmt.Println("nums1:", nums1)
	fmt.Println("")

	
	// Reference test
	nums2 := nums1[:]
	nums2[0] = 100
	
	fmt.Println("Reference test")
	fmt.Println("nums2:", nums2)
	fmt.Println("nums1:", nums1)
	fmt.Println("")


	// Append with len(slice) == cap(slice)
	foods1 := []string{"Pizza", "Hamburger", "Noodle", "Sushi", "Nugget"}
	foods2 := foods1[:]
	foods3 := append(foods2, "Fried Rice")

	fmt.Println("Append with len(slice) == cap(slice)")
	fmt.Println("foods1:", foods1)
	fmt.Println("foods2:", foods2)
	fmt.Println("foods3:", foods3)
	fmt.Println("")

	// Append with len(slice) < cap(slice)
	foods4 := foods1[0:3]
	foods5 := append(foods4, "Tomato Soup")

	fmt.Println("Append with len(slice) < cap(slice)")
	fmt.Println("foods1:", foods1)
	fmt.Println("foods4:", foods4)
	fmt.Println("foods5:", foods5)
}