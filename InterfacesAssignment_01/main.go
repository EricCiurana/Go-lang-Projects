/*
	Author: 			Eric Ciurana
	Date(dd/mm/yyyy):	04/06/2018
	License:			This is free software dude, use it as you want!
*/

package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	tr := triangle{2.5, 3.7}
	sq := square{4.1}

	printArea(tr)
	printArea(sq)
}

func (tr triangle) getArea() float64 {
	return tr.base * tr.height * 0.5
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}
