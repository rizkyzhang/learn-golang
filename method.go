package main

import (
	"fmt"
	"strings"
)

type cat struct {
	name string
	age int
	favoriteFood string
}

func (s cat) meow() {
	fmt.Println("meowwwwww")
}

func (s cat) doYouLike(food string) {
	if (strings.EqualFold(s.favoriteFood, food)) {
		fmt.Printf("I like %s :)", food)
	} else {
		fmt.Printf("I don't like %s!", food)
	}
}

func learnMethod() {
	cat1 := cat{name: "Billy", age: 2, favoriteFood: "Fish"}

	cat1.meow()
	cat1.doYouLike("fish")
}