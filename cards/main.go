package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)
}

// function declaration
func newCard() string {
	return "Five of Diamonds"
}
