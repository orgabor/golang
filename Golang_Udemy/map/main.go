package main

import (
	"fmt"
)

func main() {
	/* ++ First way to declare a map ++ */
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"blue":  "#00f",
		"white": "#fff",
	}

	//second way
	//var colors map[string]string

	//3rd way

	/*
		colors := make(map[int]string)

			colors[3] = "#fff"
			delete(colors, 3)
	*/
	printMap(colors)
}

func printMap(c map[string]string) {

	for color, hex := range c {
		fmt.Println("Hex code for ", color, "is ", hex)

	}
}
