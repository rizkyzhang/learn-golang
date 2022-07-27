package library

import "fmt"

type PublicStruct struct {
	Name string
	Age int
}

func PublicFunc() {
	fmt.Println("I am a public function")
	test()
}

func privateFunc() {
	fmt.Println("I am a private function")
}
