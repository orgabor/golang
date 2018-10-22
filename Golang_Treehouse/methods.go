/*
RewriteEngine on

RewriteBase /

RewriteRule (.*) http://www.szentdonat.hu/$1 [R=301,L]


*/
/*
A method is a function with a special receiver argument: https://golang.org/ref/spec#Method_declarations
-Allows you to define new behaviors for values of a type
-Method declaration looks just like an ordinary function declaration, except it has extra receiver parameter before the function name
-Why "receiver"? Think of calling a method on an object as sending it a message.
-No this or self, just an ordinary parameter name for the receiver
-This code defines a type named Title with a method named FixCase:
*/

package main

import (
	"fmt"
	"strings"
)

type Title string

func (t Title) FixCase() string {
	return strings.Title(string(t))
}

func main() {
	name := Title("the return of the jedi")
	fixed := name.FixCase() //name. -calling a method on that value
	fmt.Println(fixed)
	methodsMinutes()
}

/*

The dot notation for calling methods reflects the way that the method is declared.
Reciever occurs before method name, just like in the declaration
Compiler looks at type of receiver, then calls method with that name defined on that type
Don't confuse with package qualifier for exported function names; not the same thing
Here we have some code that uses our Minutes type again.
We've defined an Increment method that adds 1 to the number of minutes,
using the percent sign modulo operator to roll the number over to 0 if it's over 59.
his code sample is supposed to print "59", "00", "01" as the number of minutes rolls over to the next hour.
But it just prints "58" repeatedly.


*/

type Minutes int

func (m *Minutes) Increment() {
	*m = (*m + 1) % 60
}

func methodsMinutes() {
	minutes := Minutes(58)
	for i := 1; i <= 3; i++ {
		minutes.Increment()
		fmt.Println(minutes)
	}
}


/********** Coding Challange *****************/

/**** Task ***/

/*
We've set up an Hours type, which tracks the number of hours displayed on a clock, from 0 - 23. We've also created an Increment function that takes an Hours value, adds 1 to it (wrapping it around to 0 if the result is greater than 23), and returns the result as a new Hours value. It's used like this:

hours := Hours(23)
hours = Increment(hours)
fmt.Println(hours) // Prints "0".
Convert Increment to a method on the Hours type. Have it modify the receiver directly, so it can be used like this:

hours := Hours(23)
hours.Increment()
fmt.Println(hours) // Prints "0"
Don't forget: to modify the receiver parameter, you'll need to make it a pointer, and then modify the value at that pointer.


*/

/******* Solution *********/

package clock

type Hours int

func (h *Hours) Increment() {
  *h = (*h + 1) % 24
}

func methodsHours() {
	hours := Hours(23)
	for i := 1; i <= 3; i++ {
		hours.Increment()
		return
	}
}