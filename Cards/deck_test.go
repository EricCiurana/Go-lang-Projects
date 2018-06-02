/*
	Author: 			Eric Ciurana
	Date(dd/mm/yyyy):	02/06/2018
	License:			This is free software dude, use it as you want!
*/
package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be : Ace of Spades\nGot %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected first card to be : King of Club\nGot %v", d[0])
	}
}

func TestShuffle(t *testing.T) {
	d := newDeck()

	first := d[0]
	d.shuffle()
	if d[0] == first {
		d.shuffle()
		if d[0] == first {
			t.Errorf("Expected first card to be different after two shuffles, got the same")
		}
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
