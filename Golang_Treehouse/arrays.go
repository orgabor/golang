/****Arrays****/

/*
   So far we've only had to store single values.
   But it's much more common in computing to need to store an entire list of values,
   like items on a grocery list or phone numbers in a contact list.
   Go has a couple different structures that can do this: arrays, and slices.
   Slices are used more commonly than arrays, for reasons we'll discuss shortly,
   but since slices are built on top of arrays, it's important to understand arrays as well.
   So we'll look at arrays first.

	Array: A list of elements
	Array type is written [n]T, where n is the array size and T is the type of the array's elements
	All elements must be of same type
	Assign/access individual elements with a[0], a[1], etc.

*/

package main

import (
	"fmt"
)

/*Array version one*/

func main() {
	var months [3]string
	months[0] = "Apr"
	months[1] = "May"
	months[2] = "Jun"
	var salesByMonth [3]float64
	salesByMonth[0] = 1710.26
	salesByMonth[1] = 2245.97
	salesByMonth[2] = 3032.40
	fmt.Println(months[0], ", ", salesByMonth[0])
	fmt.Println(months[1], ", ", salesByMonth[1])
	fmt.Println(months[2], ", ", salesByMonth[2])
	arraysVer2()
	arraysLen()
	arraysRange()
	arraysBalnkIndentifier()

}

/*Array literals using curly braces*/

func arraysVer2() {

	months := [3]string{"Apr", "May", "Jun"}
	salesByMonth := [3]float64{1710.26, 2245.97, 3032.40}
	fmt.Println(months[0], salesByMonth[0])
	fmt.Println(months[1], salesByMonth[1])
	fmt.Println(months[2], salesByMonth[2])
}

/*Get length with len(a)*/

func arraysLen() {

	months := [3]string{"Apr", "May", "Jun"}
	salesByMonth := [3]float64{1710.26, 2245.97, 3032.40}
	for i := 0; i < len(months); i++ {
		fmt.Println(months[i], salesByMonth[i])
	}
}

/*Arrays with range*/
/*
The init/condition/post form of for loop is pretty brittle. Suppose we had accidentally changed the condition to i <= len(months)?
for ... = range is better.
This is essentially equivalent to "for each" in other languages.

*/

func arraysRange() {

	months := [3]string{"Apr", "May", "Jun"}
	salesByMonth := [3]float64{1710.26, 2245.97, 3032.40}
	for i, month := range months {
		fmt.Println(month, salesByMonth[i])
	}
}

/*Arrays with blank identifier*/
/*
If you're not going to use the index or the value, you can provide the blank identifier instead of a variable:

*/
func arraysBalnkIndentifier() {
	months := [3]string{"Apr", "May", "Jun"}
	//salesByMonth := [3]float64{1710.26, 2245.97, 3032.40}
	for _, month := range months {
		fmt.Println(month)
	}

}
