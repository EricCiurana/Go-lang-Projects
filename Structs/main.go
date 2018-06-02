/*
	Author: 			Eric Ciurana
	Date(dd/mm/yyyy):	01/06/2018
	License:			This is free software dude, use it as you want!
*/

package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	subject001 := person{"John", "Doe"}
	fmt.Println(subject001)

	subject002 := person{firstName: "Jane", lastName: "Doe"}
	fmt.Println(subject002)

	var subject003 person
	fmt.Println(subject003)         // Shall print '{ }'
	fmt.Printf("%+v\n", subject003) // Shall print '{firstName: lastName: }'
	subject003.firstName = "Eric"
	subject003.lastName = "Ciurana"
	fmt.Println(subject003) // Shall print '{Eric Ciurana}'
}
