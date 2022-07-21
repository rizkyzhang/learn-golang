package main

import "fmt"

func learnMap() {
	score := map[string]int{
		"Andi": 70,
		"Billy": 76,
		"Carlos": 68,
	}

	fmt.Println("score:", score)

	value, isExist := score["Ruby"]

	if (isExist) {
		fmt.Println(value)
	} else {
		fmt.Println("Key Ruby does not exist in map")
	}

	users := []map[string]string{
		{"name": "Andi", "age": "25"},
		{"name": "Ricky", "age": "23"},
		{"name": "Ferdy", "age": "33"},
	}

	for _, user := range  users {
		fmt.Println("Name:", user["name"], "Age:", user["age"])
	}
}