/*
	Author: 			Eric Ciurana
	Date(dd/mm/yyyy):	03/06/2018
	License:			This is free software dude, use it as you want!
*/

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.es")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}

	fmt.Println(resp.Body.
}
