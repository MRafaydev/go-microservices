package geometry

import "math"

// func for the calculation of the Area
func Area(length, Width float64) float64 {
	return length * Width
}

// func for the calculation of the rectangle
func Rectangle(length, Width float64) (area, parameters float64) {
	area = (length * Width)
	parameters = 2 * area
	return
}

// func for the calculation of the diagonal
func Diagonal(length, width float64) float64 {
	return math.Sqrt((length * length) * (width * width))
}
