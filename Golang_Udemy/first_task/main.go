package main

import (
	"fmt"
)

func main() {

	myNumbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, numbers := range myNumbers {

		if numbers%2 == 0 {
			fmt.Println(numbers, "is even ")
		} else {
			fmt.Println(numbers, "is odd")
		}

	}
}
