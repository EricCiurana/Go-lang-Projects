/*
	Author: 			Eric Ciurana
	Date(dd/mm/yyyy):	01/06/2018
	License:			This is free software dude, use it as you want!
*/

package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	subject001 := person{"John", "Doe", contactInfo{"john.doe@internet.com", 31415}}
	fmt.Println(subject001)

	subject002 := person{
		firstName: "Jane",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email:   "jane.doe@internet.com",
			zipCode: 11235,
		},
	}
	fmt.Printf("%+v", subject002)

	var subject003 person
	fmt.Println(subject003)         // Shall print '{ { 0}}'
	fmt.Printf("%+v\n", subject003) // Shall print '{firstName: lastName: contact:{ email: zipCode:0}}'
	subject003.firstName = "Eric"
	subject003.lastName = "Ciurana"
	//fmt.Println(subject003) // Shall print '{Eric Ciurana { 0}}'
	subject003.print()

	subject003.updateName("Erica")
	subject003.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
