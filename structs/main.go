package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// if field is not assigned, it gets the zero value for that type
/*
for string -> ""
for int -> 0
for bool -> false
for float -> 0.0
*/
func main() {

	/*	p1 := person{
			firstName: "John",
			lastName:  "Doe",
		}
		var alex person
		alex.firstName = "Alex"
		alex.lastName = "Anderson"
		fmt.Println(alex)
		//fmt.Printf("%+v", alex) // prints field names as well

		fmt.Println(p1.firstName, p1.lastName)*/

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim123@gmail.com",
			zipCode: 12345,
		},
	}

	//fmt.Printf("%+v", jim)

	//jim.updateName("Aditya") this won't work
	jimPointer := &jim
	// & gives the memory address of jim. *pointer gives the direct value
	jimPointer.updateName("Aditya")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

//func (p person) updateName(newFirstName string) {
//p.firstName = newFirstName
/*
	This won't work as p is a copy of the original struct. To modify the original struct, we need to use a pointer receiver.
*/
//}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName

	/*
	 the *person in the receiver or in front of a type is a ponter to a person struct type
	 (*p) dereferences the pointer to get the actual value

	*/
}
