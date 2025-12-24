package main

import "fmt"

func main() {
	fmt.Println("Even or Odd?")

	numbers := evenodd{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numbers.checkEvenOdd()
}
