/*Pointers point to a value in memory. That is, pointers are a value that just gives the address of another value elsewhere in memory. If you're familiar with object references in C# or Java, those work almost exactly like pointers. And if you already know Pointers from C or C++, Go pointers work almost exactly the same way, except that you can't alter an existing pointer like you can in C or C++.

To declare a variable that holds a pointer to a type of value, put a star (*) before the type. So var x float64 declares a variable that holds a float64 directly, but var x *float64 declares a variable that holds a POINTER to a a float64
&: "address of" operator
*: gets the value at an address
*/
package main

import "fmt"

/*
func main() {
	var aValue float64 = 1.23
	var aPointer *float64 = &aValue
	fmt.Println("aPointer:", aPointer)   // prints something like "aPointer: 0x1040a128"
	fmt.Println("*aPointer:", *aPointer) // prints "*aPointer: 1.23"
}
*/

/*Original*/
/*
func main() {
	myNumber := 2.6
	halve(myNumber)                              //does nothing
	fmt.Println("myNumber in 'main':", myNumber) //points 2.6
	return
}

func halve(number float64) {
	number = number / 2

	fmt.Println("number in 'halve:'", number) //print 1.3 but change is only effective here
	return
}
*/

/*pass address of myNumber*/
/*
func main() {
	myNumber := 2.6
	halve(&myNumber)                             //address
	fmt.Println("myNumber in 'main':", myNumber) //points 2.6
	return
}

func halve(number *float64) { //point
	*number = *number / 2

	fmt.Println("*number in 'halve:'", *number) //print 1.3 but change is only effective here
	return
}
*/

func main() {
	car := Car{
		Doors:         4,
		Transmissions: "automatic",
		Odometer:      36000,
	}
	describe(&car)
}

func describe(c *Car) {

	fmt.Println("A", c.Doors, "door")
	fmt.Println(c.Transmissions, "car")
	fmt.Println("with", c.Odometer, "miles")

}

/*
There are a couple circumstances where it's better to use a pointer:

Function arguments are pass-by-value: the function always receives a COPY of the value
So if functions alter the value they receive, they'll be altering the copy, not the original
But if the function receives a pointer, and alters the value that pointer points to, then the function's changes will still be effective outside the function
For complex values, passing a copy to a function wastes memory.
Passing a pointer doesn't.
*/
