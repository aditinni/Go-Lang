package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	card := deck{"Ace of Spades", "Two of Hearts"}
	card.print()
}
