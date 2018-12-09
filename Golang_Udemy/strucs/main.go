package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	/*++first option++*/
	/*
		contact   contactInfo //fieldname and type
	*/
	/*++second option++*/
	contactInfo
}

func main() {
	/*Examples for structs*/
	//	first examplre - not recommenden ->  gabor := person{"Gabor", "Gazdag"}
	/* struc examples from the first lesson
	gabor := person{firstName: "Gabor", lastName: "Gazdag"}

	var juci person

	juci.firstName = "Judit"
	juci.lastName = "Gazdag-Kovacs"

	fmt.Println(gabor)
	fmt.Println(juci)
	fmt.Printf("%v", juci)
	*/

	marci := person{

		firstName: "Marcello",
		lastName:  "di Monte-Sitacchio",
		contactInfo: contactInfo{
			email:   "cicaahegyen@mail.com",
			zipCode: 8228,
		}, //every single line needs a comma
	}

	/*Original code - not really used in practice
	marciPointer := &marci
	marciPointer.updateName("cica")
	*/
	/*Live likely code example*/
	marci.updateName("cica")
	marci.print()

}

func (p person) print() {

	fmt.Printf("%v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
