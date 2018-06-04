/*
	Author: 			Eric Ciurana
	Date(dd/mm/yyyy):	04/06/2018
	License:			This is free software dude, use it as you want!
*/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fileName := os.Args[1]
	/*
		f, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		fmt.Println(string(f))
	*/
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}
	io.Copy(os.Stdout, f)
}
