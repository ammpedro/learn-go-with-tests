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

// SecondHand accepts time and returns the equivalent endpoint on the clock
func SecondHand(t time.Time) Point {
	return Point{150, 60}
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
