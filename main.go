package main

import "fmt"

//interface for a generic type called shape that prints the area
type shape interface {
	getArea() float64
}

//triangle struct with 2 attributs/fields
type triangle struct {
	height float64
	base   float64
}

//square struct with 1 attribute/field
type square struct {
	sideLength float64
}

//getArea() specfiic for triangle
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

//getArea() specific for sqaure
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

//interface function used to print out area of a give shape
func printArea(s shape) {
	fmt.Println(s.getArea())
}

//main function that instantiates a triangle and a square and prints out the area using the interface
func main() {
	tri := triangle{height: 10, base: 10}
	sq := square{sideLength: 10}
	printArea(tri)
	printArea(sq)
}
