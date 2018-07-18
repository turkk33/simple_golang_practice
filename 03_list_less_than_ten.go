package main

import (
	"fmt"
)

func lowelements() {
	a := [11]int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	
	for i := 0; i < len(a); i++ {
		if a[i] < 5 {
			fmt.Println(a[i])
		}
	}
}

func main() {
	lowelements()
}
