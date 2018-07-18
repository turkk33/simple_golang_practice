package main

import (
	"fmt"
)

func oddoreven() {
	fmt.Print("Pick any number, I'll tell you what it is: ")

	var num int
	fmt.Scanln(&num)

	if (num % 2 == 1) {
		fmt.Println("Your number is odd.")
	} else {
		fmt.Println("Your number is even.")
	}
}

func main() {
	oddoreven()
}
