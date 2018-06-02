/*
	Author: 			Eric Ciurana
	Date(dd/mm/yyyy):	02/06/2018
	License:			This is free software dude, use it as you want!
*/

package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
foo returns a separated by commas string version of the deck
*/
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

/*
foo gets a filename and saves the deck to NVM, returns an error (if any)
*/
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

/*
foo gets a filename, reads its content and saves it to a deck
	returns the deck or an error (if any)
*/
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

/*
foo shuffles the deck
*/
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		newPos := r.Intn(len(d) - 1)

		d[index], d[newPos] = d[newPos], d[index]
	}
}
