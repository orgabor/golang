/*
Say you have several types, all of which have the same methods.
Say you want to create a method that calls the same set of methods on each of those types.
Go's static typing requires that you create a separate method for each type you want to handle,
even if the methods' logic is the same. Or does it?
"Interfaces" can help with this sort of problem.
*/
/*
Sprintf
The fmt.Sprintf function takes a formatting string, inserts one or more formatted values into it, and returns the resulting string.

Interfaces
Here are Monitor, HardDrive, and Keyboard types, each of which has its own Specs and Price methods.

package parts

import "fmt"

type Monitor struct {
    Resolution string
    Connector  string
    Value      float64
}

func (m Monitor) Specs() string {
    return fmt.Sprintf("Monitor\nResolution: %s\nConnector: %s", m.Resolution, m.Connector)
}

func (m Monitor) Price() string {
    return fmt.Sprintf("$%0.2f", m.Value)
}

type HardDrive struct {
    Type      string
    Connector string
    Value     float64
}

func (h HardDrive) Specs() string {
    return fmt.Sprintf("Hard Drive\nType: %s\nConnector: %s", h.Type, h.Connector)
}

func (h HardDrive) Price() string {
    return fmt.Sprintf("$%0.2f", h.Value)
}

type Keyboard struct {
    Keys     int
    Switches string
    Value    float64
}

func (k Keyboard) Specs() string {
    return fmt.Sprintf("Keyboard\nKeys: %d\nSwitch Type: %s", k.Keys, k.Switches)
}

func (k Keyboard) Price() string {
    return fmt.Sprintf("$%0.2f", k.Value)
}
This code creates a new interface using the Specs and Price methods:

type Part interface {
  Specs() string
  Price() string
}
Now that we have a Part interface, we can write a single method that takes a Part parameter. The method will be able to take a Monitor, a HardDrive, or a Keyboard as an argument, because all of those types fulfill the Part interface.

func Summary(part Part) string {
  return part.Specs() + "\n" + part.Price()
}
We can also declare a slice that holds Part values, which will then be able to contain Monitor, HardDrive, AND Keyboard values:

    catalog := []Part{}
    catalog = append(catalog, parts.Monitor{"HDMI", "1080p", 249.99})
    catalog = append(catalog, parts.HardDrive{"SSD", "SATA", 149.99})
    catalog = append(catalog, parts.Keyboard{108, "Mechanical", 99.99})
    for _, part := range catalog {
        fmt.Println(Summary(part))
        fmt.Println("------------------------")
    }
Have questions about this video? Start a discussion with the community and Treehouse staff.

*/

package main

import (
	"fmt"
	"parts"
)

type Part interface {
	Specs() string
	Price() string
}

func Summary(part Part) string {

	return part.Specs() + "\n" + part.Price()
}

func main() {
	catalog := []Part{}
	catalog = append(catalog, parts.Monitor{"HDMI", "1080p", 249.99})
	catalog = append(catalog, parts.HardDrive{"SSD", "SATA", 149.99})
	catalog = append(catalog, parts.Keyboard{108, "Mechanical", 99.99})
	for _, part := range catalog {
		fmt.Println(Summary(part))
		fmt.Println("------------------------")
	}

}

/* EXample code
type Dog
func (d Dog) Walk()
func (d Dog) Sit()
func (d Dog) Fetch()

type Cat
func (d Dog) Walk()
func (d Dog) Sit()
func (d Dog) Purr()

type FourLegged interface {
    Walk()
    Sit()
}

var animal FourLegged
animal= Cat{}
animal = Dog{}
animal.Walk()
animal.Sit()
*/
