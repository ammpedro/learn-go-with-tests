package mathclock

import (
	"fmt"
	"io"
	"math"
	"time"
)

//Point is a 2D coordinate
type Point struct {
	X float64
	Y float64
}

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
	clockCenterX     = 150
	clockCenterY     = 150
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`

//SVGWriter writes and SVG representation of an analog clock
func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	io.WriteString(w, svgEnd)
}

func getHandEndpoints(p Point, handLength float64) Point {
	//scale for the secondhand
	p = Point{p.X * handLength, p.Y * handLength}
	//flip, because the origin is in the top left
	p = Point{p.X, -p.Y}
	//tanslate, because the origin is in {150,150}
	return Point{p.X + clockCenterX, p.Y + clockCenterY}
}

func secondHand(w io.Writer, t time.Time) {
	p := getHandEndpoints(secondHandPoint(t), secondHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

// secondsInRadians accepts time and returns radian value
func secondsInRadians(t time.Time) float64 {
	return math.Pi / (float64(secondsInHalfClock) / float64(t.Second()))
}

// secondHandPoint accepts time and returns the endpoint on a unit circle
func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / float64(secondsInClock)) + (math.Pi / (float64(secondsInHalfClock) / float64(t.Minute())))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func minuteHand(w io.Writer, t time.Time) {
	p := getHandEndpoints(minuteHandPoint(t), float64(minuteHandLength))
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / float64(hoursInClock)) + (math.Pi / (6 / float64(t.Hour()%hoursInClock)))
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func hourHand(w io.Writer, t time.Time) {
	p := getHandEndpoints(hourHandPoint(t), float64(hourHandLength))
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}
