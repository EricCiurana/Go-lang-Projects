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
	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	fmt.Println(colours)

	/* Creation of empty maps */
	var names map[string]string
	fmt.Println(names)

	projects := make(map[string]string)
	fmt.Println(projects)

	/* Filling/reasigning maps */
	projects["WebApp"] = "TODO"
	projects["Crash Sensor"] = "FINISHED"
	fmt.Println(projects)
	projects["WebApp"] = "FINISHED"
	fmt.Println(projects)

	/* Deleting map entries */
	delete(projects, "WebApp")
	fmt.Println(projects)

	fmt.Println()
	/* Iterating over maps */
	for colour, hex := range colours {
		fmt.Println(colour + " = " + hex)
	}
}
