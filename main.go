package main

import (
	"fmt"
	"math"
)

// 1. Create a struct called 'triangle' that has two data fields: 'base' and 'height' both of type float64
type triangle struct {
	base   float64
	height float64
}

// 2. Create a method on type 'triangle' named 'area' that calculates the area of a triangle
func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

// 3. Create a method on type 'triangle' named 'perimeter' that calculates and returns the perimeter of the triangle
func (t triangle) perimeter() float64 {
	return 2*t.base + 2*t.height
}

// 6. Create a new type called 'circle' that has a single field of type float64 of name 'radius'
type circle struct {
	radius float64
}

// 7. Create a method on type 'circle' with the name 'area' that calculates and returns the area of a circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// 8. Create a method on type 'circle' with the name 'perimeter' that calculates and returns the perimeter of a circle
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// 9. Write a Go function named 'square' that accepts a parameter of type float64 that represents the sides of the square.
// The 'square' function should return the area and the perimeter of the square when called.
func square(side float64) (float64, float64) {
	return (side * side), (4 * side)
}

// 4. Inside the 'main' function create a variable of type 'triangle' and initialize it with the values for the two data fields
func main() {
	t := triangle{base: 3, height: 4}
	// 5. Call the 'area' and 'perimeter' methods and print the results
	fmt.Println(t.area())
	fmt.Println(t.perimeter())

}
