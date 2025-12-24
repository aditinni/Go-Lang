package main

import "fmt"

type evenodd []int

func (e evenodd) checkEvenOdd() {

	for _, num := range e {
		if num%2 == 0 {
			fmt.Println(fmt.Sprint(num) + " is even")
		} else {
			fmt.Println(fmt.Sprint(num) + " is odd")
		}
	}
}
