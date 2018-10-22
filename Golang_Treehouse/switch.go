/*




- To choose between several options with if / else statements, you'd have to nest them several deep, which is hard to read
- So like most programming languages, Go offers a switch statement.
- You specify the expression you want to switch on, case statements with possible values for that expression, and one or more lines of code that should run in each case.
- You can add a default case at the end, that will run if none of the other cases match.



*/

package main

import (
	"fmt"
)

func main() {
	fmt.Print("You win: ")
	doorNumber := 1
	switch doorNumber {
	case 1:
		fmt.Println("a new car!") // not printed
		fallthrough               //falltrough
	case 2:
		fmt.Println("a llama!") // printed
	default:
		fmt.Println("a goat!") // not printed

	}
}

func ShotsDescription(par int, shots int) string {
	shotsOverPar := shots - par
	// YOUR CODE HERE
	// If the shotsOverPar variable is 1, return "bogey".
	// If it's 0, return "par".
	// For -1 return "birdie".
	// For -2 return "eagle".
	// For all other cases, return "no description".

	switch shotsOverPar {
	case 1:
		return "bogey"
	case 0:
		return "par"
	case -1:
		return "birdie"
	case -2:
		return "eagle"
	default:
		return "no description"

	}

}
