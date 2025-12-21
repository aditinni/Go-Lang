package main

import "fmt"

//import "fmt"

func main() {
	//fmt.Println("Card Project")

	var card string = "Ace of Spades"
	// another way to declare variable
	// card := "Ace of Spades" -> Go will automatically infer the data type
	// for overriding variable use card = "Deck of Cards"

	/*
		var- variable declaration
		string- data type
		card- variable name
		Go is a statically typed language, so we need to declare the data type of variable
	*/
	fmt.Println(card)
}
