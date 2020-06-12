package main

import "math"

//Shape is interface
type Shape interface {
	Area() float64
}

//Triangle is struct
type Triangle struct {
	Base   float64
	Height float64
}

//Area is Triangle Method
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

//Rectangle is struct
type Rectangle struct {
	Width  float64
	Height float64
}

//Area is Rectangle Method
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//Circle is struct
type Circle struct {
	Radius float64
}

//Area is Circle Method
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Perimeter is func
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

//Area is func
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
