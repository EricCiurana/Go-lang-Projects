/*
	Author: 			Eric Ciurana
	Date(dd/mm/yyyy):	02/06/2018
	License:			This is free software dude, use it as you want!
*/

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new deck type which is a slice of strings
type deck []string

/*
foo returns an ordered deck
*/
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

/*
foo prints out the deck
*/
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/*
foo gets a deck and a hand size, splits it and returns two strings
*/
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/*
foo returns a string version of the deck
*/
func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

/*
foo gets a filename and saves the deck to NVM, returns an error (if any)
*/
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
