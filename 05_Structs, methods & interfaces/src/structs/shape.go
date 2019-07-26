package structs

import "math"

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

//둘레
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// go는 오버로딩을 허용하지 않기 때문에 rectagle과 circle에 대해 동시에 함수사용이 불가하다.
// //넓이
// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Width * rectangle.Height
// }
