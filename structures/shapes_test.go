package structures

import "testing"

func assertFloatValue(t *testing.T, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("got %g, but want %g", got, want)
	}
}

func checkArea(t *testing.T, shape Shape, want float64) {
	t.Helper()

	got := shape.Area()
	if got != want {
		t.Errorf("%#v got %g, but want %g", shape, got, want)
	}
}

func TestPerimeter(t *testing.T) {

	t.Run("should return perimeter of rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		assertFloatValue(t, got, want)

	})
}

func TestArea(t *testing.T) {

	t.Run("should return area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 3.0}
		got := rectangle.Area()
		want := 30.0

		assertFloatValue(t, got, want)
	})

	t.Run("should return area of a circle", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793

		checkArea(t, circle, want)
	})

	t.Run("should run a table driven test", func(t *testing.T) {
		areaTests := []struct {
			testCase     string
			shape        Shape
			expectedArea float64
		}{
			{testCase: "case1:Rectangle", shape: Rectangle{Width: 10.0, Height: 3.0}, expectedArea: 30.0},
			{testCase: "case2:Rectangle", shape: Rectangle{Width: 12.0, Height: 6.0}, expectedArea: 72.0},
			{testCase: "case3:Rectangle", shape: Rectangle{Width: 12.0, Height: -6.0}, expectedArea: -72.0},
			{testCase: "case4:Circle", shape: Circle{Radius: 10.0}, expectedArea: 314.1592653589793},
			{testCase: "case5:Circle", shape: Circle{Radius: 2.0}, expectedArea: 12.566370614359172},
			{testCase: "case6:Circle", shape: Circle{Radius: -2.0}, expectedArea: 12.566370614359172},
			{testCase: "case7:Triangle", shape: Triangle{Height: 12, Base: 6}, expectedArea: 36.0},
		}

		for _, tt := range areaTests {
			t.Run(tt.testCase, func(t *testing.T) {
				checkArea(t, tt.shape, tt.expectedArea)
			})

		}
	})
}
