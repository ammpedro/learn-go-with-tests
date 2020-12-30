package structures

import "math"

// Shape is a gemetric element defined by an Area
type Shape interface {
	Area() float64
}

// Rectangle is a shape defined by its width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter returns 2*(width+height)
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area returns width*height
func (r Rectangle) Area() float64 {
	return (r.Width * r.Height)
}

// Circle is a shape defined by its radius
type Circle struct {
	Radius float64
}

// Area returns Pi * Radius^2
func (c Circle) Area() float64 {
	return (math.Pi * math.Pow(c.Radius, 2))
}

// Triangle is a shape defined by the length of its base and height
type Triangle struct {
	Base   float64
	Height float64
}

// Area returns (base * height) / 2
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
