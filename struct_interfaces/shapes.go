package structinterfaces

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base float64
	Height float64
}

func (rect Rectangle) Area() float64 {
	return rect.Height * rect.Width
}

func (circle Circle) Area() float64 {
	return math.Pow(circle.Radius, 2) * math.Pi
}

func (tri Triangle) Area() float64{
	return tri.Base * tri.Height * 0.5
}