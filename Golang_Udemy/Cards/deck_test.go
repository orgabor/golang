package main

import (
	//"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 32 {
		t.Errorf("Expected length of  but got %v", len(d))
	}

	if d[0] != "Seven of Heart" {

		t.Errorf("Expected first card of  Seven of Heart, but got %v", d[0])
	}

	if d[len(d)-1] != "Ace of Acorn" {

		t.Errorf("Expected last card of  Ace of Acorn, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 32 {
		t.Errorf("Expected 32 cards in the deck  but got %v", len(loadedDeck))

	}

	os.Remove("_decktesting")

}
