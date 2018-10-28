package main

/*Create a new type of deck
which is a slice of string
*/

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Heart", "Bell", "Leaf", "Acorn"}
	cardValues := []string{"Seven", "Eigth", "Nine", "Ten", "Over", "Upper", "King", "Ace"}

	for _, suite := range cardSuits {

		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}

	}

	return cards

}

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]

}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error {

	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

func newDeckFromFile(filename string) deck {

	bs, err := ioutil.ReadFile(filename)
	if err != nil {

		//Option #1 - log the error and return a call of newDeck(0)
		//Option #2 - log the error and quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)

	}
	//very short varaible names are very common in golang s is for slice or for string
	s := strings.Split(string(bs), ",")

	return deck(s)

}

func (d deck) shuffle() {

	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]

	}
}
