/*
	Author: 			Eric Ciurana
	Date(dd/mm/yyyy):	03/06/2018
	License:			This is free software dude, use it as you want!
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	resp, err := http.Get("http://www.google.es")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body)

	if err == nil {
		fmt.Println("\nSuccess!")
	}
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Bytes written: ", len(bs))

	return len(bs), nil
}
