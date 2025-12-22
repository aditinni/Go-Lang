package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	card := newDeckFromFile("MyCards.txt")
	//card.print()
	//hand, remainingCards := deal(card, 5)

	//hand.print()
	//remainingCards.print()
	//fmt.Println(card.toString())
	//card.saveToFile("MyCards.txt")
	fmt.Println(card)
}
