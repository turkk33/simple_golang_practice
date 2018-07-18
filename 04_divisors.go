package main

import (
	"fmt"
)

func divisors() {
	fmt.Print("Give me a number, and I'll print the divisors: ")

	var num int
	fmt.Scanln(&num)

	divisors := make([]int, num)

	for i := 0; i < num; i++ {
		divisors[i] = i + 1
	}

	for i := 0; i < num; i++ {
		if (num % divisors[i]) == 0 {
			fmt.Println(divisors[i])
		}
	}
}

func main() {
	divisors()
}