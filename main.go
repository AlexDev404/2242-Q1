package main
import "math"
import "fmt"

// 1. Create a struct called 'triangle' that has two data fields: 'base' and 'height' both of type float64
type triangle struct {
	base float64
	height float64
}

// 2. Create a method on type 'triangle' named 'area' that calculates the area of a triangle
func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func main() {

}