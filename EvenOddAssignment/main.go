/*
	Author: 			Eric Ciurana
	Date(dd/mm/yyyy):	02/06/2018
	License:			This is free software dude, use it as you want!
*/

package main

import (
	"fmt"
)

func main() {
	var intList []int
	for i := 0; i < 11; i++ {
		intList = append(intList, i)
	}

	for i := range intList {
		if i%2 == 0 {
			output := fmt.Sprintf("%v is even", i)
			fmt.Println(output)
		} else {
			output := fmt.Sprintf("%v is odd", i)
			fmt.Println(output)
		}
	}

	/* Note */
	// I tried to use fmt.Println("%v is odd", i), but it wouldn't work
}
