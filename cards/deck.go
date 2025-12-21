package main

import "fmt"

// Create a new type of deck- slice of string
type deck []string

// Function to create and return a new deck of cards

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
