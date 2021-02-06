package mathclock

import (
	"bytes"
	"encoding/xml"
	"math"
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {

	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := Point{X: 150, Y: 150 - 90}
	got := SecondHand(tm)

	if got != want {
		t.Errorf("got %v, but want %v", got, want)
	}
}

func TestSecondHandAt30Seconds(t *testing.T) {

	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	want := Point{X: 150, Y: 150 + 90}
	got := SecondHand(tm)

	if got != want {
		t.Errorf("got %v, but want %v", got, want)
	}
}

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 60), 0},
	}

	for _, c := range cases {
		t.Run(getTestName(c.time), func(t *testing.T) {
			got := secondsInRadians(c.time)

			if got != c.angle {
				t.Errorf("got %v rad, but want %v", got, c.angle)
			}
		})
	}
}

func TestSecondHandVector(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(getTestName(c.time), func(t *testing.T) {
			got := secondHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Errorf("got %v, but want %v", got, c.point)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, c := range cases {
		t.Run(getTestName(c.time), func(t *testing.T) {
			got := minutesInRadians(c.time)
			if got != c.angle {
				t.Fatalf("Wanted %v, but got %v", c.angle, got)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(getTestName(c.time), func(t *testing.T) {
			got := minuteHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("Wanted %v Point, but got %v", c.point, got)
			}
		})
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(2020, time.January, 1, hours, minutes, seconds, 0, time.UTC)
}

func getTestName(t time.Time) string {
	return t.Format("15:04:05")
}

func roughlyEqualFloat(a, b float64) bool {
	const threshhold = 1e-7
	return math.Abs(a-b) < threshhold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat(a.X, b.X) && roughlyEqualFloat(a.Y, b.Y)
}

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}
type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}
type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 60},
		},
		{
			simpleTime(0, 0, 30),
			Line{150, 150, 150, 240},
		},
	}

	for _, c := range cases {
		t.Run(getTestName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)
			want := c.line

			if !containsLine(want, svg.Line) {
				t.Errorf("Expected to find the secondhand line %+v in the SVG output %+v", want, svg.Line)
			}
		})
	}
}

func TestSVGWriterMinuteHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 70},
		},
	}

	for _, c := range cases {
		t.Run(getTestName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the minutehand line %+v in the SVG output %+v", c.line, svg.Line)
			}
		})
	}
}

func containsLine(line Line, lines []Line) bool {
	for _, l := range lines {
		if l == line {
			return true
		}
	}
	return false
}
