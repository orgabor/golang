/*
Arrays, slices, and maps can hold groups of data that's all the same type:
all ints, all strings, etc. But sometimes you need to group data of different types together.
A real estate listing might need a string for the address together with a float for the price.
A catalog record might need a part name together with a part number. And so on.
Like the C programming language before it,
Go solves this sort of situation by offering "structs".




*/

/*
A struct is a group of named elements, called fields, each of which has a name and a type.
Fields accessed through dot notation

*/
/******** Original code **********/
package main

import (
"fmt"
"https://github.com/treehouse-projects/go-intro/clock"
)

type Monitor struct {
	Resolution string
	Connector  string
	Value      float64
}

/*
func main() {
	monitor := Monitor{}
	monitor.Resolution = "1080p"
	monitor.Connector = "HDMI"
	monitor.Value = 249.99
	fmt.Println(monitor.Resolution, monitor.Connector, monitor.Value)
}
*/
/*

Struct literals are similar to the literals for arrays, slices, and maps.
monitor := Monitor{"1080p", "HDMI", 249.99}
Must be in order, must include all fields.
With field names: monitor := Monitor{Value: 249.99, Resolution: "1080p", Connector: "HDMI"}
Can be in any order
Can leave a field off, and it will just get initialized to its zero value.


*/

/*modified*/
func showFields(m Monitor) {

	fmt.Println("Resolution:", m.Resolution, "Connector: ", m.Connector, "Value: ", m.Value)
}

func main() {
	//	monitor := Monitor{"1080p", "HDMI", 249.99}
	//monitor := Monitor{Value: 249.99, Resolution: "1080p"}
	monitor := Monitor{}
	showFields(monitor)
	displayClock()
}




func displayClock {
	minutes := clock.Minutes(58)
	minutes += 5 //Minutes now an invalid values!
	minutes.Display()
}