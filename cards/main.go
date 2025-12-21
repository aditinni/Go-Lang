package main

import "fmt"

func main() {

	cards := []string{"Ace of Diamonds", "Two of Hearts"}
	cards = append(cards, "Six of Spades") // it does not modify the original slice, it returns a new slice
	fmt.Println(cards)
	fmt.Println("Welcome to the Card Game!")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

//Arrays vs Slices
// An array has a fixed size, while a slice is a dynamically-sized, flexible view into the elements of an array.
