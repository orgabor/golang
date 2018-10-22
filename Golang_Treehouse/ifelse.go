package main

import (
	"fmt"
)

func main() {
	/*first part*/

	if true {
		fmt.Println("true") // printed
	}
	if false {
		fmt.Println("false") // not printed
	}
	if 1 < 2 {
		fmt.Println("1 < 2") // printed
	}
	if 1 > 2 {
		fmt.Println("1 > 2") // not printed
	}

	/*Second Part - Operators*/
	/*

	   Boolean operators:
	   !: "not"
	   &&: "and"
	   ||: "or"

	*/

	if !true {
		fmt.Println("!true") // not printed
	}
	if !false {
		fmt.Println("!false") // printed
	}
	if true && false {
		fmt.Println("true && false") // not printed
	}
	if true && true {
		fmt.Println("true && true") // printed
	}
	if true || false {
		fmt.Println("true || false") // printed
	}
	if true || true {
		fmt.Println("true || true") // printed
	}

	/*Third part - if else*/
	if false {
		fmt.Println("if") // not printed
	} else {
		fmt.Println("else") // printed
	}

	/*Fourth part - elseif*/
	if false {
		fmt.Println("if") // not printed
	} else if true {
		fmt.Println("else if") // printed
	} else {
		fmt.Println("else") // not printed
	}

	/*Fifth part Block scope rules apply to if block*/
	beforeIf := 888
	if true {
		withinIf := 999
		fmt.Println("if")
		fmt.Println(withinIf)
	}
	fmt.Println(beforeIf) // OK
	//fmt.Println(withinIf) // Error!

}
func PlayHand(score int) string {
  // YOUR CODE HERE
  if score <17 {
        fmt.Println("hit me") 
    } else if score > {
        fmt.Println("bust") 
    } else {
        fmt.Println("stand") 
    }
}

