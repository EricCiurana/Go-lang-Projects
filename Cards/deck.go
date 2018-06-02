/*
	Author: 			Eric Ciurana
	Date(dd/mm/yyyy):	02/06/2018
	License:			This is free software dude, use it as you want!
*/

package main

import "fmt"

// Creation a new type 'deck' as a slice of strings
type deck []string

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "king"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}
