
/*

Declaring a function
-func keyword
-Function name
-Parentheses
-Code block in curly braces that gets run when function is called

Naming rules

-Same naming rules as variables, types, etc.
-Must start with letter
-Capital letter means exported from the current package, lower case means unexported
-Convention is to use camelCase for multiple words

Return values

-Must specify return value type.
-Can name return values - sometimes useful as documentation.
-Can use bare returns if return values are named.

Go doesn't allow:
-Default parameter values
-Named Parameters
-Mehod overloadin (multiple names with the same but different parameters)

Goal is:

-make calling methods simple clear and fast

*/

package main

import "fmt"

func main() {
	myFunction()
	 square(3)
	// square("Hello") //Error
    add(2, 4)
    subtract(10, 3)
   	fmt.Println(squareV2(5))
    fmt.Println(addV2(8, 22))
    fmt.Println(subtractV2(56, 44))
}
	



func myFunction(){
 fmt.Println("Running myFunction")
}


func square(number int) { //parameter +type of the parameter
    fmt.Println(number * number)
}

func add(a float64, b float64) { //multiple parameter with type for each
    fmt.Println(a + b)
}

func subtract(a, b float64) { //multiple variable with the same type
    fmt.Println(a - b)
}

func squareV2(number int) int {
    return number * number
}

func addV2(a float64, b float64) (sum float64) {
    return a + b
}

func subtractV2(a, b float64) (difference float64) {
    difference = a - b
    return
}