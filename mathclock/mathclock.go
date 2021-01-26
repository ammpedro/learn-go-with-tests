package mathclock

import (
	"math"
	"time"
)

//Point is a 2D coordinate
type Point struct {
	X float64
	Y float64
}

const secondHandLength = 90
const clockCenterX = 150
const clockCenterY = 150

// SecondHand accepts time and returns the equivalent endpoint on the clock
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	//scale for the secondhand
	p = Point{p.X * secondHandLength, p.Y * secondHandLength}
	//flip, because the origin is in the top left
	p = Point{p.X, -p.Y}
	//tanslate, because the origin is in {150,150}
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}
	return p
}

// secondsInRadians accepts time and returns radian value
func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}

// secondHandPoint accepts time and returns the endpoint on a unit circle
func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
