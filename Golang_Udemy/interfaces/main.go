package main

import (
	"fmt"
)

/*
*
*
 */
type bot interface {

	//live example
	//	getGreeting(string, int) (string, error) //argument () return ()
	//getGreeting() string
	getGreeting() string
}

type englishBot struct{}
type hungarianBot struct{}

func main() {

	eb := englishBot{}
	hu := hungarianBot{}

	printGreeting(eb)
	printGreeting(hu)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())

}

func (englishBot) getGreeting() string {

	//VERY Custom Logic for generating an english greeting

	return "Hi There!"

}

func (hungarianBot) getGreeting() string {

	return "Szia!"

}
