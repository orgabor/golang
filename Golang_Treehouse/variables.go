/*
Must be declared
Must declare a type for variable
Can declare multiple variables of same type at once
Type can be inferred from initial value, if you provide one
Short variable declarations: e := 6*/
package main

import "fmt"

func main() {
  var a int
  a = 1
  var b, c int
  b = 2
  c = 3
  var d = 5
  e := 6
  fmt.Println()

}
/*
If we run this sample, we'll get a compile error for all of these variables saying that they have been "declared and not used". Just like you're required to use every package you import, you also have to use every variable you declare. An unused variable often indicates a bug, and Go helps your development team spot these bugs quicker by reporting unused variables.

Since this is just a quick example, the way we'll "use" the variables is by updating the last line to print them all out:

fmt.Println(a, b, c, d, e)
Now that we're actually accessing the variable values somewhere, the program will compile and run successfully.

You can assign multiple variables at the same time:
*/
package main

import "fmt"

func main() {
    var a, b int
    a, b = 1, 2
    var c, d = 3, 4
    e, f := 5, 6
    fmt.Println(a, b, c, d, e, f)
}
/*
The same rules apply for the names of variables, constants, functions, and custom types.
Must begin with a letter
Can be followed by any number of letters and digits
If the first letter is lower-case, the variable is unexported and can only be used within the current package. (The same is true for constants, functions, etc.)
If the first letter is upper-case, the variable is exported and can also be used OUTSIDE the current package.
Those are the only rules enforced by the compiler. But there are a couple conventions that Go programmers follow as well.
*/


func DeclareVariables() {
    // OK
    word := 1
    multipleWords := 2
    // Not OK
    multiplewords
    multiple_words
}

//You can only assign values of the declared type to a variable:

  var wholeNumber int = 1
  var fractionalNumber float64 = 2.75
  var wholeNumber2 int = fractionalNumber // ERROR
  var fractionalNumber2 float64 = wholeNumber // ERROR
  fmt.Println("wholeNumber2:", wholeNumber2)
  fmt.Println("fractionalNubmer2:", fractionalNumber2)
//Try to run the above, and you'll get compile errors. But you can assign the values if you do a conversion: write the type you want to convert to, followed by the value you want to convert in parentheses.

  var wholeNumber int = 1
  var fractionalNumber float64 = 2.75
  var wholeNumber2 int = int(fractionalNumber)
  var fractionalNumber2 float64 = float64(wholeNumber)
  fmt.Println("wholeNumber2:", wholeNumber2)
  fmt.Println("fractionalNubmer2:", fractionalNumber2)
//Must also convert both operands to the same type prior to math operations or comparisons:

  var wholeNumber int = 1
  var fractionalNumber float64 = 2.75
  fmt.Println(float64(wholeNumber) + fractionalNumber)
  fmt.Println(float64(wholeNumber) < fractionalNumber)
/*A variable's scope is limited to block in which it's defined
Blocks are a chunk of code surrounded by {} braces
They're usually associated with functions, if statements, and loops, all of which we'll see later. But it's possible to use blocks by themselves.
Even though it's not surrounded by curly braces, each package is an implicit block. Variables defined at package level are accessible from every block nested within that package. Most importantly, that includes the bodies of functions.
*/

package main

import (
    "fmt"
)

var a = "a" // Package-level variable

func main() {
    var b = "b"
    {
        var c = "c"
        {
            var d = "d"
            fmt.Println(a, b, c, d)
        }
        fmt.Println(a, b, c, d) // ERROR: "d" undefined
    }
    fmt.Println(a, b, c, d) // ERROR: "c", "d" undefined
}