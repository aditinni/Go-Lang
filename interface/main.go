package main

import "fmt"

type person1 string
type person2 string

type intro interface {
	introduction() string
}

func main() {

	p1 := person1("Aditya")
	p2 := person2("Anshu")

	introMessage(p1)
	introMessage(p2)

}

func (per1 person1) introduction() string {

	return "Hello, I am " + string(per1) + " I am learining Go Lang."

}

func (per2 person2) introduction() string {

	return "Hello, I am " + string(per2) + " I am preparing for UPSC."
}

func introMessage(i intro) {

	fmt.Println(i.introduction())
}
