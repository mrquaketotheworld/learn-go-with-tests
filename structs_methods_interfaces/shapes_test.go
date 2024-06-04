package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name string
		Shape
		want float64
	}{
		{name: "Rectangle", Shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", Shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", Shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.Shape, got, tt.want)
			}
		})
	}
}
