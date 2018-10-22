/*
Slices are good for storing collections of data,
but the only way you can get data back out of a slice is by using a numeric index.
Sometimes it's useful to be able to use other values as keys that let you look up items in a collection.
Go uses the term "map" to refer to this sort of key-value list.

Sometimes it's useful to be able to use other values as keys
that let you look up items in a collection. Go uses the term
map to refer to this sort of key-value list.

The idea is that a map "maps" a set of keys to corresponding values.
You can use any type you want as a key, as long as the keys are all the same type.
If you know the key a value is stored under, you can use the key to retrieve that value.
*/

package main

import "fmt"

func main() {
	ages := map[string]float64{}
	// Like the array/slice syntax, but you can use any value of the type you specified for the keys.
	ages["Alice"] = 12
	ages["Bob"] = 9
	fmt.Println(ages)
	mapKeysAndVals()
	mapWithValues()
	mapForLoop()
}

/*
The syntax to retrieve an individual value is just like arrays or slices,
too, except you use a key in place of a numeric index.
*/
func mapKeysAndVals() {
	ages := map[string]float64{}

	ages["Alice"] = 12
	ages["Bob"] = 9
	fmt.Println(ages["Alice"], ages["Bob"])
}

//Initialize map with values:

func mapWithValues() {
	ages := map[string]float64{"Alice": 12, "Bob": 9}
	fmt.Println(ages)
}

//for loop
func mapForLoop() {
	ages := map[string]float64{"Alice": 12, "Bob": 9}
	for name, age := range ages {
		fmt.Println(name, age/2)
	}
	//blank identifier
	for _, alter := range ages {
		fmt.Println(alter)
	}

	//just the names
	for name := range ages {
		fmt.Println(name)
	}
}
