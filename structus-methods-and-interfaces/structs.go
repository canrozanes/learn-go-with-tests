package main

import "math"

// Shape is an interface that includes all Structs with Area() methods
type Shape interface {
	Area() float64
}

// Rectangle defines a geometric rectangle with Width and Height
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle defines a circle with Radius
type Circle struct {
	Radius float64
}

// Triangle defines a triangle with Base and Height
type Triangle struct {
	Base   float64
	Height float64
}

// Perimeter calculates the perimeter of a rectangle
func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area calculates the area of a rectangle
func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2) * math.Pi
}

// Area calculates the area of a triangle
func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

func main() {

}
