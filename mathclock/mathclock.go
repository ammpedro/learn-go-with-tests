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
	clockCenterX     = 150
	clockCenterY     = 150
)

//SVGWriter writes and SVG representation of an analog clock
func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	io.WriteString(w, svgEnd)
}

func secondHand(w io.Writer, t time.Time) {
	p := secondHandPoint(t)
	//scale for the secondhand
	p = Point{p.X * secondHandLength, p.Y * secondHandLength}
	//flip, because the origin is in the top left
	p = Point{p.X, -p.Y}
	//tanslate, because the origin is in {150,150}
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`

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
	return angleToPoint(secondsInRadians(t))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / 60) + (math.Pi / (30 / float64(t.Minute())))
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
	p := minuteHandPoint(t)
	p = Point{p.X * minuteHandLength, p.Y * minuteHandLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}
