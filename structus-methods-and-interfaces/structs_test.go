package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		Width:  10.0,
		Height: 10.0,
	}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	type TestCase struct {
		name    string
		Shape   Shape
		HasArea float64
	}

	areaTests := []TestCase{
		{name: "Rectangle", Shape: Rectangle{Width: 12, Height: 6}, HasArea: 72.0},
		{name: "Circle", Shape: Circle{Radius: 10}, HasArea: 314.1592653589793},
		{name: "Triangle", Shape: Triangle{Base: 12, Height: 6}, HasArea: 36.0},
	}

	for _, testCase := range areaTests {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.Shape.Area()
			if got != testCase.HasArea {
				t.Errorf("%#v got %g want %g", testCase.Shape, got, testCase.HasArea)
			}
		})
	}
}
