package main

import "testing"

// 10. Write a test for the 'square' function

func TestSquare(t *testing.T) {
	area, perimeter := square(4)
	expected_value := 16
	if area != float64(expected_value) {
		t.Errorf("Area of square is incorrect. Got %v expected %v", area, expected_value)
	}
	if perimeter != float64(expected_value) {
		t.Errorf("Perimeter of square is incorrect. Got %v expected %v", perimeter, expected_value)
	}
}
