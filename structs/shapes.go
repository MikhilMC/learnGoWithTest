package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

type Cirle struct {
	radius float64
}

func (circle Cirle) Area() float64 {
	return math.Pi * circle.radius * circle.radius
}
