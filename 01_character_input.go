package main

import (
	"fmt"
)

func charinput() {
	var name string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	s := fmt.Sprintf("Thank you, %s", name)
	fmt.Println(s)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	s = fmt.Sprintf("Ok got it. %s, age %d", name, age)
	fmt.Println(s)
}

func main() {
	charinput()
}
